package main

import (
	"context"
	pb "go-grpc/stream"
	"log"
	"time"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/balancer/grpclb"
)

const (
	ADDRESS = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		return
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	allStr, _ := c.AllStream(context.Background())

	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data)
		}
	}()

	go func() {
		for {
			allStr.Send(&pb.StreamReqData{Data: "from client"})
			time.Sleep(time.Second)
		}
	}()

	select {}
}
