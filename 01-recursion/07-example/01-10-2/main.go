package main

import "fmt"

func main() {
	var n int = 5

	H10To2(n)
}

func H10To2(n int) {
	if n > 0 {
		sd := n % 2
		n = n / 2
		H10To2(n)
		fmt.Print(sd)
	}
}

/*

viết đệ qui chuyển cơ số 10 sang 2
- tìm quy luật : /2 và lấy số dư
- tìm điểm dừng: kết quả bằng 0 thì dừng

khi gọi đệ quy thì lưu stack
- lưu chỉ thị lệnh ngay các dòng đằng sau chỗ gọi đệ qui
- lưu giá trị các biến local

0>0-> false -> kết thúc đệ qui -> gọi stack
*/
