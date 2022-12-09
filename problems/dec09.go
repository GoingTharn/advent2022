package problems

import (
	"fmt"
	"strconv"
	"strings"
)

type tail *node

type head struct {
	t        tail
	position *node
}

type node struct {
	x       int
	y       int
	visited bool
}

func (n node) String() string {
	return fmt.Sprintf("(%d, %d, v: %t)", n.x, n.y, n.visited)
}

type grid []*node

func (g grid) String() string {
	builder := strings.Builder{}

	for i, n := range g {
		builder.WriteString(fmt.Sprintf("(%s)", (*n)))
		if i != len(g)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}

func getNode(g *grid, x int, y int) *node {
	for i, n := range *g {
		if n.x == x && n.y == y {
			return (*g)[i]
		}
	}
	// no nodes found, create one
	newNode := node{x: x, y: y}
	(*g) = append((*g), &newNode)
	return &newNode
}

func parseCoords(dir string) (x_or_y string, step int) {
	switch dir {
	case "R":
		x_or_y = "x"
		step = 1
	case "L":
		x_or_y = "x"
		step = -1
	case "U":
		x_or_y = "y"
		step = 1
	case "D":
		x_or_y = "y"
		step = -1
	}
	return x_or_y, step
}

func (h head) moveStep(g *grid, dir string) head {
	var newNode *node

	x_or_y, step := parseCoords(dir)
	oldPosition := (*&h.position)

	switch x_or_y {
	case "x":
		newNode = getNode(g, oldPosition.x+step, oldPosition.y)
	case "y":
		newNode = getNode(g, oldPosition.x, oldPosition.y+step)

	}
	h.position = newNode
	newTail := h.moveTail(g, dir)
	h.t = &newTail
	return h
}

func (h head) moveTail(g *grid, x_or_y string) node {
	tailNode := (*h.t)
	headNode := (*h.position)

	new_x, new_y := tailNode.x, tailNode.y
	if headNode.x-tailNode.x > 1 {
		new_x++
		new_y = headNode.y
	} else if tailNode.x-headNode.x > 1 {
		new_x--
		new_y = headNode.y
	} else if headNode.y-tailNode.y > 1 {
		new_y++
		new_x = headNode.x
	} else if tailNode.y-headNode.y > 1 {
		new_y--
		new_x = headNode.x
	}
	newTail := getNode(g, new_x, new_y)
	(*newTail).visited = true
	return *newTail
}

func Problem9() {
	raw, _ := getInput("input/problem9.txt")
	count := Process9(raw)
	fmt.Printf("Visited Nodes: %d", count)
}

func Process9(raw []byte) int {
	var n node = node{x: 0, y: 0, visited: true}
	tl := &n
	var h head = head{position: &n, t: tl}
	var g grid = make([]*node, 0)

	for _, line := range strings.Split(string(raw), "\n") {
		if len(line) != 0 {
			directionSet := strings.Split(line, " ")

			magnitude, _ := strconv.Atoi(directionSet[1])

			for i := 0; i < magnitude; i++ {
				h = h.moveStep(&g, directionSet[0])
			}
		}
	}
	count := 0
	for _, n := range g {
		if (*n).visited == true {
			count++
		}
	}
	return count
}
