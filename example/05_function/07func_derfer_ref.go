package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}

	fmt.Println("done")
}
