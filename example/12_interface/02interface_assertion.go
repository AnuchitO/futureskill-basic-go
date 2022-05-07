package main

import "fmt"

func show(val any) {
	i, ok := val.(int)
	if ok {
		i = i + 10
		fmt.Println(i)
	} else {
		fmt.Println("not int")
	}
}

func main() {
	var v any
	v = "gopher"
	show(v)
}
