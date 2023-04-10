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

	max := m[0]
	for i := 1; i < n; i++ {
		if m[i] > max {
			max = m[i]
		}
	}
	fmt.Printf("Phần tử lớn nhất trong mảng là: %.2f", max)
	fmt.Println()

	min := m[0]
	for i := 1; i < n; i++ {
		if m[i] < min {
			min = m[i]
		}
	}
	fmt.Printf("Phần tử nhỏ nhất trong mảng là: %.2f", min)
	fmt.Println()

	sum := 0.0
	for i := 0; i < n; i++ {
		sum += m[i]
	}
	fmt.Printf("Tổng các phần tử của mảng là: %.2f", sum)
}
