package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

func loadGraph(filename string) *utils.Graph {
	re := regexp.MustCompile(`([[:alpha:]]+) to ([[:alpha:]]+) = (\d+)`)
	lines := utils.ReadLines(filename)
	graph := new(utils.Graph)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		start := match[1]
		end := match[2]
		distance, _ := strconv.Atoi(match[3])

		startNode := graph.GetNode(start)
		if startNode == nil {
			startNode = new(utils.Node)
			startNode.SetName(start)
			graph.AddNode(startNode)
		}
		endNode := graph.GetNode(end)
		if endNode == nil {
			endNode = new(utils.Node)
			endNode.SetName(end)
			graph.AddNode(endNode)
		}

		graph.AddEdge(startNode, endNode, distance)
		graph.AddEdge(endNode, startNode, distance)
	}
	return graph
}

func nodePermutations(graph *utils.Graph) [][]*utils.Node {
	nodes := graph.GetNodes()
	interfaces := make([]interface{}, len(nodes))
	for i, val := range nodes {
		interfaces[i] = val
	}
	interfacePermutations := utils.Permutations(interfaces)
	nodePermutations := make([][]*utils.Node, len(interfacePermutations))
	for i, permutation := range interfacePermutations {
		nodePermutations[i] = make([]*utils.Node, len(permutation))
		for j, val := range permutation {
			nodePermutations[i][j] = val.(*utils.Node)
		}
	}
	return nodePermutations
}

func shortestPath(graph *utils.Graph) int {
	permutations := nodePermutations(graph)
	shortest := math.MaxInt32
	for i := range permutations {
		permutation := permutations[i]
		isConnected := true
		length := 0
		for i := 1; i < len(permutation); i++ {
			if !graph.IsNeighbor(permutation[i-1], permutation[i]) {
				isConnected = false
				break
			}
			length += graph.Weight(permutation[i-1], permutation[i])
		}
		if isConnected && length < shortest {
			shortest = length
		}
	}
	return shortest
}

func longestPath(graph *utils.Graph) int {
	permutations := nodePermutations(graph)
	longest := math.MinInt32
	for i := range permutations {
		permutation := permutations[i]
		isConnected := true
		length := 0
		for i := 1; i < len(permutation); i++ {
			if !graph.IsNeighbor(permutation[i-1], permutation[i]) {
				isConnected = false
				break
			}
			length += graph.Weight(permutation[i-1], permutation[i])
		}
		if isConnected && length > longest {
			longest = length
		}
	}
	return longest
}

func main() {
	graph := loadGraph("input.txt")
	fmt.Printf("Part 1: %d\n", shortestPath(graph))
	fmt.Printf("Part 2: %d\n", longestPath(graph))
}
