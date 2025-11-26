package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const maxGuesses int = 5

func main() {
	fmt.Println("Welcome to guess the number!")
	fmt.Printf("I have chosen a number between 0 and 100 (including)\nCan you guess it within %v tries?\n", maxGuesses)

	num := rand.Intn(101)
	guessesLeft := maxGuesses
	for guessesLeft > 0 {
		fmt.Printf("You have %v guesses left\n", guessesLeft)
		fmt.Print("What is your guess?: ")
		var guess string

		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Something went wrong!")
			fmt.Println(err.Error())
			continue
		}

		guessInt, err := strconv.Atoi(guess)
		if err != nil {
			fmt.Println("That is not a number!")
			continue
		}

		var hint string

		diff := num - guessInt

		if diff > 0 {
			hint = "That's too low"
			if diff < 10 {
				hint += " (Close though...)"
			}
		} else if diff < 0 {
			hint = "That's too high"
			if diff > -10 {
				hint += " (Close though...)"
			}
		} else {
			fmt.Printf("You guessed it! It took you %v tries", maxGuesses-guessesLeft)
			return
		}

		fmt.Println(hint)
		guessesLeft--
	}

	fmt.Printf("Too bad... You couldn't guess it in %v tries.\n", maxGuesses)
	fmt.Printf("The number was %v", num)
}
