package main

import "fmt"

func main() {
	var i, n, sum int

	fmt.Println("Tính tổng từ 1 -> n. Nhập n")
	fmt.Scan(&n)

	i = 1

	for i <= n {
		sum += i
		i++
	}

	fmt.Printf("Tổng là %d", sum)
}

/*

while
b1: expresssion đc định trị
b2: nếu kết quả là true thì statement thực thi và quay lại b1
b3: nếu kết quả là false thì thoát khỏi vòng lặp

dùng để giải quyết các bài toán với số lượng lớn ko dùng tay được
phải xác định quy luật để lặp và dừng

do...while
b1:statement đc thực hiện
b2:expression đc định trị
b3:true quay lại bươc 1
b4:false thoát vòng lặp

*/
