package main

import "fmt"

func main() {
	m, n := 7, 3
	result := gcd(m, n)
	fmt.Println(result)
}

// tìm ước số chung lớn nhất
func gcd(m, n int) int {
	var r int

	if m < n {
		return gcd(n, m)
	}

	r = m % n

	if r == 0 {
		return n
	} else {
		return gcd(n, r)
	}
}
