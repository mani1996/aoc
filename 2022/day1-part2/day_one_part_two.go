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
	foodPerMonk := strings.Split(string(dat), lineBreak+lineBreak)
	var r1, r2, r3 int64

	for _, foodString := range foodPerMonk {
		foods := strings.Split(foodString, lineBreak)
		var calories int64

		for _, food := range foods {
			foodNum, err := strconv.ParseInt(food, 10, 64)
			if err != nil {
				fmt.Printf("Unable to convert %s to integer : %v%s", food, err, lineBreak)
				return
			}

			calories += foodNum
		}

		if calories > r1 {
			if calories > r2 {
				r1 = r2
				if calories > r3 {
					r2 = r3
					r3 = calories
				} else {
					r2 = calories
				}
			} else {
				r1 = calories
			}
		}
	}

	fmt.Printf("Result : %d%s", r1+r2+r3, lineBreak)
}
