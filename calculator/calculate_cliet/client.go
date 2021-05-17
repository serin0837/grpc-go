package main

import (
	"context"
	"fmt"
	"go-grpc/calculator/calculatepb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am a client")
	//1.create connection to server
	//dial function takes two args addreess, options
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //no ssl for now(bydeafult ssl inside), open insecure connection for now only!
	if err != nil {
		log.Fatalf("could not connect:%v", err)
	}
	//3.close connection after executing finish
	defer cc.Close()
	//2.create client with connection
	c := calculatepb.NewCalculateServiceClient(cc)

	//fmt.Printf("created client: %f", c)

	doUnary(c)
}

//create function main first and create doUnary()
func doUnary(c calculatepb.CalculateServiceClient) {
	//4. Greet function
	fmt.Println("staring to do a Unray rpc")
	req := &calculatepb.CalculateRequest{
		Calculating: &calculatepb.Calculating{
			FirstNumber:  10,
			SecondNumber: 25,
		},
	}
	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling calculate RPC: %v", err)
	}
	log.Printf("Response from Calculate: %v", res.Result)
}
