package main

import "fmt"

func main() {
	var month int

	fmt.Print("Enter a month (1-12): ")
	fmt.Scanln(&month)

	days := 0

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		var y int
		fmt.Println("Nhập năm")
		fmt.Scanln(&y)

		if y%4 == 0 && y%100 != 0 {
			days = 29
		} else {
			days = 28
		}
	default:
		fmt.Println("Invalid input.")
		return
	}

	fmt.Printf("The month has %d days.\n", days)
}
