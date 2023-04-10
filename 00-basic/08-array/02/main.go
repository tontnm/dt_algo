package main

import "fmt"

func main() {
	x := 10
	m := make([]int, 5)

	for i := 0; i < len(m); i++ {
		fmt.Printf("Nhập m[%d]=", i)
		fmt.Scan(&m[i])
	}

	kq := false
	for i := 0; i < len(m); i++ {
		if x == m[i] {
			kq = true
			break
		}
	}

	if kq {
		fmt.Printf("số %d xuất hiện ở trong mảng m\n", x)
	} else {
		fmt.Printf("không tìm thấy %d trong mảng m\n", x)
	}

	count := 0
	for i := 0; i < len(m); i++ {
		if x == m[i] {
			count++
		}
	}
	fmt.Printf("%d xuất hiện %d lần trong mảng m", x, count)
}
