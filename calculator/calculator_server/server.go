package main

import (
	"context"
	"fmt"
	"grpc101/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Sum was invoked with req: %+v\n", *req)
	return &calculatorpb.SumResponse{Sum: req.FirstNumber + req.SecondNumber}, nil
}

func main() {
	fmt.Println("Calculator on duty !!!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Cannot listen to the port %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("Listening to the port 50051...")
	if s.Serve(lis) != nil {
		log.Fatalf("Fail to serve %v", err)
	}
}
