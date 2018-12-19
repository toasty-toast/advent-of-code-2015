package main

import (
	"fmt"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

const (
	left = iota
	up
	right
	down
)

type direction int

func loadDirections(filename string) []direction {
	directions := make([]direction, 0)
	line := utils.ReadLines(filename)[0]
	for _, char := range line {
		switch char {
		case '<':
			directions = append(directions, left)
			break
		case '^':
			directions = append(directions, up)
			break
		case '>':
			directions = append(directions, right)
			break
		case 'v':
			directions = append(directions, down)
			break
		}
	}
	return directions
}

func extendLeft(arr [][]int) [][]int {
	for i := range arr {
		arr[i] = append([]int{0}, arr[i]...)
	}
	return arr
}

func extendRight(arr [][]int) [][]int {
	for i := range arr {
		arr[i] = append(arr[i], 0)
	}
	return arr
}

func extendUp(arr [][]int) [][]int {
	prepend := make([][]int, 1)
	prepend[0] = make([]int, len(arr[0]))
	return append(prepend, arr...)
}

func extendDown(arr [][]int) [][]int {
	return append(arr, make([]int, len(arr[0])))
}

func completeRoute(directions []direction, numSantas int) [][]int {
	houses := make([][]int, 1)
	houses[0] = make([]int, 1)
	houses[0][0] = 1
	x, y := make([]int, numSantas), make([]int, numSantas)
	for i := 0; i < len(directions); i += numSantas {
		for j := 0; j < numSantas && (i+j) < len(directions); j++ {
			dir := directions[i+j]
			switch dir {
			case left:
				if x[j] == 0 {
					extendLeft(houses)
					for k := range x {
						if k != j {
							x[k]++
						}
					}
				} else {
					x[j]--
				}
				break
			case up:
				if y[j] == 0 {
					houses = extendUp(houses)
					for k := range y {
						if k != j {
							y[k]++
						}
					}
				} else {
					y[j]--
				}
				break
			case right:
				if x[j] == len(houses[y[j]])-1 {
					houses = extendRight(houses)
				}
				x[j]++
				break
			case down:
				if y[j] == len(houses)-1 {
					houses = extendDown(houses)
				}
				y[j]++
				break
			}
			houses[y[j]][x[j]]++
		}
	}
	return houses
}

func countWithPresents(houses [][]int) int {
	sum := 0
	for i := range houses {
		for j := range houses[i] {
			if houses[i][j] != 0 {
				sum++
			}
		}
	}
	return sum
}

func main() {
	directions := loadDirections("input.txt")
	fmt.Printf("Houses with presents for 1 santa: %d\n", countWithPresents(completeRoute(directions, 1)))
	fmt.Printf("Houses with presents for 2 santas: %d\n", countWithPresents(completeRoute(directions, 2)))
}
