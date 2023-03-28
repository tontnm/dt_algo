package main

import "fmt"

func main() {
	var so, hangChuc, hangDv int
	var chuoiChuc, chuoiDv, result string

	fmt.Println("Nhập số tối đa 2 chữ số:")
	fmt.Scanln(&so)

	if so < 0 || so > 99 {
		fmt.Println("số không hợp lệ")
	} else if so == 0 {
		fmt.Println("không")
	} else {
		hangChuc = so / 10
		hangDv = so % 10

		switch hangChuc {
		case 0:
			chuoiChuc = ""
		case 1:
			chuoiChuc = "mười"
		case 2:
			chuoiChuc = "hai mươi"
		case 3:
			chuoiChuc = "ba mươi"
		case 4:
			chuoiChuc = "bốn mươi"
		case 5:
			chuoiChuc = "năm mươi"
		case 6:
			chuoiChuc = "sáu mươi"
		case 7:
			chuoiChuc = "bảy mươi"
		case 8:
			chuoiChuc = "tám mươi"
		case 9:
			chuoiChuc = "chín mươi"
		}

		if hangDv == 0 && hangChuc > 1 {
			chuoiDv = ""
		} else if hangDv == 0 && hangChuc == 1 {
			chuoiDv = ""
		} else if hangDv == 1 {
			chuoiDv = "một"
		} else if hangDv == 2 {
			chuoiDv = "hai"
		} else if hangDv == 3 {
			chuoiDv = "ba"
		} else if hangDv == 4 {
			chuoiDv = "bốn"
		} else if hangDv == 5 {
			chuoiDv = "năm"
		} else if hangDv == 6 {
			chuoiDv = "sáu"
		} else if hangDv == 7 {
			chuoiDv = "bảy"
		} else if hangDv == 8 {
			chuoiDv = "tám"
		} else if hangDv == 9 {
			chuoiDv = "chín"
		}

		result = chuoiChuc + " " + chuoiDv
	}

	fmt.Printf(result)
}
