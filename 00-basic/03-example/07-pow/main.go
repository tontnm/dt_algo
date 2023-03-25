package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var x float64

	fmt.Println("Nhấp số nguyên: ")
	fmt.Scan(&n)

	fmt.Println("Nhấp số thực: ")
	fmt.Scan(&x)

	a := math.Pow(math.Pow(x, 2)+x+1, float64(n)) + math.Pow(math.Pow(x, 2)+x+1, float64(n))

	fmt.Printf("A=(%.1f^2+%.1f+1)^%d+(%.1f^2+%.1f+1)^%d=%.2f", x, x, n, x, x, n, a)
}
