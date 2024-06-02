package main

import "go-micro/adapters/grpc"

func main() {
	server := grpc.NewGRPCServer(":7331")
	err := server.Run()
	if err != nil {
		return
	}
}
