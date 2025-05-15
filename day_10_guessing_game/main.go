package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // safer & modern seeding

	secret := r.Intn(10) + 1 // Random number between 1 and 10
	var guess int

	fmt.Println("Guess a number between 1 and 10:")

	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanln(&guess)

		if err != nil {
			return
		}

		if guess < secret {
			fmt.Println("Too low!")
		} else if guess > secret {
			fmt.Println("Too high!")
		} else {
			fmt.Println("You got it!")
			break
		}
	}
}
