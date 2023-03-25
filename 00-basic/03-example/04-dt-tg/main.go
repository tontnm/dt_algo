package main

import "fmt"

func main() {
	var a, h float64

	fmt.Println("Nhập cạnh")
	fmt.Scan(&a)

	fmt.Println("Nhập chiều cao")
	fmt.Scan(&h)

	dt := 0.5 * a * h

	fmt.Println(dt)
}
