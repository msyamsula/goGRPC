package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/msyamsula/goGRPC/client_streaming/clientstreampb"

	"google.golang.org/grpc"
)

func StreamFromClient(c clientstreampb.GreetServiceClient) {
	stream, err := c.Greet(context.Background())
	if err != nil {
		log.Printf("%v", err)
	}

	reqs := []*clientstreampb.ClientStreamRequest{
		&clientstreampb.ClientStreamRequest{
			Name: "syamsul",
		},
		&clientstreampb.ClientStreamRequest{
			Name: "Arifin",
		},
		&clientstreampb.ClientStreamRequest{
			Name: "Enak",
		},
		&clientstreampb.ClientStreamRequest{
			Name: "Wikwik",
		},
	}

	// equivalent with for item in list in python
	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("%v", err)
	}

	fmt.Printf("%v", resp.GetRes())
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Print("error")
	}

	c := clientstreampb.NewGreetServiceClient(conn)

	StreamFromClient(c)
}
