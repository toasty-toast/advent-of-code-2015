package main

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

func findSuffix(prefix string, leadingZeroes int) int {
	target := strings.Repeat("0", leadingZeroes)
	for i := 0; ; i++ {
		input := fmt.Sprintf("%s%d", prefix, i)
		hash := md5.Sum([]byte(input))
		hex := fmt.Sprintf("%x", hash)
		if hex[:leadingZeroes] == target {
			return i
		}
	}
}

func main() {
	prefix := utils.ReadLines("input.txt")[0]
	fmt.Printf("Suffix for 5 leading zeroes: %d\n", findSuffix(prefix, 5))
	fmt.Printf("Suffix for 6 leading zeroes: %d\n", findSuffix(prefix, 6))
}
