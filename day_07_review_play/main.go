package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Review 1: Variables and types
func reviewVariables() {
	name := "Go Explorer"
	age := 30
	const pi = 3.1415

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Pi constant:", pi)
}

// Review 2: Input & Output
func reviewInputOutput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your favorite Go feature? ")
	feature, _ := reader.ReadString('\n')
	feature = strings.TrimSpace(feature)
	fmt.Println("Cool! You like:", feature)
}

// Review 3: Simple math operation
func reviewMath() {
	num1 := 15
	num2 := 7
	sum := num1 + num2
	diff := num1 - num2
	product := num1 * num2
	quotient := num1 / num2

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
}

// Fun Bonus: Convert string to integer and multiply
func bonusStringToInt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number to double: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number!")
		return
	}

	fmt.Printf("%d doubled is %d\n", num, num*2)
}

func main() {
	fmt.Println("=== Day 7: Review & Play ===")

	reviewVariables()
	reviewInputOutput()
	reviewMath()
	bonusStringToInt()
}
