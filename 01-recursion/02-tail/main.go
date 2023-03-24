package main

import "fmt"

func main() {
	m, n := 7, 3
	result := gcd(m, n)
	fmt.Println(result)
}

// tìm ước số chung lớn nhất
func gcd(m, n int) int {
	var r int

	if m < n {
		return gcd(n, m)
	}

	r = m % n

	if r == 0 {
		return n
	} else {
		return gcd(n, r)
	}
}

/*

Tail Recursion (Đệ qui đuôi)
	Gọi lại chính nó
	Có điểm dừng (Return 1 giá trị nào đó, mà stack rỗng, coi như kết thúc chương trình)
	Có sự điều hướng
	LIFO (Last In First Out), cơ chế Stack
	Khi 1 hàm gọi đệ qui, thì nó sẽ tự động lưu lại biến local và các dòng chỉ thị lệnh đằng sau hàm đệ qui đc gọi, nếu như đằng sau hàm đệ qui ko có dòng lệnh nào hết thì nó sẽ không lưu vào stack
	Khi kết thúc đệ qui sẽ gọi stack

*/
