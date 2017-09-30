package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

var version = "undefined"

const letters = "9ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func intToCharByte(i int64) byte {
	return byte(letters[i])
}

func generateRandomSeed() (string, error) {
	ints, err := generateRandomInts(81)

	if err != nil {
		return "", err
	}

	token := make([]byte, 81)

	for i, x := range ints {
		token[i] = intToCharByte(x)
	}

	return string(token), nil
}

// "interactive" usage of the program.
// we print a greeting to the user
// generate print some seeds for them
// print a friendly reminder, to back up the seed
// and pause execution so the terminal won't be closed on windows
func interactive() {
	// greeting
	fmt.Printf("Welcome to iota-seedgen v%s - your friendly IOTA wallet seed generator!\n", version)
	fmt.Println("------------------")

	// generate and print 10 wallet seeds
	for i := 0; i < 10; i++ {
		token, err := generateRandomSeed()
		if err != nil {
			fmt.Printf("an error occured: %v\n	", err.Error())
			os.Exit(1)
		}

		fmt.Println(token)
	}

	// backup reminder
	fmt.Println("------------------")
	fmt.Println("pick one of the random lines, back it up anywhere >SAFE< and use it to generate your wallet :)")

	// pause execution
	fmt.Println("")
	fmt.Println("Press the Enter Key to terminate the console screen!")
	var input string
	fmt.Scanln(&input)
	os.Exit(0)
}

// non interactive usage of the program.
// just generate and print one seed
// then exit
func nonInteractive() {
	token, err := generateRandomSeed()
	if err != nil {
		fmt.Printf("an error occured: %v\n	", err.Error())
		os.Exit(1)
	}

	fmt.Print(token)
	os.Exit(0)
}

func main() {
	// command line flag handling
	scriptedPtr := flag.Bool("s", false, "is the execution scripted?")
	flag.Parse()

	if !*scriptedPtr {
		interactive()
	} else {
		nonInteractive()
	}
}
