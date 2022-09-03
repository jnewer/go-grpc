package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type ArithmeticRequest struct {
	A int
	B int
}

type ArithmeticResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

func main() {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := ArithmeticRequest{9, 2}
	var res ArithmeticResponse

	err = conn.Call("Arithmetic.Multiply", req, &res) // 乘法运算
	if err != nil {
		log.Fatalln("arithmetic multiply error: ", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	err = conn.Call("Arithmetic.Divide", req, &res)
	if err != nil {
		log.Fatalln("arithmetic divide error: ", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}
