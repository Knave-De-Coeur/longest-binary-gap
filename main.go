package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solution(N int) int {
	// Implement your solution here

	var (
		binaryNum      string
		currentGap     int
		listOfGapSizes []int
		countOfOnes    int
		onesPassed     int
	)

	binaryNum = strconv.FormatInt(int64(N), 2)
	fmt.Printf("the binary number of %d is %s \n", N, binaryNum)

	countOfOnes = strings.Count(binaryNum, "1")

	for _, bin := range binaryNum {
		digit := fmt.Sprintf("%c", bin)
		if digit == "0" {
			currentGap++
		}

		if digit == "1" {
			onesPassed++

			if currentGap >= 1 {
				listOfGapSizes = append(listOfGapSizes, currentGap)
				currentGap = 0
			}

			// trailing 0s
			if onesPassed == countOfOnes {
				break
			}

			continue
		}

	}

	var answer int

	if len(listOfGapSizes) < 1 {
		answer = 0
	} else if len(listOfGapSizes) == 1 {
		answer = listOfGapSizes[0]
	} else {
		sort.Ints(listOfGapSizes)
		answer = listOfGapSizes[len(listOfGapSizes)-1]
	}

	return answer
}

func main() {
	examples := []int{9, 529, 20, 15, 32}
	for _, test := range examples {
		binaryGap := Solution(test)
		fmt.Printf("longest binray gap of %d is %d \n", test, binaryGap)
	}
}
