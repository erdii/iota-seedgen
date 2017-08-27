package main

import (
	"fmt"
	"os"
	"testing"
)

func TestIntDistribution(t *testing.T) {
	const targetCount = 10000000
	counts := make(map[int64]int64)
	ints, err := generateRandomInts(targetCount)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, x := range ints {
		counts[x]++
	}

	sum := int64(0)

	for i := 0; i <= 26; i++ {
		count := counts[int64(i)]
		sum += count
		fmt.Printf("%02d: %d\n", i, count)
	}

	fmt.Printf("Sum: %d\n", sum)

	if sum != targetCount {
		t.Fatal(fmt.Sprintf("sum (%d) != targetCount(%d)", sum, targetCount))
	}
}

func TestCharacterDistribution(t *testing.T) {
	const targetCount = 10000000
	const letters = "9ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	counts := make(map[byte]int64)
	ints, err := generateRandomInts(targetCount)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, x := range ints {
		char := letters[x]
		counts[char]++
	}

	sum := int64(0)

	for i := 0; i <= 26; i++ {
		count := counts[letters[i]]
		sum += count
		fmt.Printf("%s: %d\n", string(letters[i]), count)
	}

	fmt.Printf("Sum: %d\n", sum)

	if sum != targetCount {
		t.Fatal(fmt.Sprintf("sum (%d) != targetCount(%d)", sum, targetCount))
	}
}
