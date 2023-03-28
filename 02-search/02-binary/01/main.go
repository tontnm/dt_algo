package main

import "fmt"

func main() {
	a := []int{1, 5, 15, 19, 25, 27, 29, 31, 33, 45, 55, 88, 100}
	result := binarySearch(a, len(a), 15)
	fmt.Println(result)
}

func binarySearch(s []int, n, x int) int {
	l, r := 0, n-1

	for l <= r {
		m := (l + r) / 2

		if x == s[m] { // thấy x ngay ở giữa
			return m
		} else if x < s[m] { // x nằm bên trái
			r = m - 1
		} else { // x nằm bên phải
			l = m + 1
		}
	}

	return -1 // duyệt hết mảng ko tìm thấy x
}

/*

giải thuật tìm kiếm nhị phân trên mảng cơ sở

đặc trưng dữ liệu: đã sắp xếp theo 1 thứ tự nào đó (nhỏ -> lớn)
tính chất của dữ liệu đc sắp là:
	m(i-1)<=m(i)<=m(i+1)
vậy nếu x > m(i) thì x chỉ có thể xuất hiện trong đoạn [m(i+1):m(n-1)] của dãy,
ngược lại nếu x<m(i) thì x có thể xuất hiện trong đoạn [m(0):m(i-1)] của dãy

Giải thuật tìm nhị phân áp dụng nhận xét trên đây với tư tưởng chia để trị để tìm
cách giới hạn phạm vi tìm kiếm sau mỗi lần so sánh x với một phần tử trong dãy.

Cài đặt giải thuật
Bước 1: left=0, right=n-1  // tìm kiếm trên tất cả các phần tử
Bước 2:
mid=(left+right)/2 // lấy mốc so sánh
So sánh m[mid] với k, có 3 khả năng:
m[mid]==k : tìm thấy, dừng
m[mid]>k  : tìm tiếp x trong dãy con m[left]...m[mid-1] : right = mid-1
m[mid]<k  : tìm tiếp x trong dãy con m[mid+1]...m[right]: left = mid + 1
Bước 3:
Nếu left<right: Còn phần tử chưa xét -> tìm tiếp
				Lặp lại bước 2
Ngược lại: dừng // đã xét hết tất cả các phần tử, tìm không thấy


đánh giá giải thuật
tốt nhất 1 phần tử giữa của mảng có giá trị x
xấu nhất log2n không có x trong mảng
trung bình log2n/2 giả sử xác suất các phần tử trong mảng nhận giá trị x là như nhau

giải thuật nhị phân có độ phức tạp tính toán cấp n:
T(N)=O(log2N)

giải thuật tìm nhị phân dựa vào quan hệ giá rị của các phần tử mảng để định hướng
trong quá trình tìm kiếm, do vậy chỉ áp dụng được cho những dãy đã có thứ tự
giải thuật tìm nhị phân tiết kiệm thời gian hơn rất nhiều so với giải thuật
tìm kiếm tuyến tính do
Tìm nhị phân (N)=O(log2N) < Tìm tuyến tính(N)=O(N)

muốn áp dụng tìm kiếm nhị phân, phải xét thời gian sắp xếp dãy số. thời gian ko nhỏ
và khi dãy số biến động cần phải tiến hành sắp xếp lại
cần cân nhắc nhu cầu thực tế để chọn một trong hai giải thuật tìm kiếm trên sao cho
có lợi nhất

thường xuyên biến động, ít tìm kiếm -> tuần tự
ít biến động, tần suất tìm kiếm nhiều -> nhị phân

*/
