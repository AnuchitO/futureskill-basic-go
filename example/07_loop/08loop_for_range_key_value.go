
package main

import "fmt"

func main() {
	skills := [3]string{"js", "go", "python"}
	for i, val := range skills {
		fmt.Println("index:", i, "value:", val)
	}
}