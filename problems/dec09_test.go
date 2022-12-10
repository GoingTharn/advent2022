package problems

import (
	"fmt"
	"testing"
)

func TestMoveRope(t *testing.T) {
	h := buildSnake(2)
	var g grid = make(map[node]bool)

	fmt.Println("before")
	h = h.moveStep(&g, "R")
	fmt.Println("after")
	if (*h.position).x != 1 {
		fmt.Printf("Head didn't move as expected! %s\n", (*h.position))
		t.Fail()
	}
	if (*(h.t)).position.x != 0 {
		fmt.Println("Tail moved!")
		t.Fail()
	}
	if g[(*(*h.t).position)] != true {
		fmt.Printf("Tail didn't visit the node %s!\n", (*h.t).position)
		t.Fail()
	}
	h = h.moveStep(&g, "U")
	if (*h.position).y != 1 && (*h.position).x == 1 {
		fmt.Printf("Head didn't move as expected! %s\n", (*h.position))
		t.Fail()
	}
	if (*h.t).position.x != 0 {
		fmt.Println("Tail moved!")
		t.Fail()
	}
	h = h.moveStep(&g, "U")
	if (*h.position).y != 2 && (*h.position).x == 1 {
		fmt.Println("Head didn't move as expected")
		t.Fail()
	}
	if (*h.t).position.x != 1 && (*h.t).position.y != 1 {
		fmt.Printf("Tail didn't move to the right place! Expected: 1,1 Got: %s\n", (*h.t).position)
		t.Fail()
	}

	if g[(*(*h.t).position)] != true {
		fmt.Printf("Tail didn't visit the node %s!\n", (*h.t).position)
		t.Fail()
	}
	h = h.moveStep(&g, "U")
	if (*h.position).y != 3 && (*h.position).x == 1 {
		fmt.Println("Head didn't move as expected")
		t.Fail()
	}
	if (*h.t).position.x != 1 && (*h.t).position.y != 2 {
		fmt.Printf("Tail didn't move to the right place! Expected: 1,1 Got: %s\n", (*h.t).position)
		t.Fail()
	}
	h = h.moveStep(&g, "U")
	if (*h.position).y != 4 && (*h.position).x == 1 {
		fmt.Println("Head didn't move as expected")
		t.Fail()
	}
	if (*h.t).position.x != 1 && (*h.t).position.y != 3 {
		fmt.Printf("Tail didn't move to the right place! Expected: 1,1 Got: %s\n", (*h.t).position)
		t.Fail()
	}
	h = h.moveStep(&g, "U")
	if (*h.position).y != 5 && (*h.position).x == 1 {
		fmt.Println("Head didn't move as expected")
		t.Fail()
	}
	if (*h.t).position.x != 1 && (*h.t).position.y != 4 {
		fmt.Printf("Tail didn't move to the right place! Expected: 1,1 Got: %s\n", (*h.t).position)
		t.Fail()
	}
}

func TestFullProblem(t *testing.T) {
	raw := []byte(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`)
	count := Process9(raw)
	if count != 36 {
		fmt.Printf("Count wasn't right. Wanted 36, got %d\n", count)
		t.Fail()
	}
}
