package main

import (
	"fmt"

	"me/test/calculator"
)

func main() {
	var input int
	fmt.Println("Basic Calculator")
	fmt.Println("-----Options-----")
	fmt.Println("     1. Add      ")
	fmt.Println("   2. Subtract   ")
	fmt.Println("   3. Multiply   ")
	fmt.Println("    4. Divide    ")
	fmt.Println("Choose the equation")
	fmt.Scanln(&input)
	if input == 1 {
		var A, B float32
		fmt.Println("Type the first number: ")
		fmt.Scanln(&A)
		fmt.Println("Type the second number: ")
		fmt.Scanln(&B)
		fmt.Printf("Your answer is: %v", calculator.Add(A, B))
	}
	if input == 2 {
		var A, B float32
		fmt.Println("Type the first number: ")
		fmt.Scanln(&A)
		fmt.Println("Type the second number: ")
		fmt.Scanln(&B)
		fmt.Printf("Your answer is: %v", calculator.Subtract(A, B))
	}
	if input == 3 {
		var A, B float32
		fmt.Println("Type the first number: ")
		fmt.Scanln(&A)
		fmt.Println("Type the second number: ")
		fmt.Scanln(&B)
		fmt.Printf("Your answer is: %v", calculator.Multiply(A, B))
	}
	if input == 4 {
		var A, B float32
		fmt.Println("Type the first number: ")
		fmt.Scanln(&A)
		fmt.Println("Type the second number: ")
		fmt.Scanln(&B)
		fmt.Printf("Your answer is: %v", calculator.Divide(A, B))
	}
}
