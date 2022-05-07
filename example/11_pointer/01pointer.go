package main

import "fmt"

func main() {
	var price int = 9999
	var addr *int = &price
	fmt.Println(price, addr)

	*addr = 9400 // write
	fmt.Println(price, addr)
	v := *addr // read
	fmt.Println(v)
}
