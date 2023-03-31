package main

import "fmt"

func main() {
	m := []int{9, 7, 6, 3, 2, 1}
	x := timXBangNhiPhanTrongMangGiamDan(m, len(m), 94)
	if x == -1 {
		fmt.Println("Không tìm thấy x trong mảng")
	} else {
		fmt.Println(x)
	}
}

func timXBangNhiPhanTrongMangGiamDan(m []int, n, x int) int {
	l, r := 0, n-1

	for i := 0; i < n; i++ {
		if l <= r {
			mid := (l + r) / 2

			if x == m[mid] {
				return mid
			} else if x > m[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}
