package main

import (
	"context"
	pb "github.com/memochou1993/grpc-go-example"
	"google.golang.org/grpc"
	"log"
	"net"
)

type service struct {
	pb.UnimplementedHelloServiceServer
}

func (s *service) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Request received: %s", r.GetGreeting())
	return &pb.HelloResponse{Reply: "Hello, " + r.GetGreeting()}, nil
}

func main() {
	addr := "127.0.0.1:8080"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(gRPCServer, &service{})
	if err := gRPCServer.Serve(ln); err != nil {
		log.Fatalln(err.Error())
	}
}
