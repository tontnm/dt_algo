package main

import "fmt"

func main() {
	var kwh int
	const gia1, gia2, gia3, gia4 = 600, 700, 900, 1100

	fmt.Println("Tính tiền điện")
	fmt.Println("Nhập số kwh")
	fmt.Scan(&kwh)

	var result int

	if kwh <= 100 {
		result = kwh * gia1
	} else if kwh <= 150 {
		result = 100*gia1 + (kwh-100)*gia2
	} else if kwh <= 200 {
		result = 100*gia1 + 50*gia2 + (kwh-150)*gia3
	} else if kwh > 200 {
		result = 100*gia1 + 50*gia2 + 50*gia3 + (kwh-200)*gia4
	} else {
		fmt.Println("số kwh không hợp lệ")
	}

	fmt.Printf("Tiền điện cho %d kwh là %d", kwh, result)
}
