package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Println("Nhập số a,b,c. Tìm số lớn nhất")
	fmt.Scanln(&a, &b, &c)

	if a > b && a > c {
		fmt.Printf("%d là số lớn nhất", a)
	} else if b > a && b > c {
		fmt.Printf("%d là số lớn nhất", b)
	} else if c > a && c > b {
		fmt.Printf("%d là số lớn nhất", c)
	} else {
		fmt.Println("không có số lớn nhất")
	}
}
