package main

import "fmt"

func main() {
	var v interface{}
	v = 36
	fmt.Printf("%T %#v\n", v, v)
	v = "hello"
	fmt.Printf("%T %#v\n", v, v)
}
