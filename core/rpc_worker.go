package core

import (
	"github.com/nodersteam/cosmos-indexer/clients"
	"net/http"
	"sync"

	"github.com/DefiantLabs/probe/client"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	txTypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/nodersteam/cosmos-indexer/config"
	dbTypes "github.com/nodersteam/cosmos-indexer/db"
	"github.com/nodersteam/cosmos-indexer/rpc"
	"gorm.io/gorm"
)

type IndexerBlockEventData struct {
	BlockData                *ctypes.ResultBlock
	BlockResultsData         *ctypes.ResultBlockResults
	BlockEventRequestsFailed bool
	GetTxsResponse           *txTypes.GetTxsEventResponse
	TxRequestsFailed         bool
	IndexBlockEvents         bool
	IndexTransactions        bool
}

type BlockRPCWorker interface {
	Worker(wg *sync.WaitGroup, blockEnqueueChan chan *EnqueueData, outputChannel chan IndexerBlockEventData)
}

type blockRPCWorker struct {
	chainStringID string
	cfg           *config.IndexConfig
	chainClient   *client.ChainClient
	db            *gorm.DB
	rpcClient     clients.ChainRPC
}

func NewBlockRPCWorker(
	chainStringID string,
	cfg *config.IndexConfig,
	chainClient *client.ChainClient,
	db *gorm.DB,
	rpcClient clients.ChainRPC) BlockRPCWorker {
	return &blockRPCWorker{
		chainStringID: chainStringID,
		cfg:           cfg,
		chainClient:   chainClient,
		db:            db,
		rpcClient:     rpcClient,
	}
}

// Worker This function is responsible for making all RPC requests to the chain needed for later processing.
// The indexer relies on a number of RPC endpoints for full block data, including block event and transaction searches.
func (w *blockRPCWorker) Worker(wg *sync.WaitGroup, blockEnqueueChan chan *EnqueueData, outputChannel chan IndexerBlockEventData) {
	defer wg.Done()
	rpcClient := rpc.URIClient{
		Address: w.chainClient.Config.RPCAddr,
		Client:  &http.Client{},
	}

	for {
		// Get the next block to process
		block, open := <-blockEnqueueChan
		if !open {
			config.Log.Debugf("Block enqueue channel closed. Exiting RPC worker.")
			break
		}

		currentHeightIndexerData := IndexerBlockEventData{
			BlockEventRequestsFailed: false,
			TxRequestsFailed:         false,
			IndexBlockEvents:         block.IndexBlockEvents,
			IndexTransactions:        block.IndexTransactions,
		}

		// Get the block from the RPC
		blockData, err := w.rpcClient.GetBlock(block.Height)
		if err != nil {
			// This is the only response we continue on. If we can't get the block, we can't index anything.
			config.Log.Errorf("Error getting block %v from RPC. Err: %v", block, err)
			err := dbTypes.UpsertFailedEventBlock(w.db, block.Height, w.chainStringID, w.cfg.Probe.ChainName)
			if err != nil {
				config.Log.Fatal("Failed to insert failed block event", err)
			}
			err = dbTypes.UpsertFailedBlock(w.db, block.Height, w.chainStringID, w.cfg.Probe.ChainName)
			if err != nil {
				config.Log.Fatal("Failed to insert failed block", err)
			}
			continue
		}

		currentHeightIndexerData.BlockData = blockData

		if block.IndexBlockEvents {
			bresults, err := rpc.GetBlockResultWithRetry(rpcClient,
				block.Height, w.cfg.Base.RequestRetryAttempts, w.cfg.Base.RequestRetryMaxWait)

			if err != nil {
				config.Log.Errorf("Error getting block results for block %v from RPC. Err: %v", block, err)
				err := dbTypes.UpsertFailedEventBlock(w.db, block.Height, w.chainStringID, w.cfg.Probe.ChainName)
				if err != nil {
					config.Log.Fatal("Failed to insert failed block event", err)
				}
				currentHeightIndexerData.BlockResultsData = nil
				currentHeightIndexerData.BlockEventRequestsFailed = true
			} else {
				currentHeightIndexerData.BlockResultsData = bresults
			}
		}

		if block.IndexTransactions {
			txsEventResp, err := w.rpcClient.GetTxsByBlockHeight(block.Height)

			if err != nil {
				// Attempt to get block results to attempt an in-app codec decode of transactions.
				if currentHeightIndexerData.BlockResultsData == nil {

					bresults, err := rpc.GetBlockResultWithRetry(rpcClient, block.Height,
						w.cfg.Base.RequestRetryAttempts, w.cfg.Base.RequestRetryMaxWait)

					if err != nil {
						config.Log.Errorf("Error getting txs for block %v from RPC. Err: %v", block, err)
						err := dbTypes.UpsertFailedBlock(w.db, block.Height, w.chainStringID, w.cfg.Probe.ChainName)
						if err != nil {
							config.Log.Fatal("Failed to insert failed block", err)
						}
						currentHeightIndexerData.GetTxsResponse = nil
						currentHeightIndexerData.BlockResultsData = nil
						// Only set failed when we can't get the block results either.
						currentHeightIndexerData.TxRequestsFailed = true
					} else {
						currentHeightIndexerData.BlockResultsData = bresults
					}

				}
			} else {
				currentHeightIndexerData.GetTxsResponse = txsEventResp
			}
		}

		outputChannel <- currentHeightIndexerData
	}
}
