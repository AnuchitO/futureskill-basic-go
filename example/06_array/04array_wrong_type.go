package main

import "fmt"

func show(sk [3]string) {
	fmt.Printf("show: %#v\n", sk)
}

func main() {
	var skills [3]string = [3]string{"js", "go", "python"}
	show(skills)

	var xs [4]string
	show(xs)
}