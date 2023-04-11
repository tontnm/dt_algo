package main

import "fmt"

func main() {
	var n, x int
	fmt.Println("Nhập số phần tử:")
	fmt.Scan(&n)
	m := make([]int, n)
	NhapM(m, n)

	fmt.Println("Mảng sau khi nhập:")
	XuatM(m, n)

	fmt.Println("Nhập x:")
	fmt.Scan(&x)
	kq := TimXTrongM(m, n, x)
	if kq != -1 {
		fmt.Printf("%d có tồn tại trong mảng\n", x)
	} else {
		fmt.Printf("%d không tồn tại trong mảng\n", x)
	}

	SapXepMTangDan(m, n)
	XuatM(m, n)
}

func NhapM(m []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("m[%d]=", i)
		fmt.Scan(&m[i])
	}
}

func XuatM(m []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("m[%d]=%d\t", i, m[i])
	}
	fmt.Println()
}

func TimXTrongM(m []int, n, x int) int {
	for i := 0; i < n; i++ {
		if m[i] == x {
			return i
		}
	}

	return -1
}

func SapXepMTangDan(m []int, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if m[i] > m[j] {
				t := m[i]
				m[i] = m[j]
				m[j] = t
			}
		}
	}
}
