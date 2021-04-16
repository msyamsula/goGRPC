package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/msyamsula/goGRPC/maximum/protobuf/newfindmaxpb"
	"google.golang.org/grpc"
)

func GetMax(c newfindmaxpb.MaxServiceClient) {
	stream, err := c.GetMax(context.Background())
	if err != nil {
		log.Fatalf("%v", err)
	}

	waitc := make(chan struct{})
	// client send data
	go func() {
		arr := []int64{1, 2, 3, 4, 2, 1, 10, 1, 3, 5, 6, 100, 2, 3}
		for _, num := range arr {
			req := &newfindmaxpb.Request{
				Num: num,
			}

			stream.Send(req)
			time.Sleep(5 * time.Second)
		}

		stream.CloseSend()
	}()

	// client get data
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
			}
			if err != nil {
				log.Fatalf("%v", err)
				close(waitc)
			}

			mn := msg.GetNum()

			log.Printf("maximum: %v", mn)
		}
	}()

	<-waitc
}

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}

	c := newfindmaxpb.NewMaxServiceClient(conn)

	GetMax(c)
}
