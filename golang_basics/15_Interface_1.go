package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	//cannot use (circle literal) (value of type circle) as Shape value in array or slice literal: circle does not implement Shape (missing method shape)
// 	//shapes :=  []Shape{circle{}}

// 	shapes := []Shape{circle{}, square{
// 		side: 2,
// 	}}

// 	for _, v := range shapes {
// 		fmt.Println(v.area())
// 	}

// 	//0
// 	//4
// }

// type Shape interface {
// 	//Method signature to be implemented.

// 	area() float64
// 	shape(name string) string

// 	//We cant put variables in interface as we cant initialize variables using constructor.
// 	//message string
// }

// type square struct {
// 	side float64
// }

// type circle struct {
// 	radius float64
// }

// func (s square) area() float64 {
// 	return s.side * s.side
// }

// func (s square) shape(name string) string {
// 	return "Square"
// }

// func (c circle) area() float64 {
// 	return c.radius * c.radius * math.Pi
// }

// func (s circle) shape(name string) string {
// 	return "Square"
// }
