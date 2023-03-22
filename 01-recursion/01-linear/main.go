package main

import "fmt"

func main() {
	n := 4
	result := factorial(n)
	fmt.Println(fmt.Sprintf("%d!=%d", n, result))
}

// linear recursion (đệ qui tuyến tính: hàm tự gọi đến nó 1 lần và giảm theo cấp bậc)
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
