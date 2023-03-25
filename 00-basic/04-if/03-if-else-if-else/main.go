package main

import "fmt"

func main() {
	var dtb float64

	fmt.Println("Nhập điểm trung bình:")
	fmt.Scan(&dtb)

	if dtb < 0 || dtb > 10 {
		fmt.Println("nhập sai")
	} else {
		if dtb >= 5 {
			fmt.Println("Đậu")
		} else {
			fmt.Println("Rớt")
		}
	}
}

/*

dùng if...else lồng nhàu thì else sẽ kết hợp với if gần nhất chưa có else

*/
