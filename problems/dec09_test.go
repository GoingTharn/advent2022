package problems

import (
	"fmt"
	"testing"
)

func TestAddGrid(t *testing.T) {
	var g grid = make([]*node, 0)
	nN := getNode(&g, 0, 0)
	if len(g) != 1 {
		fmt.Printf("Node %s not added\n", (*nN))
		t.Fail()
	}

	nN = getNode(&g, 0, 0)
	if len(g) != 1 {
		fmt.Printf("Node %s not added\n", (*nN))
		t.Fail()
	}

	if (*nN).x != 0 && (*nN).y != 0 {
		fmt.Printf("Wrong Node recieved. Got: %s Expected: (0, 0)\n", (*nN))
		t.Fail()
	}

	nN = getNode(&g, 2, 3)
	if len(g) != 2 {
		fmt.Printf("Node %s not added\n", (*nN))
		t.Fail()
	}

	if (*nN).x != 2 && (*nN).y != 3 {
		fmt.Printf("Wrong Node recieved. Got: %s Expected: (2, 3)\n", (*nN))
		t.Fail()
	}
}

func TestMoveRope(t *testing.T) {
	var n node = node{x: 0, y: 0, visited: true}
	tl := &n
	var h head = head{position: &n, t: tl}
	var g grid = make([]*node, 0)

	h = h.moveStep(&g, "R")
	if (*h.position).x != 1 {
		fmt.Printf("Head didn't move as expected! %s\n", (*h.position))
		t.Fail()
	}
	if (*(h.t)).x != 0 {
		fmt.Println("Tail moved!")
		t.Fail()
	}
	if (*h.t).visited != true {
		fmt.Printf("Tail didn't visit the node %s!\n", (*h.t))
		t.Fail()
	}
	h = h.moveStep(&g, "U")
	if (*h.position).y != 1 && (*h.position).x == 1 {
		fmt.Printf("Head didn't move as expected! %s\n", (*h.position))
		t.Fail()
	}
	if (*h.t).x != 0 {
		fmt.Println("Tail moved!")
		t.Fail()
	}
	h = h.moveStep(&g, "U")
	if (*h.position).y != 2 && (*h.position).x == 1 {
		fmt.Println("Head didn't move as expected")
		t.Fail()
	}
	if (*h.t).x != 1 && (*h.t).y != 1 {
		fmt.Printf("Tail didn't move to the right place! Expected: 1,1 Got: %s\n", (*h.t))
		t.Fail()
	}

	if (*h.t).visited != true {
		fmt.Printf("Tail didn't visit the node %s!\n", (*h.t))
		t.Fail()
	}
	h = h.moveStep(&g, "U")
	if (*h.position).y != 3 && (*h.position).x == 1 {
		fmt.Println("Head didn't move as expected")
		t.Fail()
	}
	if (*h.t).x != 1 && (*h.t).y != 2 {
		fmt.Printf("Tail didn't move to the right place! Expected: 1,1 Got: %s\n", (*h.t))
		t.Fail()
	}
	h = h.moveStep(&g, "U")
	if (*h.position).y != 4 && (*h.position).x == 1 {
		fmt.Println("Head didn't move as expected")
		t.Fail()
	}
	if (*h.t).x != 1 && (*h.t).y != 3 {
		fmt.Printf("Tail didn't move to the right place! Expected: 1,1 Got: %s\n", (*h.t))
		t.Fail()
	}
	h = h.moveStep(&g, "U")
	if (*h.position).y != 5 && (*h.position).x == 1 {
		fmt.Println("Head didn't move as expected")
		t.Fail()
	}
	if (*h.t).x != 1 && (*h.t).y != 4 {
		fmt.Printf("Tail didn't move to the right place! Expected: 1,1 Got: %s\n", (*h.t))
		t.Fail()
	}
}
func TestFullProblem(t *testing.T) {
	raw := []byte(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`)
	count := Process9(raw)
	if count != 13 {
		fmt.Printf("Count wasn't right. Wanted 13, got %d\n", count)
		t.Fail()
	}
}
