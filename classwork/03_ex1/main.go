package main

import (
	"fmt"
	"time"
)

func format_time() {
	// 01 02 03 04 05
	// January 02 15:04.05 2006 | Monday

	fmt.Println(time.Now().Format("2006 01 02 15:04"))
	fmt.Println(time.Now().Format("January, 02 15:04:05"))
	fmt.Println(time.Now().Format("Mon, 02.01.06 15:04 MST"))
	fmt.Println(time.Now().Unix())
}

func is_even(x int) bool {
	return (x%2 == 0)
}

func generate_even_str(x int) string {
	ret := "is "
	if !is_even(x) {
		ret += "not "
	}
	ret += "even"
	return ret
}

func generate_sign_str(x int) string {
	ret := "is "
	if x > 0 {
		ret += "positive"
	} else if x < 0 {
		ret += "negative"
	} else {
		ret += "zero"
	}
	return ret
}

func parse_number() {
	x := 0
	for {
		fmt.Print("Enter any integer: ")
		fmt.Scan(&x)
		msg := "Your number " + generate_even_str(x) + " and " + generate_sign_str(x)
		fmt.Println(msg)
	}
}

func main() {

}
