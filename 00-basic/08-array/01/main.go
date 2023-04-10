package main

import "fmt"

func main() {
	m := make([]int, 5)

	for i := 0; i < len(m); i++ {
		fmt.Printf("Nhập m[%d]=", i)
		fmt.Scan(&m[i])
	}

	fmt.Println("Xuất mảng chiều xuôi")
	for i := 0; i < len(m); i++ {
		fmt.Printf("m[%d]=%d\t", i, m[i])
	}
	fmt.Println()

	fmt.Println("Xuất mảng chiều ngược")
	for i := len(m) - 1; i >= 0; i-- {
		fmt.Printf("m[%d]=%d\t", i, m[i])
	}
}

/*

Khái niệm
Mảng là một tập hợp các biến có cùng kiểu dữ liệu nằm liên tiếp nhau trong bộ nhớ
và được tham chiếu bởi 1 tên chung chính là tên mảng
Tại sao dùng mảng? khai báo 100 biến kiểu float
Mỗi phần tử của mảng được tham chiếu thông qua chỉ mục index
nếu mảng có n phần tử thì phần tử đầu tiên có chỉ mục là 0 và phần tử cuối có chỉ mục n-1
để tham chiếu đến một phần tử ta dùng tên mảng và chỉ mục của phần tử được đặt trong cắp dấu []
mỗi phần tử manggr là một biến thông thường

*/
