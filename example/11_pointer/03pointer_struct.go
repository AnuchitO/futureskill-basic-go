package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func (c *course) discount() int {
	c.price = c.price - 599
	fmt.Println("discount:", c.price)
	return c.price
}

func main() {
	c := &course{"Basic Go", "AnuchitO", 9999}

	d := c.discount() // เข้าถึงค่าในตัวแปรเหมือนเดิมไม่ต้อง de-reference อีกที
	fmt.Println("discount price:", d)
	fmt.Println("price:", c.price)
}
