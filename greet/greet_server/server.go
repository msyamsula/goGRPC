package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/msyamsula/goGRPC/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	first_name := req.GetGreeting().GetFirstName()

	// simulate 3 second delay
	for i := 0; i < 3; i++ {
		// check context error for each second
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("the client canceled the request")
			return nil, status.Error(codes.Canceled, "client canceled the request")
		}

		time.Sleep(1 * time.Second)
	}

	result := "Hello " + first_name + "!!!!\n"
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
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
