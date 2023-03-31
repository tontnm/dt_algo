package main

import "fmt"

type Book struct {
	Id   int
	Name string
	Year int
}

func main() {
	myBook := []Book{
		{
			Id:   1,
			Name: "1",
			Year: 1,
		},
		{
			Id:   2,
			Name: "2",
			Year: 3,
		},
		{
			Id:   3,
			Name: "3",
			Year: 3,
		},
	}

	result, count := printBook(myBook, 0)

	if count == 0 {
		fmt.Println("Không có sách xuất bản năm này")
	} else {
		fmt.Println(result)
	}
}

func printBook(m []Book, year int) ([]Book, int) {
	var foundBook []Book

	var count int

	for _, book := range m {
		if book.Year == year {
			foundBook = append(foundBook, book)
			count++
		}
	}

	return foundBook, count
}
