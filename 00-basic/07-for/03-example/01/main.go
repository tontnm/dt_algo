package main

import "fmt"

func main() {
	var sumUoc int
	n := 6
	i := 1

	for i < n {
		if n%i == 0 { // i là số chia hết hay không
			sumUoc += i
		}
		i++
	}

	if sumUoc == n {
		fmt.Printf("%d là số hoàn thiện", n)
	} else {
		fmt.Printf("%d không phải số hoàn thiện", n)
	}
}
