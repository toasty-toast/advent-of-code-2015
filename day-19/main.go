package main

import (
	"fmt"
	"index/suffixarray"
	"regexp"
	"strings"
	"unicode"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

type transition struct {
	Input, Output string
}

func parseInput(lines []string) (string, []*transition) {
	transitionRegex := regexp.MustCompile(`(?P<Input>[[:alpha:]]+) => (?P<Output>[[:alpha:]]+)`)
	transitions := make([]*transition, 0)
	for _, line := range lines {
		if strings.Compare(line, "") == 0 {
			break
		}
		match := transitionRegex.FindStringSubmatch(line)
		next := new(transition)
		next.Input = match[1]
		next.Output = match[2]
		transitions = append(transitions, next)
	}
	return lines[len(lines)-1], transitions
}

func getDistnctOutputs(value string, transitions []*transition) []string {
	outputs := make([]string, 0)
	outputSet := make(map[string]struct{})
	lookup := suffixarray.New([]byte(value))
	for _, transition := range transitions {
		indices := lookup.Lookup([]byte(transition.Input), -1)
		for _, index := range indices {
			prefix := value[:index]
			suffix := value[index+len(transition.Input):]
			output := strings.Join([]string{prefix, transition.Output, suffix}, "")
			if _, ok := outputSet[output]; ok {
				continue
			}
			outputSet[output] = struct{}{}
			outputs = append(outputs, output)
		}
	}
	return outputs
}

func countStepsForMolecule(value string) int {
	numUpperCase := 0
	for _, char := range value {
		if unicode.IsUpper(char) {
			numUpperCase++
		}
	}

	lookup := suffixarray.New([]byte(value))
	numRn := len(lookup.Lookup([]byte("Rn"), -1))
	numAr := len(lookup.Lookup([]byte("Ar"), -1))
	numY := len(lookup.Lookup([]byte("Y"), -1))

	return numUpperCase - numRn - numAr - 2*numY - 1
}

func main() {
	value, transitions := parseInput(utils.ReadLines("input.txt"))
	outputs := getDistnctOutputs(value, transitions)
	fmt.Printf("Part 1: %d\n", len(outputs))
	fmt.Printf("Part 2: %d\n", countStepsForMolecule(value))
}
