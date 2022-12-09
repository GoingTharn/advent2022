package problems

import (
	"fmt"
	"testing"
)

func TestProblem8(t *testing.T) {
	testInput := ProblemInput(`30373
25512
65332
33549
35390
`)

	tg := BuildTreeGrid(testInput)
	if tg[0][2].Height != 3 {
		fmt.Println(tg)
		t.Fail()
	}
	if tg[2][0].Height != 6 {
		fmt.Println(tg)
		t.Fail()
	}

	// check edge up
	y, x := 1, 2
	actual := (&tg).CheckUp(y, x)
	if actual != 1 {
		fmt.Println(tg)
		fmt.Printf("Looking up got: %d Expected: 1\n", actual)
		t.Fail()
	}
	actual = (&tg).CheckDown(y, x)
	if actual != 2 {
		fmt.Println(tg)
		fmt.Printf("Looking down got: %d Expected: 2\n", actual)
		t.Fail()
	}
	actual = (&tg).CheckLeft(y, x)
	if actual != 1 {
		fmt.Println(tg)
		fmt.Printf("Looking left got: %d Expected: 1\n", actual)
		t.Fail()
	}
	actual = (&tg).CheckRight(y, x)
	if actual != 2 {
		fmt.Println(tg)
		fmt.Printf("Looking right got: %d Expected: 2\n", actual)
		t.Fail()
	}

	y, x = 3, 2
	actual = (&tg).CheckUp(y, x)
	if actual != 2 {
		fmt.Println(tg)
		fmt.Printf("Looking up got: %d Expected: 2\n", actual)
		t.Fail()
	}
	actual = (&tg).CheckDown(y, x)
	if actual != 1 {
		fmt.Println(tg)
		fmt.Printf("Looking down got: %d Expected: 1\n", actual)
		t.Fail()
	}
	actual = (&tg).CheckLeft(y, x)
	if actual != 2 {
		fmt.Println(tg)
		fmt.Printf("Looking left got: %d Expected: 2\n", actual)
		t.Fail()
	}
	actual = (&tg).CheckRight(y, x)
	if actual != 2 {
		fmt.Println(tg)
		fmt.Printf("Looking right got: %d Expected: 2\n", actual)
		t.Fail()
	}

	totVis := tg.MaxScenicView()
	if totVis != 8 {
		fmt.Printf("Wrong number of trees. Got: %d Wanted: 21", totVis)
		t.Fail()
	}
	if tg[1][2].ScenicScore != 4 {
		fmt.Println("got wrong scenic score for 2,1")
		fmt.Println(tg[2][1].ScenicScore)
		t.Fail()
	}
}
