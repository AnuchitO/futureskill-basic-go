-- create learn-basic-go folder --
-- create learn-basic-go/hello.go --

-- hello.go --
// go run hello.go
// Output:
/*
hello.go:1:1: expected 'package', found 'EOF'
*/

-- hello.go --
package main

// go run hello.go
// Output:
/*
# command-line-arguments
runtime.main_main·f: function main is undeclared in the main package
*/

-- hello.go --
package main



func main() {
	
}
// go run hello.go
// Output:

-- hello.go --
package main



func main() {
	fmt.Println("Hello Gopher!!!")
}
// go run hello.go
// Output:
/*
# command-line-arguments
./hello.go:4:2: undefined: fmt
*/

-- hello.go --
package main

import "fmt"

func main() {
	fmt.Println("Hello Gopher!!!")
}
// go run hello.go
// Output:
/* Hello Gopher!!! */

/* จุดเริ่มต้นรันโปรแกรมของ Go จะเริ่มจาก ฟังก์ชัน main ใน package main เท่านั้น */
/* เป็นกติกาที่ Go กำหนดไว้ (convention) */
-- hello.go --
package hello

import "fmt"

func main() {
	fmt.Println("Hello Gopher!!!")
}
// go run hello.go
// Output:
/* package command-line-arguments is not a main package */

-- hello.go --
package main

import "fmt"

func main() {
	fmt.Println("Hello Gopher!!!")
}
// go run hello.go
// Output:
/* Hello Gopher!!! */


/* จุดเริ่มต้นรันโปรแกรมของ Go จะเริ่มจาก ฟังก์ชัน main ใน package main เท่านั้น */
-- rename: hello.go to xxx.go --
// go run xxx.go
// Output:
/* Hello Gopher!!! */

-- rename: xxx.go to main.go --
// go run main.go
// Output:
/* Hello Gopher!!! */

// ; is optional ไม่บังคับ best practice ไม่ต้องใส่
// อะไรที่ไม่จำเป็น Go เอาออกหมด
-- main.go --
package main

import "fmt"

func main() {
	fmt.Println("Hello Gopher!!!");
}
// go run main.go
// Output:
/* Hello Gopher!!! */

-- main.go --
package main

import "fmt"

func main() 
{
	fmt.Println("Hello Gopher!!!");
}
// go run main.go
// Output:
// # command-line-arguments
// ./main.go:6:1: syntax error: unexpected semicolon or newline before {

-- main.go --
package main

import "fmt"

func main() {
	var msg string = "Hello Gopher!!!"
	fmt.Println(msg)
}
// go run main.go
// Output:
/* Hello Gopher!!! */

// ; is optional just show.
-- main.go --
package main

import "fmt"

func main() {
	var msg string = "Hello Gopher!!!";
	fmt.Println(msg);
}
// go run main.go
// Output:
/* Hello Gopher!!! */

-- main.go --
package main

import "fmt"

func main() {
	var msg = "Hello Gopher!!!"
	fmt.Println(msg)
}
// go run main.go
// Output:
/* Hello Gopher!!! */

-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
}
// go run main.go
// Output:
/* Hello Gopher!!! */

-- main.go --
package main

import "fmt"

msg := "Hello Gopher!!!"

func main() {
	fmt.Println(msg)
}
// go run main.go
// Output:
/* 
# command-line-arguments
./main.go:5:1: syntax error: non-declaration statement outside function body 
*/

-- main.go --
package main

import "fmt"

var msg = "Hello Gopher!!!"

func main() {
	fmt.Println(msg)
}
// go run main.go
// Output:
/* Hello Gopher!!! */

-- main.go --
package main

import "fmt"

var msg string = "Hello Gopher!!!"

func main() {
	fmt.Println(msg)
}
// go run main.go
// Output:
/* Hello Gopher!!! */

// ทำไมฟังก์ชันปริ้นถึงมองเห็นตัวแปร msg 
-- main.go --
package main

import "fmt"

var msg string = "Hello Gopher!!!" // ประกาศในระดับ package เป็น global variable

func main() {
	fmt.Println(msg)
}
// go run main.go
// Output:
/* Hello Gopher!!! */

-- main.go --
package main

import "fmt"

func main() {
	{
		var msg string = "Hello Gopher!!!"
		fmt.Println(msg)
	}
	fmt.Println(msg)
}
// go run main.go
// Output:
/* 
# command-line-arguments
./main.go:10:14: undefined: msg
 */


// go is statically type language
-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
	msg = 55
	fmt.Println(msg)
}
// go run main.go
// Output:
/*
# command-line-arguments
./main.go:8:9: cannot use 55 (untyped int constant) as string value in assignment
*/

-- main.go --
package main

import "fmt"

func main() {
	var msg string = "Hello Gopher!!!"
	fmt.Println(msg)
	msg = "AnuchitO"
	fmt.Println(msg)
}
// go run main.go
// Output:
/*
Hello Gopher!!!
AnuchitO
*/

-- main.go --
package main

import "fmt"

func main() {
	var msg = "Hello Gopher!!!"
	fmt.Println(msg)
	msg = "AnuchitO"
	fmt.Println(msg)
}
// go run main.go
// Output:
/*
Hello Gopher!!!
AnuchitO
*/

// string concat
-- main.go --
package main

import "fmt"

func main() {
	var msg = "Hello Gopher!!!"
	fmt.Println(msg)
	msg = "AnuchitO" + " Nong"
	fmt.Println(msg)
}

// go run main.go
// Output:
// Hello Gopher!!!
// AnuchitO Nong


// เน้นเรื่อง := กับ = 
-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
	msg = "AnuchitO" + " Nong"
	fmt.Println(msg)
}

// go run main.go
// Output:
// Hello Gopher!!!
// AnuchitO Nong

-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
	age := 55
}
// go run main.go
// Output:
// #  command-line-arguments
// ./main.go:8:2: age declared but not used

// comment
-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
	// age := 55
}
// go run main.go
// Output:
// Hello Gopher!!!

// block comment
-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
	/*
		comment หลายบรรทัด
	*/
	// age := 55
}
// go run main.go
// Output:
// Hello Gopher!!!

-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
	age := 55
	fmt.Println(age)
}
// go run main.go
// Output:
// Hello Gopher!!!
// 55

// not only variable import also need to use
-- main.go --
package main

import "fmt"
import "math"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
	age := 55
	fmt.Println(age)
}
// go run main.go
// Output:
// #  command-line-arguments
// ./main.go:4:8: imported and not used: "math"

// comment can be anywhere
-- main.go --
package main

import "fmt"
// import "math"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
	age := 55
	fmt.Println(age)
}
// go run main.go
// Output:
// #  command-line-arguments
// ./main.go:4:8: imported and not used: "math"

-- main.go --
package main

import "fmt"

var price = 399.99

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
	age := 55
	fmt.Println(age)
}
// go run main.go
// Output:
// Hello Gopher!!!
// 55

-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
	age := 55
	fmt.Println(age)
}
// go run main.go
// Output:
// Hello Gopher!!!
// 55

-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
	var age = 55
	fmt.Println(age)
}
// go run main.go
// Output:
// Hello Gopher!!!
// 55

-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println(msg)
	var age int = 55
	fmt.Println(age)
}
// go run main.go
// Output:
// Hello Gopher!!!
// 55


// pass many argument to Println
// open doc: https://pkg.go.dev/fmt#Println
// อธิบาย parameters vs arguments
-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println("msg:", msg)
	var age int = 55
	fmt.Println("age:", age)
}
// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55

-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	fmt.Println("msg:", msg)
	var age int = 55
	fmt.Println("age:", age)
	var price float64 = 22.52
	fmt.Println("price:", price)
	var check bool = true
	fmt.Println("check:", check)
}
// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true

-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	var age int = 55
	var price float64 = 22.52
	var check bool = true

	fmt.Println("msg:", msg)
	fmt.Println("age:", age)
	fmt.Println("price:", price)
	fmt.Println("check:", check)
}
// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true

-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	var age = 55
	var price = 22.52
	var check = true

	fmt.Println("msg:", msg)
	fmt.Println("age:", age)
	fmt.Println("price:", price)
	fmt.Println("check:", check)
}
// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true

-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	age := 55
	price := 22.52
	check := true

	fmt.Println("msg:", msg)
	fmt.Println("age:", age)
	fmt.Println("price:", price)
	fmt.Println("check:", check)
}
// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true

// multiple_declaration.go 
-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	age := 55
	price, check := 22.52, true

	fmt.Println("msg:", msg)
	fmt.Println("age:", age)
	fmt.Println("price:", price)
	fmt.Println("check:", check)
}
// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true

-- main.go --
package main

import "fmt"

func main() {
	msg := "Hello Gopher!!!"
	age, price, check :=55, 22.52, true

	fmt.Println("msg:", msg)
	fmt.Println("age:", age)
	fmt.Println("price:", price)
	fmt.Println("check:", check)
}
// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true

-- main.go --
package main

import "fmt"

func main() {
	msg, age, price, check := "Hello Gopher!!!", 55, 22.52, true

	fmt.Println("msg:", msg)
	fmt.Println("age:", age)
	fmt.Println("price:", price)
	fmt.Println("check:", check)
}
// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true

-- main.go --
package main

import "fmt"

func main() {
	var msg, age, price, check = "Hello Gopher!!!", 55, 22.52, true

	fmt.Println("msg:", msg)
	fmt.Println("age:", age)
	fmt.Println("price:", price)
	fmt.Println("check:", check)
}
// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true

// # workshop: Println
// problem: https://go.dev/play/p/WPtOfb_YsS3
// solution 1 literal: https://go.dev/play/p/jhfzQ9jjsin
// solution 2 var: https://go.dev/play/p/PQHt1WbwyL0
// solution 3 raw: https://go.dev/play/p/E6prOIF414E
// explain raw string raw string ``
title := "Avengers: Endgame"
year := 2019
ratings := 8.4
genre := "Sci-Fi"
isSuperHero := true

// เรื่อง: Avengers: Endgame
// ปี: 2019
// เรตติ้ง: 8.4
// ประเภท: Sci-Fi
// ซุปเปอร์ฮีโร่: true


// #  rune and formatting
// recap all the type of variables that we know.
-- main.go --
package main

import "fmt"

func main() {
	var msg, age, price, check = "Hello Gopher!!!", 55, 22.52, true

	fmt.Println("msg:", msg)
	fmt.Println("age:", age)
	fmt.Println("price:", price)
	fmt.Println("check:", check)
}
// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true

-- main.go --
package main

import "fmt"

func main() {
	var msg, age, price, check = "Hello Gopher!!!", 55, 22.52, true

	var r rune = '词' // 'A'  😷

	fmt.Println("msg:", msg)
	fmt.Println("age:", age)
	fmt.Println("price:", price)
	fmt.Println("check:", check)
	fmt.Println("r:", r)
}
// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true
// r: 35789

-- main.go --
package main

import "fmt"

func main() {
	var msg, age, price, check = "Hello Gopher!!!", 55, 22.52, true

	var r rune = '😷' // 'A'  😷

	fmt.Println("msg:", msg)
	fmt.Println("age:", age)
	fmt.Println("price:", price)
	fmt.Println("check:", check)
	fmt.Printf("r: %c\n", r)
}

// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true
// r: 😷

-- main.go --
package main

import "fmt"

func main() {
	var msg, age, price, check = "Hello Gopher!!!", 55, 22.52, true

	var r rune = '😷' // 'A'  😷

	fmt.Printf("msg: %s\n", msg)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %f\n", price)
	fmt.Printf("check: %t\n", check)
	fmt.Printf("r: %c\n", r)
}

// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.520000
// check: true
// r: 😷

-- main.go --
package main

import "fmt"

func main() {
	var msg, age, price, check = "Hello Gopher!!!", 55, 22.52, true

	var r rune = '😷' // 'A'  😷

	fmt.Printf("msg: %s\n", msg)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %.2f\n", price)
	fmt.Printf("check: %t\n", check)
	fmt.Printf("r: %c\n", r)
}

// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true
// r: 😷

-- main.go --
package main

import "fmt"

func main() {
	var msg, age, price, check = "Hello Gopher!!!", 55, 22.52, true

	var r rune = '😷' // 'A'  😷

	fmt.Printf("msg: %#v\n", msg)
	fmt.Printf("age: %#v\n", age)
	fmt.Printf("price: %#v\n", price)
	fmt.Printf("check: %#v\n", check)
	fmt.Printf("r: %c\n", r)
}

// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true
// r: 😷

// print type of variable
-- main.go --
package main

import "fmt"

func main() {
	msg, age, price, check := "Hello Gopher!!!", 55, 22.52, true

	var r rune = '😷' // 'A'  😷

	fmt.Printf("type: %T -- msg: %#v\n", msg, msg)
	fmt.Printf("type: %T -- age: %#v\n", age, age)
	fmt.Printf("type: %T -- price: %#v\n", price, price)
	fmt.Printf("type: %T -- check: %#v\n", check, check)
	fmt.Printf("type: %T -- r: %c\n", r, r)
}

// go run main.go
// Output:
// type: string -- msg: "Hello Gopher!!!"
// type: int -- age: 55
// type: float64 -- price: 22.52
// type: bool -- check: true
// type: int32 -- r: 😷

// # workshop: Printf
// problem: https://go.dev/play/p/0WRuhh-jFY1
// solution 1 joke raw-string: https://go.dev/play/p/Vx6DEyRL_Wx
// solution 2 joke ไก่ย่างห้าดาว: https://go.dev/play/p/j40F8V2pSn9

// เรื่อง: Avengers: Endgame
// ปี: 2019
// เรตติ้ง: 8.4
// ประเภท: Sci-Fi
// ซุปเปอร์ฮีโร่: true
// เรื่องโปรด: ⭐


// # zero value
-- main.go -- 
package main

import "fmt"

func main() {
	var msg string = "Hello Gopher!!!"
	var age int = 55
	var price float64 = 22.52
	var check bool = true
	var r rune = '😷'

	fmt.Printf("msg: %s\n", msg)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %.2f\n", price)
	fmt.Printf("check: %t\n", check)
	fmt.Printf("r: %c\n", r)
}

// go run main.go
// Output:
// msg: Hello Gopher!!!
// age: 55
// price: 22.52
// check: true
// r: 😷

-- main.go -- 
package main

import "fmt"

func main() {
	var msg string
	var age int
	var price float64
	var check bool
	var r rune

	fmt.Printf("msg: %s\n", msg)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %.2f\n", price)
	fmt.Printf("check: %t\n", check)
	fmt.Printf("r: %c\n", r)
}

// go run main.go
// Output:
// msg: 
// age: 0
// price: 0.00
// check: false
// r: 

-- main.go -- 
package main

import "fmt"

func main() {
	var msg string
	var age int
	var price float64
	var check bool
	var r rune

	fmt.Printf("msg: %q\n", msg)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %.2f\n", price)
	fmt.Printf("check: %t\n", check)
	fmt.Printf("r: %d\n", r)
}

// go run main.go
// Output:
// msg: ""
// age: 0
// price: 0.00
// check: false
// r: 0

// # Condition
-- main.go --
package main

import "fmt"

func main() {
	num := 34
	if (num == 34) { // () is optional
		fmt.Println("Yes!! it's Thirty four")
	}
}
// go run main.go
// Output:
// Yes!! it's Thirty three

-- main.go --
package main

import "fmt"

func main() {
	num := 34
	if num == 34 {
		fmt.Println("Yes!! it's Thirty four")
	}
}
// go run main.go
// Output:
// Yes!! it's Thirty three

-- main.go --
package main

import "fmt"

func main() {
	num := 34
	if num == 34 && (num > 30 || num < 39)  { // หน่องยังไงเลย
		fmt.Println("Yes!! it's Thirty four.")
	}
}
// go run main.go
// Output:
// Yes!! it's Thirty four.

-- main.go --
package main

import "fmt"

func main() {
	num := 34
	if num%2 == 0 {
		fmt.Println("เลขคู่")
	}
}
// go run main.go
// Output:
// เลขคู่

-- main.go --
package main

import "fmt"

func main() {
	num := 35
	if num%2 == 0 {
		fmt.Println("เลขคู่")
	}
}
// go run main.go
// Output:

-- main.go --
package main

import "fmt"

func main() {
	num := 35
	if num%2 == 0 {
		fmt.Println("เลขคู่")
	}else {
		fmt.Println()
	}
}
// None
// Output:

-- main.go --
package main

import "fmt"

func main() {
	num := 35
	if num%2 == 0 {
		fmt.Println("เลขคู่")
	}else {
		fmt.Println("เลขโสด")
	}
}
// go run main.go
// Output:
// เลขโสด

-- main.go --
package main

import "fmt"

func main() {
	num := 35
	if num%2 == 0 {
		fmt.Println("เลขคู่")
	} else if num == 35 {
		fmt.Println("")
	} else {
		fmt.Println("เลขโสด")
	}
}
// None
// Output:

-- main.go --
package main

import "fmt"

func main() {
	num := 35
	if num%2 == 0 {
		fmt.Println("เลขคู่")
	} else if num == 35 {
		fmt.Println("คนมีคู่ไม่รู้หรอก")
	} else {
		fmt.Println("เลขโสด")
	}
}
// go run main.go
// Output:
// คนมีคู่ไม่รู้หรอก

-- main.go -- // short if
package main

import	"fmt"
import	"math"

func main() {
	lim := 225.0
	v := math.Pow(10, 2)
	if v < lim {
		fmt.Println("x power n is:", v)
	} else {
		fmt.Printf("x power n is %g over limit %g.\n", v, lim)
	}
}
// go run main.go
// Output:
// x power n is: 100

-- main.go -- // best practice for multiple import
package main

import (
	"fmt"
	"math"
)

func main() {
	lim := 225.0
	v := math.Pow(10, 2)
	if v < lim {
		fmt.Println("x power n is:", v)
	} else {
		fmt.Printf("x power n is %g over limit %g.\n", v, lim)
	}
}
// go run main.go
// Output:
// x power n is: 100

-- main.go -- 
package main

import (
	"fmt"
	"math"
)

func main() {
	lim := 225.0

	if v := math.Pow(10, 2); v < lim {
		fmt.Println("x power n is:", v)
	} else {
		fmt.Printf("x power n is %g over limit %g.\n", v, lim)
	}
}
// go run main.go
// Output:
// x power n is: 100

-- main.go --
package main

import (
	"fmt"
	"math"
)

func main() {
	lim := 225.0

	if v := math.Pow(10, 2); v < lim {
		fmt.Println("x power n is:", v)
	} else {
		fmt.Printf("x power n is %g over limit %g.\n", v, lim)
	}
	fmt.Println(v)
}

// go run main.go
// Output:
// # command-line-arguments
// ./main.go:16:14: undefined: v


// # workshop: if
// problem: https://go.dev/play/p/5qBXcXFYOJ5
// ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการแสดงผลคำว่า "Disappointed 😞"
// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการแสดงผลคำว่า "Normal 😐"
// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการแสดงผลคำว่า "Good 🥰"
// กรณีอื่นๆ ให้ทำการแสดงผลคำว่า "🤷🤷🤷🤷"
// solution 1: https://go.dev/play/p/ZsnHXfHZmNh

// # Condition : Switch case
-- main.go --
package main

import "fmt"

func main() {
	today := "Saturday"

	if today == "Saturday" {
		fmt.Println("today is Saturday")
	} else if today == "Monday" {
		fmt.Println("today is weekdays")
	} else {
		fmt.Printf("สวัสดีวัน %v อยู่ดีมีแฮงเดย์\n", today)
	}
}

// go run main.go
// Output:
// today is Saturday

// there is no break
-- main.go --
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

// go run main.go
// Output:
// today is Saturday

-- main.go --
package main

import "fmt"

func main() {
	today := "Saturday"

	switch today {
	case "Saturday":
		fmt.Println("today is Saturday")
		fallthrough
	case "Monday":
		fmt.Println("today is weekdays")
	default:
		fmt.Printf("สวัสดีวัน %v อยู่ดีมีแฮงเดย์\n", today)
	}
}

// go run main.go
// Output:
// today is Saturday
// today is weekdays

// short statement
-- main.go --
package main

import "fmt"

func main() {


	switch today := "Saturday"; today {
	case "Saturday":
		fmt.Println("today is Saturday")
	case "Monday":
		fmt.Println("today is weekdays")
	default:
		fmt.Printf("สวัสดีวัน %v อยู่ดีมีแฮงเดย์\n", today)
	}
}

// go run main.go
// Output:
// today is Saturday
// today is weekdays

// multiple match
-- main.go --
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

// go run main.go
// Output:
// today is weekdays

-- main.go --
package main

import "fmt"

func main() {


	switch today := "Tuesday"; today {
	case "Saturday":
		fmt.Println("today is Saturday")
	case "Monday", "Tuesday":
		fmt.Println("today is weekdays")
	default:
		fmt.Printf("สวัสดีวัน %#v อยู่ดีมีแฮงเดย์\n", today)
	}
}

// go run main.go
// Output:
// today is weekdays

// Switch with no expression
-- main.go --
package main

import "fmt"

func main() {
	num := 1
	switch {
	case num < 0:
		fmt.Println("nagative number.")
	case num == 0:
		fmt.Println("zero.")
	case num > 0:
		fmt.Println("positive number.")
	default:
		fmt.Println("🤷🤷🤷🤷")
	}
}

// go run main.go
// Output:
// positive number.

// # workshop: Switch case
// problem: https://go.dev/play/p/AKlYvUycUWj
// ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการแสดงผลคำว่า "Disappointed 😞"
// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการแสดงผลคำว่า "Normal 😐"
// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการแสดงผลคำว่า "Good 🥰"
// กรณีอื่นๆ ให้ทำการแสดงผลคำว่า "🤷🤷🤷🤷"

// solution 1: https://go.dev/play/p/UmtyGq2vuKI

// TODO: # function
-- main.go --
func discount() {

}


// Search
// Vote 8 for moivieID
// Ticket Movie
// awards

// # workshop: function
// problem: https://go.dev/play/p/Lpl1dtmHbXn
	// กำหนด: 1.สร้างฟังก์ชันชื่อ emote และรับพารามิเตอร์หนึ่งตัวชื่อ ratings เป็นตัวแปรชนิด float64 และคืนค่าเป็น string
	// 	   ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการคืนค่าคำว่า "Disappointed 😞"
	// 	   ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการคืนค่าคำว่า "Normal 😐"
	// 	   ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการคืนค่าคำว่า "Good 🥰"
	//  	   กรณีอื่นๆ ให้ทำการคืนค่าคำว่า "🤷🤷🤷🤷"

// solution 1: https://go.dev/play/p/b8QAB5cJVlc


// # array
-- main.go --
package main

import "fmt"

func main() {
	var skills string
	fmt.Printf("%#v\n", skills)
}

// go run main.go
// Output:
// ""

-- main.go --
package main

import "fmt"

func main() {
	var skills [3]string
	fmt.Printf("%#v\n", skills)
}

// go run main.go
// Output:
// [3]string{"", "", ""}

// zero value of array
// don't name the array skillArray
-- main.go --
package main

import "fmt"

func main() {
	var skills [3]int
	fmt.Printf("%#v\n", skills)
}

// go run main.go
// Output:
// [3]int{0, 0, 0}

-- main.go --
package main

import "fmt"

func main() {
	var skills [3]bool
	fmt.Printf("%#v\n", skills)
}

// go run main.go
// Output:
// [3]bool{false, false, false}


// ready to use.
-- main.go --
package main

import "fmt"

func main() {
	var skills [3]string
	fmt.Printf("%#v\n", skills)

	s := skills[1]
	fmt.Printf("%#v\n", s)
}

// go run main.go
// Output:
// [3]string{"", "", ""}
// ""

// index go from 0 to n-1
-- main.go --
package main

import "fmt"

func main() {
	var skills [3]string = [3]string{"", "", ""}
	fmt.Printf("%#v\n", skills)

	s := skills[1]
	fmt.Printf("%#v\n", s)
}

// None
// Output:

// read array
-- main.go --
package main

import "fmt"

func main() {
	var skills [3]string = [3]string{"js", "go", "python"}
	fmt.Printf("%#v\n", skills)

	s := skills[1]
	fmt.Printf("%#v\n", s)
}

// go run main.go
// Output:
// [3]string{"js", "go", "python"}
// "go"

-- main.go --
package main

import "fmt"

func main() {
	var skills [3]string = [3]string{"js", "go", "python"}
	fmt.Printf("%#v\n", skills)

	s := skills[0]
	fmt.Printf("%#v\n", s)
}

// go run main.go
// Output:
// [3]string{"js", "go", "python"}
// "js"

// write array
-- main.go --
package main

import "fmt"

func main() {
	var skills [3]string = [3]string{"js", "go", "python"}
	fmt.Printf("%#v\n", skills)

	skills[1] = "golang"

	s := skills[1] // change index back to 1
	fmt.Printf("%#v\n", s)
}

// go run main.go
// Output:
// [3]string{"js", "go", "python"}
// "golang"

-- main.go --
package main

import "fmt"

func main() {
	var skills [3]string = [3]string{"js", "go", "python"}
	fmt.Printf("%#v\n", skills)

	skills[1] = "golang"

	s := skills[1] // change index back to 1
	fmt.Printf("%#v\n", s)

	l := len(skills)
	fmt.Printf("length: %#v\n", l)
}
// go run main.go
// Output:
// [3]string{"js", "go", "python"}
// "golang"
// length: 3

-- main.go --
package main

import "fmt"

func show(sk [3]string) {
	fmt.Printf("show: %#v\n", sk)
}

func main() {
	var skills [3]string = [3]string{"js", "go", "python"}
	show(skills)


}

// go run main.go
// Output:
// show: [3]string{"js", "go", "python"}

-- main.go --
package main

import "fmt"

func show(sk [3]string) {
	fmt.Printf("show: %#v\n", sk)
}

func main() {
	var skills [3]string = [3]string{"js", "go", "python"}
	show(skills)

	var xs [4]string
	show(xs)
}

// go run main.go
// Output:
// # command-line-arguments
// ./main.go:14:7: cannot use xs (variable of type [4]string) as type [3]string in argument to show

// การประกาศแบบกระชับ
-- main.go --
package main

import "fmt"

func show(sk [3]string) {
	fmt.Printf("show: %#v\n", sk)
}

func main() {
	var skills [3]string = [3]string{"js", "go", "python"}
	// var skills = [3]string{"js", "go", "python"}
	// skills := [3]string{"js", "go", "python"}
	// skills := [3]string{}
	show(skills)
}

// # workshop: array
// problem: https://go.dev/play/p/ynSrKmKWBt-
// กำหนด: ให้ประกาศตัวแปรอาร์เรย์ประเภทหนัง(genres)ที่เก็บค่าเป็น "Action", "Adventure" และ "Fantasy" ตามลำดับ
// 	  ให้อ่านค่าของอาร์เรย์แต่ละช่องเพื่อแสดงผล ให้แสดงผลทั้งหมดตามตัวอย่าง Output ด้านล่าง
// 	  หลังจากนั้นเปลี่ยนแปลงค่าในอาร์เรย์ index ที่ 1 ให้เป็น Sci-Fi พร้อมทั้งแสดงผล เพื่อยืนยันว่าค่าเปลี่ยนจริง

// Output:
// "genres[0]: Action"
// "genres[1]: Adventure"
// "genres[2]: Fantasy"
// "genres[1]: Sci-Fi"

// solution 1: https://go.dev/play/p/MC8ICOf4rqy
// solution 2 [...]: https://go.dev/play/p/Xk2PmBw6nVk


// # loop
-- main.go --
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}
}

// go run main.go
// Output:
// i: 0
// i: 1
// i: 2
// i: 3
// i: 4

-- main.go --
package main

import "fmt"

func main() {
	for i := 0; i < 5; i = i + 2 {
		fmt.Println("i:", i)
	}
}

// go run main.go
// Output:
// i: 0
// i: 2
// i: 4


-- main.go --
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 5; i++ {
		sum = sum + i
		fmt.Println("i:", i, "sum:", sum)
	}
	fmt.Println("sum:", sum)
}

// go run main.go
// Output:
// i: 0 sum: 0
// i: 1 sum: 1
// i: 2 sum: 3
// i: 3 sum: 6
// i: 4 sum: 10
// sum: 10

// note: += 
-- main.go --
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
		fmt.Println("i:", i, "sum:", sum)
	}
	fmt.Println("sum:", sum)
}

// go run main.go
// Output:
// i: 0 sum: 0
// i: 1 sum: 1
// i: 2 sum: 3
// i: 3 sum: 6
// i: 4 sum: 10
// sum: 10

// scope of variable
-- main.go --
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		sum := 0
		sum += i
		fmt.Println("i:", i, "sum:", sum)
	}
}

// go run main.go
// Output:
// i: 0 sum: 0
// i: 1 sum: 1
// i: 2 sum: 2
// i: 3 sum: 3
// i: 4 sum: 4

// prepare for while loop
// i not relavent
-- main.go --
package main

import "fmt"

func main() {
	sum := 1
	for i := 0; sum < 5; i++ {
		sum += sum
		fmt.Println("i:", i, "sum:", sum)
	}
	fmt.Println("sum done:", sum)
}

// go run main.go
// Output:
// i: 0 sum: 2
// i: 1 sum: 4
// i: 2 sum: 8
// sum done: 8

-- main.go --
package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 5; {
		sum += sum
		fmt.Println("sum:", sum)
	}
	fmt.Println("sum done:", sum)
}

// go run main.go
// Output:
// sum: 2
// sum: 4
// sum: 8
// sum done: 8

// remove ;
-- main.go --
package main

import "fmt"

func main() {
	sum := 1
	for sum < 5 {
		sum += sum
		fmt.Println("sum:", sum)
	}
	fmt.Println("sum done:", sum)
}

// go run main.go
// Output:
// sum: 2
// sum: 4
// sum: 8
// sum done: 8

// forever loop
-- main.go --
package main

import "fmt"

func main() {
	for true {
		fmt.Println("print forever.")
	}
}

// go run main.go
// Output:
// print forever.
// print forever.
// ^Csignal: interrupt

// we will need this in future when we learn go concurrency
-- main.go --
package main

import "fmt"

func main() {
	for {
		fmt.Println("print forever.")
	}
}

// go run main.go
// Output:
// print forever.
// print forever.
// ^Csignal: interrupt


// for-range
-- main.go -- 
package main

import "fmt"

func main() {
	skills := [3]string{"js", "go", "python"}
	for i := range skills {
		fmt.Println("index:", i)
	}
}

// go run main.go
// Output:
// index: 0
// index: 1
// index: 2

-- main.go --
package main

import "fmt"

func main() {
	skills := [3]string{"js", "go", "python"}
	for i := range skills {
		val := skills[i]
		fmt.Println("index:", i, "value:", val)
	}
}

// go run main.go
// Output:
// index: 0 value: js
// index: 1 value: go
// index: 2 value: python

-- main.go --
package main

import "fmt"

func main() {
	skills := [3]string{"js", "go", "python"}
	for i, val := range skills {
		fmt.Println("index:", i, "value:", val)
	}
}

// go run main.go
// Output:
// index: 0 value: js
// index: 1 value: go
// index: 2 value: python


-- main.go --
package main

import "fmt"

func main() {
	skills := [3]string{"js", "go", "python"}
	for i, val := range skills {
		fmt.Println("value:", val)
	}
}

// go run main.go
// # command-line-arguments
// ./main.go:7:6: i declared but not used

// _ underscore
-- main.go --
package main

import "fmt"

func main() {
	skills := [3]string{"js", "go", "python"}
	for _, val := range skills {
		fmt.Println("value:", val)
	}
}

// go run main.go
// Output:
// value: js
// value: go
// value: python

// # workshop: for, for-range
// problem: https://go.dev/play/p/MCDLMUHsoVV
// กำหนด: 1. ให้ใช้ for loop ทำการเปลี่ยนค่าในอาร์เรย์ genres โดยเพิ่มคำนำหน้า "Movie: " ทุกค่า ดังนี้ "Movie: Action", "Movie: Adventure", "Movie: Fantasy"
// 	    	2. ให้แสดงผลค่า genres ทีละค่า โดยใช้ for-range

// Output:
// before for loop: [3]string{"Action", "Adventure", "Fantasy"}
// after  for loop: [3]string{"Movie: Action", "Movie: Adventure", "Movie: Fantasy"}
// genres[0]: Movie: Action
// genres[1]: Movie: Adventure
// genres[2]: Movie: Fantasy

// solution 1: https://go.dev/play/p/dXS8IwBEmMe

// # slice
-- main.go --
package main

import "fmt"

func main() {
	skills := [3]string{"js", "go", "python"}
	fmt.Printf("%T => %v\n", skills, skills)
}

// go run main.go
// Output:
// [3]string => [js go python]

-- main.go --
package main

import "fmt"

func main() {
	skills := []string{"js", "go", "python"}
	fmt.Printf("%T => %v\n", skills, skills)
}

// go run main.go
// Output:
// []string => [js go python]

// %v value only 
// %#v value and type
-- main.go --
package main

import "fmt"

func main() {
	skills := []string{"js", "go", "python"}
	fmt.Printf("%T => %#v\n", skills, skills)
}

// go run main.go
// Output:
// []string => []string{"js", "go", "python"}

// ใช้อ่าน เขียน ได้เหมือน array เลย
-- main.go --
package main

import "fmt"

func main() {
	skills := []string{"js", "go", "python"}
	fmt.Printf("%#v\n", skills)

	skills[1] = "golang"
	fmt.Printf("%#v\n", skills)

	s := skills[0]
	fmt.Printf("%#v\n", s)
}

// go run main.go
// Output:
// []string{"js", "go", "python"}
// []string{"js", "golang", "python"}
// "js"


// ใช้ loop ได้เหมือน array เลย 
-- main.go --
package main

import "fmt"

func main() {
	skills := []string{"js", "go", "python"}
	fmt.Printf("%#v\n", skills)

	skills[1] = "golang"
	fmt.Printf("%#v\n", skills)

	s := skills[0]
	fmt.Printf("%#v\n", s)

	for i := 0; i < len(skills); i++ {
		fmt.Printf("for => %#v\n", skills[i])
	}

	for _, val := range skills {
		fmt.Printf("for-range value: %#v\n", val)
	}
}

// go run main.go
// Output:
// []string{"js", "go", "python"}
// []string{"js", "golang", "python"}
// "js"
// for => "js"
// for => "golang"
// for => "python"
// for-range value: "js"
// for-range value: "golang"
// for-range value: "python"

// แล้วมันต่างจากอาร์เรย์ตรงไหนหน่อง?
// note: append คือต่อท้าย
// go idiomatic way to replace slice with new value from appends
-- main.go --
package main

import "fmt"

func main() {
	skills := []string{"js", "go", "python"}
	fmt.Printf("length: %d -- val:%#v\n", len(skills), skills)

	skills = append(skills, "ruby")
	fmt.Printf("length: %d -- val:%#v\n", len(skills), skills)
}

// go run main.go
// Output:
// length: 3 -- val:[]string{"js", "go", "python"}
// length: 4 -- val:[]string{"js", "go", "python", "ruby"}


// no side effect from append
-- main.go --
package main

import "fmt"

func main() {
	skills := []string{"js", "go", "python"}
	fmt.Printf("length: %d -- val:%#v\n", len(skills), skills)

	_ = append(skills, "ruby")
	fmt.Printf("length: %d -- val:%#v\n", len(skills), skills)
}

// go run main.go
// Output:
// length: 3 -- val:[]string{"js", "go", "python"}
// length: 3 -- val:[]string{"js", "go", "python"}

-- main.go --
package main

import "fmt"

func main() {
	skills := []string{"js", "go", "python"}
	fmt.Printf("length: %d -- val:%#v\n", len(skills), skills)

	ss := append(skills, "ruby")
	fmt.Printf("length: %d -- val:%#v\n", len(skills), skills)
	fmt.Printf("length: %d -- ss:%#v\n", len(ss), ss)
}

// go run main.go
// Output:
// length: 3 -- val:[]string{"js", "go", "python"}
// length: 3 -- val:[]string{"js", "go", "python"}
// length: 4 -- ss:[]string{"js", "go", "python", "ruby"}


// multiple append values
-- main.go --
package main

import "fmt"

func main() {
	skills := []string{"js", "go", "python"}
	fmt.Printf("length: %d -- val:%#v\n", len(skills), skills)

	skills = append(skills, "ruby", "java", "erlang")
	fmt.Printf("length: %d -- val:%#v\n", len(skills), skills)
}

// go run main.go
// Output:
// length: 3 -- val:[]string{"js", "go", "python"}
// length: 6 -- val:[]string{"js", "go", "python", "ruby", "java", "erlang"}

// there is no problem with difference size
-- main.go --
package main

import "fmt"

func show(sk []string) { // it is a slice now not an array
	fmt.Printf("show: %#v\n", sk)
}

func main() {
	skills := []string{"js", "go", "python"}
	show(skills)

}

// go run main.go
// Output:
// show: []string{"js", "go", "python"}

-- main.go --
package main

import "fmt"

func show(sk []string) {
	fmt.Printf("show: %#v\n", sk)
}

func main() {
	skills := []string{"js", "go", "python"}
	show(skills)

	genres := []string{"Action", "Adventure", "Fantasy", "Sci-Fi"}
	show(genres)
}

// go run main.go
// Output:
// show: []string{"js", "go", "python"}
// show: []string{"Action", "Adventure", "Fantasy", "Sci-Fi"}


// note: why they called it "slice"? => เพราะว่ามันหั่นได้
-- main.go -- 
package main

import "fmt"

func show(sk []string) {
	fmt.Printf("show: %#v\n", sk)
}

func main() {
	skills := []string{"js", "go", "python"}
	show(skills)

	s1 := skills[0:2]
	show(s1)
}

// go run main.go
// Output:
// show: []string{"js", "go", "python"}
// show: []string{"js", "go"}

//NOTE: เปิด slide hafl-open range

// the length of slice after sliced
-- main.go --
package main

import "fmt"

func show(sk []string) {
	l := len(sk)
	fmt.Printf("length: %d -- show: %#v\n", l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}
	show(skills)

	s1 := skills[0:2]
	show(s1)
}

// go run main.go
// Output:
// length: 3 -- show: []string{"js", "go", "python"}
// length: 2 -- show: []string{"js", "go"}


// อยากได้ go, python
-- main.go -- 
package main

import "fmt"

func show(sk []string) {
	l := len(sk)
	fmt.Printf("length: %d -- show: %#v\n", l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}
	show(skills)

	s1 := skills[0:2]
	show(s1)

	s2 := skills[1:3]
	show(s2)
}

// go run main.go
// Output:
// length: 3 -- show: []string{"js", "go", "python"}
// length: 2 -- show: []string{"js", "go"}
// length: 2 -- show: []string{"go", "python"}

// note: inline
-- main.go --
package main

import "fmt"

func show(sk []string) {
	l := len(sk)
	fmt.Printf("length: %d -- show: %#v\n", l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}
	show(skills)

	show(skills[0:2])
	show(skills[1:3])
}

// go run main.go
// Output:
// length: 3 -- show: []string{"js", "go", "python"}
// length: 2 -- show: []string{"js", "go"}
// length: 2 -- show: []string{"go", "python"}

// ถ้าอยากได้ทั้งหมดสมาชิกทุกตัวละ
-- main.go -- 
package main

import "fmt"

func show(sk []string) {
	l := len(sk)
	fmt.Printf("length: %d -- show: %#v\n", l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}
	show(skills)

	show(skills[0:2])
	show(skills[1:3])


	show(skills[0:3])
}

// go run main.go
// Output:
// length: 3 -- show: []string{"js", "go", "python"}
// length: 2 -- show: []string{"js", "go"}
// length: 2 -- show: []string{"go", "python"}
// length: 3 -- show: []string{"js", "go", "python"}

// เราจะรู้ได้ไงว่าของข้างใน slice มีสมาชิกทั้งหมดเท่าไร?
-- main.go -- 
package main

import "fmt"

func show(sk []string) {
	l := len(sk)
	fmt.Printf("length: %d -- show: %#v\n", l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}
	show(skills)

	show(skills[0:2])
	show(skills[1:3])

	l := len(skills)
	show(skills[0:l])
}

// go run main.go
// Output:
// length: 3 -- show: []string{"js", "go", "python"}
// length: 2 -- show: []string{"js", "go"}
// length: 2 -- show: []string{"go", "python"}
// length: 3 -- show: []string{"js", "go", "python"}

// show line by line
-- main.go -- 
package main

import "fmt"

func show(sk []string) {
	l := len(sk)
	fmt.Printf("length: %d -- show: %#v\n", l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}
	show(skills)

	show(skills[0:2])
	show(skills[1:3])

	l := len(skills)
	show(skills[0:l])
	show(skills[0:])
	show(skills[:l])
	show(skills[:])
}

// go run main.go
// Output:
// length: 3 -- show: []string{"js", "go", "python"}
// length: 2 -- show: []string{"js", "go"}
// length: 2 -- show: []string{"go", "python"}
// length: 3 -- show: []string{"js", "go", "python"}
// length: 3 -- show: []string{"js", "go", "python"}
// length: 3 -- show: []string{"js", "go", "python"}
// length: 3 -- show: []string{"js", "go", "python"}

// หน่อง นั่นมันหมายความว่า [1:3] => [1:] และ [0:2] => [:2]
-- main.go --
package main

import "fmt"

func show(sk []string) {
	l := len(sk)
	fmt.Printf("length: %d -- show: %#v\n", l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}
	show(skills)

	show(skills[:2])
	show(skills[1:])

	l := len(skills)
	show(skills[0:l])
	show(skills[0:])
	show(skills[:l])
	show(skills[:])
}

// go run main.go
// Output:
// length: 3 -- show: []string{"js", "go", "python"}
// length: 2 -- show: []string{"js", "go"}
// length: 2 -- show: []string{"go", "python"}
// length: 3 -- show: []string{"js", "go", "python"}
// length: 3 -- show: []string{"js", "go", "python"}
// length: 3 -- show: []string{"js", "go", "python"}
// length: 3 -- show: []string{"js", "go", "python"}


// หน่อง zero value ของ slice คืออะไร?
-- main.go --
package main

import "fmt"

func main() {
	var ss []int
	fmt.Printf("%#v\n", ss)
	if ss == nil {
		fmt.Println("nil")
	}
}

// go run main.go
// Output:
// []int(nil)
// nil


// แล้วจะเกิดอะไรขึ้นถ้าหน่อง append ใส่ nil slice?
-- main.go --
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

// go run main.go
// Output:
// []int(nil)
// nil
// []int{33}

// แล้วจะเกิดอะไรขึ้นถ้าหน่อง อ่านของใน nil slice?
// runtime error นะ (ถ้าเป็น array compile error)
-- main.go --
package main

import "fmt"

func main() {
	var ss []int
	fmt.Printf("%#v\n", ss)
	if ss == nil {
		fmt.Println("nil")
	}

	v := ss[0]
	fmt.Printf("%#v\n", v)
}

// go run main.go
// Output:
// []int(nil)
// nil
// panic: runtime error: index out of range [0] with length 0

// เช่นกันถ้าเขียนค่าใน nil slice 
-- main.go -- 
package main

import "fmt"

func main() {
	var ss []int
	fmt.Printf("%#v\n", ss)
	if ss == nil {
		fmt.Println("nil")
	}

	ss[0] = 33
	fmt.Printf("%#v\n", ss)
}
// go run main.go
// Output:
// []int(nil)
// nil
// panic: runtime error: index out of range [0] with length 0

// เราสามารถประกาศ slice แล้วพร้อมใช้เหมือน array ได้ไหม?
// เราต้องทำการ allowcate membery หรือ ให้ Go สร้าง underlying array ให้ ณ ขณะสร้าง
// zero initialize เรียบร้อย
-- main.go --
package main

import "fmt"

func main() {
	ss := make([]int, 3)
	fmt.Printf("%#v\n", ss)
}

// go run main.go
// Output:
// []int{0, 0, 0}

// จะเขียน จะอ่านก็สบายใจ
-- main.go --
package main

import "fmt"

func main() {
	ss := make([]int, 3)
	fmt.Printf("%#v\n", ss)

	ss[1] = 36
	fmt.Printf("%#v\n", ss)
	v := ss[1]
	fmt.Printf("%#v\n", v)
}

// go run main.go
// Output:
// []int{0, 0, 0}
// []int{0, 36, 0}
// 36


// note: ตอนนี้พวกเราเข้าใจการจัดการ slice แล้ว มันเหมือน array เลยต่างแค่ ขนาดของมันยืดหยุ่น
// คราวนี้หน่องจะทำอะไร งงๆ ให้ดู
// สิ่งที่หน่องจะเล่าต่อไปนี้ ไม่จำเป็นต้องจำ เพียงแค่รู้ไว้ผ่านๆ ก็พอ เพราะทุกอย่าง Go จัดการให้
-- main.go --
package main

import "fmt"

func show(sk []string) {
	l := len(sk)
	fmt.Printf("length: %d -- show: %#v\n", l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}

	s1 := skills[0:2]
	show(s1)

	s2 := skills[1:3]
	show(s2)
}

// go run main.go
// Output:
// length: 2 -- show: []string{"js", "go"}
// length: 2 -- show: []string{"go", "python"}

// note: ขอปรับ ฟังก์ชัน show นิดหนึ่งเพื่อดูง่าย
-- main.go --
package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	fmt.Printf("%s: len: %d -- show: %v\n", tag, l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}

	s1 := skills[0:2]
	show("s1", s1)

	s2 := skills[1:3]
	show("s1", s2)
}

// go run main.go
// Output:
// s1: len: 2 -- show: [js go]
// s1: len: 2 -- show: [go python]



-- main.go -- 
package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	fmt.Printf("%s: len: %d -- show: %v\n", tag, l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}

	s1 := skills[0:2]
	show("s1", s1)

	s2 := skills[1:3]
	show("s2", s2)

	fmt.Println()

	s1[1] = "gopher"
	show("s1", s1)
	show("s2", s2)
}

// go run main.go
// Output:
// s1: len: 2 -- show: [js go]
// s2: len: 2 -- show: [go python]
//
// s1: len: 2 -- show: [js gopher]
// s2: len: 2 -- show: [gopher python]

// s2[0] = "gopher"
-- main.go --
package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	fmt.Printf("%s: len: %d -- show: %v\n", tag, l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}

	s1 := skills[0:2]
	show("s1", s1)

	s2 := skills[1:3]
	show("s2", s2)

	fmt.Println()

	s2[0] = "gopher"
	show("s1", s1)
	show("s2", s2)
}

// go run main.go
// Output:
// s1: len: 2 -- show: [js go]
// s2: len: 2 -- show: [go python]
//
// s1: len: 2 -- show: [js gopher]
// s2: len: 2 -- show: [gopher python]


// ทำไมถึงเป็นอย่างนั้นหน่อง?
// เพราะจริงๆ เวลาเราประกาศตัวแปร slice Go จะทำการสร้างอาร์เรย์ แล้วเอาค่าที่เราระบุไปเก็บไว้ที่ อาร์เรย์นั้น
// นั่นมันหมายความว่าจริงๆ ถ้าเรามี อาร์เรย์ อยู่เราก็สามารถสร้างตัวแปรสไลด์จากอาร์เรย์นั้นได้เลย
// ถ้าหน่องเปลี่ยน skills กลับไปเป็น อาร์เรย์ ขนาด 3
-- main.go -- 
package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	fmt.Printf("%s: len: %d -- show: %v\n", tag, l, sk)
}

func main() {
	skills := [3]string{"js", "go", "python"}

	s1 := skills[0:2]
	show("s1", s1)

	s2 := skills[1:3]
	show("s2", s2)

	fmt.Println()

	s2[0] = "gopher"
	show("s1", s1)
	show("s2", s2)
}

// go run main.go
// Output:
// s1: len: 2 -- show: [js go]
// s2: len: 2 -- show: [go python]
//
// s1: len: 2 -- show: [js gopher]
// s2: len: 2 -- show: [gopher python]


// note: ค่าในอาร์เรย์เปลี่ยนด้วย
// s1, s2 อ้างอิงไปที่อาร์เรย์ตัวเดียวกัน 
// prinf อาร์เรย์ skills ออกมาดู
-- main.go --
package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	fmt.Printf("%s: len: %d -- show: %v\n", tag, l, sk)
}

func main() {
	skills := [3]string{"js", "go", "python"}

	s1 := skills[0:2]
	show("s1", s1)

	s2 := skills[1:3]
	show("s2", s2)

	fmt.Println()

	s2[0] = "gopher"
	show("s1", s1)
	show("s2", s2)

	fmt.Printf("skills: %#v\n", skills)
}

// go run main.go
// Output:
// s1: len: 2 -- show: [js go]
// s2: len: 2 -- show: [go python]
//
// s1: len: 2 -- show: [js gopher]
// s2: len: 2 -- show: [gopher python]
// skills: [3]string{"js", "gopher", "python"}

// จะเกิดอะไรขึ้นถ้าเรา append ต่อท้าย s1 เพื่อเพิ่มค่าเข้าไป?
// ผลลัพธ์ s2 เปลี่ยนจาก python เป็น c++
-- main.go --
package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	fmt.Printf("%s: len: %d -- show: %v\n", tag, l, sk)
}

func main() {
	skills := [3]string{"js", "go", "python"}

	s1 := skills[0:2]
	s1 = append(s1, "c++")
	show("s1", s1)

	s2 := skills[1:3]
	show("s2", s2)

	fmt.Println()

	s2[0] = "gopher"
	show("s1", s1)
	show("s2", s2)

	fmt.Printf("skills: %#v\n", skills)
}

// go run main.go
// Output:
// s1: len: 3 -- show: [js go c++]
// s2: len: 2 -- show: [go c++]
//
// s1: len: 3 -- show: [js gopher c++]
// s2: len: 2 -- show: [gopher c++]
// skills: [3]string{"js", "gopher", "c++"}

// แล้วถ้าเรา append ต่อท้าย s2 แทนละ มันเหมือนกันไหม
// ผลลัพธ์ s2 ไปสร้าง array ตัวใหม่เพื่อ จัดการค่าของ s2 เอง
-- main.go --
package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	fmt.Printf("%s: len: %d -- show: %v\n", tag, l, sk)
}

func main() {
	skills := [3]string{"js", "go", "python"}

	s1 := skills[0:2]
	show("s1", s1)

	s2 := skills[1:3]
	s2 = append(s2, "c++")
	show("s2", s2)

	fmt.Println()

	s2[0] = "gopher"
	show("s1", s1)
	show("s2", s2)

	fmt.Printf("skills: %#v\n", skills)
}

// go run main.go
// Output:
// s1: len: 2 -- show: [js go]
// s2: len: 3 -- show: [go python c++]
//
// s1: len: 2 -- show: [js go]
// s2: len: 3 -- show: [gopher python c++]
// skills: [3]string{"js", "go", "python"}

// จะรู้ได้ไงว่า Go จะไปสร้างอาร์เรย์ตัวใหม่ให้ slice เมื่อไหร่?
// หน่องขอแนะนำตัวคร ตัวใหม่ ชื่อ ว่า capacity หรือ ความจุของ slice นั่นเอง
-- main.go -- 
package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	fmt.Printf("%s: len: %d -- show: %v\n", tag, l, sk)
}

func main() {
	skills := []string{"js", "go", "python"}

	s1 := skills[0:4]
	show("s1", s1)
}

// go run main.go
// Output:
// panic: runtime error: slice bounds out of range [:4] with capacity 3

-- main.go -- 
package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	c := cap(sk)
	fmt.Printf("%s: len: %d cap: %d -- show: %v\n", tag, l, c, sk)
}

func main() {
	skills := []string{"js", "go", "python"}

	s1 := skills[0:3]
	show("s1", s1)
}

// go run main.go
// Output:
// s1: len: 3 cap: 3 -- show: [js go python]

// skills[1:3]  cap จะเหลือ 2
-- main.go -- 
package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	c := cap(sk)
	fmt.Printf("%s: len: %d cap: %d -- show: %v\n", tag, l, c, sk)
}

func main() {
	skills := []string{"js", "go", "python"}

	s1 := skills[1:3]
	show("s1", s1)
}

// go run main.go
// Output:
// s1: len: 2 cap: 2 -- show: [go python]


// NOTE: เปิด slice visualize ไปอธิบายเรื่องการ slice และ capacity

// # workshop: slice
// joke: https://go.dev/play/p/QSiQgIx4-qn  -- ยากจัด หยอกๆ
// problem: https://go.dev/play/p/C90FxMb1bQm
// กำหนด: 1. เราได้เก็บสะสมคะแนนโหวตผู้ชมไว้เป็น 2 ชุด ที่เก็บอยู่ในตัวแปร xs และ ys เรียบร้อยแล้ว
// 	  		2. ให้ทำการรวมคะแนนโหวตที่อยู่ในตัวแปร xs และ ys ไปเก็บไว้ในตัวแปร votes ตามลำดับ (ค่าใน xs ทั้งหมดแล้วต่อด้วย ys).
//	  		3. ทำการแสดงผลคะแนนโหวตของผู้ชมที่อยู่ในตำแหน่ง(index)ที่ 5 ถึง 8 ของ votes ออกมาทางหน้าจอ
//
// solution: https://go.dev/play/p/STHSciSQMb4

// # struct
-- main.go --
package main

import "fmt"

func main() {
	name := "Basic Go"
	instructor := "AnuchitO"
	price := 9999

	fmt.Println("name:", name)
	fmt.Println("instructor:", instructor)
	fmt.Println("price:", price)
}

// go run main.go
// Output:
// name: Basic Go
// instructor: AnuchitO
// price: 9999

// นิยมโครงสร้างข้อมูล course
-- main.go --
package main

import "fmt"

type course struct {
	name       string
	instructor string
	price      float64
}

func main() {
	name := "Basic Go"
	instructor := "AnuchitO"
	price := 9999

	fmt.Println("name:", name)
	fmt.Println("instructor:", instructor)
	fmt.Println("price:", price)
}

// None
// Output:

// ห่อไว้ใน course
// comma ทุกบรรทัด
-- main.go --
package main

import "fmt"

type course struct {
	name       string
	instructor string
	price      float64
}

func main() {
	c := course{
		name:       "Basic Go",
		instructor: "AnuchitO",
		price:      9999,
	}

	fmt.Println("name:", c.name)
	fmt.Println("instructor:", c.instructor)
	fmt.Println("price:", c.price)
}

// go run main.go
// Output:
// name: Basic Go
// instructor: AnuchitO
// price: 9999

// read / write struct
-- main.go --
package main

import "fmt"

type course struct {
	name       string
	instructor string
	price      float64
}

func main() {
	c := course{
		name:       "Basic Go",
		instructor: "AnuchitO",
		price:      9999,
	}

	n := c.name
	c.instructor = "Nong"

	fmt.Println("name:", n)
	fmt.Println("instructor:", c.instructor)
	fmt.Println("price:", c.price)
}

// go run main.go
// Output:
// name: Basic Go
// instructor: Nong
// price: 9999

// การประกาศ struct มีกี่แบบ?
-- main.go --
package main

import "fmt"

type course struct {
	name       string
	instructor string
	price      float64
}

func main() {
	var c1 = course{name: "Basic Go", instructor: "AnuchitO", price: 9999}
	var c2 = course{"Basic Go", "AnuchitO", 9999}
	var c3 = course{instructor: "AnuchitO"}
	var c4 = course{}

	fmt.Printf("c1: %#v\n", c1)
	fmt.Printf("c2: %#v\n", c2)
	fmt.Printf("c3: %#v\n", c3)
	fmt.Printf("c4: %#v\n", c4) // zero value
}

// go run main.go
// Output:
// c1: main.course{name:"Basic Go", instructor:"AnuchitO", price:9999}
// c2: main.course{name:"Basic Go", instructor:"AnuchitO", price:9999}
// c3: main.course{name:"", instructor:"AnuchitO", price:0}
// c4: main.course{name:"", instructor:"", price:0}

// # workshop: struct
// problem: https://go.dev/play/p/GsAErtuEKDD
// กำหนด: 1. ให้นิยามโครงสร้างข้อมูล movie เพื่อเก็บ ชื่อเรื่อง(string) ปี(ตัวเลข) เรตติ้ง(ตัวเลขทศนิยม) ประเภท(slice ของ string) และ isSuperHero(bool)
	// 	  2. ให้ประกาศตัวแปรเพื่อเก็บหนัง Avengers: Endgame, ปี:2019, เรตติ้ง:8.4, ประเภท:["Action", "Drama"] และ isSuperHero:true
	//	  3. ให้ประกาศตัวแปรเพื่อเก็บหนัง Infinity War, ปี:2018, เรตติ้ง:8.4, ประเภท:["Action", "Sci-Fi"] และ isSuperHero:true
	//	  4. ให้เก็บหนังทั้งสองเรื่องไว้ในตัวแปร mvs
	// 	  5. ทำการวนลูปเพื่อแสดงผลหนังทีละเรื่อง
	//
	// Output:
	// main.movie{title:"Avengers: Endgame", year:2019, rating:8.4, genres:[]string{"Action", "Drama"}, isSuperHero:true}
	// main.movie{title:"Avengers: Infinity War", year:2018, rating:8.4, genres:[]string{"Action", "Sci-Fi"}, isSuperHero:true}
// solution: https://go.dev/play/p/2yEHW0R92M8

// # method : normal function discount
-- main.go -- 
package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func discount(c course) int {
	p := c.price - 599
	fmt.Println("discount:", p)
	return p
}

func main() {
	c := course{"Basic Go", "AnuchitO", 9999}
	fmt.Printf("%#v\n", c)

	d := discount(c)
	fmt.Println("discount price:", d)
}

// go run main.go
// Output:
// main.course{name:"Basic Go", instructor:"AnuchitO", price:9999}
// discount: 9400
// discount price: 9400

// error undefined discount
-- main.go --
package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func (c course) discount() int {
	p := c.price - 599
	fmt.Println("discount:", p)
	return p
}

func main() {
	c := course{"Basic Go", "AnuchitO", 9999}
	fmt.Printf("%#v\n", c)

	d := discount(c)
	fmt.Println("discount price:", d)
}

// go run main.go
// Output:
// ./main.go:20:7: undefined: discount

// method discount ตัวแปรเป็น receiver
-- main.go --
package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func (c course) discount() int {
	p := c.price - 599
	fmt.Println("discount:", p)
	return p
}

func main() {
	c := course{"Basic Go", "AnuchitO", 9999}
	fmt.Printf("%#v\n", c)

	d := c.discount()
	fmt.Println("discount price:", d)
}

// go run main.go
// Output:
// main.course{name:"Basic Go", instructor:"AnuchitO", price:9999}
// discount: 9400
// discount price: 9400


// we can have many methods
-- main.go --
package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func (c course) discount() int {
	p := c.price - 599
	fmt.Println("discount:", p)
	return p
}

func (c course) info() {
	fmt.Println("name:", c.name)
	fmt.Println("instructor:", c.instructor)
	fmt.Println("price:", c.price)
}

func main() {
	c := course{"Basic Go", "AnuchitO", 9999}
	fmt.Printf("%#v\n", c)

	d := c.discount()
	fmt.Println("discount price:", d)
	c.info()
}

// go run main.go
// Output:
// main.course{name:"Basic Go", instructor:"AnuchitO", price:9999}
// discount: 9400
// discount price: 9400
// name: Basic Go
// instructor: AnuchitO
// price: 9999

// # workshop: method
// problem: https://go.dev/play/p/cOEbj0wB9DO
// solution 1: https://go.dev/play/p/gErFHYR3dhY


// # pointer
-- main.go --
package main

import "fmt"

func main() {
	var price int = 9999
	fmt.Println(price, &price)
}

// go run main.go
// Output:
// 9999 0xc00012a008

// short declaration
-- main.go --
package main

import "fmt"

func main() {
	var price int = 9999
	addr := &price
	fmt.Println(price, addr)
}

// go run main.go
// Output:
// 9999 0xc00012a008

// print type of variable addr
-- main.go --
package main

import "fmt"

func main() {
	var price int = 9999
	addr := &price
	fmt.Println(price, addr)
	fmt.Printf("%T\n", addr)
}

// go run main.go
// Output:
// 9999 0xc00012a008
// *int

// full declaration
-- main.go --
package main

import "fmt"

func main() {
	var price int = 9999
	var addr *int = &price
	fmt.Println(price, addr)
	fmt.Printf("%T\n", addr)
}

// go run main.go
// Output:
// 9999 0xc00012a008
// *int

// go is statically typed language.
-- main.go --
package main
package main

import "fmt"

func main() {
	var price int = 9999
	var addr *int = &price
	fmt.Println(price, addr)

	addr = 9400
	fmt.Println(price, addr)
}

// go run main.go
// Output:
// ./main.go:10:9: cannot use 9400 (untyped int constant) as *int value in assignment

// change value of price via addr
// สิ่งนี้เรียกว่า (de-reference)
-- main.go --
package main

import "fmt"

func main() {
	var price int = 9999
	var addr *int = &price
	fmt.Println(price, addr)

	*addr = 9400 // write
	fmt.Println(price, addr)
	v := *addr // read
	fmt.Println(v)
}

// go run main.go
// Output:
// 9999 0xc000014160
// 9400 0xc000014160
// 9400

// NOTE: open slide pointer visualization

// หลังจาก changePrice จะได้ผลลัพธ์เป็นอย่างไร?
-- main.go --
package main

import "fmt"

func changePrice(price int) {
	price = price - 599
}

func main() {
	var price int = 9999
	var addr *int = &price

	changePrice(price)
	fmt.Println("after ", price, addr)
}

// go run main.go
// Output:
// after  9999 0xc000014160

// ค่าก็เปลี่ยนนะ ทำไมถึงเป็นอย่างนั้น
// เพราะมันเป็นตัวแปรคนละตัว go always pass by value
-- main.go --
package main

import "fmt"

func changePrice(price int) {
	price = price - 599
	fmt.Println("change", price, &price)
}

func main() {
	var price int = 9999
	var addr *int = &price

	changePrice(price)
	fmt.Println("after ", price, addr)
}

// go run main.go
// Output:
// change 9400 0xc000014168
// after  9999 0xc000014160

// changePrice ไปเป็น pointer
-- main.go --
package main

import "fmt"

func changePrice(p *int) {
	*p = *p - 599
	fmt.Println("change", *p, &p)
}

func main() {
	var price int = 9999
	var addr *int = &price

	changePrice(&price)
	fmt.Println("after ", price, addr)
}

// go run main.go
// Output:
// change 9400 0xc00000e028
// after  9400 0xc000014160

// ถึงอย่างนั้น Go ยังคง pass by value เสมอ
-- main.go --
package main

import "fmt"

func changePrice(p *int) {
	fmt.Println("value in p is", p)
	*p = *p - 599
	fmt.Println("change", *p, &p)
}

func main() {
	var price int = 9999
	var addr *int = &price

	changePrice(&price)
	fmt.Println("after ", price, addr)
}

// go run main.go
// Output:
// value in p is 0xc000014160
// change 9400 0xc00000e028
// after  9400 0xc000014160

// หน่องครับถ้าแบบเป็น struct ไปที่ ฟังก์ชันใดๆ แล้วภายในฟังก์ชันเปลี่ยนค่าของ struct จะเป็นยังไง?
// เป็นเพราะอะไรครับ? เก่งมาก
-- main.go --
package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func discount(c course) int {
	c.price = c.price - 599
	fmt.Println("discount:", c.price)
	return c.price
}

func main() {
	c := course{"Basic Go", "AnuchitO", 9999}

	d := discount(c)
	fmt.Println("discount price:", d)
	fmt.Println("price:", c.price)
}

// go run main.go
// Output:
// discount: 9400
// discount price: 9400
// price: 9999

// แล้วทำยังไงค่า price ใน course ถึงจะเปลี่ยน?
-- main.go -- 
package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func discount(c *course) int {
	c.price = c.price - 599
	fmt.Println("discount:", c.price)
	return c.price
}

func main() {
	c := course{"Basic Go", "AnuchitO", 9999}

	d := discount(&c)
	fmt.Println("discount price:", d)
	fmt.Println("price:", c.price)
}

// go run main.go
// Output:
// discount: 9400
// discount price: 9400
// price: 9400

// หน่องแต่ c มันเป็น pointer นะ ทำไมเราถึงไม่ต้อง (*c).price ละ? (de-reference)
-- main.go --
package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func discount(c *course) int {
	(*c).price = (*c).price - 599
	fmt.Println("discount:", c.price)
	return c.price
}

func main() {
	c := course{"Basic Go", "AnuchitO", 9999}

	d := discount(&c)
	fmt.Println("discount price:", d)
	fmt.Println("price:", c.price)
}

// go run main.go
// Output:
// discount: 9400
// discount price: 9400
// price: 9400

// ท่าการประการ struct pointer ทั่วไป
-- main.go --
package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func discount(c *course) int {
	c.price = c.price - 599
	fmt.Println("discount:", c.price)
	return c.price
}

func main() {
	c := &course{"Basic Go", "AnuchitO", 9999}

	d := discount(c)
	fmt.Println("discount price:", d)
	fmt.Println("price:", c.price)
}

// go run main.go
// Output:
// discount: 9400
// discount price: 9400
// price: 9400

// เข้าถึงค่าในตัวแปรเหมือนเดิมไม่ต้อง de-reference อีกที
-- main.go --
package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func discount(c *course) int {
	c.price = c.price - 599
	fmt.Println("discount:", c.price)
	return c.price
}

func main() {
	c := &course{"Basic Go", "AnuchitO", 9999}
	c.price = 5599
	d := discount(c)
	fmt.Println("discount price:", d)
	fmt.Println("price:", c.price)
}

// go run main.go
// Output:
// discount: 5000
// discount price: 5000
// price: 5000

// หน่องแล้ว method มีผลเหมือนกันไหม?
-- main.go --
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

// go run main.go
// Output:
// discount: 9400
// discount price: 9400
// price: 9400


// แล้วถ้าประการเป็น pointer แต่ reciever ไม่เป็น pointer ละ?
-- main.go --
package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func (c course) discount() int {
	c.price = c.price - 599
	fmt.Println("discount:", c.price)
	return c.price
}

func main() {
	c := &course{"Basic Go", "AnuchitO", 9999}

	d := c.discount()
	fmt.Println("discount price:", d)
	fmt.Println("price:", c.price)
}

// go run main.go
// Output:
// discount: 9400
// discount price: 9400
// price: 9999

// คำถาม zero value ของตัวแปร pointer คืออะไร
-- main.go --
package main

import "fmt"

func main() {
	var addr *int
	fmt.Println("zero value:", addr)
}

// go run main.go
// Output:
// zero value: <nil>

// # workshop: pointer of struct
// problem: https://go.dev/play/p/-J6kojiy-r1
	// กำหนด: 1. ให้สร้าง method addVote รับพารามิเตอร์ rating เป็น float64
	// 	  		2. ให้ method addVote เพิ่มค่า rating เข้าไปใน slice votes ของ struct movie
// solution 1: https://go.dev/play/p/EwSe14HAU4P


// # interface
-- main.go --
package main

import "fmt"

func main() {
	var v interface{}
	v = 36
	fmt.Printf("%T %#v\n", v, v)
	v = "hello"
	fmt.Printf("%T %#v\n", v, v)
}

// go run main.go
// Output:
// int 36
// string "hello"

-- main.go --
package main

import "fmt"

func main() {
	var v any
	v = 36
	fmt.Printf("%T %#v\n", v, v)
	v = "hello"
	fmt.Printf("%T %#v\n", v, v)
}

// go run main.go
// Output:
// int 36
// string "hello"

-- main.go --
package main

import "fmt"

func show(val int){
	fmt.Println(val)
}

func main() {
	var v any
	v = 36
	show(v)
}

// go run main.go
// Output:
// ./main.go:12:7: cannot use v (variable of type any) as type int in argument to show:
//         need type assertion

// type assertion
-- main.go --
package main

import "fmt"

func show(val int){
	fmt.Println(val)
}

func main() {
	var v any
	v = 36
	show(v.(int))
}

// go run main.go
// Output:
// int 36

// ถ้าเราไม่อยากทำ type assertion ทำไง?
package main

import "fmt"

func show(val any) {
	fmt.Println(val)
}

func main() {
	var v any
	v = 36
	show(v)
}

// go run main.go
// Output:
// 36

// แต่ข้างใน show function ก็ไม่รู้ type อยู่ดี
-- main.go --
package main

import "fmt"

func show(val any) {
	i := val.(int)
	i = i + 10
	fmt.Println(i)
}

func main() {
	var v any
	v = 36
	show(v)
}

// go run main.go
// Output:
// 46

// จะเกิดอะไรขึ้นถ้าไม่ใช่ int?
-- main.go --
package main

import "fmt"

func show(val any) {
	i := val.(int)
	i = i + 10
	fmt.Println(i)
}

func main() {
	var v any
	v = "gopher"
	show(v)
}

// go run main.go
// Output:
// panic: interface conversion: interface {} is string, not int

-- main.go --
package main

import "fmt"

func show(val any) {
	i, ok := val.(int)
	if ok {
		i = i + 10
		fmt.Println(i)
	} else {
		fmt.Println("not int")
	}
}

func main() {
	var v any
	v = "gopher"
	show(v)
}

// go run main.go
// Output:
// not int

// แบบนี้ก็ if else บานเลยอะดิ
// interface : type switch
-- main.go --
package main

import "fmt"

func show(val any) {
	switch v := val.(type) {
	case int:
		i := v + 10
		fmt.Println("int:", i)
	case string:
		s := v + ", Hi!"
		fmt.Println("string:", s)
	default:
		fmt.Println("not handle type.")
	}
}

func main() {
	var v any
	v = "gopher"
	show(v)
}

// go run main.go
// Output:
// string: gopher, Hi!

// อย่างที่เราเห็น interface มันคือชนิดของข้อมูลแบบหนึ่ง แล้วมันเกี่ยวกับ method ที่หน่องว่ายังไงอะ?
// สมมุติ หน่องมี ฟังก์ชัน sale
-- main.go --
package main

import "fmt"

func sale(val any) {
	fmt.Println("sale:", val)
}

func main() {
	v := "gopher"
	sale(v)
}

// go run main.go
// Output:
// sale: gopher

-- main.go -- 
package main

import "fmt"

func sale(val interface{}) {
	fmt.Println("sale:", val)
}

func main() {
	v := "gopher"
	sale(v)
}

// go run main.go
// Output:
// sale: gopher

// นิยาม interface promotion
-- main.go -- 
package main

import "fmt"

type promotion interface{}

func sale(val promotion) {
	fmt.Println("sale:", val)
}

func main() {
	v := "gopher"
	sale(v)
}

// go run main.go
// Output:
// sale: gopher

-- main.go --
package main

import "fmt"

type promotion interface {
	discount() int
}

func sale(val promotion) {
	fmt.Println("sale:", val)
}

func main() {
	v := "gopher"
	sale(v)
}

// go run main.go
// Output:
// ./main.go:15:7: cannot use v (variable of type string) as type promotion in argument to sale:
//         string does not implement promotion (missing discount method)

// สนใจแค่ method signature ตรงกัน ไม่ได้ สนใจว่ามี fields อะไรบ้าง
-- main.go -- 
package main

import "fmt"

type promotion interface {
	discount() int
}

func sale(val promotion) {
	fmt.Println("sale:", val)
}

type course struct{}

func (c course) discount() int {
	return 599
}

func main() {
	v := course{}
	sale(v)
}

// go run main.go
// Output:
// sale: {}

// อย่างน้อยต้องมี method ตามที่ interface ต้องการ
// interface คือข้อตกลงหรือการกำหนดว่าชนิดของข้อมูลนั้นต้องมีพฤติกรรม(method) แบบนี้ถึงจะยอมรับ
-- main.go --
package main

import "fmt"

type promotion interface {
	discount() int
}

func sale(val promotion) {
	fmt.Println("sale:", val)
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info:", c)
}

func main() {
	v := course{}
	sale(v)
}

// go run main.go
// Output:
// sale: {}

// ภายในฟังก์ชัน sale จะมองเห็นเท่าที่ interface มี
-- main.go -- 
package main

import "fmt"

type promotion interface {
	discount() int
}

func sale(val promotion) {
	fmt.Println("sale:", val.discount())
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info:", c)
}

func main() {
	v := course{}
	sale(v)
}

// go run main.go
// Output:
// sale: 599

// มองไม่เห็น method info นะ
-- main.go -- 
package main

import "fmt"

type promotion interface {
	discount() int
}

func sale(val promotion) {
	fmt.Println("sale:", val.discount())
	val.info()
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info:", c)
}

func main() {
	v := course{}
	sale(v)
}

// go run main.go
// Output:
// ./main.go:11:6: val.info undefined (type promotion has no field or method info)

// method info still there.
-- main.go --
package main

import "fmt"

type promotion interface {
	discount() int
}

func sale(val promotion) {
	fmt.Println("sale:", val.discount())
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info:", c)
}

func main() {
	v := course{}
	sale(v)
	v.info()
}

// ถ้ามี interface อีกตัวละ?
// many-to-many not one to many
-- main.go --
package main

import "fmt"

type promotion interface {
	discount() int
}

func sale(val promotion) {
	fmt.Println("sale:", val.discount())
}

type display interface {
	info()
}

func show(val display) {
	val.info()
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info:", c)
}

func main() {
	v := course{}
	sale(v)
	show(v)
}

// go run main.go
// Output:
// sale: 599
// info: {}

// ถ้าอยากมี function summary ที่ มีทั้ง method discount and info ด้วยได้ไหม
-- main.go --
package main

import "fmt"

func summary(val xxx) {
	fmt.Println("discount price:", val.discount())
	val.info()
}

-- main.go --
package main

import "fmt"

type xxx interface {
	discount() int
	info()
}

func summary(val xxx) {
	fmt.Println("discount price:", val.discount())
	val.info()
}

-- main.go --
package main

import "fmt"

type presenter interface {
	discount() int
	info()
}

func summary(val presenter) {
	fmt.Println("discount price:", val.discount())
	val.info()
}

-- main.go --
package main

import "fmt"

type presenter interface {
	discount() int
	info()
}

func summary(val presenter) {
	fmt.Println("discount price:", val.discount())
	val.info()
}

func main() {
	v := course{}
	summary(v)
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info:", c)
}

// go run main.go
// Output:
// discount price: 599
// info: {}

// แล้วมันยังเข้ากันได้กับ promotion inteface อยู่ไหม
-- main.go --
package main

import "fmt"

type presenter interface {
	discount() int
	info()
}

type promotion interface {
	discount() int
}

func summary(val presenter) {
	fmt.Println("discount price:", val.discount())
	val.info()
}

func sale(val promotion) {
	fmt.Println("sale:", val.discount())
}

func main() {
	v := course{}
	summary(v)
	sale(v)
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info:", c)
}

// go run main.go
// Output:
// discount price: 599
// info: {}
// sale: 599

// แล้วมันยังเข้ากันได้กับ display inteface อยู่ไหม
-- main.go -- 
package main

import "fmt"

type presenter interface {
	discount() int
	info()
}

type promotion interface {
	discount() int
}

type display interface {
	info()
}

func summary(val presenter) {
	fmt.Println("discount price:", val.discount())
	val.info()
}

func sale(val promotion) {
	fmt.Println("sale:", val.discount())
}

func show(val display) {
	val.info()
}

func main() {
	v := course{}
	summary(v)
	sale(v)
	show(v)
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info:", c)
}

// go run main.go
// Output:
// discount price: 599
// info: {}
// sale: 599
// info: {}

// embedding interface
-- main.go --
package main

import "fmt"

type presenter interface {
	promotion
	display
}

type promotion interface {
	discount() int
}

type display interface {
	info()
}

func summary(val presenter) {
	fmt.Println("discount price:", val.discount())
	val.info()
}

func sale(val promotion) {
	fmt.Println("sale:", val.discount())
}

func show(val display) {
	val.info()
}

func main() {
	v := course{}
	summary(v)
	sale(v)
	show(v)
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info:", c)
}

// go run main.go
// Output:
// discount price: 599
// info: {}
// sale: 599
// info: {}

// # workshop: interface
// problem: https://go.dev/play/p/Edrx94i72PK
// solution 1: https://go.dev/play/p/53uNT-ReSgG


// TODO: # error


// TODO: # error : creation

// go run main.go
// Output:

// TODO: # error : handling

// go run main.go
// Output:

// TODO: # maps

// go run main.go
// Output:

// TODO: # project creation

// go run main.go
// Output:

// TODO: # project creation: package separation

// go run main.go
// Output:

// TODO: # project creation: public / private

// go run main.go
// Output:

// TODO: # project : library public to github

// TODO: # type convertion : number -> number

// go run main.go
// Output:

// TODO: # type convertion : string -> number

// go run main.go
// Output:

// TODO: # json

// go run main.go
// Output:

// TODO: # json : json to struct -- Unmarshal

// go run main.go
// Output:

// TODO: # json : struct to json -- Marshal

// go run main.go
// Output:

// TODO: # Constants
-- main.go --
package main

import "fmt"

type day int

func main() {
	const (
		Sunday day = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Sunday :", Sunday)
	fmt.Println("Monday :", Monday)
	fmt.Println("Tuesday :", Tuesday)
	fmt.Println("Wednesday :", Wednesday)
	fmt.Println("Thursday :", Thursday)
	fmt.Println("Friday :", Friday)
	fmt.Println("Saturday :", Saturday)
}

// go run main.go
// Output:

// TODO: # Constants : iota

// go run main.go
// Output:

// TODO: # variadic function

// go run main.go
// Output:

// TODO: # Naked return function

// go run main.go
// Output:

// TODO: # First-class function

// go run main.go
// Output:

// TODO: # time

// go run main.go
// Output:

// TODO: # http protocol

// go run main.go
// Output:

// TODO: # rest api

// go run main.go
// Output:

// TODO: simple api in go
-- main.go --
package main

import (
	"fmt"
	"log"
	"net/http"
)

func moviesHandler(w http.ResponseWriter, req *http.Request) {
	method := "GET"
	fmt.Fprintf(w, "hello %s movies", method)
}

func main() {
	http.HandleFunc("/movies", moviesHandler)

	log.Fatal(http.ListenAndServe("localhost:2565", nil))
}

-- main.go --
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Movie struct {
	ImdbID      string   `json:"imdbID"`
	Title       string   `json:"title"`
	Year        int      `json:"year"`
	Rating      float32  `json:"rating"`
	IsSuperHero bool     `json:"isSuperHero"`
}

var movies []Movie

func moviesHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		t := Movie{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			fmt.Fprintf(w, "error: ", err)
			return
		}

		movies = append(movies, t)
		fmt.Printf("% #v\n", movies)

		fmt.Fprintf(w, "hello %s created movies", "POST")
		return
	}

	if req.Method == "GET" {
		b, err := json.Marshal(movies)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

func main() {
	http.HandleFunc("/movies", moviesHandler)

	log.Fatal(http.ListenAndServe(":2565", nil))
}


// TODO: # rest api : GET /mmm

// go run main.go
// Output:

// TODO: # rest api : POST /mmm

// go run main.go
// Output:

// TODO: # workshop: Certificate

// go run main.go
// Output:

// TODO: # testing

// TODO: generic bonus
// https://go.dev/play/p/00L1eUwODUO?v=gotip


// Theme: 
// example -- Course
// workshop -- Movie
// m.AddWatchlist
// m.AddRating
// m.GetAverageRating
// m.GetTopRated
// m.UserReviews
// m.AddUserReview


// movie votes: [9, 5, 7, 8, 3, 8, 10, 7, 8, 10, 9, 7]
// movie ratings: votes / len(votes)
// go run main.go
// Output:


package main

import (
	"play.ground/foo"
)

func main() {
	foo.Bar()
}

-- go.mod --
module play.ground

-- foo/foo.go --
package foo

import "fmt"

func Bar() {
	fmt.Println("The Go playground now has support for multiple files!")
}
