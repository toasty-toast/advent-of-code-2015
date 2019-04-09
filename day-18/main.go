package main

import (
	"fmt"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

func parseLights(lines []string) [][]bool {
	lights := make([][]bool, 0)
	for _, line := range lines {
		row := make([]bool, 0)
		for _, char := range line {
			if char == '#' {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		lights = append(lights, row)
	}
	return lights
}

func isOn(lights [][]bool, x, y int) bool {
	if y < 0 || y >= len(lights) {
		return false
	}
	if x < 0 || x >= len(lights[y]) {
		return false
	}
	return lights[y][x]
}

func countNeighbors(lights [][]bool, x, y int) int {
	count := 0
	if isOn(lights, x-1, y-1) {
		count++
	}
	if isOn(lights, x-1, y) {
		count++
	}
	if isOn(lights, x-1, y+1) {
		count++
	}
	if isOn(lights, x, y-1) {
		count++
	}
	if isOn(lights, x, y+1) {
		count++
	}
	if isOn(lights, x+1, y-1) {
		count++
	}
	if isOn(lights, x+1, y) {
		count++
	}
	if isOn(lights, x+1, y+1) {
		count++
	}
	return count
}

func step(lights [][]bool) [][]bool {
	nextState := make([][]bool, len(lights))
	for i := range nextState {
		nextState[i] = make([]bool, len(lights[i]))
	}

	for i := range lights {
		for j := range lights[i] {
			nextState[i][j] = lights[i][j]
			neighbors := countNeighbors(lights, j, i)
			if lights[i][j] && (neighbors < 2 || neighbors > 3) {
				nextState[i][j] = false
			} else if !lights[i][j] && neighbors == 3 {
				nextState[i][j] = true
			}
		}
	}

	return nextState
}

func stepWithCornersLocked(lights [][]bool) [][]bool {
	nextState := make([][]bool, len(lights))
	for i := range nextState {
		nextState[i] = make([]bool, len(lights[i]))
	}

	for i := range lights {
		for j := range lights[i] {
			if (i == 0 && j == 0) || (i == 0 && j == len(lights[i])-1) || i == len(lights)-1 && j == 0 || i == len(lights)-1 && j == len(lights)-1 {
				nextState[i][j] = true
				continue
			}

			nextState[i][j] = lights[i][j]
			neighbors := countNeighbors(lights, j, i)
			if lights[i][j] && (neighbors < 2 || neighbors > 3) {
				nextState[i][j] = false
			} else if !lights[i][j] && neighbors == 3 {
				nextState[i][j] = true
			}
		}
	}

	return nextState
}

func countLights(lights [][]bool) int {
	count := 0
	for _, row := range lights {
		for _, light := range row {
			if light {
				count++
			}
		}
	}
	return count
}

func main() {
	originalLights := parseLights(utils.ReadLines("input.txt"))

	lights := originalLights
	for i := 0; i < 100; i++ {
		lights = step(lights)
	}
	fmt.Printf("Part 1: %d\n", countLights(lights))

	lights = originalLights
	for i := 0; i < 100; i++ {
		lights = stepWithCornersLocked(lights)
	}
	fmt.Printf("Part 2: %d\n", countLights(lights))
}
