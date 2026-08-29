package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	a := 321
	b := 123
	fmt.Printf("Sum: %d\n", add(a, b))
}
