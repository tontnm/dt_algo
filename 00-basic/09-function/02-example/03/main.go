package main

import "fmt"

func main() {
	var n int

	fmt.Println("Nhập n:")
	fmt.Scan(&n)

	fmt.Println(fibonacci(n))

	fmt.Println("Day so chi tiet la:")
	xuatdayfib(n)
}

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func xuatdayfib(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}
