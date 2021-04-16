package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/msyamsula/goGRPC/bidirectional/protobuf/greetEveryOnePB"
	"google.golang.org/grpc"
)

func Greet(c greetEveryOnePB.GreetEveryOneServiceClient) {
	// Greet(ctx context.Context, opts ...grpc.CallOption) (GreetEveryOneService_GreetClient, error)
	// create stream in client
	stream, err := c.Greet(context.Background())
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	waitc := make(chan struct{})

	// send go routine
	go func() {
		arr := []string{"a", "b", "c", "D"}

		for _, name := range arr {
			req := &greetEveryOnePB.GreetEveryOneRequest{
				Name: name,
			}
			stream.Send(req)
			time.Sleep(2000 * time.Millisecond)
		}

		// close the stream
		stream.CloseSend()
	}()

	// receive go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
			}
			if err != nil {
				log.Fatalf("%v", err)
				close(waitc)
			}

			fmt.Printf("response: %v", res.GetRes())
		}
	}()

	<-waitc

}

func main() {
	// create conn
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}

	c := greetEveryOnePB.NewGreetEveryOneServiceClient(conn)

	Greet(c)
}
