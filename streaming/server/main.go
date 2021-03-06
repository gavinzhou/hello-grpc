package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/gavinzhou/hello-grpc/streaming/pb"
)

var (
	port = flag.String("p", ":8972", "port")
)

type server struct{}

// func (s *server) SayHello1(in *pb.HelloRequest, gs pb.Greeter_SayHello1Server) error {
// 	name := in.Name

// 	for i := 0; i < 100; i++ {
// 		gs.Send(&pb.HelloReply{Message: "Hello " + name + strconv.Itoa(i)})
// 	}
// 	return nil
// }

func (s *server) SayHello2(gs pb.Greeter_SayHello2Server) error {
	var names []string

	for {
		in, err := gs.Recv()
		if err == io.EOF {
			gs.SendAndClose(&pb.HelloReply{Message: "Hello " + strings.Join(names, ",")})
			return nil
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
			return err
		}
		names = append(names, in.Name)
	}

	return nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
