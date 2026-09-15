package main

import (
	"fmt"
	"math"
	"strconv"
)

func TruncCount(v float64, count float64) float64 {
	mult := math.Pow(10, count)
	return math.Trunc(v*mult) / mult
}

// strconv.Itoa | *.Atoi ASCII <-> int

func AnyBaseIntParsing() {
	integer, err := strconv.ParseInt("KRT", 32, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(integer)
	}
}

func BoolToString() {
	// Ternary operator :D
	parse := func(v bool) string {
		if v {
			return "Yes"
		} else {
			return "No"
		}
	}
	fmt.Println(parse(true))
	fmt.Println(parse(false))
}

func main() {
	// fmt.Println(TruncCount(10.9999999, 4))
	AnyBaseIntParsing()
	BoolToString()
}
