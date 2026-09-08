package main

import (
	"fmt"
	"math"
)

func constants() {
	const pi float64 = 3.1415
	const n = 2 // int

	const (
		pi2 float64 = 3.14
		e   float64 = 2.72
	)

	const (
		a = 1
		b // 1
		c // 1
		d = 3
	)

	const ( // enum!
		flag1 = iota // Starts from zero
		flag2        // 1
		flag3        // 2
		flag4        // 3
	)
}

func floats() {
	if 0 == 0.0 { // won't work for variables
		fmt.Println("0 == 0.0")
	}

	// value := 0.1
	// value += 0.2 //0.30000000000400000000000
	// fmt.Println(value)
	fmt.Printf("%.55f\n", 0.1)
	fmt.Printf("%.55f\n", 0.2)
}

func math_package() {
	x := math.Trunc(1.123)
	fmt.Printf("Trunc: %f\n", x)
}

func main() {
	constants()
	floats()
	math_package()
}
