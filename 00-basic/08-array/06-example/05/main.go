package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

	var countChan, countLe int

	fmt.Println("Số chẳn trong mảng:")
	for i := 0; i < n; i++ {
		if m[i]%2 == 0 {
			fmt.Printf("%d\t", m[i])
			countChan++
		}
	}
	fmt.Println()
	fmt.Printf("Có %d số chẵn", countChan)
	fmt.Println()

	fmt.Println("Số lẻ trong mảng:")
	for i := 0; i < n; i++ {
		if m[i]%2 != 0 {
			fmt.Printf("%d\t", m[i])
			countLe++
		}
	}
	fmt.Println()
	fmt.Printf("Có %d số lẻ", countLe)
	fmt.Println()

	fmt.Println("Số nguyên tố trong mảng:")
	for i := 0; i < n; i++ {
		var dem int
		for j := 1; j <= m[i]; j++ {
			if m[i]%j == 0 {
				dem++
			}
		}

		if dem == 2 {
			fmt.Printf("%d\t", m[i])
		}
	}
}
