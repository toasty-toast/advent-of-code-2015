package utils

import (
	"bufio"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// ReadLines reads all lines from a file and returns them in a slice
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
