package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

type gate interface {
	OutputName() string
	OutputValue() uint16
	OverrideOutput(uint16)
	SetInputs([]gate)
}

type sourceGate struct {
	input, overrideValue uint16
	outputName           string
	override             bool
}

func (g *sourceGate) OutputName() string {
	return g.outputName
}

func (g *sourceGate) OutputValue() uint16 {
	if g.override {
		return g.overrideValue
	}
	return g.input
}

func (g *sourceGate) OverrideOutput(output uint16) {
	g.override = true
	g.overrideValue = output
}

func (g *sourceGate) SetInputs(inputs []gate) {}

type passthroughGate struct {
	input                         gate
	outputName                    string
	computedOutput, overrideValue uint16
	computed, override            bool
}

func (g *passthroughGate) OutputName() string {
	return g.outputName
}

func (g *passthroughGate) OutputValue() uint16 {
	if g.override {
		return g.overrideValue
	}
	if !g.computed {
		g.computedOutput = g.input.OutputValue()
		g.computed = true
	}
	return g.computedOutput
}

func (g *passthroughGate) OverrideOutput(output uint16) {
	g.override = true
	g.overrideValue = output
}

func (g *passthroughGate) SetInputs(inputs []gate) {
	g.input = inputs[0]
}

type lshiftGate struct {
	input                                      gate
	shiftAmount, computedOutput, overrideValue uint16
	outputName                                 string
	computed, override                         bool
}

func (g *lshiftGate) OutputName() string {
	return g.outputName
}

func (g *lshiftGate) OutputValue() uint16 {
	if g.override {
		return g.overrideValue
	}
	if !g.computed {
		g.computedOutput = g.input.OutputValue() << g.shiftAmount
		g.computed = true
	}
	return g.computedOutput
}

func (g *lshiftGate) OverrideOutput(output uint16) {
	g.override = true
	g.overrideValue = output
}

func (g *lshiftGate) SetInputs(inputs []gate) {
	g.input = inputs[0]
}

type rshiftGate struct {
	input                                      gate
	shiftAmount, computedOutput, overrideValue uint16
	outputName                                 string
	computed, override                         bool
}

func (g *rshiftGate) OutputName() string {
	return g.outputName
}

func (g *rshiftGate) OutputValue() uint16 {
	if g.override {
		return g.overrideValue
	}
	if !g.computed {
		g.computedOutput = g.input.OutputValue() >> g.shiftAmount
		g.computed = true
	}
	return g.computedOutput
}

func (g *rshiftGate) OverrideOutput(output uint16) {
	g.override = true
	g.overrideValue = output
}

func (g *rshiftGate) SetInputs(inputs []gate) {
	g.input = inputs[0]
}

type andGate struct {
	input1, input2                gate
	outputName                    string
	computedOutput, overrideValue uint16
	computed, override            bool
}

func (g *andGate) OutputName() string {
	return g.outputName
}

func (g *andGate) OutputValue() uint16 {
	if g.override {
		return g.overrideValue
	}
	if !g.computed {
		g.computedOutput = g.input1.OutputValue() & g.input2.OutputValue()
		g.computed = true
	}
	return g.computedOutput
}

func (g *andGate) OverrideOutput(output uint16) {
	g.override = true
	g.overrideValue = output
}

func (g *andGate) SetInputs(inputs []gate) {
	g.input1 = inputs[0]
	g.input2 = inputs[1]
}

type numericAndGate struct {
	inputValue, computedOutput, overrideValue uint16
	inputGate                                 gate
	outputName                                string
	computed, override                        bool
}

func (g *numericAndGate) OutputName() string {
	return g.outputName
}

func (g *numericAndGate) OutputValue() uint16 {
	if g.override {
		return g.overrideValue
	}
	if !g.computed {
		g.computedOutput = g.inputValue & g.inputGate.OutputValue()
		g.computed = true
	}
	return g.computedOutput
}

func (g *numericAndGate) OverrideOutput(output uint16) {
	g.override = true
	g.overrideValue = output
}

func (g *numericAndGate) SetInputs(inputs []gate) {
	g.inputGate = inputs[0]
}

type orGate struct {
	input1, input2                gate
	outputName                    string
	computedOutput, overrideValue uint16
	computed, override            bool
}

func (g *orGate) OutputName() string {
	return g.outputName
}

func (g *orGate) OutputValue() uint16 {
	if g.override {
		return g.overrideValue
	}
	if !g.computed {
		g.computedOutput = g.input1.OutputValue() | g.input2.OutputValue()
		g.computed = true
	}
	return g.computedOutput
}

func (g *orGate) OverrideOutput(output uint16) {
	g.override = true
	g.overrideValue = output
}

func (g *orGate) SetInputs(inputs []gate) {
	g.input1 = inputs[0]
	g.input2 = inputs[1]
}

type notGate struct {
	input                         gate
	outputName                    string
	computedOutput, overrideValue uint16
	computed, override            bool
}

func (g *notGate) OutputName() string {
	return g.outputName
}

func (g *notGate) OutputValue() uint16 {
	if g.override {
		return g.overrideValue
	}
	if !g.computed {
		g.computedOutput = ^g.input.OutputValue()
		g.computed = true
	}
	return g.computedOutput
}

func (g *notGate) OverrideOutput(output uint16) {
	g.override = true
	g.overrideValue = output
}

func (g *notGate) SetInputs(inputs []gate) {
	g.input = inputs[0]
}

func strToUint16(s string) uint16 {
	val, _ := strconv.Atoi(s)
	return uint16(val)
}

func loadGates(filename string) []gate {
	sourceRe := regexp.MustCompile(`^(\d+) -> (.*)$`)
	passthroughRe := regexp.MustCompile(`^([[:alpha:]]+) -> ([[:alpha:]]+)$`)
	lshiftRe := regexp.MustCompile(`^([[:alpha:]]+) LSHIFT (\d+) -> ([[:alpha:]]+)$`)
	rshiftRe := regexp.MustCompile(`^([[:alpha:]]+) RSHIFT (\d+) -> ([[:alpha:]]+)$`)
	andRe := regexp.MustCompile(`^([[:alpha:]]+) AND ([[:alpha:]]+) -> ([[:alpha:]]+)$`)
	numericAndRe := regexp.MustCompile(`^(\d+) AND ([[:alpha:]]+) -> ([[:alpha:]]+)$`)
	orRe := regexp.MustCompile(`^([[:alpha:]]+) OR ([[:alpha:]]+) -> ([[:alpha:]]+)$`)
	notRe := regexp.MustCompile(`^NOT ([[:alpha:]]+) -> ([[:alpha:]]+)$`)

	lines := utils.ReadLines(filename)
	gates := make([]gate, 0)
	sourceToGate := make(map[string]gate)
	gateToRequiredInputs := make(map[gate][]string)
	for _, line := range lines {
		if sourceRe.MatchString(line) {
			match := sourceRe.FindStringSubmatch(line)
			wire := new(sourceGate)
			wire.input = strToUint16(match[1])
			wire.outputName = match[2]
			gates = append(gates, wire)
			sourceToGate[wire.outputName] = wire
		} else if passthroughRe.MatchString(line) {
			match := passthroughRe.FindStringSubmatch(line)
			passthrough := new(passthroughGate)
			passthrough.outputName = match[2]
			gates = append(gates, passthrough)
			sourceToGate[passthrough.outputName] = passthrough
			gateToRequiredInputs[passthrough] = []string{match[1]}
		} else if lshiftRe.MatchString(line) {
			match := lshiftRe.FindStringSubmatch(line)
			gate := new(lshiftGate)
			gate.shiftAmount = strToUint16(match[2])
			gate.outputName = match[3]
			gates = append(gates, gate)
			sourceToGate[gate.outputName] = gate
			gateToRequiredInputs[gate] = []string{match[1]}
		} else if rshiftRe.MatchString(line) {
			match := rshiftRe.FindStringSubmatch(line)
			gate := new(rshiftGate)
			gate.shiftAmount = strToUint16(match[2])
			gate.outputName = match[3]
			gates = append(gates, gate)
			sourceToGate[gate.outputName] = gate
			gateToRequiredInputs[gate] = []string{match[1]}
		} else if andRe.MatchString(line) {
			match := andRe.FindStringSubmatch(line)
			gate := new(andGate)
			gate.outputName = match[3]
			gates = append(gates, gate)
			sourceToGate[gate.outputName] = gate
			gateToRequiredInputs[gate] = []string{match[1], match[2]}
		} else if numericAndRe.MatchString(line) {
			match := numericAndRe.FindStringSubmatch(line)
			gate := new(numericAndGate)
			gate.inputValue = strToUint16(match[1])
			gate.outputName = match[3]
			gates = append(gates, gate)
			sourceToGate[gate.outputName] = gate
			gateToRequiredInputs[gate] = []string{match[2]}
		} else if orRe.MatchString(line) {
			match := orRe.FindStringSubmatch(line)
			gate := new(orGate)
			gate.outputName = match[3]
			gates = append(gates, gate)
			sourceToGate[gate.outputName] = gate
			gateToRequiredInputs[gate] = []string{match[1], match[2]}
		} else if notRe.MatchString(line) {
			match := notRe.FindStringSubmatch(line)
			gate := new(notGate)
			gate.outputName = match[2]
			gates = append(gates, gate)
			sourceToGate[gate.outputName] = gate
			gateToRequiredInputs[gate] = []string{match[1]}
		} else {
			panic("Unknown gate input format")
		}
	}

	for curGate, requiredInputs := range gateToRequiredInputs {
		inputs := make([]gate, 0)
		for _, req := range requiredInputs {
			inputs = append(inputs, sourceToGate[req])
		}
		curGate.SetInputs(inputs)
	}

	return gates
}

func gateByName(gates []gate, name string) gate {
	for i := range gates {
		if strings.Compare(gates[i].OutputName(), name) == 0 {
			return gates[i]
		}
	}
	return nil
}

func main() {
	var override uint16
	{
		gates := loadGates("input.txt")
		override = gateByName(gates, "a").OutputValue()
		fmt.Printf("Part 1: gate a = %d\n", override)
	}
	{
		gates := loadGates("input.txt")
		gateB := gateByName(gates, "b")
		gateB.OverrideOutput(override)
		fmt.Printf("Part 2: gate a = %d\n", gateByName(gates, "a").OutputValue())
	}
}
