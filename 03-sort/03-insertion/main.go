package main

import "fmt"

func main() {
	m := []int{84, 20, 71, 61, 66}
	insertionSort(m)
	fmt.Println(m)
}

func insertionSort(m []int) {
	// lưu vị trí cần chèn pos
	// lưu trữ giá trị x = m[i] tránh bị ghi đè khi dời chỗ các phần tử
	var pos, x int

	// xem đoạn m(0) đã sắp xếp
	for i := 1; i < len(m); i++ {
		x = m[i]

		for pos = i; (pos > 0) && (m[pos-1] > x); pos-- {
			m[pos] = m[pos-1]
		}
		m[pos] = x // chèn x vào dãy
	}

}

/*

giải thuật chèn trực tiếp

xét dãy n phần tử: m(0), m(1), ... , m(n-1)
xem dãy gồm 1 phần tử là m(0) dãy có thứ tự
thêm m(1) vào dãy có thứ tự m(0) sao cho dãy mới m(0),m(1) là dãy có thứ tự. nếu m[1]<m[0] ta đổi chỗ
m(1) với m(0)
thêm m(2) vào dãy có thứ tự m(0),m(1) sao cho dãy mới m(0),m(1),m(2) là dãy có thứ tự
tiếp tực như thế đến n-1 bước ta sẽ có dãy có thứ tự m(0),m(1),...m(n-1)

xếp hàng theo thứ tự thấp đến cao rồi, người được chèn vào phải tự quan sát để đi vào, và khi họ đi vào
phải giữ được trật tự từ thấp đến cao

mả giã

Bước 1: i=1 // giả sử m(0) đã đc sắp xếp
Bước 2: x=m[i] tìm vị trí pos thích hợp trong đoạn m[0] đến m[i-1] để chèn m(i) vào
Bước 3: dời chỗ các phần tử từ m[pos] đến m[i-1] sang phải 1 vị trí để dành chỗ cho m[i]
Bước 4: m[pos]=x // có đoạn m(0)..m(i-1) đã đc sắp xếp
Bước 5: i++
	nếu i<n: lặp lại bước 2
	ngược lại: dừng

*/
