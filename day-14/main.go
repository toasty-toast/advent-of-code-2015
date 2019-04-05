package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

var timeLimit = 2503

type reindeer struct {
	Name                     string
	Speed, FlyTime, RestTime int
}

func parseReindeer(lines []string) []*reindeer {
	lineRegex := regexp.MustCompile(`(?P<Name>[[:alpha:]]+) can fly (?P<Speed>\d+) km/s for (?P<FlyTime>\d+) seconds, but then must rest for (?P<RestTime>\d+) seconds.`)
	list := make([]*reindeer, 0)
	for _, line := range lines {
		lineMatch := lineRegex.FindStringSubmatch(line)
		deer := new(reindeer)
		deer.Name = lineMatch[1]
		deer.Speed, _ = strconv.Atoi(lineMatch[2])
		deer.FlyTime, _ = strconv.Atoi(lineMatch[3])
		deer.RestTime, _ = strconv.Atoi(lineMatch[4])
		list = append(list, deer)
	}
	return list
}

func distance(deer *reindeer) int {
	distance := 0
	for i := 0; i < timeLimit; i++ {
		if i%(deer.FlyTime+deer.RestTime) < deer.FlyTime {
			distance += deer.Speed
		}
	}
	return distance
}

func longestDistance(list []*reindeer) int {
	longest := 0
	for _, deer := range list {
		distance := distance(deer)
		if distance > longest {
			longest = distance
		}
	}
	return longest
}

func mostPoints(list []*reindeer) int {
	points := make(map[*reindeer]int)
	distance := make(map[*reindeer]int)
	for _, deer := range list {
		points[deer] = 0
		distance[deer] = 0
	}

	for i := 0; i < timeLimit; i++ {
		furthest := 0
		for _, deer := range list {
			if i%(deer.FlyTime+deer.RestTime) < deer.FlyTime {
				distance[deer] += deer.Speed
			}
			if distance[deer] > furthest {
				furthest = distance[deer]
			}
		}

		for _, deer := range list {
			if distance[deer] == furthest {
				points[deer]++
			}
		}
	}

	highest := 0
	for _, deer := range list {
		if points[deer] > highest {
			highest = points[deer]
		}
	}
	return highest
}

func main() {
	reindeer := parseReindeer(utils.ReadLines("input.txt"))
	fmt.Printf("Part 1: %d km\n", longestDistance(reindeer))
	fmt.Printf("Part 2: %d points\n", mostPoints(reindeer))
}
