package main

import "fmt"

func main() {
	n := 4
	result := factorial(n)
	fmt.Println(fmt.Sprintf("%d!=%d", n, result))
}

// linear recursion (đệ qui tuyến tính: hàm tự gọi đến nó 1 lần và giảm theo cấp bậc)
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

/*

Linear Recursion (Đệ qui tuyến tính: hàm gọi đến nó 1 lần thôi và giảm theo cấp bậc)
	Hàm đệ qui là hàm gọi lại chính nó
	Phải tìm được điểm dừng cho đệ quy (Return 1 giá trị nào đó, mà stack rỗng, coi như kết thúc chương trình)
	Cơ chế hoạt động LIFO (Last In First Out), cơ chế Stack
	Khi 1 hàm gọi đệ qui, thì nó sẽ tự động lưu lại biến local và các dòng chỉ thị lệnh đằng sau hàm đệ qui đc gọi, nếu như đằng sau hàm đệ qui ko có dòng lệnh nào hết thì nó sẽ không lưu vào stack
	Khi kết thúc đệ qui sẽ gọi stack

*/
