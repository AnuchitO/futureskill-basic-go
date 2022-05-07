package main

import "fmt"

func main() {


	switch today := "Monday"; today {
	case "Saturday":
		fmt.Println("today is Saturday")
	case "Monday", "Tuesday":
		fmt.Println("today is weekdays")
	default:
		fmt.Printf("สวัสดีวัน %#v อยู่ดีมีแฮงเดย์\n", today)
	}
}