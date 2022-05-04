package main

import "fmt"

func main() {
	var skills [3]string = [3]string{"js", "go", "python"}
	fmt.Printf("%#v\n", skills)

	s := skills[1]
	fmt.Printf("%#v\n", s)
}