package grpc

import (
	"go-micro/handler"
	"go-micro/services"
	rpc "google.golang.org/grpc"
	"net"
)

type GRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *GRPCServer {
	return &GRPCServer{
		addr: addr,
	}
}

func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	grpcServer := rpc.NewServer()
	calcService := services.NewCalculationService()
	handler.NewGrpcHandler(grpcServer, calcService)

	return grpcServer.Serve(lis)
}
