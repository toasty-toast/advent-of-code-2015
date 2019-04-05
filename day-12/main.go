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

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	max := 0
	for i := range height {
		for j := i; i < len(height); j++ {
			volume := minInt(height[i], height[j]) * (j - i)
			if volume > max {
				max = volume
			}
		}
	}
	return max
}
