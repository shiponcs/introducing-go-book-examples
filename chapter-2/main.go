package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var input float64
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		return
	}
	output := input * 2

	fmt.Println(output)
}
