package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	r float64
}

type rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func (c *circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *rectangle) area() float64 {
	return distance(r.x1, r.y1, r.x1, r.y2) * distance(r.x1, r.y1, r.x2, r.y1)
}

func (r *rectangle) perimeter() float64 {
	return 2*distance(r.x1, r.y1, r.x1, r.y2) + 2*distance(r.x1, r.y1, r.x2, r.y1)
}

func totalArea(s ...Shape) float64 {
	var area float64
	for _, s := range s {
		area += s.area()
	}
	return area
}

func totalPerimeter(s ...Shape) float64 {
	var perimeter float64
	for _, s := range s {
		perimeter += s.perimeter()
	}
	return perimeter
}

type MutiShape struct {
	shapes []Shape
}

func (m *MutiShape) area() float64 {
	var area float64
	for _, shape := range m.shapes {
		area += shape.area()
	}
	return area
}

func (m *MutiShape) perimeter() float64 {
	var perimeter float64
	for _, shape := range m.shapes {
		perimeter += shape.perimeter()
	}
	return perimeter
}

func main() {
	c := circle{r: 5}
	r := rectangle{1, 2, 3, 4}
	fmt.Println(totalArea(&c, &r))
	fmt.Println(totalPerimeter(&c, &r))
	multiShape := MutiShape{[]Shape{&c, &r}}
	fmt.Println(totalArea(&multiShape))
}
