package main

import "fmt"

func main() {
	result := choose(5, 3)
	fmt.Println(result)
}

func choose(n, k int) int {
	if k == 0 || n == k {
		return 1
	}

	return choose(n-1, k) + choose(n-1, k-1)
}
