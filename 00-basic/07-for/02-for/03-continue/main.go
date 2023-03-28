package main

import "fmt"

func main() {
	var n int

	fmt.Println("Tính tổng số nguyên từ 1 -> n, bỏ qua số 2, 4:")
	fmt.Scan(&n)

	sum := 0

	for i := 1; i <= n; i++ {
		if i == 2 || i == 4 {
			continue
		}
		sum += i
	}

	fmt.Println(sum)
}

/*

Lệnh continue dùng để kết thúc lần lặp hiên tại và bắt đầu lần lặp tiếp theo
Lệnh continue chỉ được dùng bên trong lệnh for
Lện continue thường đi cùng if

*/
