package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to the Go Calculator!")

	fmt.Print("Please enter the first number: ")
	num1, err := strconv.ParseFloat(getInput(), 64)
	if err != nil {
		fmt.Println("Invalid number entered")
		os.Exit(1)
	}

	fmt.Print("Please enter the second number: ")
	num2, err := strconv.ParseFloat(getInput(), 64)
	if err != nil {
		fmt.Println("Invalid number entered")
		os.Exit(1)
	}

	fmt.Print("Please choose an operation (+, -, *, /): ")
	op := getInput()

	var result float64
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Invalid operator entered")
		os.Exit(1)
	}

	fmt.Printf("Result: %f\n", result)
}

func getInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}
