package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Enter three sides of the triangle:")
	fmt.Scanln(&a, &b, &c)

	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || b+c <= a || a+c <= b {
		// The three sides do not create a valid triangle
		fmt.Println("The given values do not create a valid triangle")
	} else {
		perimeter := a + b + c
		s := perimeter / 2.0
		diameter := math.Sqrt(s * (s - a) * (s - b) * (s - c))

		fmt.Printf("The given values create a triangle with perimeter %.2f and diameter %.2f\n", perimeter, diameter)
	}
}
