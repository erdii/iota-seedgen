package main

import (
	"fmt"
	"os"
	"testing"
)

func TestIntDistribution(*testing.T) {
	counts := make(map[int64]int64)
	ints, err := generateRandomInts(10000000)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, x := range ints {
		counts[x]++
	}

	for i := 0; i <= 26; i++ {
		fmt.Printf("%02d: %d\n", i, counts[int64(i)])
	}
}

func TestCharacterDistribution(*testing.T) {
	const letters = "9ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	counts := make(map[byte]int64)
	ints, err := generateRandomInts(10000000)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, x := range ints {
		char := letters[x]
		counts[char]++
	}

	for i := 0; i <= 26; i++ {
		count := counts[letters[i]]
		fmt.Printf("%s: %d\n", string(letters[i]), count)
	}
}
