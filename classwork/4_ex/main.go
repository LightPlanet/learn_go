package main

import (
	"fmt"
	"math"
)

func degToRad(deg float64) float64 {
	return deg / 180 * math.Pi
}

func RadToDeg(rad float64) float64 {
	return rad / math.Pi * 180
}

func main() {
	fmt.Printf("sin 30: %f\n", math.Sin(degToRad(30)))
	fmt.Printf("Round 1.5: %f\n", math.Round(1.5))
	fmt.Printf("Round 1.5: %f\n", math.Round(10))
	fmt.Printf("Round 1.5: %f\n", math.Round(1000))
	fmt.Printf("Round 1.5: %f\n", math.Round(0.00000000000001))
	fmt.Printf("e^2: %f\n", math.Pow(math.E, 2))
}
