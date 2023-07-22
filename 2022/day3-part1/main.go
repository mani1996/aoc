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

	for _, sack := range rucksacks {
		l := len(sack)
		tracker := map[rune]bool{}
		for _, c := range sack[:l/2] {
			tracker[c] = true
		}

		for _, c := range sack[l/2:] {
			if _, ok := tracker[c]; ok {
				if c >= 'a' && c <= 'z' {
					total += int64(c - 'a' + 1)
				} else {
					total += (27 + int64(c-'A'))
				}
				break
			}
		}
	}

	fmt.Printf("Score: %d%s", total, lineBreak)
}
