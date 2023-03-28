package main

import "fmt"

func main() {
	n := 9
	kq := isEven(n)

	if kq {
		fmt.Printf("%d is even", n)
	} else {
		fmt.Printf("%d is odd", n)
	}
}

func isEven(n int) bool {
	if n == 0 {
		return true
	} else {
		return isOdd(n - 1)
	}
}

func isOdd(n int) bool {
	return !isEven(n)
}

/*

là loại đệ qui không gọi đệ qui trực tiếp chính nó, mà gọi một hàm khác,
trong hàm khác lại gọi lại nó

*/
