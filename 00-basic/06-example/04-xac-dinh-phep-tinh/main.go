package main

import "fmt"

func main() {
	var a, b int
	var op rune

	fmt.Print("Enter value for a: ")
	fmt.Scanln(&a)

	fmt.Print("Enter value for b: ")
	fmt.Scanln(&b)

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanf("%c\n", &op)

	var result int

	switch op {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		result = a / b
	default:
		fmt.Printf("Invalid operator: %c\n", op)
		return
	}

	fmt.Printf("%d %c %d = %d\n", a, op, b, result)
}
