package main

import (
	"fmt"
	"os"
	"strconv"
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

	assignments := strings.Split(string(dat), lineBreak)
	var total int64

	for _, a := range assignments {
		parts := strings.Split(a, ",")
		p1 := strings.Split(parts[0], "-")
		p2 := strings.Split(parts[1], "-")
		l1, err := strconv.ParseInt(p1[0], 10, 64)
		if err != nil {
			fmt.Printf("Error in parsing assignment %s : %v%s", a, err, lineBreak)
		}
		r1, err := strconv.ParseInt(p1[1], 10, 64)
		if err != nil {
			fmt.Printf("Error in parsing assignment %s : %v%s", a, err, lineBreak)
		}
		l2, err := strconv.ParseInt(p2[0], 10, 64)
		if err != nil {
			fmt.Printf("Error in parsing assignment %s : %v%s", a, err, lineBreak)
		}
		r2, err := strconv.ParseInt(p2[1], 10, 64)
		if err != nil {
			fmt.Printf("Error in parsing assignment %s : %v%s", a, err, lineBreak)
		}

		if l1 <= r2 && r1 >= l2 {
			total++
		}
	}

	fmt.Printf("Total: %d%s", total, lineBreak)
}
