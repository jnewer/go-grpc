package main

import (
	pb "stream/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

const PORT = ":50051"

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) AllStream(allStr pb.Greeter_AllStreamServer) error {

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data)
		}
	}()

	go func() {
		for {
			allStr.Send(&pb.StreamResData{Data: "from server"})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		return
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})

	s.Serve(lis)
}
