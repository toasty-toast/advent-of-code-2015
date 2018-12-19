package utils

import (
	"bufio"
	"math"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// ReadLines reads all lines from a file and returns them in a slice.
func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// MinInt returns the int with the lowest value from the slice.
func MinInt(vals []int) int {
	min := math.MaxInt32
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}
