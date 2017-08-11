package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomSeed() (string, error) {
	const letters = "9ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes, err := generateRandomBytes(81)

	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes), nil
}

func main() {
	fmt.Println("Welcome to iota-seedgen - your friendly IOTA wallet seed generator!")
	fmt.Println("------------------")

	// generate 10 wallet seeds
	for i := 0; i < 10; i++ {
		token, err := generateRandomSeed()
		if err != nil {
			fmt.Printf("an error occured: %v\n	", err.Error())
			os.Exit(1)
		}

		fmt.Println(token)
	}

	fmt.Println("------------------")
	fmt.Println("pick one of the random lines, back it up anywhere >SAFE< and use it to generate your wallet :)")

	fmt.Println("")
	fmt.Println("Press the Enter Key to terminate the console screen!")
	var input string
	fmt.Scanln(&input)

	os.Exit(0)
}
