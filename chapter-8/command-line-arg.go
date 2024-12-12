package main

import (
	"flag"
	"fmt"
)

func main() {
	maxp := flag.Int("maxp", 5, "maximum number of ports to open")
	flag.Parse()
	fmt.Printf("maxp: %d\n", *maxp)
	x := flag.Args()
	fmt.Println(x)
}
