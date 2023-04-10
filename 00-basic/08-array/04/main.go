package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var r, c int

	fmt.Println("nhập số hàng và cột:")
	fmt.Scanln(&r, &c)

	slice := make([][]int, r)
	for i := range slice {
		slice[i] = make([]int, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			slice[i][j] = rand.Intn(101)
		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf("%d\t", slice[i][j])
		}
		fmt.Println()
	}
}
