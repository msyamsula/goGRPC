package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/msyamsula/goGRPC/calculator/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) Calculator(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	fmt.Println("calculator is invoke")
	a := req.GetReq().GetA()
	b := req.GetReq().GetB()

	sum := a + b

	res := &pb.CalculatorResponse{
		Res: sum,
	}

	return res, nil
}

func (*server) SquareRoot(ctx context.Context, req *pb.SquareRootRequest) (*pb.SquareRootResponse, error) {
	num := req.GetNum()
	ans := math.Sqrt(float64(num))

	// return error invalid argument
	if num < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Receive negative number: %v", num),
		)
	}

	res := &pb.SquareRootResponse{
		Root: float32(ans),
	}

	return res, nil
}

func main() {
	fmt.Println("Hello")

	// create listener
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Printf("error %v:", err)
	}

	// create server
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal("error")
	}
}
