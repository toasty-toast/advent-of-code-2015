package main

import "fmt"

const puzzleInput = 29000000

func main() {
	{
		houses := make([]int, puzzleInput/10)
		for i := 1; i < puzzleInput/10; i++ {
			for j := i; j < puzzleInput/10; j += i {
				houses[j] += i * 10
			}
		}
		for i := 0; i < puzzleInput/10; i++ {
			if houses[i] >= puzzleInput {
				fmt.Printf("Part 1: %d\n", i)
				break
			}
		}
	}
	{
		houses := make([]int, puzzleInput/10)
		elves := make([]int, puzzleInput/10)
		for i := 1; i < puzzleInput/10; i++ {
			for j := i; j < puzzleInput/10; j += i {
				if elves[i] >= 50 {
					break
				}
				houses[j] += i * 11
				elves[i]++
			}
		}
		for i := 0; i < puzzleInput/10; i++ {
			if houses[i] >= puzzleInput {
				fmt.Printf("Part 2: %d\n", i)
				break
			}
		}
	}
}
