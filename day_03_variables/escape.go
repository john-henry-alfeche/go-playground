package main

import "fmt"

func escapeSequenceDemo() {
	fmt.Printf("Hello\nGolang\n")    // \n Newline
	fmt.Printf("Hello\tGolang\n")    // \t Tab
	fmt.Printf("Hello \"Golang\"\n") // \" Quote
	fmt.Printf("Hello \\ Golang\n")  // \\ Backslash
}
