package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type ArithmeticRequest struct {
	A int
	B int
}

type ArithmeticResponse struct {
	Pro int
	Quo int
	Rem int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalln("dialing error: ", err)
	}

	req := ArithmeticRequest{9, 2}
	var res ArithmeticResponse

	err = conn.Call("Arithmetic.Multiply", req, &res)
	if err != nil {
		log.Fatalln("arithmetic multiply error", err)
	}

	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	err = conn.Call("Arithmetic.Divide", req, &res)
	if err != nil {
		log.Fatalln("arithmetic divide error", err)
	}

	fmt.Printf("%d / %d ,quo is %d, rem is %d", req.A, req.B, res.Quo, res.Rem)
}
