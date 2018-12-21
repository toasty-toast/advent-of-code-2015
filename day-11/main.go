package main

import (
	"fmt"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

func containsThreeStraightLetters(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i+1] == s[i]+1 && s[i+2] == s[i]+2 {
			return true
		}
	}
	return false
}

func containsBadCharacters(s string) bool {
	for i := range s {
		if s[i] == 'i' || s[i] == 'o' || s[i] == 'l' {
			return true
		}
	}
	return false
}

func containsAtLeastTwoPairs(s string) bool {
	pairs := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			pairs++
			i++
			if pairs == 2 {
				return true
			}
		}
	}
	return pairs == 2
}

func increment(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'z' {
			s = utils.ReplaceAtIndex(s, i, 'a')
		} else {
			s = utils.ReplaceAtIndex(s, i, rune(s[i]+1))
			return s
		}
	}
	s = fmt.Sprintf("a%s", s)
	return s
}

func findNextPassword(input string) string {
	for {
		input = increment(input)
		if containsThreeStraightLetters(input) && !containsBadCharacters(input) && containsAtLeastTwoPairs(input) {
			return input
		}
	}
}

func main() {
	input := utils.ReadLines("input.txt")[0]
	first := findNextPassword(input)
	second := findNextPassword(first)
	fmt.Printf("Part 1: %s\n", first)
	fmt.Printf("Part 2: %s\n", second)
}
