package main

import "fmt"

func main() {
	var num int

	fmt.Println("XÁC ĐỊNH SỐ CHẴN LẺ")
	fmt.Println("Nhập số:")
	fmt.Scan(&num)

	switch num % 2 {
	case 0:
		fmt.Println("Chẵn")
	case 1:
		fmt.Println("Lẻ")
	default:
		fmt.Println("Lẻ")
	}
}

/*

switch là một cấu trúc lựa chọn có nhiều nhánh, được sử dụng khi có nhiều lựa chọn

Giải thích
Expression đã đc định trị
Nếu giá trị của expression = value-1 thì thực hiện statement1 và thoát khỏi switch
Nếu giá trị của expression # value-1 thì so sánh với value-2, nếu bằng value-2
thì thực hiện statement2 và thoát khỏi switch, so sánh đến value_n
Nếu tất cả các phép so sánh đều sai thì thực hiện statement của default
Dùng switch khi 1 biến có nhiều giá trị khác nhau để so sánh

*/
