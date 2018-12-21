package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

func sumAllNumbers(s string) int {
	re := regexp.MustCompile(`(-?\d+)`)
	numbers := re.FindAllString(s, -1)
	sum := 0
	for _, number := range numbers {
		value, _ := strconv.Atoi(number)
		sum += value
	}
	return sum
}

func main() {
	input := utils.ReadLines("input.txt")[0]
	fmt.Printf("Part 1: %d\n", sumAllNumbers(input))
}
