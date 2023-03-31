package main

import "fmt"

func main() {
	n := 123456
	var sum int

	for n > 0 {
		sd := n % 10
		sum += sd

		n = n / 10
	}

	fmt.Println(sum)
}
