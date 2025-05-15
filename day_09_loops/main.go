package main

import "fmt"

func main() {
	// Basic for loop
	fmt.Println("Basic for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println("Count:", i)
	}

	// For loop like a while loop
	fmt.Println("\nWhile-style loop:")
	x := 1
	for x < 5 {
		fmt.Println("x is", x)
		x++
	}

	// Infinite loop with break
	fmt.Println("\nInfinite loop with break:")
	y := 1
	for {
		fmt.Println("y =", y)
		if y == 3 {
			break
		}
		y++
	}

	// Using continue
	fmt.Println("\nUsing continue:")
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println("Odd number:", i)
	}

	// Loop over slice
	fmt.Println("\nLooping through a slice:")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("Index %d: %s\n", index, fruit)
	}

	// Loop over map
	fmt.Println("\nLooping through a map:")
	person := map[string]string{"name": "Alice", "city": "Paris"}
	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}
}
