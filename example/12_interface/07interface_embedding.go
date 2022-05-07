package main

import "fmt"

type presenter interface {
	promotion
	infoer
}

type promotion interface {
	discount() int
}

type infoer interface {
	info()
}

func summary(val presenter) {
	fmt.Println("discount price:", val.discount())
	val.info()
}

func sale(val promotion) {
	fmt.Println("sale:", val.discount())
}

func show(val infoer) {
	val.info()
}

func main() {
	v := course{}
	summary(v)
	sale(v)
	show(v)
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info:", c)
}
