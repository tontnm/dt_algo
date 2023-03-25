package main

import "fmt"

func main() {
	var giay int

	fmt.Println("Nhập giây:")
	fmt.Scan(&giay)

	result := calculateHMS(giay)
	fmt.Println(result)
}

func calculateHMS(t int) string {
	h := (t / 3600) % 24
	m := (t % 3600) / 60
	s := (t % 3600) % 60

	if h > 12 {
		h = h - 12
		return fmt.Sprintf("%d:%d:%d PM", h, m, s)
	}

	return fmt.Sprintf("%d:%d:%d AM", h, m, s)
}
