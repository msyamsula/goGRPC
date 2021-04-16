package main

import (
	"io"
	"log"
	"net"

	"github.com/msyamsula/goGRPC/maximum/protobuf/newfindmaxpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) GetMax(stream newfindmaxpb.MaxService_GetMaxServer) error {

	var m int64
	m = -1
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return err
		}

		if err != nil {
			return err
		}

		cnum := msg.GetNum()

		log.Printf("get %v", cnum)

		if cnum > m {
			res := &newfindmaxpb.Response{
				Num: cnum,
			}

			stream.Send(res)
			m = cnum
		}
	}
}

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("%v", err)
	}

	// create server
	s := grpc.NewServer()
	newfindmaxpb.RegisterMaxServiceServer(s, &server{})
	// run server
	s.Serve(lis)
}
