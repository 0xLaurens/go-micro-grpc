package handler

import (
	"context"
	"go-micro/proto"
	"go-micro/types"
	rpc "google.golang.org/grpc"
)

type CalculationGrpcHandler struct {
	calculatorService types.CalculatorService
	proto.UnimplementedCalculatorServer
}

func NewGrpcHandler(grpc *rpc.Server, calculatorService types.CalculatorService) *CalculationGrpcHandler {
	gRPCHandler := &CalculationGrpcHandler{
		calculatorService: calculatorService,
	}

	proto.RegisterCalculatorServer(grpc, gRPCHandler)

	return gRPCHandler
}

func (h *CalculationGrpcHandler) Add(ctx context.Context, req *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	return h.calculatorService.Add(ctx, req)
}
func (h *CalculationGrpcHandler) Subtract(ctx context.Context, req *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	return h.calculatorService.Subtract(ctx, req)
}
func (h *CalculationGrpcHandler) Multiply(ctx context.Context, req *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	return h.calculatorService.Multiply(ctx, req)
}
func (h *CalculationGrpcHandler) Divide(ctx context.Context, req *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	return h.calculatorService.Divide(ctx, req)
}
