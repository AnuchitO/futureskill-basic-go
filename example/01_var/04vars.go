package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	var age int = 55
	var price float64 = 22.52
	var check bool = true

	fmt.Println("msg:", msg)
	fmt.Println("age:", age)
	fmt.Println("price:", price)
	fmt.Println("check:", check)
}

// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true
