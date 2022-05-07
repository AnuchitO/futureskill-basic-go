
package main

import "fmt"

func main() {
	skills := [3]string{"js", "go", "python"}
	for i := range skills {
		val := skills[i]
		fmt.Println("index:", i, "value:", val)
	}
}