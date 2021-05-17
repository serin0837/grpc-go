package main

import (
	"context"
	"fmt"
	"go-grpc/calculator/calculatepb"
	"log"
	"net"

	"google.golang.org/grpc"
)

//3. create server struct
type server struct{}

// 4. server unary inferface implementing
func (*server) Calculate(ctx context.Context, req *calculatepb.CalculateRequest) (*calculatepb.CalculateResponse, error) {
	fmt.Printf("Calculate func invoked with %v", req)
	firstNumber := req.GetCalculating().GetFirstNumber()
	//also can do req.FirstNumber
	//req.SecondNumber
	secondNumber := req.GetCalculating().GetSecondNumber()

	result := firstNumber + secondNumber
	res := &calculatepb.CalculateResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("calculate server")
	//1. make listener//open tcp connection port for grpc 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	//2. register service with server
	calculatepb.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
