package main

import (
	"context"
	"io"
	"log"

	"github.com/msyamsula/goGRPC/prime/primepb"

	"google.golang.org/grpc"
)

func Decompose(c primepb.PrimeServiceClient, num int64) {
	// make request
	req := &primepb.PrimeRequest{
		Num: num,
	}
	// run remote call
	res, err := c.Decompose(context.Background(), req)

	if err != nil {
		log.Print("error")
	}

	// get streaming response
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

	// create connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Print("error")
	}

	defer conn.Close()

	// create client in that connection
	c := primepb.NewPrimeServiceClient(conn)

	Decompose(c, 123)
	Decompose(c, 254)
	Decompose(c, 12340)
}
