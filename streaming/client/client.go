package main

import (
	"context"
	"io"
	"log"

	"github.com/msyamsula/goGRPC/streaming/pb"
	"google.golang.org/grpc"
)

func greet(c pb.GreetServiceClient, name string) {
	// make request
	req := &pb.GreetRequest{
		Name: name,
	}

	// run the procedure call
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Print("error")
	}

	// streaming processing
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error")
		}
		log.Printf("%v", msg)
	}

}

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Print("error")
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	greet(c, "syamsul")
}
