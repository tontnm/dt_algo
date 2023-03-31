package main

import "fmt"

func main() {
	m := []int{1, 2, 3, 4, 5}
	result1 := timSoChanLonNhatTrongMang(m, len(m))
	if result1 == -99999 {
		fmt.Println("Khong tìm thấy")
	} else {
		fmt.Printf("max chẵn là %d", result1)
	}

	result2 := timSoLeLonNhatTrongMang(m, len(m))
	if result2 == -99999 {
		fmt.Println("Khong tìm thấy")
	} else {
		fmt.Printf("max lẻ là %d", result2)
	}
}

func timSoChanLonNhatTrongMang(m []int, n int) int {
	max := -99999

	for i := 0; i < n; i++ {
		if m[i] > max && m[i]%2 == 0 {
			max = m[i]
		}
	}

	return max
}

func timSoLeLonNhatTrongMang(m []int, n int) int {
	max := -99999

	for i := 0; i < n; i++ {
		if m[i] > max && m[i]%2 != 0 {
			max = m[i]
		}
	}

	return max
}
