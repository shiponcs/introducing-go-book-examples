package main

import "fmt"

func f(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
}

func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)
}
