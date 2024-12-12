package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func server() {
	//listen on a port
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accept a connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleServerConnection(conn)
	}

}

func handleServerConnection(conn net.Conn) {
	// recv message
	var msg string
	err := gob.NewDecoder(conn).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	err = conn.Close()
	if err != nil {
		return
	}
}

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client error", err)
		return
	}
	msg := "Hello, Server"
	err = gob.NewEncoder(conn).Encode(&msg)
	if err != nil {
		fmt.Println(err)
	}
	err = conn.Close()
	if err != nil {
		return
	}
}

func main() {
	go server()
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		go client()
	}

	var input string
	fmt.Scanln(&input)
}
