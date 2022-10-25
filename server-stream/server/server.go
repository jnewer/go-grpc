package main

import (
	"fmt"
	pb "server-stream/pb"
	"google.golang.org/grpc"
	"net"
	"time"
)

const PORT = ":50051"

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) GetStream(req *pb.StreamReqData, res pb.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		res.Send(&pb.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(1 * time.Second)

		if i > 10 {
			break
		}
	}

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
