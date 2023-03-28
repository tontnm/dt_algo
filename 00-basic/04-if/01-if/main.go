package main

import "fmt"

func main() {
	var a int

	fmt.Println("Nhập số a: ")
	fmt.Scan(&a)

	if a >= 0 {
		fmt.Println("a là số dương")
	}
}

/*

if expression {
	statement
}

b1: expression là true, statment được thực thi, chạy các dòng lệnh sau if
b2: expression là false, thoát khỏi if, chạy các dòng lệnh sau if

*/
