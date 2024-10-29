package main

import (
	"context"
	"fmt"
	"log"

	pb "client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	connString := "127.0.0.1:50051"
	clientConn, err := grpc.NewClient(connString, grpc.WithTransportCredentials(
		insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		err := clientConn.Close()
		fmt.Printf("close error: %v\n", err)
	}()

	devClient := pb.NewDeviceClient(clientConn)
	{
		req := &pb.GetIdentityRequest{
			TransactionId: 10,
			DeviceId:      4411442211,
			CommandName:   "get_identity",
		}
		res, err := devClient.GetIdentity(ctx, req)
		if err != nil {
			log.Fatalf("response error: %v\n", err)
		}

		// ***

		fmt.Printf("TransactionId: %v\n", res.TransactionId)
		fmt.Printf("DeviceId: %v\n", res.DeviceId)
		fmt.Printf("DeviceType: %v\n", res.DeviceType)
		fmt.Printf("Serial: %v\n", res.Serial)
	}

	{
		_, err := devClient.GenerateError(ctx, &pb.Empty{})
		if err != nil {
			log.Fatalf("response error: %v\n", err)
		}
	}
}
