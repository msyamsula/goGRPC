package main

import (
	"io"
	"log"
	"net"

	"github.com/msyamsula/goGRPC/average/pb/averagepb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Average(stream averagepb.Service_AverageServer) error {

	var sum int64
	var count int64
	var average int64
	sum = 0
	count = 0
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			average = sum / count
			res := &averagepb.Response{
				Num: average,
			}

			log.Printf("sum: %v", sum)
			log.Printf("average: %v", average)
			return stream.SendAndClose(res)
		}

		if err != nil {
			log.Fatalf("error")
		}

		sum += msg.GetNum()
		count++
		log.Printf("%v, sum %v, average %v", msg.GetNum(), sum, sum/count)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Printf("%v", err)
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	checkErr(err)
	s := grpc.NewServer()
	averagepb.RegisterServiceServer(s, &server{})
	s.Serve(lis)
}
