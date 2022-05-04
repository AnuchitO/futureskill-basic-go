package main

import (
	"fmt"
	"math"
)

func main() {
	lim := 225.0

	if v := math.Pow(10, 2); v < lim {
		fmt.Println("x power n is:", v)
	} else {
		fmt.Printf("x power n is %g over limit %g.\n", v, lim)
	}
}