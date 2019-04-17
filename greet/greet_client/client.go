package main

import (
	"context"
	"fmt"
	"grpc101/greet/greetpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
	doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Start doing unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Duong",
			LastName:  "Pham",
		},
	}

	resq, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Println("Response from Greet:", resq.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Start doing server streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Duong",
			LastName:  "Pham",
		},
	}

	respStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := respStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}
		log.Println("Response from GreetManyTimes: ", msg.Result)
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Start doing client streaming RPC ..")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Duong",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Duong dz",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Duong pro",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Duong 9x",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Duong khung",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	for _, req := range requests {
		fmt.Printf("Sending a request %v\n", req)
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Error while sending message: %v", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving: %v", err)
	}
	fmt.Println(res.Result)
}
