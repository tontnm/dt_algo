package main

import "fmt"

func main() {
	m := []int{10, 2, 5}
	printPermutations(m, 3, 0)
}

func printPermutations(arr []int, n, i int) {
	var j, swap int

	printArr(arr, n)

	for j = i + 1; j < n; j++ {
		if arr[i] > arr[j] {
			swap = arr[i]
			arr[i] = arr[j]
			arr[j] = swap
		}
		printPermutations(arr, n, i+1)
	}
}

func printArr(arr []int, n int) {
	for _, v := range arr {
		fmt.Print(fmt.Sprintf("%d\t", v))
	}
	fmt.Println()
}

/*

Exponential Recursion (Đệ qui đa tuyến)

1 hàm có 1 vòng lặp nào đó, trong vòng lặp gọi lại đệ qui của chính nó
Loại vừa vòng lặp vừa đệ qui
độ phức tạp O(a^n)

*/
