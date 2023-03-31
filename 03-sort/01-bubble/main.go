package main

import "fmt"

func main() {
	m := []int{9, 7, 1, 2, 0, 5, 3, 4, 1}
	bubbleSort(m)
	fmt.Println(m)
}

func bubbleSort(m []int) {

	for i := 0; i < len(m)-1; i++ {
		for j := len(m) - 1; j > i; j-- {
			if m[j] < m[j-1] {
				var temp = m[j]
				m[j] = m[j-1]
				m[j-1] = temp
			}
		}
	}
}

/*

Bubble Sort (O(n2))

ý tưởng chính:
Xuất phát từ cuối (đầu) dãy, đổi chỗ các cặp phần tử kế cận để đưa phần tử nhỏ (lớn)
hơn trong cặp phần tử đó về vị trí đứng đầu (cuối) dãy hiện hành, sau đó sẽ không
xét đến nó ở bước tiếp theo
ở lần xử lý thứ i có vị trí đầu dãy là i
lặp lại xử lý trên cho đến khi không còn cặp phần tử nào để xét

ý tưởng: nặng thì chìm dưới đáy, nhẹ thì nổi lên

Mã giả
Bước 1: i=0 // lần xử lý đầu tiên
Bước 2: j=n-1 // duyệt từ cuối dãy ngược về vị trí i
	Trong khi j>i thực hiện:
		nếu m[j]<m[j-1]: m[j] <-> m[j-1] // xét cặp phần tử kể cận
		j=j-1
Bước 3: i=i+1 // lần xử lý kế tiếp
	Nếu i>=n-1 hết dãy. dừng
	ngược lại: lặp lại bước 2

Khuyết điểm:
không nhận diện được tình trạng dãy đã có thứ tự hay có thứ tự từng phần
các phần tử nhỏ được đưa về vị trí đúng rất nhanh, trong khi các phần
tử lớn lại được đưa về vị trí đúng rất chậm
Yêu cầu: cái tiến giải thuật bubble sort

Tại sao phải sắp xếp dữ liệu

Sắp xếp là quá trình xử lý một danh sách các phần tử (hoặc các mẫu tin) để đặt chúng
theo một thứ tự thỏa mãn một số tiêu chuẩn nào đó dựa trên nội dung thông tin lưu trữ
tại mỗi phần tử

VD:
sắp xếp mảng số nguyên tăng dần
sắp xếp sinh viên theo tuổi tăng dần
sắp xếp sinh viên theo tuổi tăng dần, nếu tuổi bằng nhau thì sắp xếp theo mã giảm dần
sắp xếp sản phẩm theo đơn giá giảm dần, nếu đơn giá bằng nhau thì sắp xếp theo ngày
nhập tăng dần

Khái niệm nghịch thế:
xét một danh sách các số m(0) m(1) ... m(n-1)
nếu có i<j và m(i) > m(j) thì ta gọi đó là 1 nghịch thế
danh sách chưa sắp xếp sẽ có nghịch thế
danh sách đã có thứ tự sẽ không chứa nghịch thế
m(0) <= m(1) <= ... <= m(n-1)
khi nào có nghịch thế thì cần sắp xếp

kiểu dữ liệu cơ sở, kiểu dữ liệu cấu trúc (struct, class) sẽ có tư tưởng sắp xếp
khác nhau để đáp ứng nhu cầu thực tế
kiểu cơ sở: sắp xếp trực tiếp trên giá trị của các phần tử, tăng dần hoặc giảm d
kiểu cấu trúc: sắp xếp theo từng tiêu chí của đối tượng, vì 1 đối tượng có
nhiều thuộc tính nên cần sắp xếp theo tiêu chí nào
tùy vào mục đích mà ta sử dụng các giải thuật khác nhau

*/
