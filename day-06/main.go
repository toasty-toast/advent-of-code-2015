package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

const (
	turnOn = iota
	turnOff
	toggle
)

type action int

type coordinate struct {
	X, Y int
}

type direction struct {
	TopLeft, BottomRight coordinate
	Action               action
}

func loadDirections(filename string) []*direction {
	re := regexp.MustCompile(`^(.*) (\d+),(\d+) through (\d+),(\d+)$`)
	lines := utils.ReadLines(filename)
	directions := make([]*direction, 0)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		cur := new(direction)
		switch match[1] {
		case "turn off":
			cur.Action = turnOff
			break
		case "turn on":
			cur.Action = turnOn
			break
		case "toggle":
			cur.Action = toggle
			break
		default:
			panic("Unknown instruction")
		}
		cur.TopLeft.X, _ = strconv.Atoi(match[2])
		cur.TopLeft.Y, _ = strconv.Atoi(match[3])
		cur.BottomRight.X, _ = strconv.Atoi(match[4])
		cur.BottomRight.Y, _ = strconv.Atoi(match[5])
		directions = append(directions, cur)
	}
	return directions
}

func followOnOffDirections(directions []*direction) [][]bool {
	lights := make([][]bool, 1000)
	for i := range lights {
		lights[i] = make([]bool, 1000)
	}
	for _, direction := range directions {
		for i := direction.TopLeft.Y; i <= direction.BottomRight.Y; i++ {
			for j := direction.TopLeft.X; j <= direction.BottomRight.X; j++ {
				switch direction.Action {
				case turnOff:
					lights[i][j] = false
					break
				case turnOn:
					lights[i][j] = true
					break
				case toggle:
					lights[i][j] = !lights[i][j]
					break
				}
			}
		}
	}
	return lights
}

func followBrightnessDirections(directions []*direction) [][]int {
	lights := make([][]int, 1000)
	for i := range lights {
		lights[i] = make([]int, 1000)
	}
	for _, direction := range directions {
		for i := direction.TopLeft.Y; i <= direction.BottomRight.Y; i++ {
			for j := direction.TopLeft.X; j <= direction.BottomRight.X; j++ {
				switch direction.Action {
				case turnOff:
					if lights[i][j] > 0 {
						lights[i][j]--
					}
					break
				case turnOn:
					lights[i][j]++
					break
				case toggle:
					lights[i][j] += 2
					break
				}
			}
		}
	}
	return lights
}

func countOnLights(lights [][]bool) int {
	count := 0
	for i := range lights {
		for j := range lights[i] {
			if lights[i][j] {
				count++
			}
		}
	}
	return count
}

func totalBrightness(lights [][]int) int {
	count := 0
	for i := range lights {
		for j := range lights[i] {
			count += lights[i][j]
		}
	}
	return count
}

func main() {
	directions := loadDirections("input.txt")
	fmt.Printf("Part 1: %d lights are on\n", countOnLights(followOnOffDirections(directions)))
	fmt.Printf("Part 2: total brightness is %d\n", totalBrightness(followBrightnessDirections(directions)))
}
