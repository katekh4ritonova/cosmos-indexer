// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: blocks.proto

package blocks

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	BlocksService_BlockInfo_FullMethodName               = "/blocks.BlocksService/BlockInfo"
	BlocksService_BlockInfoByHash_FullMethodName         = "/blocks.BlocksService/BlockInfoByHash"
	BlocksService_BlockValidators_FullMethodName         = "/blocks.BlocksService/BlockValidators"
	BlocksService_TxChartByDay_FullMethodName            = "/blocks.BlocksService/TxChartByDay"
	BlocksService_TxByHash_FullMethodName                = "/blocks.BlocksService/TxByHash"
	BlocksService_TotalTransactions_FullMethodName       = "/blocks.BlocksService/TotalTransactions"
	BlocksService_Transactions_FullMethodName            = "/blocks.BlocksService/Transactions"
	BlocksService_TotalBlocks_FullMethodName             = "/blocks.BlocksService/TotalBlocks"
	BlocksService_GetBlocks_FullMethodName               = "/blocks.BlocksService/GetBlocks"
	BlocksService_BlockSignatures_FullMethodName         = "/blocks.BlocksService/BlockSignatures"
	BlocksService_TxsByBlock_FullMethodName              = "/blocks.BlocksService/TxsByBlock"
	BlocksService_TransactionRawLog_FullMethodName       = "/blocks.BlocksService/TransactionRawLog"
	BlocksService_TransactionSigners_FullMethodName      = "/blocks.BlocksService/TransactionSigners"
	BlocksService_CacheTransactions_FullMethodName       = "/blocks.BlocksService/CacheTransactions"
	BlocksService_CacheGetBlocks_FullMethodName          = "/blocks.BlocksService/CacheGetBlocks"
	BlocksService_CacheAggregated_FullMethodName         = "/blocks.BlocksService/CacheAggregated"
	BlocksService_SearchHashByText_FullMethodName        = "/blocks.BlocksService/SearchHashByText"
	BlocksService_ChartTransactionsByHour_FullMethodName = "/blocks.BlocksService/ChartTransactionsByHour"
	BlocksService_ChartTransactionsVolume_FullMethodName = "/blocks.BlocksService/ChartTransactionsVolume"
	BlocksService_BlockUpTime_FullMethodName             = "/blocks.BlocksService/BlockUpTime"
	BlocksService_UptimeByBlocks_FullMethodName          = "/blocks.BlocksService/UptimeByBlocks"
	BlocksService_GetVotes_FullMethodName                = "/blocks.BlocksService/GetVotes"
)

// BlocksServiceClient is the client API for BlocksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlocksServiceClient interface {
	BlockInfo(ctx context.Context, in *GetBlockInfoRequest, opts ...grpc.CallOption) (*GetBlockInfoResponse, error)
	BlockInfoByHash(ctx context.Context, in *BlockInfoByHashRequest, opts ...grpc.CallOption) (*BlockInfoByHashResponse, error)
	BlockValidators(ctx context.Context, in *GetBlockValidatorsRequest, opts ...grpc.CallOption) (*GetBlockValidatorsResponse, error)
	TxChartByDay(ctx context.Context, in *TxChartByDayRequest, opts ...grpc.CallOption) (*TxChartByDayResponse, error)
	TxByHash(ctx context.Context, in *TxByHashRequest, opts ...grpc.CallOption) (*TxByHashResponse, error)
	TotalTransactions(ctx context.Context, in *TotalTransactionsRequest, opts ...grpc.CallOption) (*TotalTransactionsResponse, error)
	Transactions(ctx context.Context, in *TransactionsRequest, opts ...grpc.CallOption) (*TransactionsResponse, error)
	TotalBlocks(ctx context.Context, in *TotalBlocksRequest, opts ...grpc.CallOption) (*TotalBlocksResponse, error)
	GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksResponse, error)
	BlockSignatures(ctx context.Context, in *BlockSignaturesRequest, opts ...grpc.CallOption) (*BlockSignaturesResponse, error)
	TxsByBlock(ctx context.Context, in *TxsByBlockRequest, opts ...grpc.CallOption) (*TxsByBlockResponse, error)
	TransactionRawLog(ctx context.Context, in *TransactionRawLogRequest, opts ...grpc.CallOption) (*TransactionRawLogResponse, error)
	TransactionSigners(ctx context.Context, in *TransactionSignersRequest, opts ...grpc.CallOption) (*TransactionSignersResponse, error)
	CacheTransactions(ctx context.Context, in *TransactionsRequest, opts ...grpc.CallOption) (*TransactionsResponse, error)
	CacheGetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksResponse, error)
	CacheAggregated(ctx context.Context, in *CacheAggregatedRequest, opts ...grpc.CallOption) (*CacheAggregatedResponse, error)
	SearchHashByText(ctx context.Context, in *SearchHashByTextRequest, opts ...grpc.CallOption) (*SearchHashByTextResponse, error)
	ChartTransactionsByHour(ctx context.Context, in *ChartTransactionsByHourRequest, opts ...grpc.CallOption) (*ChartTransactionsByHourResponse, error)
	ChartTransactionsVolume(ctx context.Context, in *ChartTransactionsVolumeRequest, opts ...grpc.CallOption) (*ChartTransactionsVolumeResponse, error)
	BlockUpTime(ctx context.Context, in *BlockUpTimeRequest, opts ...grpc.CallOption) (*BlockUpTimeResponse, error)
	UptimeByBlocks(ctx context.Context, in *UptimeByBlocksRequest, opts ...grpc.CallOption) (*UptimeByBlocksResponse, error)
	GetVotes(ctx context.Context, in *GetVotesRequest, opts ...grpc.CallOption) (*GetVotesResponse, error)
}

type blocksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlocksServiceClient(cc grpc.ClientConnInterface) BlocksServiceClient {
	return &blocksServiceClient{cc}
}

func (c *blocksServiceClient) BlockInfo(ctx context.Context, in *GetBlockInfoRequest, opts ...grpc.CallOption) (*GetBlockInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBlockInfoResponse)
	err := c.cc.Invoke(ctx, BlocksService_BlockInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) BlockInfoByHash(ctx context.Context, in *BlockInfoByHashRequest, opts ...grpc.CallOption) (*BlockInfoByHashResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockInfoByHashResponse)
	err := c.cc.Invoke(ctx, BlocksService_BlockInfoByHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) BlockValidators(ctx context.Context, in *GetBlockValidatorsRequest, opts ...grpc.CallOption) (*GetBlockValidatorsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBlockValidatorsResponse)
	err := c.cc.Invoke(ctx, BlocksService_BlockValidators_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) TxChartByDay(ctx context.Context, in *TxChartByDayRequest, opts ...grpc.CallOption) (*TxChartByDayResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TxChartByDayResponse)
	err := c.cc.Invoke(ctx, BlocksService_TxChartByDay_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) TxByHash(ctx context.Context, in *TxByHashRequest, opts ...grpc.CallOption) (*TxByHashResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TxByHashResponse)
	err := c.cc.Invoke(ctx, BlocksService_TxByHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) TotalTransactions(ctx context.Context, in *TotalTransactionsRequest, opts ...grpc.CallOption) (*TotalTransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TotalTransactionsResponse)
	err := c.cc.Invoke(ctx, BlocksService_TotalTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) Transactions(ctx context.Context, in *TransactionsRequest, opts ...grpc.CallOption) (*TransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionsResponse)
	err := c.cc.Invoke(ctx, BlocksService_Transactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) TotalBlocks(ctx context.Context, in *TotalBlocksRequest, opts ...grpc.CallOption) (*TotalBlocksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TotalBlocksResponse)
	err := c.cc.Invoke(ctx, BlocksService_TotalBlocks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBlocksResponse)
	err := c.cc.Invoke(ctx, BlocksService_GetBlocks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) BlockSignatures(ctx context.Context, in *BlockSignaturesRequest, opts ...grpc.CallOption) (*BlockSignaturesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockSignaturesResponse)
	err := c.cc.Invoke(ctx, BlocksService_BlockSignatures_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) TxsByBlock(ctx context.Context, in *TxsByBlockRequest, opts ...grpc.CallOption) (*TxsByBlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TxsByBlockResponse)
	err := c.cc.Invoke(ctx, BlocksService_TxsByBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) TransactionRawLog(ctx context.Context, in *TransactionRawLogRequest, opts ...grpc.CallOption) (*TransactionRawLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionRawLogResponse)
	err := c.cc.Invoke(ctx, BlocksService_TransactionRawLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) TransactionSigners(ctx context.Context, in *TransactionSignersRequest, opts ...grpc.CallOption) (*TransactionSignersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionSignersResponse)
	err := c.cc.Invoke(ctx, BlocksService_TransactionSigners_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) CacheTransactions(ctx context.Context, in *TransactionsRequest, opts ...grpc.CallOption) (*TransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionsResponse)
	err := c.cc.Invoke(ctx, BlocksService_CacheTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) CacheGetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBlocksResponse)
	err := c.cc.Invoke(ctx, BlocksService_CacheGetBlocks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) CacheAggregated(ctx context.Context, in *CacheAggregatedRequest, opts ...grpc.CallOption) (*CacheAggregatedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CacheAggregatedResponse)
	err := c.cc.Invoke(ctx, BlocksService_CacheAggregated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) SearchHashByText(ctx context.Context, in *SearchHashByTextRequest, opts ...grpc.CallOption) (*SearchHashByTextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchHashByTextResponse)
	err := c.cc.Invoke(ctx, BlocksService_SearchHashByText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) ChartTransactionsByHour(ctx context.Context, in *ChartTransactionsByHourRequest, opts ...grpc.CallOption) (*ChartTransactionsByHourResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChartTransactionsByHourResponse)
	err := c.cc.Invoke(ctx, BlocksService_ChartTransactionsByHour_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) ChartTransactionsVolume(ctx context.Context, in *ChartTransactionsVolumeRequest, opts ...grpc.CallOption) (*ChartTransactionsVolumeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChartTransactionsVolumeResponse)
	err := c.cc.Invoke(ctx, BlocksService_ChartTransactionsVolume_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) BlockUpTime(ctx context.Context, in *BlockUpTimeRequest, opts ...grpc.CallOption) (*BlockUpTimeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockUpTimeResponse)
	err := c.cc.Invoke(ctx, BlocksService_BlockUpTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) UptimeByBlocks(ctx context.Context, in *UptimeByBlocksRequest, opts ...grpc.CallOption) (*UptimeByBlocksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UptimeByBlocksResponse)
	err := c.cc.Invoke(ctx, BlocksService_UptimeByBlocks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksServiceClient) GetVotes(ctx context.Context, in *GetVotesRequest, opts ...grpc.CallOption) (*GetVotesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVotesResponse)
	err := c.cc.Invoke(ctx, BlocksService_GetVotes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlocksServiceServer is the server API for BlocksService service.
// All implementations must embed UnimplementedBlocksServiceServer
// for forward compatibility
type BlocksServiceServer interface {
	BlockInfo(context.Context, *GetBlockInfoRequest) (*GetBlockInfoResponse, error)
	BlockInfoByHash(context.Context, *BlockInfoByHashRequest) (*BlockInfoByHashResponse, error)
	BlockValidators(context.Context, *GetBlockValidatorsRequest) (*GetBlockValidatorsResponse, error)
	TxChartByDay(context.Context, *TxChartByDayRequest) (*TxChartByDayResponse, error)
	TxByHash(context.Context, *TxByHashRequest) (*TxByHashResponse, error)
	TotalTransactions(context.Context, *TotalTransactionsRequest) (*TotalTransactionsResponse, error)
	Transactions(context.Context, *TransactionsRequest) (*TransactionsResponse, error)
	TotalBlocks(context.Context, *TotalBlocksRequest) (*TotalBlocksResponse, error)
	GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error)
	BlockSignatures(context.Context, *BlockSignaturesRequest) (*BlockSignaturesResponse, error)
	TxsByBlock(context.Context, *TxsByBlockRequest) (*TxsByBlockResponse, error)
	TransactionRawLog(context.Context, *TransactionRawLogRequest) (*TransactionRawLogResponse, error)
	TransactionSigners(context.Context, *TransactionSignersRequest) (*TransactionSignersResponse, error)
	CacheTransactions(context.Context, *TransactionsRequest) (*TransactionsResponse, error)
	CacheGetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error)
	CacheAggregated(context.Context, *CacheAggregatedRequest) (*CacheAggregatedResponse, error)
	SearchHashByText(context.Context, *SearchHashByTextRequest) (*SearchHashByTextResponse, error)
	ChartTransactionsByHour(context.Context, *ChartTransactionsByHourRequest) (*ChartTransactionsByHourResponse, error)
	ChartTransactionsVolume(context.Context, *ChartTransactionsVolumeRequest) (*ChartTransactionsVolumeResponse, error)
	BlockUpTime(context.Context, *BlockUpTimeRequest) (*BlockUpTimeResponse, error)
	UptimeByBlocks(context.Context, *UptimeByBlocksRequest) (*UptimeByBlocksResponse, error)
	GetVotes(context.Context, *GetVotesRequest) (*GetVotesResponse, error)
	mustEmbedUnimplementedBlocksServiceServer()
}

// UnimplementedBlocksServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlocksServiceServer struct {
}

func (UnimplementedBlocksServiceServer) BlockInfo(context.Context, *GetBlockInfoRequest) (*GetBlockInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockInfo not implemented")
}
func (UnimplementedBlocksServiceServer) BlockInfoByHash(context.Context, *BlockInfoByHashRequest) (*BlockInfoByHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockInfoByHash not implemented")
}
func (UnimplementedBlocksServiceServer) BlockValidators(context.Context, *GetBlockValidatorsRequest) (*GetBlockValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockValidators not implemented")
}
func (UnimplementedBlocksServiceServer) TxChartByDay(context.Context, *TxChartByDayRequest) (*TxChartByDayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxChartByDay not implemented")
}
func (UnimplementedBlocksServiceServer) TxByHash(context.Context, *TxByHashRequest) (*TxByHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxByHash not implemented")
}
func (UnimplementedBlocksServiceServer) TotalTransactions(context.Context, *TotalTransactionsRequest) (*TotalTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalTransactions not implemented")
}
func (UnimplementedBlocksServiceServer) Transactions(context.Context, *TransactionsRequest) (*TransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transactions not implemented")
}
func (UnimplementedBlocksServiceServer) TotalBlocks(context.Context, *TotalBlocksRequest) (*TotalBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalBlocks not implemented")
}
func (UnimplementedBlocksServiceServer) GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedBlocksServiceServer) BlockSignatures(context.Context, *BlockSignaturesRequest) (*BlockSignaturesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockSignatures not implemented")
}
func (UnimplementedBlocksServiceServer) TxsByBlock(context.Context, *TxsByBlockRequest) (*TxsByBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxsByBlock not implemented")
}
func (UnimplementedBlocksServiceServer) TransactionRawLog(context.Context, *TransactionRawLogRequest) (*TransactionRawLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactionRawLog not implemented")
}
func (UnimplementedBlocksServiceServer) TransactionSigners(context.Context, *TransactionSignersRequest) (*TransactionSignersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactionSigners not implemented")
}
func (UnimplementedBlocksServiceServer) CacheTransactions(context.Context, *TransactionsRequest) (*TransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheTransactions not implemented")
}
func (UnimplementedBlocksServiceServer) CacheGetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheGetBlocks not implemented")
}
func (UnimplementedBlocksServiceServer) CacheAggregated(context.Context, *CacheAggregatedRequest) (*CacheAggregatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheAggregated not implemented")
}
func (UnimplementedBlocksServiceServer) SearchHashByText(context.Context, *SearchHashByTextRequest) (*SearchHashByTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchHashByText not implemented")
}
func (UnimplementedBlocksServiceServer) ChartTransactionsByHour(context.Context, *ChartTransactionsByHourRequest) (*ChartTransactionsByHourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChartTransactionsByHour not implemented")
}
func (UnimplementedBlocksServiceServer) ChartTransactionsVolume(context.Context, *ChartTransactionsVolumeRequest) (*ChartTransactionsVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChartTransactionsVolume not implemented")
}
func (UnimplementedBlocksServiceServer) BlockUpTime(context.Context, *BlockUpTimeRequest) (*BlockUpTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUpTime not implemented")
}
func (UnimplementedBlocksServiceServer) UptimeByBlocks(context.Context, *UptimeByBlocksRequest) (*UptimeByBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UptimeByBlocks not implemented")
}
func (UnimplementedBlocksServiceServer) GetVotes(context.Context, *GetVotesRequest) (*GetVotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVotes not implemented")
}
func (UnimplementedBlocksServiceServer) mustEmbedUnimplementedBlocksServiceServer() {}

// UnsafeBlocksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlocksServiceServer will
// result in compilation errors.
type UnsafeBlocksServiceServer interface {
	mustEmbedUnimplementedBlocksServiceServer()
}

func RegisterBlocksServiceServer(s grpc.ServiceRegistrar, srv BlocksServiceServer) {
	s.RegisterService(&BlocksService_ServiceDesc, srv)
}

func _BlocksService_BlockInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).BlockInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_BlockInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).BlockInfo(ctx, req.(*GetBlockInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_BlockInfoByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockInfoByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).BlockInfoByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_BlockInfoByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).BlockInfoByHash(ctx, req.(*BlockInfoByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_BlockValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).BlockValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_BlockValidators_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).BlockValidators(ctx, req.(*GetBlockValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_TxChartByDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxChartByDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).TxChartByDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_TxChartByDay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).TxChartByDay(ctx, req.(*TxChartByDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_TxByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).TxByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_TxByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).TxByHash(ctx, req.(*TxByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_TotalTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TotalTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).TotalTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_TotalTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).TotalTransactions(ctx, req.(*TotalTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_Transactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).Transactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_Transactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).Transactions(ctx, req.(*TransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_TotalBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TotalBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).TotalBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_TotalBlocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).TotalBlocks(ctx, req.(*TotalBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_GetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).GetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_GetBlocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).GetBlocks(ctx, req.(*GetBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_BlockSignatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockSignaturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).BlockSignatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_BlockSignatures_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).BlockSignatures(ctx, req.(*BlockSignaturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_TxsByBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxsByBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).TxsByBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_TxsByBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).TxsByBlock(ctx, req.(*TxsByBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_TransactionRawLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRawLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).TransactionRawLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_TransactionRawLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).TransactionRawLog(ctx, req.(*TransactionRawLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_TransactionSigners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionSignersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).TransactionSigners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_TransactionSigners_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).TransactionSigners(ctx, req.(*TransactionSignersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_CacheTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).CacheTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_CacheTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).CacheTransactions(ctx, req.(*TransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_CacheGetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).CacheGetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_CacheGetBlocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).CacheGetBlocks(ctx, req.(*GetBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_CacheAggregated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheAggregatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).CacheAggregated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_CacheAggregated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).CacheAggregated(ctx, req.(*CacheAggregatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_SearchHashByText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchHashByTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).SearchHashByText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_SearchHashByText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).SearchHashByText(ctx, req.(*SearchHashByTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_ChartTransactionsByHour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChartTransactionsByHourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).ChartTransactionsByHour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_ChartTransactionsByHour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).ChartTransactionsByHour(ctx, req.(*ChartTransactionsByHourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_ChartTransactionsVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChartTransactionsVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).ChartTransactionsVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_ChartTransactionsVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).ChartTransactionsVolume(ctx, req.(*ChartTransactionsVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_BlockUpTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockUpTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).BlockUpTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_BlockUpTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).BlockUpTime(ctx, req.(*BlockUpTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_UptimeByBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UptimeByBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).UptimeByBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_UptimeByBlocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).UptimeByBlocks(ctx, req.(*UptimeByBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksService_GetVotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksServiceServer).GetVotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlocksService_GetVotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksServiceServer).GetVotes(ctx, req.(*GetVotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlocksService_ServiceDesc is the grpc.ServiceDesc for BlocksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlocksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blocks.BlocksService",
	HandlerType: (*BlocksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BlockInfo",
			Handler:    _BlocksService_BlockInfo_Handler,
		},
		{
			MethodName: "BlockInfoByHash",
			Handler:    _BlocksService_BlockInfoByHash_Handler,
		},
		{
			MethodName: "BlockValidators",
			Handler:    _BlocksService_BlockValidators_Handler,
		},
		{
			MethodName: "TxChartByDay",
			Handler:    _BlocksService_TxChartByDay_Handler,
		},
		{
			MethodName: "TxByHash",
			Handler:    _BlocksService_TxByHash_Handler,
		},
		{
			MethodName: "TotalTransactions",
			Handler:    _BlocksService_TotalTransactions_Handler,
		},
		{
			MethodName: "Transactions",
			Handler:    _BlocksService_Transactions_Handler,
		},
		{
			MethodName: "TotalBlocks",
			Handler:    _BlocksService_TotalBlocks_Handler,
		},
		{
			MethodName: "GetBlocks",
			Handler:    _BlocksService_GetBlocks_Handler,
		},
		{
			MethodName: "BlockSignatures",
			Handler:    _BlocksService_BlockSignatures_Handler,
		},
		{
			MethodName: "TxsByBlock",
			Handler:    _BlocksService_TxsByBlock_Handler,
		},
		{
			MethodName: "TransactionRawLog",
			Handler:    _BlocksService_TransactionRawLog_Handler,
		},
		{
			MethodName: "TransactionSigners",
			Handler:    _BlocksService_TransactionSigners_Handler,
		},
		{
			MethodName: "CacheTransactions",
			Handler:    _BlocksService_CacheTransactions_Handler,
		},
		{
			MethodName: "CacheGetBlocks",
			Handler:    _BlocksService_CacheGetBlocks_Handler,
		},
		{
			MethodName: "CacheAggregated",
			Handler:    _BlocksService_CacheAggregated_Handler,
		},
		{
			MethodName: "SearchHashByText",
			Handler:    _BlocksService_SearchHashByText_Handler,
		},
		{
			MethodName: "ChartTransactionsByHour",
			Handler:    _BlocksService_ChartTransactionsByHour_Handler,
		},
		{
			MethodName: "ChartTransactionsVolume",
			Handler:    _BlocksService_ChartTransactionsVolume_Handler,
		},
		{
			MethodName: "BlockUpTime",
			Handler:    _BlocksService_BlockUpTime_Handler,
		},
		{
			MethodName: "UptimeByBlocks",
			Handler:    _BlocksService_UptimeByBlocks_Handler,
		},
		{
			MethodName: "GetVotes",
			Handler:    _BlocksService_GetVotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blocks.proto",
}
