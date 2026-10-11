package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(101)
	var guessNumber int
	for {
		fmt.Print("Enter your guess (0-100): ")
		fmt.Scanf("%d", &guessNumber)
		if guessNumber == randomNumber {
			fmt.Println("Congratulations! You guessed the number.")
			break
		} else if guessNumber < randomNumber {
			fmt.Println("Too low. Try again.")
		} else {
			fmt.Println("Too high. Try again.")
		}
	}
}
