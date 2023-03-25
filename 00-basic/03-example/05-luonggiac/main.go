package main

import (
	"fmt"
	"math"
)

func main() {
	angle := 30.0                  // in degrees
	rad := angle * math.Pi / 180.0 // convert to radians

	sin := math.Sin(rad)
	cos := math.Cos(rad)
	tan := math.Tan(rad)
	cot := 1.0 / math.Tan(rad)

	fmt.Printf("sin(%.f) = %.2f\n", angle, sin)
	fmt.Printf("cos(%.f) = %.2f\n", angle, cos)
	fmt.Printf("tan(%.f) = %.2f\n", angle, tan)
	fmt.Printf("cot(%.f) = %.2f\n", angle, cot)
}
