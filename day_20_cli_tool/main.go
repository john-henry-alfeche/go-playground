package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var todos []string

func main() {
	fmt.Println("Simple To-Do Command Line Tool")
	fmt.Println("Type 'add <task>', 'list', or 'exit'")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\n")
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		input := scanner.Text()
		command := strings.Fields(input)

		if len(command) == 0 {
			continue
		}

		switch command[0] {
		case "add":
			if len(command) < 2 {
				fmt.Println("Please provide a task.")
				continue
			}

			task := strings.Join(command[1:], " ")
			todos = append(todos, task)
			fmt.Println("Task Added!")

		case "list":
			if len(command) == 0 {
				fmt.Println("No tasks yet.")
			} else {
				fmt.Println("Your tasks:")
				for i, task := range todos {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}

		case "exit":
			fmt.Println("Exiting. Have a productive day!")
			return

		default:
			fmt.Println("Unknown command. Use 'add', 'list', or 'exit'")
		}
	}
}
