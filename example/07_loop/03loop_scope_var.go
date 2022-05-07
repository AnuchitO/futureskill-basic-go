package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		sum := 0
		sum += i
		fmt.Println("i:", i, "sum:", sum)
	}
}