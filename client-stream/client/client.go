package main

import (
	"context"
	pb "client-stream/pb"
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

	putRes, _ := c.PutStream(context.Background())

	i := 1
	for {
		i++
		putRes.Send(&pb.StreamReqData{Data: "put"})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
}
