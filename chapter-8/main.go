package main

import (
	"fmt"
	"go-book/chapter-8/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Printf("Average of xs: %v\n", avg)
}
