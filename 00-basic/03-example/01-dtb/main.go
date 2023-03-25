package main

import "fmt"

func main() {
	var t, l, h float64

	fmt.Println("Nhập điểm toán: ")
	fmt.Scan(&t)

	fmt.Println("Nhập điểm toán: ")
	fmt.Scan(&l)

	fmt.Println("Nhập điểm toán: ")
	fmt.Scan(&h)

	dtb := (t + l + h) / 3.0

	fmt.Println(fmt.Sprintf("%.2f", dtb))
}
