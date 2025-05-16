package main

import (
	"fmt"
)

// Basic function
func greet(name string) {
	fmt.Println("Hello,", name)
}

// Function with return value
func add(a int, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Variadic function (takes any number of ints)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// Call basic function
	greet("Go Explorer")

	// Call function with return value
	sumResult := add(10, 5)
	fmt.Println("Sum:", sumResult)

	// Call function with multiple return values
	q, r := divide(17, 4)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)

	// Call variadic function
	totalSum := sum(1, 2, 3, 4, 5)
	fmt.Println("Total Sum:", totalSum)
}
