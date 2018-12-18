package main

import (
	"fmt"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

func finalFloor(input string) int {
	floor := 0
	for _, char := range input {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}
	return floor
}

func positionAtBasement(input string) int {
	floor := 0
	for i, char := range input {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return 0
}

func main() {
	input := utils.ReadLines("input.txt")[0]
	fmt.Printf("Final floor: %d\n", finalFloor(input))
	fmt.Printf("Position when entering basement: %d\n", positionAtBasement(input))
}
