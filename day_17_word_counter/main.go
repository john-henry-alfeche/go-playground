package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Prompt the user for input
	fmt.Print("Enter a sentence: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Clean and split input
	input = strings.TrimSpace(input)
	words := strings.Fields(input)

	// Count word frequencies
	wordCount := make(map[string]int)
	for _, word := range words {
		cleaned := strings.ToLower(strings.Trim(word, ".,!?:;"))
		wordCount[cleaned]++
	}

	// Display results
	fmt.Println("\nWord Frequencies:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
