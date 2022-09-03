package main

import (
	pb "client-stream"
	"google.golang.org/grpc"
	"log"
	"net"
)

const PORT = ":50051"

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) PutStream(cliStr pb.Greeter_PutStreamServer) error {

	for {
		if tem, err := cliStr.Recv(); err == nil {
			log.Println(tem)
		} else {
			log.Println("break, err :", err)
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
