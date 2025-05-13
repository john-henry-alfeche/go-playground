package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Example 1: Simple input using fmt.Scan
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello,", name)

	// Example 2: Read full line input using bufio.NewReader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your favorite color: ")
	color, _ := reader.ReadString('\n')
	color = strings.TrimSpace(color)
	fmt.Println("Your favorite color is:", color)

	// Example 3: Read integer input
	fmt.Print("Enter your age: ")
	var ageInput string
	ageInput, _ = reader.ReadString('\n')
	ageInput = strings.TrimSpace(ageInput)
	age, err := strconv.Atoi(ageInput)
	if err != nil {
		fmt.Println("Invalid age input")
	} else {
		fmt.Println("You are", age, "years old")
	}

	// Example 4: Read two numbers and add them
	fmt.Print("Enter first number: ")
	firstStr, _ := reader.ReadString('\n')
	firstStr = strings.TrimSpace(firstStr)
	firstNum, err1 := strconv.ParseFloat(firstStr, 64)

	fmt.Print("Enter second number: ")
	secondStr, _ := reader.ReadString('\n')
	secondStr = strings.TrimSpace(secondStr)
	secondNum, err2 := strconv.ParseFloat(secondStr, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid number input")
	} else {
		sum := firstNum + secondNum
		fmt.Printf("The sum of %.2f and %.2f is %.2f\n", firstNum, secondNum, sum)
	}
}
