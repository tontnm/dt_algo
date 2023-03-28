package main

import "fmt"

func main() {
	var n int

	fmt.Println("Tính tổng số nguyên từ 1 -> n, nếu tổng >=9 thì kết thúc:")
	fmt.Scan(&n)

	sum := 0

	for i := 1; i <= n; i++ {
		if sum >= 9 {
			break
		}
		sum += i
	}

	fmt.Println(sum)
}

/*

Lệnh break dùng để thoát 1 cấu trúc điều khiển mà ko chờ đến biểu thức
điều kiện đc định trị
Khi break đc thực hiện bên trong 1 cấu trúc lặp, điều khiển (control flow)
tự động nhảy đến lệnh đầu tiên ngay sau cấu trúc lặp đó
Không sử dụng break bên ngoài các cấu trúc lặp như while, for hay switch

*/
