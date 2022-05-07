package main

import "fmt"

func main() {
	skills := []string{"js", "go", "python"}
	fmt.Printf("length: %d -- val:%#v\n", len(skills), skills)

	skills = append(skills, "ruby", "java", "erlang")
	fmt.Printf("length: %d -- val:%#v\n", len(skills), skills)
}
