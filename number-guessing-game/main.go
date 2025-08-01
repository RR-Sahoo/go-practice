package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	retry := "y"
	for retry == "y" {
		var difficulty int
		var difficultyString string
		var guess int
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(100) + 1
		attempts := 0
		maxAttempts := 0

		fmt.Println("Welcome to the Number Guessing Game!")
		fmt.Println("I'm thinking of a number between 1 and 100.")
		fmt.Println("Please select the difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&difficulty)
		fmt.Println()

		switch difficulty {
		case 1:
			difficultyString = "Easy"
			maxAttempts = 10
		case 2:
			difficultyString = "Medium"
			maxAttempts = 5
		case 3:
			difficultyString = "Hard"
			maxAttempts = 3
		default:
			fmt.Println("Invalid difficulty level. Please try again.")
			return
		}

		fmt.Printf("Great! You have selected the %s difficulty level.\n", difficultyString)
		fmt.Println("Let's start the game!")

		for attempts < maxAttempts {
			fmt.Printf("Attempt %d - Enter your guess: ", attempts+1)
			fmt.Scan(&guess)

			if guess < randomNumber {
				fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
			} else if guess > randomNumber {
				fmt.Printf("Incorrect! The number is less than %d.\n", guess)
			} else {
				fmt.Printf("🎉 Congratulations! You guessed the correct number in %d attempts.\n", attempts+1)
				break
			}
			attempts++
		}

		if guess != randomNumber {
			fmt.Printf("❌ Sorry, you've used all your attempts. The correct number was %d.\n", randomNumber)
		}

		fmt.Print("Would you like to play again? (y/n): ")
		fmt.Scan(&retry)
		fmt.Println()
	}
}
