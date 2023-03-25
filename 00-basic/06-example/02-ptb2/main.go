package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("giải phương trình bậc 2: ax^2+bx+c=0")

	var a, b, c int

	fmt.Println("Nhập a")
	fmt.Scan(&a)

	fmt.Println("Nhập b")
	fmt.Scan(&b)

	fmt.Println("Nhập c")
	fmt.Scan(&c)

	if a == 0 {
		if b == 0 && c == 0 {
			fmt.Println("phương trình vô số nghiệm")
		} else if b == 0 && c != 0 {
			fmt.Println("phương trình vô nghiệm")
		} else {
			//bx=-c => x=-c/b
			var x float64
			x = float64(-c) / float64(b)
			fmt.Printf("phương trình có 1 nghiệm: x=%f", x)
		}
	} else {
		delta := math.Pow(float64(b), 2) - 4*float64(a)*float64(c)

		if delta < 0 {
			fmt.Println("phương trình vô nghiệm")
		} else if delta == 0 {
			nkep := float64(-b) / (2 * float64(a))
			fmt.Printf("PT có nghiệm kép x1=x2=%f", nkep)
		} else {
			x1 := (float64(-b) - math.Sqrt(delta)) / (2 * float64(a))
			x2 := (float64(-b) + math.Sqrt(delta)) / (2 * float64(a))
			fmt.Printf("PT có 2 nghiệm phân biệt x1=%f, x2=%f", x1, x2)
		}
	}

}
