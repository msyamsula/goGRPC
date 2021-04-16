package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/msyamsula/goGRPC/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("hello i'm client")

	// connection to server
	// remove with insercure when moving to prod
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("could not connect")
	}
	defer conn.Close() // close when done

	// create client
	c := greetpb.NewGreetServiceClient(conn)
	// fmt.Printf("Created client: %f", c)
	doUnary(c)
	// GreetWithDeadline(c, 1*time.Second)
	// GreetWithDeadline(c, 2*time.Second)
	GreetWithDeadline(c, 1*time.Second)

}

// function to call greet
func doUnary(c greetpb.GreetServiceClient) {
	greeting := &greetpb.Greeting{
		FirstName: "syamsul",
		LastName:  "arifin",
	}

	req := &greetpb.GreetRequest{
		Greeting: greeting,
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Printf("Error %v", err)
	}

	log.Printf("Response from Greet: %v", res)
}

// unary with deadline
func GreetWithDeadline(c greetpb.GreetServiceClient, d time.Duration) {
	// GreetWithDeadline(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "syamsul",
			LastName:  "arifin",
		},
	}
	// create request with context deadline
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()

	// make procedure call and handle response
	resp, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			if respErr.Code() == codes.DeadlineExceeded {
				fmt.Printf("Deadline Exceeded: %v\n", respErr.Message())
			}
		} else {
			fmt.Println("unexpected error")
		}
		log.Fatalf("%v", err)
		return
	}

	fmt.Printf("get response: %v", resp.GetResult())
}
