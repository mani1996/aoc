package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	lineBreak string = "\r\n"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v%s", err, lineBreak)
		return
	}

	rucksacks := strings.Split(string(dat), lineBreak)
	var total int64

	for i := 0; i < len(rucksacks); i += 3 {
		t1 := map[rune]bool{}
		t2 := map[rune]bool{}
		t3 := map[rune]bool{}

		for _, c := range rucksacks[i] {
			t1[c] = true
		}

		for _, c := range rucksacks[i+1] {
			t2[c] = true
		}

		for _, c := range rucksacks[i+2] {
			t3[c] = true
		}

		for c := 'A'; c <= 'Z'; c++ {
			if t1[c] && t2[c] && t3[c] {
				total += int64(27 + c - 'A')
			}
		}

		for c := 'a'; c <= 'z'; c++ {
			if t1[c] && t2[c] && t3[c] {
				total += int64(1 + c - 'a')
			}
		}
	}

	fmt.Printf("Score: %d%s", total, lineBreak)
}
