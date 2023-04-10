package main

import "fmt"

func main() {
	m := []int{1, 2, 6, 5, 4, 3}

	for i := 0; i < len(m)-1; i++ {
		for j := i + 1; j < len(m); j++ {
			if m[i] > m[j] {
				t := m[i]
				m[i] = m[j]
				m[j] = t
			}
		}
	}

	fmt.Println(m)
}
