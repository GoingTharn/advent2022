package problems

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type segment struct {
	t        *segment
	position *node
}

func (s segment) String() string {
	return fmt.Sprintf("(%s) - myAddress: %p tailAddress: %p", (*s.position), &s, s.t)
}

type node struct {
	x int
	y int
}

func (n node) String() string {
	return fmt.Sprintf("(%d, %d)", n.x, n.y)
}

type grid map[node]bool

func (g grid) String() string {
	builder := strings.Builder{}

	for k, v := range g {
		builder.WriteString(fmt.Sprintf("(%s) Visited: %t\n", k, v))
	}
	return builder.String()
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

func (h segment) moveStep(g *grid, dir string) segment {
	var newNode node

	x_or_y, step := parseCoords(dir)
	oldPosition := (*&h.position)

	switch x_or_y {
	case "x":
		newNode = node{x: oldPosition.x + step, y: oldPosition.y}
		(*g)[newNode] = (*g)[newNode]
	case "y":
		newNode = node{x: oldPosition.x, y: oldPosition.y + step}
		(*g)[newNode] = (*g)[newNode]

	}
	h.position = &newNode
	_ = h.moveTail(g)
	return h
}

func (h *segment) moveTail(g *grid) *segment {
	//	if h.t == nil {
	//		err := errors.New("No more segments")
	//		return h, err
	//	}
	leadS := (*h)
	trailS := (*h.t)
	trailingNode := *(trailS.position)
	currentNode := *(leadS.position)

	new_x, new_y := trailingNode.x, trailingNode.y
	if touching(trailingNode, currentNode) == false {
		move_x, move_y := getMove(currentNode, trailingNode)
		new_x += move_x
		new_y += move_y
	}
	newPosition := node{x: new_x, y: new_y}
	(*h.t).position = &newPosition
	if trailS.t == nil {
		(*g)[(*(*h.t).position)] = true
		return h.t
	}

	return (*h.t).moveTail(g)
}

func getMove(h, t node) (x, y int) {
	// only X
	if h.x-t.x > 1 && h.y == t.y {
		x = 1
		return
	} else if h.x-t.x < -1 && h.y == t.y {
		x = -1
		return
	}
	// only y
	if h.y-t.y > 1 && h.x == t.x {
		y = 1
		return
	} else if h.y-t.y < -1 && h.x == t.x {
		y = -1
		return
	}

	if h.y > t.y {
		y = 1
	} else {
		y = -1
	}
	if h.x > t.x {
		x = 1
	} else {
		x = -1
	}
	return x, y
}

func touching(a, b node) bool {
	return math.Abs(float64(a.x-b.x)) <= 1 && math.Abs(float64(a.y-b.y)) <= 1
}

func printGrid(g grid, h segment) string {
	var min_x, min_y, max_x, max_y int
	min_x, min_y, max_x, max_y = getSnakePadding(h)

	fullBuilder := strings.Builder{}
	for i := min_y; i <= max_y; i++ {
		builder := strings.Builder{}
		for j := min_x; j <= max_x; j++ {
			builder.WriteString(checkSnake(h, 0, i, j, "."))
		}
		builder.WriteString("\n")
		fullBuilder.WriteString(builder.String())
	}
	return fullBuilder.String()
}

func getSnakePadding(h segment) (min_x, min_y, max_x, max_y int) {
	for {
		n := h.position
		if (*n).x < min_x {
			min_x = (*n).x
		}
		if (*n).x > max_x {
			max_x = (*n).x
		}
		if (*n).y < min_y {
			min_y = (*n).y
		}
		if (*n).y > max_y {
			max_y = (*n).y
		}
		if h.t == nil {
			return
		}
		h = (*h.t)
	}
}

func checkSnake(h segment, i int, y int, x int, curr string) string {
	if (h.position).x == x && (h.position).y == y {
		if i == 0 {
			curr = "H"
		} else {
			curr = fmt.Sprintf("%d", i)
		}
	}
	if h.t == nil {
		return curr
	}
	i++
	return checkSnake((*h.t), i, y, x, curr)
}

func Problem9() {
	raw, _ := getInput("input/problem9.txt")
	count := Process9(raw)
	fmt.Printf("Visited Nodes: %d", count)
}

func buildSnake(length int) segment {
	fmt.Println("Building Snake!")
	var n node = node{x: 0, y: 0}
	var tail segment = segment{position: &n}
	var s *segment

	s = addSegment(&tail)
	for i := 2; i < length; i++ {
		s = addSegment(s)
	}

	fmt.Println("Snake Built!")
	return *s
}

func addSegment(t *segment) *segment {
	s := segment{position: (*t).position, t: t}
	return &s
}

func Process9(raw []byte) int {
	var g grid = make(map[node]bool)
	h := buildSnake(10)
	g[(*h.position)] = true
	for _, line := range strings.Split(string(raw), "\n") {
		if len(line) != 0 {
			directionSet := strings.Split(line, " ")
			magnitude, _ := strconv.Atoi(directionSet[1])
			for i := 0; i < magnitude; i++ {
				h = h.moveStep(&g, directionSet[0])
			}
		}
	}
	count := countGrid(g)
	fmt.Printf("Grid Size: %d", len(g))
	return count
}

func countGrid(g grid) (count int) {
	for _, val := range g {
		if val == true {
			count++
		}
	}
	return count
}
