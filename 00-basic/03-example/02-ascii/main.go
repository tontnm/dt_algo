package main

import "fmt"

func main() {
	var ch rune

	fmt.Println("Nhập ký tự:")
	fmt.Scanf("%c", &ch)

	fmt.Printf("Mã ASCII của %c là %d\n", ch, ch)
	nextCh := ch + 1

	fmt.Printf("Ký tự ASCII tiếp theo của %c là %c", ch, nextCh)
}
