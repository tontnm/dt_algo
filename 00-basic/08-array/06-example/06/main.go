package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 52

func main() {
	rand.Seed(time.Now().UnixNano())

	var n int

	fmt.Println("Nhập n:")
	fmt.Scanln(&n)

	m := make([]int, n)

	for i := 0; i < n; i++ {
		m[i] = rand.Intn(101)
	}

	fmt.Println("Mảng:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d\t", m[i])
	}
	fmt.Println()
}
