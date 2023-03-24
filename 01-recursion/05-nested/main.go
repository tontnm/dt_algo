package main

import "fmt"

func main() {
	m, n := 2, 1
	result := ackerman(m, n)
	fmt.Println(result)
}

func ackerman(m, n int) int {
	if m == 0 {
		return n + 1
	} else if n == 0 {
		return ackerman(m-1, 1)
	} else {
		return ackerman(m-1, ackerman(m, n-1))
	}
}

/*

là loại đệ qui gọi đệ qui vào đối số của hàm đệ qui

*/
