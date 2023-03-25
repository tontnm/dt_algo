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

*/
