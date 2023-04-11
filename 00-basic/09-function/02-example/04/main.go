package main

import (
	"fmt"
	"math"
)

// 63
func main() {
	a, b, c := 4, 5, 6

	if kiemtracanhhople(a, b, c) {
		fmt.Println("Tam giác không hợp lệ")
		return
	}

	fmt.Printf("chu vi tam giac= %d\n", chuvitamgiac(a, b, c))
	fmt.Printf("dien tich tam giac= %f", dientichtamgiac(a, b, c))
}

func chuvitamgiac(a, b, c int) int {
	return a + b + c
}

func dientichtamgiac(a, b, c int) float64 {
	p := float64(chuvitamgiac(a, b, c)) / 2.0

	dt := math.Sqrt(p * (p - float64(a)) * (p - float64(b)) * (p - float64(c)))

	return dt
}

func kiemtracanhhople(a, b, c int) bool {
	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || b+c <= a || a+c <= b {
		return false
	}

	return true
}
