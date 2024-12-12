package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	age  int
}

type ByName []Person

func (a ByName) Len() int {
	return len(a)
}
func (a ByName) Less(i, j int) bool {
	return a[i].age < a[j].age
}
func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	kids := ByName{
		{"m", 26},
		{"n", 22},
		{"sd", 32},
	}
	sort.Sort(kids)
	fmt.Println(kids)
}
