package main

import "fmt"

func main() {
	var n int
	fmt.Println("Nhập số phần tử trong mảng:")
	fmt.Scan(&n)

	m := make([]float64, n)

	i := 0
	for i < n {
		fmt.Printf("Nhập phần tử m[%d]=", i)
		fmt.Scan(&m[i])
		if i > 0 && m[i] <= m[i-1] {
			continue
		}
		i++
	}

	fmt.Println("Mảng vừa nhập:")
	for i := 0; i < n; i++ {
		fmt.Printf("%.2f\t", m[i])
	}
	fmt.Println()
}
