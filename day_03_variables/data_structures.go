package main

import "fmt"

func dataStructuresDemo() {
	// Integer
	var age int = 30
	fmt.Println("Age:", age)

	// Float
	var price float64 = 99.99
	fmt.Println("Price:", price)

	// String
	var name string = "Alice"
	fmt.Println("Name:", name)

	// Array
	var scores [3]int = [3]int{85, 90, 95}
	fmt.Println("Scores:", scores)

	// Slice
	var colors []string = []string{"Red", "Green", "Blue"}
	fmt.Println("Colors:", colors)

	// Map
	var personDetails map[string]string = map[string]string{"name": "Alice", "age": "25"}
	fmt.Println("Person Details:", personDetails)

	// Pointer
	var number int = 42
	var pointer *int = &number
	fmt.Println("Pointer:", pointer, "Value:", *pointer)

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	var person Person = Person{Name: "Alice", Age: 30}
	fmt.Println("Person Struct:", person)
}
