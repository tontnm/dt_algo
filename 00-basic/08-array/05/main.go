package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var r, c int

	fmt.Println("Nhập dòng & cột:")
	fmt.Scanln(&r, &c)

	m := make([][]int, r)
	for i := range m {
		m[i] = make([]int, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m[i][j] = rand.Intn(101)
		}
	}

	fmt.Println("Mảng 2 chiều:")
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf("%d\t", m[i][j])
		}
		fmt.Println()
	}

	var sr int

	fmt.Println("Chọn dòng bạn muốn in:")
	fmt.Scan(&sr)

	for j := 0; j < c; j++ {
		fmt.Printf("%d\t", m[sr][j])
	}
	fmt.Println()

	var sc int

	fmt.Println("Chọn cột bạn muốn in:")
	fmt.Scan(&sc)

	for i := 0; i < r; i++ {
		fmt.Printf("%d\n", m[i][sc])
	}

	if r == c {
		fmt.Println("Đường chéo chính")
		for i := 0; i < r; i++ {
			fmt.Printf("%d\t", m[i][i])
		}
		fmt.Println()

		fmt.Println("Đường chéo phụ")
		for i := 0; i < r; i++ {
			fmt.Printf("%d\t", m[i][r-i-1])
		}
	}
}
