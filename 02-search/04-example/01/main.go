package main

import "fmt"

func main() {
	m := []float64{9.0, 2.0, 3.0, 2.0, 9.0, 1.0, 9.0}

	result := demXTrongMang(m, len(m), 9.0)
	fmt.Println(result)
}

func demXTrongMang(m []float64, n int, x float64) int {
	d := 0

	for i := 0; i < n; i++ {
		if x == m[i] {
			d++
		}
	}

	return d
}

/*

tìm x xuất hiện bao nhiêu lần trong mảng

*/
