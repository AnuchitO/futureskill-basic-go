package main

import "fmt"

func add(x, y float64) (float64, string) {
	return x + y, "good"
}

func main() {
	v, s := add(42, 13)
	fmt.Println(v, s)
}
