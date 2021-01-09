package main

import (
	"context"
	"fmt"
	pb "github.com/memochou1993/grpc-go-example"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addr := "127.0.0.1:9999"
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	c := pb.NewHelloServiceClient(conn)
	r, err := c.SayHello(ctx, &pb.HelloRequest{Greeting: "World!"})
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(fmt.Sprintf("Response received: %s", r.GetReply()))
}
