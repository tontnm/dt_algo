package main

import "fmt"

func main() {
	n := 3
	move(n, 'A', 'B', 'C')
}

func move(n int, a, b, c rune) {
	if n == 1 {
		fmt.Printf("%c -> %c\n", a, c)
	} else {
		move(n-1, a, c, b)
		fmt.Printf("%c -> %c\n", a, c)
		move(n-1, b, a, c)
	}
}
