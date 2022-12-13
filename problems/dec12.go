package problems

import (
	"fmt"
	"strings"

	"github.com/yourbasic/graph"
)

type ElevationMap map[int][]elevationNode

type elevationNode struct {
	x, y    int
	value   string
	ele     int
	address int
}

func makeCharMap() map[string]int {
	myMap := make(map[string]int)
	for i, r := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		myMap[r] = i + 1
	}
	myMap["S"] = 1
	myMap["E"] = 26
	return myMap
}

func getInt(a string, charMap map[string]int) int {
	return charMap[a]
}

func buildArray(raw []byte) (output ElevationMap, vertexCount int, S elevationNode, E elevationNode) {
	charMap := makeCharMap()
	output = make(ElevationMap)
	i := 0
	for y, line := range strings.Split(string(raw), "\n") {
		if line == "" {
			continue
		}
		output[y] = make([]elevationNode, 0)
		for x, r := range strings.Split(line, "") {
			newEN := elevationNode{x: x, y: y, value: r, ele: charMap[r], address: i}
			output[y] = append(output[y], newEN)
			i++
			if r == "S" {
				S = newEN
			} else if r == "E" {
				E = newEN
			}
		}
	}
	return output, i, S, E
}

func buildGraph(em ElevationMap, vertexCount int) (eg *graph.Mutable) {
	eg = graph.New(vertexCount + 1)
	for _, line := range em {
		for _, val := range line {
			neighbors := checkNeighbors(val, em)
			for _, neighbor := range neighbors {
				eg.AddCost(val.address, neighbor.address, 1)
			}
		}
	}
	return eg
}

func checkNeighbors(en elevationNode, em ElevationMap) []elevationNode {
	neighbors := make([]elevationNode, 0)
	// left, right
	lnX := en.x - 1
	if lnX >= 0 {
		leftNeighbor := em[en.y][lnX]
		if leftNeighbor.ele-1 <= en.ele {
			neighbors = append(neighbors, leftNeighbor)
		}
	}
	rnX := en.x + 1
	if rnX < len(em[en.y]) {
		rightNeighbor := em[en.y][rnX]
		if rightNeighbor.ele-1 <= en.ele {
			neighbors = append(neighbors, rightNeighbor)
		}
	}
	dnY := en.y - 1
	if dnY >= 0 {
		downNeighbor := em[dnY][en.x]
		if downNeighbor.ele-1 <= en.ele {
			neighbors = append(neighbors, downNeighbor)
		}
	}
	unY := en.y + 1
	if unY < len(em) {
		upNeighbor := em[unY][en.x]
		if upNeighbor.ele-1 <= en.ele {
			neighbors = append(neighbors, upNeighbor)
		}
	}
	return neighbors
}

func getShortest(raw []byte) int64 {
	em, vertices, _, E := buildArray(raw)

	allAs := make([]elevationNode, 0)
	for _, line := range em {
		for _, v := range line {
			if strings.ContainsAny(v.value, "Sa") == true {
				allAs = append(allAs, v)
			}
		}
	}
	eg := buildGraph(em, vertices)
	shortest := int64(490)
	for _, a := range allAs {
		_, dist := graph.ShortestPath(eg, a.address, E.address)
		if dist <= shortest && dist != int64(-1) {
			shortest = dist
		}
	}
	return shortest
}

func Problem12() {
	raw, _ := getInput("input/problem12.txt")
	dist := getShortest(raw)
	fmt.Printf("Shortest Path: %d", dist)
}
