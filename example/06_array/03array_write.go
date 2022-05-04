package main

import "fmt"

func main() {
	var skills [3]string = [3]string{"js", "go", "python"}
	fmt.Printf("%#v\n", skills)

	skills[1] = "golang"

	s := skills[1] // change index back to 1
	fmt.Printf("%#v\n", s)

	l := len(skills)
	fmt.Printf("length: %#v\n", l)
}