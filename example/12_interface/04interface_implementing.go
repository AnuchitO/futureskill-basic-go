package main

import "fmt"

type promotion interface {
	discount() int
}

func sale(val promotion) {
	fmt.Println("sale:", val.discount())
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info:", c)
}

func main() {
	v := course{}
	sale(v)
	v.info()
}
