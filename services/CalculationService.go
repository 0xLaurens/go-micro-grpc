package services

import (
	"context"
	"go-micro/proto"
)

type CalculationService struct {
}

func NewCalculationService() *CalculationService {
	return &CalculationService{}
}

func (c *CalculationService) Add(ctx context.Context, request *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	return &proto.CalculatorResponse{Result: request.A + request.B}, nil
}

func (c *CalculationService) Subtract(ctx context.Context, request *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	return &proto.CalculatorResponse{Result: request.A - request.B}, nil
}

func (c *CalculationService) Multiply(ctx context.Context, request *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	return &proto.CalculatorResponse{Result: request.A * request.B}, nil
}

func (c *CalculationService) Divide(ctx context.Context, request *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	return &proto.CalculatorResponse{Result: request.A / request.B}, nil
}

func (c *CalculationService) mustEmbedUnimplementedCalculatorServer() {
	//TODO implement me
	panic("implement me")
}
