package main

import "fmt"

func main() {
	// -------- Number Types --------
	// Integer
	var age int = 30
	// Float
	var price float64 = 99.99

	fmt.Println("Number Types:")
	fmt.Println("Age:", age)
	fmt.Println("Price:", price)
	fmt.Println()

	// -------- String Operations --------
	var greeting string = "Hello"
	var name string = "John"
	full := greeting + ", " + name + "!" // String concatenation

	fmt.Println("String Operations:")
	fmt.Println(full)
	fmt.Println()

	// -------- Boolean Values --------
	var isLearning bool = true
	var isSleeping bool = false

	fmt.Println("Boolean Values:")

	fmt.Println("Is learning:", isLearning)
	fmt.Println("Is sleeping:", isSleeping)
	fmt.Println()

	// -------- Math Operations --------
	a := 10
	b := 3

	fmt.Println("Math Operations:")
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b) // Integer division
	fmt.Println("Modulus:", a%b)  // Remainder
	fmt.Println()

	// -------- Comparison Operators --------
	x := 10
	y := 20

	fmt.Println("Comparison Operators:")
	fmt.Println("x == y:", x == y)
	fmt.Println("x != y:", x != y)
	fmt.Println("x > y:", x > y)
	fmt.Println("x < y:", x < y)
	fmt.Println("x >= y:", x >= y)
	fmt.Println("x <= y:", x <= y)
}
