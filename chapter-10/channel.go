package main

import (
	"fmt"
	"time"
)

func ping(c chan<- string) {
	for {
		c <- "ping"
	}
}

func pong(c chan<- string) {
	for {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		fmt.Println(<-c)
		time.Sleep(time.Second)
	}
}

func main() {
	var c chan string = make(chan string)
	go ping(c)
	go printer(c)
	go pong(c)

	var input string
	fmt.Scanln(&input)
}
