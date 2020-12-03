package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-client/proto"
)

const (
	grpcAddr = "127.0.0.1:6655"
)

func main()  {
	c, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		panic(c)
	}

	rsp, err := proto.NewDemoClient(c).
		GetData(context.Background(), &proto.DemoReq{
			A: "123",
	})

	fmt.Println(rsp)
}
