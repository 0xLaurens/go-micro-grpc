package main

import (
	"context"
	"go-micro/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.NewClient(":7331", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewCalculatorClient(conn)
	res, err := client.Add(ctx, &proto.CalculatorRequest{A: 10, B: 5})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)

	res, _ = client.Subtract(ctx, &proto.CalculatorRequest{A: 10, B: 5})
	log.Println(res)

	res, _ = client.Multiply(ctx, &proto.CalculatorRequest{A: 10, B: 5})
	log.Println(res)

	res, _ = client.Divide(ctx, &proto.CalculatorRequest{A: 10, B: 5})
	log.Println(res)
}
