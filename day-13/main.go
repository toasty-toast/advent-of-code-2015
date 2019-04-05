package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

type person struct {
	Name      string
	Neighbors map[string]int
}

func parsePeople(lines []string) []*person {
	lineRegex := regexp.MustCompile(`(?P<Person>[[:alpha:]]+) would (?P<Change>gain|lose) (?P<Points>\d+) happiness units by sitting next to (?P<Neighbor>[[:alpha:]]+).`)
	people := make([]*person, 0)
	peopleMap := make(map[string]*person)
	for _, line := range lines {
		stringMatch := lineRegex.FindStringSubmatch(line)
		name := stringMatch[1]
		curPerson, _ := peopleMap[name]
		if curPerson == nil {
			curPerson = new(person)
			curPerson.Name = name
			curPerson.Neighbors = make(map[string]int)
			peopleMap[name] = curPerson
			people = append(people, curPerson)
		}
		points, _ := strconv.Atoi(stringMatch[3])
		if strings.Compare("lose", stringMatch[2]) == 0 {
			points *= -1
		}
		curPerson.Neighbors[stringMatch[4]] = points
	}
	return people
}

func permutations(people []*person) [][]*person {
	interfaces := make([]interface{}, len(people))
	for i, val := range people {
		interfaces[i] = val
	}
	interfacePermutations := utils.Permutations(interfaces)
	permutations := make([][]*person, len(interfacePermutations))
	for i, permutation := range interfacePermutations {
		permutations[i] = make([]*person, len(permutation))
		for j, val := range permutation {
			permutations[i][j] = val.(*person)
		}
	}
	return permutations
}

func getScore(people []*person) int {
	score := 0
	for i, person := range people {
		prevIndex := i - 1
		if prevIndex < 0 {
			prevIndex = len(people) - 1
		}
		nextIndex := (i + 1) % len(people)

		prev := people[prevIndex]
		next := people[nextIndex]
		score += person.Neighbors[prev.Name] + person.Neighbors[next.Name]
	}
	return score
}

func arrange(people []*person) []*person {
	maxScore := math.MinInt32
	arrangement := people
	for _, permuation := range permutations(people) {
		score := getScore(permuation)
		if score > maxScore {
			maxScore = score
			arrangement = permuation
		}
	}
	return arrangement
}

func main() {
	people := parsePeople(utils.ReadLines("input.txt"))
	optimalScoreWithoutMe := getScore(arrange(people))

	me := new(person)
	me.Name = "Me"
	me.Neighbors = make(map[string]int)
	for _, person := range people {
		me.Neighbors[person.Name] = 0
		person.Neighbors[me.Name] = 0
	}
	people = append(people, me)
	optimalScoreWithMe := getScore(arrange(people))

	fmt.Printf("Part 1: %d\n", optimalScoreWithoutMe)
	fmt.Printf("Part 2: %d\n", optimalScoreWithMe)
}
