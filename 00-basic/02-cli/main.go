package main

import (
	"fmt"
)

func main() {
	var toan, van, dtb float32

	fmt.Println("Nhập điểm Toán: ")
	fmt.Scan(&toan)

	fmt.Println("Nhập điểm Văn: ")
	fmt.Scan(&van)

	dtb = (toan + van) / 2

	fmt.Println(fmt.Sprintf("Điểm trung bình là %f", dtb))
}
