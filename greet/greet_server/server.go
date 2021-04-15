package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/msyamsula/goGRPC/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("greet function was invoke")
	firstName := req.GetGreeting().GetFirstName()
	// lastName := req.GetGreeting().GetFirstName()

	result := "Hello " + firstName
	res := greetpb.GreetResponse{
		Result: result,
	}

	return &res, nil
}

func main() {
	fmt.Println("hello")

	// listener creation
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	// err checking when creating server
	if err != nil {
		log.Fatal("failed to listen")
	}

	//server creation
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err != nil {
		log.Fatal("failed to serve")
	} else {
		s.Serve(lis) // server running
	}
}
