package main

import "fmt"

func variableDemo() {
	// Short variable declaration
	message := "Hello"
	fmt.Println(message)

	// Full variable declaration with type
	var age int = 25
	fmt.Println(age)

	// Type inference
	var score = 95.5
	fmt.Println(score)

	// Constant
	const Pi = 3.14
	fmt.Println(Pi)
}
