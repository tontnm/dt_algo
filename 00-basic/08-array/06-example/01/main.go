package main

import "fmt"

func main() {
	var n int
	fmt.Println("Nhập số phần tử trong mảng:")
	fmt.Scan(&n)

	m := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Nhập phần tử m[%d]=", i)
		fmt.Scan(&m[i])
	}

	fmt.Println("Mảng vừa nhập:")
	for i := 0; i < n; i++ {
		fmt.Printf("%.2f\t", m[i])
	}
	fmt.Println()

	fmt.Println("Mảng sắp xếp theo thứ tự giảm dần:")
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if m[i] < m[j] {
				t := m[i]
				m[i] = m[j]
				m[j] = t
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%.2f\t", m[i])
	}
}
