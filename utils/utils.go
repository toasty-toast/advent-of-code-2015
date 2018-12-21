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

// Permutations returns all permutations of the provide slice.
func Permutations(arr []interface{}) [][]interface{} {
	var helper func([]interface{}, int)
	res := [][]interface{}{}
	helper = func(arr []interface{}, n int) {
		if n == 1 {
			tmp := make([]interface{}, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
