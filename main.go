package main

import (
	"fmt"
)

func main() {
	var first, second float64
	var op string

	fmt.Println("Welcome! Please enter first number.")
	fmt.Print("> ")
	fmt.Scanln(&first)
	fmt.Println("Please enter second number.")
	fmt.Print("> ")
	fmt.Scanln(&second)
	fmt.Println("Please enter operator.")
	fmt.Print("> ")
	fmt.Scanln(&op)

	if op == "+" {
		fmt.Println(first + second)
	} else if op == "-" {
		fmt.Println(first - second)
	} else if op == "*" {
		fmt.Println(first * second)
	} else if op == "/" {
		fmt.Println(first / second)
	} else {
		fmt.Println("Invalid operator.")
	}
}
