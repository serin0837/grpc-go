package main

import (
	"fmt"
	"go-grpc/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //no ssl for now, open insecure connection
	if err != nil {
		log.Fatalf("could not connect:%v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("created client: %f", c)
}
