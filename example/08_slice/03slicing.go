package main

import "fmt"

func show(sk []string) {
	l := len(sk)
	fmt.Printf("length: %d -- show: %#v\n", l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}
	show(skills)

	show(skills[0:2])
	show(skills[1:3])

	l := len(skills)
	show(skills[0:l])
	show(skills[0:])
	show(skills[:l])
	show(skills[:])
}