package main

import (
	"fmt"
	"strings"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}
var illegalStrings = []string{"ab", "cd", "pq", "xy"}

func isVowel(r rune) bool {
	for _, vowel := range vowels {
		if r == vowel {
			return true
		}
	}
	return false
}

func hasThreeVowels(s string) bool {
	count := 0
	for _, char := range s {
		if isVowel(char) {
			count++
		}
		if count == 3 {
			return true
		}
	}
	return false
}

func hasRepeatedLetter(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
	}
	return false
}

func hasIllegalSubstring(s string) bool {
	for _, illegal := range illegalStrings {
		if strings.Contains(s, illegal) {
			return true
		}
	}
	return false
}

func hasRepeatedLetterPair(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}
	return false
}

func hasRepeatedLetterSandwich(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func countNiceStringsPart1(strs []string) int {
	count := 0
	for _, str := range strs {
		if hasThreeVowels(str) && hasRepeatedLetter(str) && !hasIllegalSubstring(str) {
			count++
		}
	}
	return count
}

func countNiceStringsPart2(strs []string) int {
	count := 0
	for _, str := range strs {
		if hasRepeatedLetterPair(str) && hasRepeatedLetterSandwich(str) {
			count++
		}
	}
	return count
}

func main() {
	inputs := utils.ReadLines("input.txt")
	fmt.Printf("Part 1: %d nice strings\n", countNiceStringsPart1(inputs))
	fmt.Printf("Part 2: %d nice strings\n", countNiceStringsPart2(inputs))
}
