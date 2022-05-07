package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	t := time.Date(2022, 07, 30, 20, 40, 58, 0, time.UTC)
	format := t.Format("02/01/2006 15:04:05")
	fmt.Println(format) // Output: 30/07/2022 20:40:58

	format = "2006-01-02T15:04:05"
	tt, _ := time.Parse(format, "2022-07-30T20:40:58")
	fmt.Println(tt)

}
