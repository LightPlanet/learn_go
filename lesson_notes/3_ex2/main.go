package main

import "fmt"

func factorial(n uint64) (ret uint64) {
	ret = 1
	for i := n; i > 1; i-- {
		ret *= i
	}
	return
}

func main() {
	var n uint64
	fmt.Print("Enter the natural number: ")
	fmt.Scan(&n)
	fmt.Printf("%d! is %d\n", n, factorial(n))
}
