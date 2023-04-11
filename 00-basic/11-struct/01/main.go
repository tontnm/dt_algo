package main

import "fmt"

type SinhVien struct {
	MaSV     string
	TenSV    string
	GioiTinh bool
}

type DiaChi struct {
	SoNha     string
	TenDuong  string
	Quan      string
	TinhThanh string
}

func main() {
	teo := SinhVien{
		MaSV:     "1",
		TenSV:    "a",
		GioiTinh: false,
	}

	ty := SinhVien{
		MaSV:     "2",
		TenSV:    "b",
		GioiTinh: false,
	}

	fmt.Println(teo.TenSV, ty.TenSV)
}

/*

struct: kiểu dữ liệu phức hợp do ta tạo ra

*/
