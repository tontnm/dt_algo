package main

import "fmt"

func main() {
	var day, month, year int

	fmt.Println("Nhấp ngày tháng năm theo định dạng dd/mm/yyyy")
	fmt.Scanf("%d/%d/%d", &day, &month, &year)

	isLeapYear := (year%4 == 0 && year%100 != 0) || (year%400 == 0)

	maxDaysInMonth := 31

	switch month {
	case 2:
		if isLeapYear {
			maxDaysInMonth = 29
		} else {
			maxDaysInMonth = 28
		}
	case 4, 6, 9, 11:
		maxDaysInMonth = 30
	}

	if day > maxDaysInMonth {
		fmt.Println("ngày không hợp lệ")
	} else if day < maxDaysInMonth {
		day++
	} else {
		day = 1
		if month == 12 {
			month = 1
			year++
		} else if month == 2 && day == 29 {
			day = 1
			month = 3
		} else {
			month++
		}
	}

	nextDay := fmt.Sprintf("%02d/%02d/%d", day, month, year)
	fmt.Println("Ngày kế tiếp là: ", nextDay)

}
