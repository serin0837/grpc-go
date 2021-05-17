package main

import (
	"context"
	"fmt"
	"go-grpc/greet/greetpb"
	"io"
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
	c := greetpb.NewGreetServiceClient(cc)

	//fmt.Printf("created client: %f", c)

	doUnary(c)
	doServerStreaming(c)
}

//create function main first and create doUnary()
func doUnary(c greetpb.GreetServiceClient) {
	//4. Greet function
	fmt.Println("staring to do a Unray rpc")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephane",
			LastName:  "Maarek",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do a server streaming rpc...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephane",
			LastName:  "Maarek",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greetmanytimes: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("response from greet many times: %v", msg.GetResult())
	}

}
