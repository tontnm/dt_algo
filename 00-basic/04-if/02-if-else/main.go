package main

import "fmt"

func main() {
	var a int

	fmt.Println("XÁC ĐỊNH SỐ ÂM DƯƠNG")
	fmt.Println("Nhập số a: ")
	fmt.Scan(&a)

	if a >= 0 {
		fmt.Println("a là số dương")
	} else {
		fmt.Println("a là số âm")
	}
}

/*

if expression {
	statement1
} else {
	statement2
}

*/
