package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		if z != math.Sqrt(x) {
			z -= (z*z - x) / (2 * z)
		} else {
			return z
		}
	}
}

func main() {
	var i float64
	fmt.Print("Write a number: ")
	fmt.Scan(&i)
	fmt.Println("Square root of your number is ", Sqrt(i))
}
