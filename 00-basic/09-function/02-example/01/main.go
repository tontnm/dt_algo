package main

import (
	"fmt"
	"math"
)

func main() {
	ptb2(5, 7, -12)
	ptb2(1, -4, 2)
}

func ptb2(a, b, c int) {
	if a == 0 {
		if b == 0 && c == 0 { // bx+c=0
			fmt.Println("PT Vo So Nghiem")
		} else if b == 0 && c != 0 {
			fmt.Println("PT Vo Nghiệm")
		} else {
			// bx=-c ==> c=-c/b
			x := -c * 1.0 / b
			fmt.Printf("PT co 1 nghiem x=%.2f", float64(x))
		}
	} else {
		delta := math.Pow(float64(b), 2) - 4*float64(a)*float64(c)

		if delta < 0 {
			fmt.Println("PT Vo Nghiem")
		} else if delta == 0 {
			x := -b / (2 * a)
			fmt.Printf("PT Co Nghiem Kep x1=x2=%.2f", float64(x))
		} else {
			x1 := (-float64(b) - math.Sqrt(delta)) / (2 * float64(a))
			x2 := (-float64(b) + math.Sqrt(delta)) / (2 * float64(a))
			fmt.Println("2 nghiem phan biet")
			fmt.Printf("\nx1=%.2f", x1)
			fmt.Printf("\nx2=%.2f", x2)
		}
	}
}
