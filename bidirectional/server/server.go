package main

import (
	"io"
	"log"
	"net"

	"github.com/msyamsula/goGRPC/bidirectional/protobuf/greetEveryOnePB"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(stream greetEveryOnePB.GreetEveryOneService_GreetServer) error {

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("%v", err)
			return err
		}

		log.Printf("get %v!! \n", msg.GetName())

		resp := &greetEveryOnePB.GreetEveryOneResponse{
			Res: "Hello " + msg.GetName() + "!!\n",
		}

		sendErr := stream.Send(resp)

		if sendErr != nil {
			log.Fatalf("%v", err)
			return err
		}
	}
}

func main() {
	// create listener
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("%v", err)
	}

	// create server
	s := grpc.NewServer()
	// register function
	greetEveryOnePB.RegisterGreetEveryOneServiceServer(s, &server{})

	// run server
	s.Serve(lis)
}
