package main

import "fmt"

func main() {
	m := []int{1, -24, -44, 6, 5}
	selectionSort(m)
	fmt.Println(m)
}

func selectionSort(m []int) {
	var min int

	for i := 0; i < len(m)-1; i++ {
		min = i // giả sử số đầu trong dãy đang xét là nhỏ nhất

		for j := i + 1; j < len(m); j++ { // tìm số nhỏ nhất trong dãy đang xét
			if m[j] < m[min] {
				min = j
			}
		}

		if min != i { // nếu có số nào nhỏ hơn số đầu trong dãy đang xét thì thay đổi vị trí cho nhau
			var temp = m[i]
			m[i] = m[min]
			m[min] = temp
		}
	}
}

/*

Giải thuật chọn trực tiếp

Mảng có thứ tự tăng dần thì m[i]=min[m(i),m(i+1)...,m(n-1)]

Ý tưởng: mô phỏng một trong những cách sắp xếp tự nhiêu nhất trong thực tế
	Chọn phần tử nhỏ nh trong n phần tử ban đầu, đưa phần tử này về vị trí đúng
	là đầu dãy hiện hành
	Xem dãy hiện hành chỉ còn n-1 phần tử của dãy ban đầu, bắt đầu từ vị trí thứ 2
	lặp lại quá trình trên cho dãy hiện hành... đến khi dãy hiện hành chỉ còn 1 phần tử

ý tưởng xếp hàng từ thấp đến cao

Mả giả:

Bước 1: i=0
Bước 2: tìm phần tử m[min] nhỏ nhất trong dãy hiện hành từ m[i] đến m[n-1]
Bước 3: Nếu min#i, hoán vị m[min] và m[i]
Bước 4: nếu i<n-1 thì i++; lặp lại bước 2
Ngược lại: dừng // n-1 phần tử đã nằm đúng vị trí

*/
