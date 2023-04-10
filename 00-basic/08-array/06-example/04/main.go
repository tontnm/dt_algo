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

	nm := make([][]int, c)
	for i := range nm {
		nm[i] = make([]int, r)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			nm[j][i] = m[i][j]
		}
	}

	fmt.Println("Ma trận hoán vị:")
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			fmt.Printf("%d\t", nm[i][j])
		}
		fmt.Println()
	}
}
