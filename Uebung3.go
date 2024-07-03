package main

import (
	"fmt"
)

// Define the rectangle struct
type rectangle struct {
	length int
	width  int
}

// Define the square struct
type square struct {
	length int
}

// Implement the area method for rectangle
func (r rectangle) area() int {
	return r.length * r.width
}

// Implement the area method for square
func (s square) area() int {
	return s.length * s.length
}

// Define the shape interface
type shape interface {
	area() int
}

// Define the shape_Value struct for dictionary passing
type shape_Value struct {
	val  interface{}
	area func(interface{}) int
}

// Implement the area function for square in dictionary passing
func squareArea(val interface{}) int {
	return val.(square).area()
}

// Implement the area function for rectangle in dictionary passing
func rectangleArea(val interface{}) int {
	return val.(rectangle).area()
}

// Extend the sumArea_Dict function to handle type assertions
func sumArea_Dict(x, y shape_Value) int {
	switch v := y.val.(type) {
	case square:
		return x.area(x.val) + y.area(y.val) + v.length
	case rectangle:
		return x.area(x.val) + y.area(y.val) + v.length
	default:
		// Handle the error case when y is not a known shape
		panic("y is not a known shape")
	}
}

// Test function for dictionary variant
func testSumArea_Dict() {
	x := shape_Value{val: square{1}, area: squareArea}
	y := shape_Value{val: square{2}, area: squareArea}
	fmt.Printf("%d\n", sumArea_Dict(x, y))

	a := shape_Value{val: rectangle{1, 2}, area: rectangleArea}
	b := shape_Value{val: rectangle{3, 4}, area: rectangleArea}
	fmt.Printf("%d\n", sumArea_Dict(a, b))
}

// Implement the sumAreaVariant function to handle type assertions
func sumAreaVariant(x, y shape) int {
	if ySquare, ok := y.(square); ok {
		return x.area() + y.area() + ySquare.length
	} else if yRectangle, ok := y.(rectangle); ok {
		return x.area() + y.area() + yRectangle.length
	} else {
		// Handle the error case when y is not a known shape
		panic("y is not a known shape")
	}
}

// Test function for runtime variant
func testSumAreaVariant() {
	fmt.Printf("%d\n", sumAreaVariant(square{1}, square{2}))
	fmt.Printf("%d\n", sumAreaVariant(rectangle{1, 2}, rectangle{3, 4}))
}

func main() {
	testSumAreaVariant()
	testSumArea_Dict()
}
