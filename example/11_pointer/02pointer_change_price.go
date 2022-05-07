package main

import "fmt"

func changePrice(p *int) {
	fmt.Println("value in p is", p)
	*p = *p - 599
	fmt.Println("change", *p, &p)
}

func main() {
	var price int = 9999
	var addr *int = &price

	changePrice(&price)
	fmt.Println("after ", price, addr)
}
