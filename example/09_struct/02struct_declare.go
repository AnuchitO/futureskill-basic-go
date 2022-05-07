package main

import "fmt"

type course struct {
	name       string
	instructor string
	price      float64
}

func main() {
	var c1 = course{name: "Basic Go", instructor: "AnuchitO", price: 9999}
	var c2 = course{"Basic Go", "AnuchitO", 9999}
	var c3 = course{instructor: "AnuchitO"}
	var c4 = course{}

	fmt.Printf("c1: %#v\n", c1)
	fmt.Printf("c2: %#v\n", c2)
	fmt.Printf("c3: %#v\n", c3)
	fmt.Printf("c4: %#v\n", c4) // zero value
}
