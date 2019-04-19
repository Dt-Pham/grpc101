package main

import (
	"context"
	"fmt"
	"grpc101/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to the port %v", err)
	}

	c := calculatorpb.NewCalculatorServiceClient(cc)
	doSum(c)
}

func doSum(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 9,
	}

	resp, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Println("Error while receiving response from server", err)
		return
	}

	fmt.Println("SumResponse =", resp.Sum)
}
