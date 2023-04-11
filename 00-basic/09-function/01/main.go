package main

import "fmt"

func main() {
	a := 5
	ham1(a)
	fmt.Println("a ở hàm main ko bị ảnh hưởng bưởi ham1:", a)
	ham2(&a)
	fmt.Println("a ở hàm main bị ảnh hưởng bưởi ham2:", a)
}

func ham1(a int) {
	a = 10
	fmt.Println("a ở trong ham1:", a)
}

func ham2(a *int) {
	*a = 10
	fmt.Println("a ở trong ham2:", *a)
}

/*

hàm là 1 khối lệnh thực hiện 1 công việc hoàn chỉnh (module), đc đặt tên và
được gọi thực thi nhiều lần tại nhiều vị trí trong chương trình

tại sao phải viết hàm?

hàm có thể được gọi từ chương trình chính (hàm main) hoặc từ 1 hàm khác
hàm có giá trị trả về hoặc không

có 2 lại hàm:
Hàm thư viện : là hàm do các nhà phát triển làm sẵn, chỉ cần gọi ra là dùng
Hàm do người dùng định nghĩa

nguyên tắc hoạt động của hàm
LIFO là nguyên tắc hoạt động của hàm

khi 1 hàm đc gọi, chương trình sẽ lưu lại giá trị các biến local và chỉ thị lệnh tiếp theo.
Khi kết thúc quá trình gọi hàm, chương trình sẽ trả về đúng chỉ thị lệnh kế tiếp đó cùng
với giá trị của các biến local

truyền tham trị và tham biến

truyền tham trị: sau khi thoát khỏi hàm nó vẫn giữ giá trị gốc
truyền tham biến: sau khi thoát khỏi hàm, nó sẽ lấy giá trị thay đổi trong hàm

truyền tham trị (call by value)
sao chép giá trị của đối số vào tham số hình thức của hàm
những thay đổi của tham số không ảnh hưởng đến giá trị của đối số

truyền tham chiếu (call by reference)
Sao chép địa chỉ của đối số vào tham số hình thức. do đó, những thay đổi đối với tham số
sẽ có tác dụng trên đối số

*/
