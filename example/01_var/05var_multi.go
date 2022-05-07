package main

import "fmt"

func main() {
	msg, age, price, check := "Hello Gopher!!!", 55, 22.52, true

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
