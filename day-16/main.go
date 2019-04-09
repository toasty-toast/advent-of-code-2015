package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

var numRegex = regexp.MustCompile(`Sue (?P<Number>\d+)`)
var childrenRegex = regexp.MustCompile(`children: (?P<Children>\d+)`)
var catsRegex = regexp.MustCompile(`cats: (?P<Cats>\d+)`)
var samoyedsRegex = regexp.MustCompile(`samoyeds: (?P<Samoyeds>\d+)`)
var pomeraniansRegex = regexp.MustCompile(`pomeranians: (?P<Pomeranians>\d+)`)
var akitasRegex = regexp.MustCompile(`akitas: (?P<Akitas>\d+)`)
var vizslasRegex = regexp.MustCompile(`vizslas: (?P<Vizslas>\d+)`)
var goldfishRegex = regexp.MustCompile(`goldfish: (?P<Goldfish>\d+)`)
var treesRegex = regexp.MustCompile(`trees: (?P<Trees>\d+)`)
var carsRegex = regexp.MustCompile(`cars: (?P<Cars>\d+)`)
var perfumesRegex = regexp.MustCompile(`perfumes: (?P<Perfumes>\d+)`)

const unknownValue = -1
const targetChildren = 3
const targetCats = 7
const targetSamoyeds = 2
const targetPomeranians = 3
const targetAkitas = 0
const targetVizslas = 0
const targetGoldfish = 5
const targetTrees = 3
const targetCars = 2
const targetPerfumes = 1

type aunt struct {
	Number, Children, Cats, Samoyeds, Pomeranians, Akitas, Vizslas, Goldfish, Trees, Cars, Perfumes int
}

func parseAunt(line string) *aunt {
	numMatch := numRegex.FindStringSubmatch(line)
	aunt := new(aunt)
	aunt.Number, _ = strconv.Atoi(numMatch[1])

	childrenMatch := childrenRegex.FindStringSubmatch(line)
	if len(childrenMatch) > 0 {
		aunt.Children, _ = strconv.Atoi(childrenMatch[1])
	} else {
		aunt.Children = unknownValue
	}

	catsMatch := catsRegex.FindStringSubmatch(line)
	if len(catsMatch) > 0 {
		aunt.Cats, _ = strconv.Atoi(catsMatch[1])
	} else {
		aunt.Cats = unknownValue
	}

	samoyedsMatch := samoyedsRegex.FindStringSubmatch(line)
	if len(samoyedsMatch) > 0 {
		aunt.Samoyeds, _ = strconv.Atoi(samoyedsMatch[1])
	} else {
		aunt.Samoyeds = unknownValue
	}

	pomeraniansMatch := pomeraniansRegex.FindStringSubmatch(line)
	if len(pomeraniansMatch) > 0 {
		aunt.Pomeranians, _ = strconv.Atoi(pomeraniansMatch[1])
	} else {
		aunt.Pomeranians = unknownValue
	}

	akitasMatch := akitasRegex.FindStringSubmatch(line)
	if len(akitasMatch) > 0 {
		aunt.Akitas, _ = strconv.Atoi(akitasMatch[1])
	} else {
		aunt.Akitas = unknownValue
	}

	vizslasMatch := vizslasRegex.FindStringSubmatch(line)
	if len(vizslasMatch) > 0 {
		aunt.Vizslas, _ = strconv.Atoi(vizslasMatch[1])
	} else {
		aunt.Vizslas = unknownValue
	}

	goldfishMatch := goldfishRegex.FindStringSubmatch(line)
	if len(goldfishMatch) > 0 {
		aunt.Goldfish, _ = strconv.Atoi(goldfishMatch[1])
	} else {
		aunt.Goldfish = unknownValue
	}

	treesMatch := treesRegex.FindStringSubmatch(line)
	if len(treesMatch) > 0 {
		aunt.Trees, _ = strconv.Atoi(treesMatch[1])
	} else {
		aunt.Trees = unknownValue
	}

	carsMatch := carsRegex.FindStringSubmatch(line)
	if len(carsMatch) > 0 {
		aunt.Cars, _ = strconv.Atoi(carsMatch[1])
	} else {
		aunt.Cars = unknownValue
	}

	perfumesMatch := perfumesRegex.FindStringSubmatch(line)
	if len(perfumesMatch) > 0 {
		aunt.Perfumes, _ = strconv.Atoi(perfumesMatch[1])
	} else {
		aunt.Perfumes = unknownValue
	}

	return aunt
}

func parseAunts(lines []string) []*aunt {
	aunts := make([]*aunt, 0)
	for _, line := range lines {
		aunts = append(aunts, parseAunt(line))
	}
	return aunts
}

func testAunt(aunt *aunt) bool {
	if aunt.Children != unknownValue && aunt.Children != targetChildren {
		return false
	}
	if aunt.Cats != unknownValue && aunt.Cats != targetCats {
		return false
	}
	if aunt.Samoyeds != unknownValue && aunt.Samoyeds != targetSamoyeds {
		return false
	}
	if aunt.Pomeranians != unknownValue && aunt.Pomeranians != targetPomeranians {
		return false
	}
	if aunt.Akitas != unknownValue && aunt.Akitas != targetAkitas {
		return false
	}
	if aunt.Vizslas != unknownValue && aunt.Vizslas != targetVizslas {
		return false
	}
	if aunt.Goldfish != unknownValue && aunt.Goldfish != targetGoldfish {
		return false
	}
	if aunt.Trees != unknownValue && aunt.Trees != targetTrees {
		return false
	}
	if aunt.Cars != unknownValue && aunt.Cars != targetCars {
		return false
	}
	if aunt.Perfumes != unknownValue && aunt.Perfumes != targetPerfumes {
		return false
	}
	return true
}

func testAuntImproved(aunt *aunt) bool {
	if aunt.Children != unknownValue && aunt.Children != targetChildren {
		return false
	}
	if aunt.Cats != unknownValue && aunt.Cats <= targetCats {
		return false
	}
	if aunt.Samoyeds != unknownValue && aunt.Samoyeds != targetSamoyeds {
		return false
	}
	if aunt.Pomeranians != unknownValue && aunt.Pomeranians >= targetPomeranians {
		return false
	}
	if aunt.Akitas != unknownValue && aunt.Akitas != targetAkitas {
		return false
	}
	if aunt.Vizslas != unknownValue && aunt.Vizslas != targetVizslas {
		return false
	}
	if aunt.Goldfish != unknownValue && aunt.Goldfish >= targetGoldfish {
		return false
	}
	if aunt.Trees != unknownValue && aunt.Trees <= targetTrees {
		return false
	}
	if aunt.Cars != unknownValue && aunt.Cars != targetCars {
		return false
	}
	if aunt.Perfumes != unknownValue && aunt.Perfumes != targetPerfumes {
		return false
	}
	return true
}

func main() {
	aunts := parseAunts(utils.ReadLines("input.txt"))
	for _, aunt := range aunts {
		if testAunt(aunt) {
			fmt.Printf("Part 1: %d\n", aunt.Number)
			break
		}
	}
	for _, aunt := range aunts {
		if testAuntImproved(aunt) {
			fmt.Printf("Part 2: %d\n", aunt.Number)
			break
		}
	}
}
