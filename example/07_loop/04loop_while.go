package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 5; {
		sum += sum
		fmt.Println("sum:", sum)
	}
	fmt.Println("sum done:", sum)
}