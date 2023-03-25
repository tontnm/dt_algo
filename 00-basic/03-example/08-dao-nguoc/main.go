package main

import "fmt"

func main() {
	var input int

	fmt.Println("Nhập 1 số nguyên dương có 3 số:")
	fmt.Scan(&input)

	f := input / 100
	l := input % 10
	m := (input % 100) / 10

	fmt.Printf("số đảo ngược là: %d%d%d", l, m, f)
}
