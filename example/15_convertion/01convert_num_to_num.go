package main

import (
	"fmt"
)

func main() {
	var i int = 256
	fmt.Printf("type:%T, val: %d\n", i, i)

	var f float64 = float64(i)
	fmt.Printf("type:%T, val: %f\n", f, f)

	var u uint8 = uint8(f)
	fmt.Printf("type:%T, val: %d\n", u, u)
}
