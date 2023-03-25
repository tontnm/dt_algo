package main

import (
	"fmt"
)

const Pi = 3.14

func main() {

	var r float64

	fmt.Println("Nhập r")
	fmt.Scan(&r)

	cv := 2 * Pi * r
	dt := Pi * r

	fmt.Println(fmt.Sprintf("cv = %f, dt = %f", cv, dt))
}
