package main

import "fmt"

func main() {
	var n int

	fmt.Println("Tính tổng số nguyên từ 1 -> n:")
	fmt.Scan(&n)

	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println(sum)
}

/*

fori

for ex1;ex2;ex3
 statement

ex1: khởi tạo
ex2: điều kiện
ex3: biểu thức điều kiện lặp

*/
