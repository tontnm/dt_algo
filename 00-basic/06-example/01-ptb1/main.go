package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Giải phương trình bậc 1: ax+b=0")
	fmt.Println("Nhập a:")
	fmt.Scan(&a)

	fmt.Println("Nhập b:")
	fmt.Scan(&b)

	if a == 0 && b == 0 {
		fmt.Println("phương trình vô số nghiệm")
	} else if a == 0 && b != 0 {
		fmt.Println("phương trình vô nghiệm")
	} else {
		// ax=-b => x=-b/a
		var x float64
		x = float64(-b) / float64(a)
		fmt.Printf("phương trình có 1 nghiệm %f", x)
	}
}
