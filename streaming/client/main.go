package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"

	pb "github.com/gavinzhou/hello-grpc/streaming/pb"
	"google.golang.org/grpc"
)

var (
	address = flag.String("addr", "localhost:8972", "address")
	name    = flag.String("n", "world", "name")
	example = flag.Int("e", 0, "0 for replay streaming, 1 for request streaming, 2 for bi-streaming")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	switch *example {
	// case 0:
	// 	log.Println("Test Reply streaming")
	// 	sayHello1(c)
	case 0:
		log.Println("\n\n\nTest Request streaming")
		sayHello2(c)
		// case 2:
		// 	log.Println("\n\n\nTest Bidirectional streaming")
		// 	sayHello3(c)
	}

}

// func sayHello1(c pb.GreeterClient) {
// 	stream, err := c.SayHello1(context.Background(), &pb.HelloRequest{Name: *name})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}

// 	for {
// 		reply, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Printf("failed to recv: %v", err)
// 		}

// 		log.Printf("Greeting: %s", reply.Message)
// 	}
// }

func sayHello2(c pb.GreeterClient) {
	var err error
	stream, err := c.SayHello2(context.Background())

	for i := 0; i < 100; i++ {
		if err != nil {
			log.Printf("failed to call: %v", err)
			break
		}
		stream.Send(&pb.HelloRequest{Name: *name + strconv.Itoa(i)})
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("failed to recv: %v", err)
	}

	log.Printf("Greeting: %s", reply.Message)
}
