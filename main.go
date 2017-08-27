package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func generateRandomInts(n int) ([]int64, error) {
	ints := make([]int64, n)

	for i := range ints {
		randomInt, err := rand.Int(rand.Reader, big.NewInt(27))

		if err != nil {
			return nil, err
		}

		ints[i] = randomInt.Int64()
	}

	return ints, nil
}

func generateRandomSeed() (string, error) {
	const letters = "9ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ints, err := generateRandomInts(81)

	if err != nil {
		return "", err
	}

	token := make([]byte, 81)

	for i, x := range ints {
		token[i] = byte(letters[x])
	}

	return string(token), nil
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
