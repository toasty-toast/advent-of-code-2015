package main

import (
	"fmt"
	"strings"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

func lookAndSay(input string, iterations int) string {
	value := input
	for i := 0; i < iterations; i++ {
		var builder strings.Builder
		cur := value[0]
		count := 1
		for j := 1; j < len(value); j++ {
			if value[j] != cur {
				builder.WriteString(fmt.Sprintf("%d%c", count, cur))
				cur = value[j]
				count = 0
			}
			count++
		}
		builder.WriteString(fmt.Sprintf("%d%c", count, cur))
		value = builder.String()
	}
	return value
}

func main() {
	input := utils.ReadLines("input.txt")[0]
	fmt.Printf("Part 1: %d\n", len(lookAndSay(input, 40)))
	fmt.Printf("Part 1: %d\n", len(lookAndSay(input, 50)))
}
