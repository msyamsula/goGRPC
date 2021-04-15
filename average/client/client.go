package main

import (
	"context"
	"log"

	"github.com/msyamsula/goGRPC/average/pb/averagepb"
	"google.golang.org/grpc"
)

func Average(arr []int64, c averagepb.ServiceClient) {

	// create stream in client
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Printf("%v", err)
	}

	// stream the data
	for _, num := range arr {
		req := &averagepb.Request{
			Num: num,
		}
		stream.Send(req)
	}

	// close the stream
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("%v", err)
	}

	log.Printf("%v", resp.GetNum())
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("%v", err)
	}

	c := averagepb.NewServiceClient(conn)

	arr := []int64{1, 2, 3, 4, 5, 4, 3, 10, 2, 9}

	Average(arr, c)
}
