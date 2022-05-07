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
	data := []byte(`{
      		"name": "basic go",
      		"level": "normal",
      		"views": 7239,
      		"instructor": "อนุชิโตะ",
      		"full_price": 9999
   		}`)
	var c Course
	err := json.Unmarshal(data, &c)
	fmt.Printf("% #v\n", c)
	fmt.Println(err)
}
