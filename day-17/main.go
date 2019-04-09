package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

const totalLiters = 150

func parseContainers(lines []string) []int {
	containers := make([]int, 0)
	for _, line := range lines {
		container, _ := strconv.Atoi(line)
		containers = append(containers, container)
	}
	return containers
}

func getContainerCombinations(containers []int) [][]int {
	combinations := make([][]int, 0)
	for i := 1; i <= (1 << (uint)(len(containers))); i++ {
		capacity := 0
		combination := make([]int, 0)
		for j := range containers {
			if i&(1<<(uint)(j)) != 0 {
				capacity += containers[j]
				combination = append(combination, containers[j])
			}
		}
		if capacity == totalLiters {
			combinations = append(combinations, combination)
		}
	}
	return combinations
}

func countBestCombinations(combinations [][]int) int {
	smallestSize := math.MaxInt32
	sizeToCount := make(map[int]int)
	for _, combination := range combinations {
		size := len(combination)

		if _, ok := sizeToCount[size]; !ok {
			sizeToCount[size] = 0
		}
		sizeToCount[size]++

		if size < smallestSize {
			smallestSize = size
		}
	}
	return sizeToCount[smallestSize]
}

func main() {
	containers := parseContainers(utils.ReadLines("input.txt"))
	combinations := getContainerCombinations(containers)
	fmt.Printf("Part 1: %d\n", len(combinations))
	fmt.Printf("Part 2: %d\n", countBestCombinations(combinations))
}
