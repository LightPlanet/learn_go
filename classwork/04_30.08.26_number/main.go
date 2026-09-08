package main

import (
	"fmt"
	"math"
)

// Notes --------------------------------------------------------------------------------

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

// Ex1 ----------------------------------------------------------------------------------

func degToRad(deg float64) float64 {
	return deg / 180 * math.Pi
}

func RadToDeg(rad float64) float64 {
	return rad / math.Pi * 180
}

func ex1() {
	fmt.Printf("sin 30: %f\n", math.Sin(degToRad(30)))
	fmt.Printf("Round 1.5: %f\n", math.Round(1.5))
	fmt.Printf("Round 1.5: %f\n", math.Round(10))
	fmt.Printf("Round 1.5: %f\n", math.Round(1000))
	fmt.Printf("Round 1.5: %f\n", math.Round(0.00000000000001))
	fmt.Printf("e^2: %f\n", math.Pow(math.E, 2))
}

func main() {
	constants()
	floats()
	math_package()
	ex1()
}
