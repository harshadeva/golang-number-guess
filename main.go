package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	startNewGame()
}

// Generate a random number between range
func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func startNewGame() {
	var min, max int = 1, 100
	var target int64 = int64(randRange(min, max))
	attempts := 0

	fmt.Println("Welcome to number guessing game!")
	fmt.Println("Number is between %v and %v. Try to guess it in minimum attempts.", min, max)
	guessNumber(&attempts, target)
}

func guessNumber(attemptRef *int, target int64) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your guess: ")
	guessInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again.")
		guessNumber(attemptRef, target)
	}
	guessInput = strings.TrimSpace(guessInput)
	guess, err := strconv.ParseInt(guessInput, 10, 64)
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again.")
		guessNumber(attemptRef, target)
	}
	*attemptRef++

	if guess < target {
		fmt.Println("Too low! Try again.")
		guessNumber(attemptRef, target)
	} else if guess > target {
		fmt.Println("Too high! Try again.")
		guessNumber(attemptRef, target)
	} else {
		fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", *attemptRef)
		handleWin()
	}
}

func handleWin() {
	fmt.Println("Pres 'y' to start the game again, Press 'n' to exit the game: ")
	reader := bufio.NewReader(os.Stdin)
	userPrompt, _ := reader.ReadString('\n')
	userPrompt = strings.TrimSpace(userPrompt)
	switch userPrompt {
	case "y":
		startNewGame()
	case "n":
		fmt.Println("Thank you for playing!")
	default:
		fmt.Println("Invalid entry")
		handleWin()
	}

}
