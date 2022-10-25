package main

import (
	"context"
	pb "server-stream/pb"
	"google.golang.org/grpc"
	"log"
)

const ADDRESS = "localhost:50051"

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		return
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	reqStreamData := &pb.StreamReqData{Data: "golang"}
	res, _ := c.GetStream(context.Background(), reqStreamData)
	for {
		r, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(r)
	}
}
