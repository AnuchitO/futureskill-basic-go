package ticket

import "fmt"

func init() {
	fmt.Println("init: ticket")
}

func BuyTicket(movie string) {
	fmt.Printf("I bought tickets to %s\n", movie)
}
