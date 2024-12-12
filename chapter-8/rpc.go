package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go rpc.ServeConn(c)
	}
}

func client() {
	c, err := rpc.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", 9, &result)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(result)
	}
}
func main() {
	go server()
	time.Sleep(1 * time.Second)
	client()
}
