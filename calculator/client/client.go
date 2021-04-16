package main

import (
	"context"
	"fmt"
	"log"

	"github.com/msyamsula/goGRPC/calculator/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func calculator(c pb.CalculatorServiceClient, a int64, b int64) {

	calculator := &pb.Calculator{
		A: a,
		B: b,
	}

	req := &pb.CalculatorRequest{
		Req: calculator,
	}

	res, err := c.Calculator(context.Background(), req)

	if err != nil {
		log.Printf("error: %v", err)
	} else {
		log.Printf("%v", res.GetRes())
	}

}

func sqrt(c pb.CalculatorServiceClient, num int64) {
	req := &pb.SquareRootRequest{
		Num: num,
	}

	resp, err := c.SquareRoot(context.Background(), req)

	// error handling in client
	if err != nil {
		respErr, ok := status.FromError(err)

		if ok {
			// actual error from grpc
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println(respErr.Message())
				return
			}
		} else {
			log.Fatalf("%v", respErr)
			return
		}
	}

	log.Printf("%v", resp.GetRoot())
}

func main() {
	fmt.Println("hello i'm client")

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		log.Printf("error: %v", err)
	}

	c := pb.NewCalculatorServiceClient(conn)
	calculator(c, 2, 3)
	calculator(c, 5, 10)
	calculator(c, 100, 3)
	sqrt(c, 12)
	sqrt(c, 64)
	sqrt(c, -1)
}
