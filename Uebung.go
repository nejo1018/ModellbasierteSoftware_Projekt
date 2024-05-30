package main

import (
	"fmt"
	"time"
)

type rectangle struct {
	length int
	width  int
}

type square struct {
	length int
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (s square) area() int {
	return s.length * s.length
}

func (r *rectangle) scale(x int) {
	r.length = r.length * x
	r.width = r.width * x
}

func (s *square) scale(x int) {
	s.length = s.length * x
}

type shape interface {
	area() int
}

type shapeExt interface {
	shape
	scale(int)
}

func area_Rec(r rectangle) int {
	return r.length * r.width
}

func area_Sq(s square) int {
	return s.length * s.length
}

func area_Lookup(x interface{}) int {
	var y int

	switch v := x.(type) {
	case square:
		y = area_Sq(v)
	case rectangle:
		y = area_Rec(v)
	}
	return y

}

func sumArea_Lookup(x, y interface{}) int {
	return area_Lookup(x) + area_Lookup(y)
}

type shape_Value struct {
	val  interface{}
	area func(interface{}) int
}

func sumArea_Dict(x, y shape_Value) int {
	return x.area(x.val) + y.area(y.val)
}

func main() {

	startTime := time.Now()

	rect_Lookup := rectangle{length: 4, width: 5}
	sq_Lookup := square{length: 3}

	for i := 0; i < 10000; i++ {
		sumArea_Lookup(rect_Lookup, sq_Lookup)
	}

	rect_Dict := shape_Value{
		val:  rectangle{length: 4, width: 5},
		area: func(x interface{}) int { return area_Rec(x.(rectangle)) },
	}

	sq_Dict := shape_Value{
		val:  square{length: 3},
		area: func(x interface{}) int { return area_Sq(x.(square)) },
	}

	for i := 0; i < 1000000; i++ {
		sumArea_Dict(rect_Dict, sq_Dict)
	}

	elapsedTime := time.Since(startTime)

	fmt.Println("Dauer der Funktion: ", elapsedTime)

}
