package main

import (
	"errors"
	"fmt"
)

// divide performs division and returns an error if dividing by zero
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// validateAge checks if age is valid, using fmt.Errorf for detailed error
func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("invalid age %d: must be non-negative", age)
	}
	if age < 18 {
		return errors.New("age must be at least 18")
	}
	return nil
}

func main() {
	// Example 1: Division with error handling
	fmt.Println(" == Division Example ==")
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}

	// Example 2: Validating age
	fmt.Println("\n == Age Validation Example ==")
	ages := []int{25, -1, 15}
	for _, age := range ages {
		err := validateAge(age)
		if err != nil {
			fmt.Printf("Age %d is invalid: %v\n", age, err)
		} else {
			fmt.Printf("Age %d is valid\n", age)
		}
	}

	// Example 3: Custom inline error
	fmt.Println("\n== Inline Error Example ==")
	var name string
	if name == "" {
		err := errors.New("name cannot be empty")
		fmt.Println("Error:", err)
	}
}
