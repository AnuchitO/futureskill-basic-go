package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name       string `json:"name"`
	Level      string `json:"level"`
	Views      int    `json:"views"`
	Instructor string `json:"instructor"`
	FullPrice  int    `json:"full_price"`
}

func main() {
	c := Course{
		Name:       "basic go",
		Level:      "normal",
		Views:      7239,
		Instructor: "อนุชิโตะ",
		FullPrice:  9999,
	}

	b, err := json.Marshal(c)
	fmt.Printf("type : %T \n", b)
	fmt.Printf("byte : %v  \n", b)
	fmt.Printf("string: %s \n", b)
	fmt.Println(err)
}
