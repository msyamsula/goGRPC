package main

import (
	"context"
	"fmt"
	"log"

	"github.com/msyamsula/goGRPC/calculator/pb"
	"google.golang.org/grpc"
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
}
