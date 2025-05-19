package main

import "fmt"

func main() {
	// Creating a map
	person := map[string]string{
		"name":    "John",
		"city":    "New York",
		"country": "USA",
	}

	fmt.Println("Person map:", person)

	// Accessing a value
	fmt.Println("Name:", person["name"])

	// Adding a new key-value pair
	person["job"] = "Software Developer"
	fmt.Println("After adding job:", person)

	// Updating a value
	person["city"] = "San Francisco"
	fmt.Println("After updating city:", person)

	// Deleting a key
	delete(person, "country")
	fmt.Println("After deleting country:", person)

	// Checking if a key exists
	if value, exists := person["job"]; exists {
		fmt.Println("Job exists on value:", value)
	} else {
		fmt.Println("Job key not exists")
	}

	// Looping over a map
	fmt.Println("\nLooping through map:")
	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}
}
