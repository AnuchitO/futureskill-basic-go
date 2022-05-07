package main

import "fmt"

func main() {
	today := "Saturday"

	switch today {
	case "Saturday":
		fmt.Println("today is Saturday")
	case "Monday":
		fmt.Println("today is weekdays")
	default:
		fmt.Printf("สวัสดีวัน %v อยู่ดีมีแฮงเดย์\n", today)
	}
}