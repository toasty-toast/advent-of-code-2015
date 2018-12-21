package main

import (
	"fmt"
	"regexp"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

func decode(s string) string {
	backslashRe := regexp.MustCompile(`\\\\`)
	quoteRe := regexp.MustCompile(`\\"`)
	asciiRe := regexp.MustCompile(`\\x[0123456789abcdef][0123456789abcdef]`)
	mod := s[1 : len(s)-1]
	mod = backslashRe.ReplaceAllString(mod, `\`)
	mod = quoteRe.ReplaceAllString(mod, `"`)
	mod = asciiRe.ReplaceAllString(mod, `_`)
	return mod
}

func encode(s string) string {
	backslashRe := regexp.MustCompile(`\\`)
	quoteRe := regexp.MustCompile(`"`)
	mod := s
	mod = backslashRe.ReplaceAllString(mod, `\\`)
	mod = quoteRe.ReplaceAllString(mod, `\"`)
	return fmt.Sprintf(`"%s"`, mod)
}

func main() {
	lines := utils.ReadLines("input.txt")

	original, decoded, encoded := 0, 0, 0
	for _, line := range lines {
		original += len(line)
		decoded += len(decode(line))
		encoded += len(encode(line))
	}
	fmt.Printf("Part 1: %d\n", original-decoded)
	fmt.Printf("Part 2: %d\n", encoded-original)
}
