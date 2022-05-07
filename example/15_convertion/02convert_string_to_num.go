package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "72"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	i := 10
	s := strconv.Itoa(i) // ตัวเลข ไปเป็นstring
	fmt.Printf("%T, %v\n", s, s)
}
