package main

import "fmt"

func main() {
	m := []int{5, 1, 3, 4, 2}
	result := lSearch(m, 5, 3)
	fmt.Println(result)
}

func lSearch(m []int, n, x int) int {
	for i := 0; i < n; i++ {
		if m[i] == x {
			return i
		}
	}
	return -1
}

/*

giải thuật tìm kiếm tuyến tính trên mảng cơ sở, tức là chạy từ đầu đến cuối ko bỏ sót
phần tử nào cả

Mã giả
bước 1: i=0 =>bắt đầu từ phần tử đầu tiên
bước 2: so sánh m[i] với x, có 2 khả năng:
	+ m[i] == x: tìm thấy, dừng
	+ m[i] # x: sang bước 3
bước 3: i=i+1 => xét tiếp phần tử kế trong mảng
	nếu i>=n: hết mảng, ko tìm thấy, dừng
	ngược lại: lặp lại bước 2

tìm kiếm để làm gì?

ta muốn lọc một số phần tử "có những đặc điểm nào đó" từ một tập dữ liệu đã có trước
để xử lý

tập dữ liệu có trước
1.dữ liệu bên trong để không theo thứ tự ->tìm kiếm tuần tự
2.dữ liệu bên trong có để tho thứ tự -> tìm kiểm tuần tự, tìm kiếm nhị phân

*/
