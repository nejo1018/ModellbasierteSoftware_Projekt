package main

import (
	"fmt"
)


type Show interface {
	show() string
}


type ShowInt int

func (i ShowInt) show() string {
	return fmt.Sprintf("%d", i)
}


type ShowString string

func (s ShowString) show() string {
	return string(s)
}


type node[T any] struct {
	val  T
	next *node[T]
}


func showNode[T Show](n *node[T]) string {
	var s string
	for n != nil {
		s = s + n.val.show() + " -> "
		n = n.next
	}
	s = s + "nil"
	return s
}


type ShowFunc func(interface{}) string


type node_Value struct {
	val  interface{}
	next *node_Value
	show ShowFunc
}

n
func showNode_Dict(n *node_Value) string {
	var s string
	for n != nil {
		s = s + n.show(n.val) + " -> "
		n = n.next
	}
	s = s + "nil"
	return s
}

// Test function for dictionary variant
func testShowNode_Dict() {
	n1 := &node_Value{val: ShowInt(1), show: func(val interface{}) string { return val.(ShowInt).show() }}
	n2 := &node_Value{val: ShowInt(2), show: func(val interface{}) string { return val.(ShowInt).show() }}
	n3 := &node_Value{val: ShowInt(3), show: func(val interface{}) string { return val.(ShowInt).show() }}
	n1.next = n2
	n2.next = n3

	fmt.Println(showNode_Dict(n1))

	n4 := &node_Value{val: ShowString("a"), show: func(val interface{}) string { return val.(ShowString).show() }}
	n5 := &node_Value{val: ShowString("b"), show: func(val interface{}) string { return val.(ShowString).show() }}
	n6 := &node_Value{val: ShowString("c"), show: func(val interface{}) string { return val.(ShowString).show() }}
	n4.next = n5
	n5.next = n6

	fmt.Println(showNode_Dict(n4))
}

// Test function for runtime variant
func testShowNode() {
	n1 := &node[ShowInt]{val: ShowInt(1)}
	n2 := &node[ShowInt]{val: ShowInt(2)}
	n3 := &node[ShowInt]{val: ShowInt(3)}
	n1.next = n2
	n2.next = n3

	fmt.Println(showNode(n1))

	n4 := &node[ShowString]{val: ShowString("a")}
	n5 := &node[ShowString]{val: ShowString("b")}
	n6 := &node[ShowString]{val: ShowString("c")}
	n4.next = n5
	n5.next = n6

	fmt.Println(showNode(n4))
}

func main() {
	testShowNode()
	testShowNode_Dict()
}
