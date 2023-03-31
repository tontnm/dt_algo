package main

import "fmt"

func main() {

	for {
		m := 81
		fmt.Println("Chương trình game đoán số")
		solandoan := 0
		win := false

		for {
			var n int
			fmt.Println("Máy đã đoán số [1-100], mời bạn đoán: ")
			fmt.Scan(&n)

			solandoan++
			fmt.Printf("Bạn đoàn lần thứ %d\n", solandoan)

			if n == m {
				fmt.Printf("HAHA bạn tài thật %d\n", m)
				win = true
				break // thắng và ngưng vòng lặp
			} else if n < m {
				fmt.Println("Bạn đoán sai số người < số máy")
			} else {
				fmt.Println("Bạn đoán sai số người > số máy")
			}

			if solandoan == 7 {
				break // thua thì ngưng
			}
		}

		if win != true {
			fmt.Println("GAME OVER")
			fmt.Printf("Số của máy %d\n", m)
		}

		var status string
		fmt.Println("Tiếp không thím?(c/k)")
		fmt.Scan(&status)

		if status == "k" {
			break // ngưng game ko chơi nữa
		}
	}
	fmt.Println("Goodbye")
}
