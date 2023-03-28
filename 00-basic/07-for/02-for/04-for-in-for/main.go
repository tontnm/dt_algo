package main

import "fmt"

func main() {
	var h int

	fmt.Println("Nhập chiều cao:")
	fmt.Scan(&h)

	for i := 0; i < h; i++ {
		for j := 0; j < h; j++ {
			if j == 0 || i == j || i == h-1 {
				fmt.Print("*\t")
			} else {
				fmt.Print(" \t")
			}
		}
		fmt.Println()
	}
}

/*

dùng ma trận vuông để giải bài này

		j=0		j=1		j=2		j=3		j=4
i=0		*
i=1		*		*
i=2		*				*
i=3		*						*
i=4		*		*		*		*		*

*/
