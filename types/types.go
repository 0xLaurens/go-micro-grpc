package types

import (
	"context"
	"go-micro/proto"
)

type CalculatorService interface {
	Add(ctx context.Context, request *proto.CalculatorRequest) (*proto.CalculatorResponse, error)
	Subtract(ctx context.Context, request *proto.CalculatorRequest) (*proto.CalculatorResponse, error)
	Multiply(ctx context.Context, request *proto.CalculatorRequest) (*proto.CalculatorResponse, error)
	Divide(ctx context.Context, request *proto.CalculatorRequest) (*proto.CalculatorResponse, error)
}
