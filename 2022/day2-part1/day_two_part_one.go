package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	lineBreak string = "\r\n"
)

var outcomes = map[string]int64{
	"A X": 4, // 1 + 3
	"A Y": 8, // 2 + 6
	"A Z": 3, // 3 + 0
	"B X": 1, // 1 + 0
	"B Y": 5, // 2 + 3
	"B Z": 9, // 3 + 6
	"C X": 7, // 1 + 6
	"C Y": 2, // 2 + 0
	"C Z": 6, // 3 + 3
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v%s", err, lineBreak)
		return
	}

	rounds := strings.Split(string(dat), lineBreak)
	var total int64

	for _, round := range rounds {
		total += outcomes[round]
	}

	fmt.Printf("Score: %d%s", total, lineBreak)
}
