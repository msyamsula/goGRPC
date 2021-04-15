package main

import (
	"log"
	"net"

	"github.com/msyamsula/goGRPC/prime/primepb"

	"google.golang.org/grpc"
)

type server struct{}

// decompose implementation
func (*server) Decompose(req *primepb.PrimeRequest, stream primepb.PrimeService_DecomposeServer) error {
	num := req.GetNum()
	var p int64
	p = 2
	for p*p <= num {
		cp := 0
		for num%p == 0 {
			num /= p
			cp++
			res := &primepb.PrimeResponse{
				Prime: p,
			}

			stream.Send(res)
		}

		p++
	}

	if num > 1 {
		res := &primepb.PrimeResponse{
			Prime: num,
		}
		stream.Send(res)
	}

	return nil
}

func main() {

	// create listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Print("error")
	}

	// create srever
	s := grpc.NewServer()

	// register service
	primepb.RegisterPrimeServiceServer(s, &server{})

	// run server
	s.Serve(lis)
}
