package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("Please input your guess")
	var guess int
	for {
		_, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}

		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
