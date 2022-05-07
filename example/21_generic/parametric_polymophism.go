package main

import (
	"fmt"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

type Number interface {
	int | float64
}

func min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	var a, b = 1, 2
	fmt.Println("minInt:", minInt(a, b))

	c, d := 1.1, 2.2
	fmt.Println("minFloat64:", minFloat64(c, d))

	fmt.Println("min:", min(a, b))
	fmt.Println("min:", min(c, d))
}
