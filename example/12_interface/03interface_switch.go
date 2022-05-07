package main

import "fmt"

func show(val any) {
	switch v := val.(type) {
	case int:
		i := v + 10
		fmt.Println("int:", i)
	case string:
		s := v + ", Hi!"
		fmt.Println("string:", s)
	default:
		fmt.Println("not handle type.")
	}
}

func main() {
	var v any
	v = "gopher"
	show(v)
}
