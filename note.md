doing again today 17.05.21

I moved folder from src to just home directory 
and did go get -u github.com/golang/protobuf/protoc-gen-go, go get -u google.golang.org/grpc


go mod init go-grpc
go mod tidy

looking good


 go run greet/greet_server/server.go



- grpc set up boilerplate server
1. create greet_server
2. make listener
3. make server 
4. bind port to the server
5. register service

go run greet/greet_server/server.go
server run waitin for connection
```go 
func main() {
	fmt.Println("hello world")
	//1. make listener//open tcp connection port for grpc 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	//2. register service with server
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```
- grpc set up boilerplate client
```go
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
```

- unary connection!
    - Greet API
    - our message is Greeting and contains first name and last name string field
    - it will take a greetrequest that containes a greeitng
    - it will return a greet response that contains a result string

- in greet.proto
before had only service GreetService{} that was it
now add
```proto
message Greeting{
    string first_name = 1;
    string last_name = 2;
}

message GreetRequest {
    Greeting greeting = 1;
}

message GreetResponse{
    string result = 1;
}

service GreetService{
    //unary input request result response
    rpc Greet(GreetRequest) returns (GreetResponse) {};
}
```

- implement server GreetServiceServer
```go
// 4. server unary inferface implementing
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet func invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello" + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}
```

- implement client 
done!

- create sum API- exercise
done it!!