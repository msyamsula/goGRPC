package main

import (
	"context"
	"fmt"
	"log"

	"github.com/msyamsula/goGRPC/greet/greetpb"
	"google.golang.org/grpc"
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
