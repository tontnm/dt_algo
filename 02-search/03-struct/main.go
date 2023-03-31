package main

import "fmt"

type SanPham struct {
	MSP string
	Ten string
	Gia float64
}

func main() {
	sanPham := []SanPham{
		{
			MSP: "1",
			Ten: "t",
			Gia: 10,
		},
		{
			MSP: "2",
			Ten: "y",
			Gia: 20,
		},
		{
			MSP: "3",
			Ten: "z",
			Gia: 30,
		},
		{
			MSP: "4",
			Ten: "a",
			Gia: 40,
		},
		{
			MSP: "5",
			Ten: "k",
			Gia: 50,
		},
	}
	gia := 40.0

	sp := findSPWithGia(sanPham, gia)

	if sp == nil {
		fmt.Println("Không tìm thấy sản phẩm")
	} else {
		fmt.Printf("%s-%s-%f", sp.MSP, sp.Ten, sp.Gia)
	}

	fmt.Println(sp)
}

func findSPWithGia(spSlice []SanPham, gia float64) *SanPham {
	l, r := 0, len(spSlice)-1

	for l <= r {
		m := (l + r) / 2

		sp := spSlice[m]

		if sp.Gia == gia {
			return &sp
		} else if gia < sp.Gia {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return nil
}
