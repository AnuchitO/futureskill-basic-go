package main

import "fmt"

func main() {
	var ss []int
	fmt.Printf("%#v\n", ss)
	if ss == nil {
		fmt.Println("nil")
	}

	ss = append(ss, 33)
	fmt.Printf("%#v\n", ss)
}
