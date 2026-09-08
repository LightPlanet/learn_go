package main

// go mod init <project name> - creates go.mod (imports won't work otherwise)
// Uppercase is exportable by default

import (
	"1/greeting"
	"fmt"
)

// Passing by reference
func incr(v *int) {
	*v++
}

// Return tuple
func get_tuple() (int, float32, string) {
	return 42, 1.23, "text"
}

// Error report via tuple
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can't divide by zero!")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Greetings")
	greeting.SayHi()

	fmt.Println("\nPassing by reference")
	x := 10
	incr(&x)
	fmt.Println(x)

	fmt.Println("\nError handling via tuple")
	handle_error := func(r int, e error) {
		if e != nil {
			fmt.Println("Error:", e)
		} else {
			fmt.Println("Result:", r)
		}
	}
	handle_error(divide(10, 2))
	handle_error(divide(10, 0))
}
