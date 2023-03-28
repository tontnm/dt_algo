package main

import "fmt"

type SinhVien struct {
	Ma       string
	Ten      string
	GioiTinh string
}

func main() {
	sinhVien := []SinhVien{
		{
			Ma:       "MSV01",
			Ten:      "Nam",
			GioiTinh: "M",
		},
		{
			Ma:       "MSV02",
			Ten:      "Nu",
			GioiTinh: "F",
		},
		{
			Ma:       "MSV03",
			Ten:      "Nem",
			GioiTinh: "M",
		},
	}

	result := timSVTheoMaSV(sinhVien, "MSV03")

	fmt.Println(result)
}

func timSVTheoMaSV(sv []SinhVien, maSV string) []SinhVien {
	var result []SinhVien

	for _, v := range sv {
		if v.Ma == maSV {
			result = append(result, v)
		}
	}

	return result
}

/*

giải thuật tìm kiems tuyến tính trên mảng cấu trúc

*/
