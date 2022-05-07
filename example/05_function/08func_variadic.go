package main

import "fmt"

func skills(xs ...string) {
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])

	for _, s := range xs {
		fmt.Println("skill:", s)
	}
}

func main() {
	skills("js", "go", "python")
}
