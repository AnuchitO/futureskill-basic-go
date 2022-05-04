package main

import "fmt"

func main() {
	num := 1
	switch {
	case num < 0:
		fmt.Println("nagative number.")
	case num == 0:
		fmt.Println("zero.")
	case num > 0:
		fmt.Println("positive number.")
	default:
		fmt.Println("🤷🤷🤷🤷")
	}
}