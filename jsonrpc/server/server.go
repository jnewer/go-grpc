package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Arithmetic struct {
}

type ArithmeticRequest struct {
	A int
	B int
}

type ArithmeticResponse struct {
	Pro int
	Quo int
	Rem int
}

func (a *Arithmetic) Multiply(req ArithmeticRequest, res *ArithmeticResponse) error {
	res.Pro = req.A * req.B
	return nil
}

func (a *Arithmetic) Divide(req ArithmeticRequest, res *ArithmeticResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}
func main() {
	rpc.Register(new(Arithmetic))

	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}

	fmt.Fprintf(os.Stdout, "%s", "start connection\n")

	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}

		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s", "new client in coming\n")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
