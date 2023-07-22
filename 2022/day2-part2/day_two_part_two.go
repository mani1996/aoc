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
	"A X": 3, // Scissors(3) + Loss(0)
	"A Y": 4, // Rock(1) + Draw(3)
	"A Z": 8, // Paper(2) + Win(6)
	"B X": 1, // Rock(1) + Loss(0)
	"B Y": 5, // Paper(2) + Draw(3)
	"B Z": 9, // Scissors(3) + Win(6)
	"C X": 2, // Paper(2) + Loss(0)
	"C Y": 6, // Scissors(3) + Draw(3)
	"C Z": 7, // Rock(1) + Win(6)
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
