package main

import "fmt"

func main() {
	n := 7
	result := fib(n)
	fmt.Println(result)
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

/*

n  : 1 2 3 4 5 6 7  8
fib: 1 1 2 3 5 8 13 21

Binary Recursion (Đệ qui nhị phân)
	Kỹ thuật dùng trong tìm kiếm ,sắp xếp, cây
	Đệ qui nhị phân là dạng đệ qui gọi 2 lần chính nó


*/
