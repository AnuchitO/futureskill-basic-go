package main

import "fmt"

func main() {
	var msg, age, price, check = "Hello Gopher!!!", 55, 22.52, true

	var r rune = '😷' // 'A'  😷

	fmt.Printf("msg: %#v\n", msg)
	fmt.Printf("age: %#v\n", age)
	fmt.Printf("price: %#v\n", price)
	fmt.Printf("check: %#v\n", check)
	fmt.Printf("r: %c\n", r)
}

// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true
// r: 😷
