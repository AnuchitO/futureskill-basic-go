package main

import "fmt"

func main() {
	skills := [3]string{"js", "go", "python"}
	for _, val := range skills {
		fmt.Println("value:", val)
	}
}