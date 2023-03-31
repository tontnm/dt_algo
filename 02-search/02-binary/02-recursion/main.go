package main

import "fmt"

func main() {
	a := []int{1, 5, 15, 19, 25, 27, 29, 31, 33, 45, 55, 88, 100}
	result := binarySearch(a, 27, 0, 15)
	fmt.Println(result)
}

func binarySearch(a []int, x, l, r int) int {
	if l > r {
		return -1
	}

	m := (l + r) / 2

	if x == a[m] {
		return m
	}

	if x < a[m] {
		return binarySearch(a, x, l, m-1)
	}

	return binarySearch(a, x, m+1, r)
}

/*

tìm kiếm nhị phân bằng đệ qui

Đệ qui: quy luật + điểm dừng

*/
