package main

import "fmt"

type course struct {
	name       string
	instructor string
	price      float64
}

func main() {
	c := course{
		name:       "Basic Go",
		instructor: "AnuchitO",
		price:      9999,
	}

	n := c.name
	c.instructor = "Nong"

	fmt.Println("name:", n)
	fmt.Println("instructor:", c.instructor)
	fmt.Println("price:", c.price)
}
