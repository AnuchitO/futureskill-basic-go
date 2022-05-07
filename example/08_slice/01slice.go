package main

import "fmt"

func main() {
	skills := []string{"js", "go", "python"}
	fmt.Printf("%#v\n", skills)

	skills[1] = "golang"
	fmt.Printf("%#v\n", skills)

	s := skills[0]
	fmt.Printf("%#v\n", s)

	for i := 0; i < len(skills); i++ {
		fmt.Printf("for => %#v\n", skills[i])
	}

	for _, val := range skills {
		fmt.Printf("for-range value: %#v\n", val)
	}
}
