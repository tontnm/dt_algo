package main

import (
	"fmt"
	"math"
)

func main() {
	var x, n int

	fmt.Println("Nhập x và n")
	fmt.Scanln(&x, &n)

	var s float64
	var mau float64 = 1

	for i := 1; i <= n; i++ {
		tu := math.Pow(float64(x), float64(i))

		mau = mau * float64(i)

		s += tu / mau
	}

	fmt.Printf("s=%f", s)
}
