package main

import "fmt"

func auto_memory_handling() {
	var ptr *int = nil
	{
		scoped := 42
		ptr = &scoped
	}
	fmt.Println(ptr, *ptr)
}

func operators() {
	is_even := func(x int) bool {
		return (x%2 == 0)
	}
	x := 0
	for {
		fmt.Print("Enter any integer: ")
		fmt.Scan(&x)
		// fmt.Printf("Your number is even: %t\n", is_even(x))
		msg := "Your number is "
		if is_even(x) {
			msg += "even"
		} else {
			msg += "not even"
		}
		fmt.Println(msg)
	}
}

func main() {
	operators()
}
