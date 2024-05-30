package main

// Example

type rectangle struct {
	length int
	width  int
}

type square struct {
	length int
}

type shape_Value struct {
	val  interface{}
	area func(interface{}) int
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (s square) area() int {
	return s.length * s.length
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

func sumArea_Dict(x, y shape_Value) int {
	return x.area(x.val) + y.area(y.val)
}

func main() {

}
