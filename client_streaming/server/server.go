package main

import (
	"io"
	"log"
	"net"

	"github.com/msyamsula/goGRPC/client_streaming/clientstreampb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(stream clientstreampb.GreetService_GreetServer) error {

	result := "Hello \n"
	// process stream from client
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			res := &clientstreampb.ClientStreamResponse{
				Res: result,
			}
			return stream.SendAndClose(res)
		}

		if err != nil {
			log.Fatalf("error when reading")
		}

		name := msg.GetName()
		result += name + "! \n"
		log.Printf("%v", result)
	}
}

func main() {

	// create listener
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Print("error")
	}

	// create server
	s := grpc.NewServer()

	// register function
	clientstreampb.RegisterGreetServiceServer(s, &server{})

	s.Serve(lis)
}
