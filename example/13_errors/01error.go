package main

import (
	"errors"
	"fmt"
)

var myErr1 = errors.New("my custom error message")
var myErr2 = fmt.Errorf("user %q (id %d) not found", "AnuchitO", 123)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide '%.2f' by zero", a)
	}
	return a / b, nil
}

func main() {
	r, err := Divide(1, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
