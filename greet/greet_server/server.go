package main

import (
	"context"
	"fmt"
	"grpc101/greet/greetpb"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

const port = 50051

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", req)

	firstName := req.GetGreeting().GetFirstName()
	resp := &greetpb.GreetResponse{Result: "Hello " + firstName}
	return resp, nil
}

func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes was invoked with %v\n", req)

	for i := 1; i <= 10; i++ {
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		resp := &greetpb.GreetManyTimesResponse{Result: result}
		stream.Send(resp)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (s *server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Println("LongGreet was invoked with a stream")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{Result: result})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Printf("Listening on port %d...\n", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
