package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Exercise 1: Print your name
func printName() {
	fmt.Println("My name is John Henry Alfeche")
}

// Exercise 2: Add two numbers entered by the user
func addTwoNumbers() {
	reader := bufio.NewReader(os.Stdin)

	// Read first number
	fmt.Print("Enter the first number: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	num1, err1 := strconv.Atoi(input1)
	if err1 != nil {
		fmt.Println("Invalid input for first number")
		return
	}

	// Read second number
	fmt.Print("Enter the second number: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	num2, err2 := strconv.Atoi(input2)
	if err2 != nil {
		fmt.Println("Invalid input for second number")
		return
	}

	// Calculate and display the sum
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)
}

func main() {
	fmt.Println("=== Day 6: Mini Practice ===")

	// Exercise 1
	printName()

	// Exercise 2
	addTwoNumbers()
}
