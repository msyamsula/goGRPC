package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/msyamsula/goGRPC/streaming/pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(req *pb.GreetRequest, stream pb.GreetService_GreetServer) error {
	log.Print("stream greeting")
	name := req.GetName()

	for i := 0; i < 10; i++ {
		m := "Hello " + name + " number " + strconv.Itoa(i)
		res := &pb.GreetResponse{
			Res: m,
		}

		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func main() {
	fmt.Println("hello")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Print("error")
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})

	s.Serve(lis)
}
