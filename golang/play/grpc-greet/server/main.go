package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	pb "server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// -----------------------------------------------------------------------

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

type deviceServer struct {
	pb.UnimplementedDeviceServer
}

type chatServer struct {
	pb.UnimplementedChatServer
}

// Greeter
// -----------------------------------------------------------------------

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (
	*pb.HelloReply, error) {
	message := fmt.Sprintf("Hello, %s!", req.GetName())
	return &pb.HelloReply{
		Message: message,
	}, nil
}

func (s *greeterServer) Ping(ctx context.Context, req *pb.PingReqRes) (
	*pb.PingReqRes, error) {
	return &pb.PingReqRes{
		Value: "pong",
	}, nil
}

// Device
// -----------------------------------------------------------------------

func (s *deviceServer) GetIdentity(ctx context.Context, req *pb.GetIdentityRequest) (
	*pb.GetIdentityResponse, error) {

	log.Printf("DeviceId: %v", req.DeviceId)
	log.Printf("TransactionId: %v", req.TransactionId)
	log.Printf("CommandName: %v", req.CommandName)

	// ***

	res := pb.GetIdentityResponse{
		TransactionId: req.TransactionId,
		DeviceId:      req.DeviceId,

		DeviceType: uint64(rand.Intn(10) + 1), // if 0 then empty!
		Serial:     uint64(rand.Intn(10) + 1),
	}
	return &res, nil
}

var (
	GlobalNumber uint64 = 1000
)

func (s *deviceServer) SetNumberToRandom(context.Context, *pb.Empty) (*pb.Empty, error) {
	GlobalNumber = rand.Uint64()%100 + 1
	return &pb.Empty{}, nil
}

func (s *deviceServer) GetNumber(context.Context, *pb.Empty) (*pb.GetNumberRespons, error) {
	return &pb.GetNumberRespons{
		Number: int64(GlobalNumber),
	}, nil
}

func (s *deviceServer) GenerateError(context.Context, *pb.Empty) (*pb.Empty, error) {
	return nil,
		status.Errorf(codes.Unimplemented, "method GenerateError not implemented")
}

// Chat
// -----------------------------------------------------------------------

func (s *chatServer) SendMessages(stream grpc.ClientStreamingServer[pb.ChatMessage, pb.ChatResponse]) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("Received message from %s: %s",
			msg.User, msg.Message)
	}

	return stream.SendAndClose(
		&pb.ChatResponse{Confirmation: "Messages received"})
}

func (s *chatServer) StreamMessages(req *pb.ChatRequest, stream grpc.ServerStreamingServer[pb.ChatMessage]) error {
	for i := 0; i < 5; i++ {
		msg := &pb.ChatMessage{
			User:    req.User,
			Message: fmt.Sprintf("Message %d from %s", i+1, req.User),
		}
		if err := stream.Send(msg); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *chatServer) ExecuteChat(stream grpc.BidiStreamingServer[pb.ChatMessage, pb.ChatResponse]) error {
	counter := 0
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("Received message from %s: `%s`", msg.User, msg.Message)
		counter++

		if counter > 5 {
			reply := &pb.ChatResponse{Confirmation: "OK"}
			if err := stream.Send(reply); err != nil {
				return err
			}

			counter = 0
		}
	}
	return nil
}

// -----------------------------------------------------------------------

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreeterServer(grpcServer, &greeterServer{})
	pb.RegisterDeviceServer(grpcServer, &deviceServer{})
	pb.RegisterChatServer(grpcServer, &chatServer{})

	reflection.Register(grpcServer) // !

	log.Println("The server is running on: 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// -----------------------------------------------------------------------
