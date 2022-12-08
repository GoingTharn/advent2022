package problems_test

import (
	"fmt"
	"testing"

	"com.advent2022.goingtharn/problems"
)

func TestProblem8(t *testing.T) {
	testInput := problems.ProblemInput(`30373
25512
65332
33549
35390
`)

	tg := problems.BuildTreeGrid(testInput)
	fmt.Println(tg)
	if tg[0][2].Height != 3 {
		fmt.Println(tg)
		t.Fail()
	}
	if tg[2][0].Height != 6 {
		fmt.Println(tg)
		t.Fail()
	}

	// check edge up
	y, x := 4, 2
	(&tg).CheckUp(y, x)
	if tg[y][x].Visible != true {
		fmt.Println(tg)
		fmt.Println("Edge not marked visible")
		fmt.Println(tg[y][x].Print())
		t.Fail()
	}
	y, x = 1, 2
	(&tg).CheckUp(y, x)
	if tg[y][x].Visible != true {
		fmt.Println(tg)
		fmt.Printf("tree (%d, %d) should have been true but isn't\n", y, x)
		fmt.Println(tg[y][x].Print())
		t.Fail()
	}

	y, x = 3, 2
	(&tg).CheckUp(y, x)
	if tg[y][x].Visible != false {
		fmt.Println(tg)
		fmt.Printf("tree (%d, %d) should have been false but isn't\n", y, x)
		fmt.Println(tg[y][x].Print())
		t.Fail()
	}

	y, x = 0, 4
	(&tg).CheckUp(y, x)
	if tg[y][x].Visible != true {
		fmt.Println(tg)
		fmt.Printf("tree (%d, %d) should have been true but isn't\n", y, x)
		fmt.Println(tg[y][x].Print())
		t.Fail()
	}

	// check edge down
	y, x = 0, 2
	(&tg).CheckDown(y, x)
	if tg[y][x].Visible != true {
		fmt.Println(tg)
		fmt.Println("Edge not marked visible")
		fmt.Println(tg[y][x].Print())
		t.Fail()
	}
	y, x = 3, 2
	(&tg).CheckDown(y, x)
	if tg[y][x].Visible != true {
		fmt.Println(tg)
		fmt.Printf("tree (%d, %d) should have been true from CheckUp\n", y, x)
		fmt.Println(tg[y][x].Print())
		t.Fail()
	}

	y, x = 4, 2
	(&tg).CheckDown(y, x)
	if tg[y][x].Visible != true {
		fmt.Println(tg)
		fmt.Printf("tree (%d, %d) should have been true but isn't\n", y, x)
		fmt.Println(tg[y][x].Print())
		t.Fail()
	}

	y, x = 3, 4
	(&tg).CheckLeft(y, x)
	if tg[y][x].Visible != true {
		fmt.Println(tg)
		fmt.Printf("Edge tree (%d, %d) should have been true but isn't\n", y, x)
		fmt.Println(tg[y][x].Print())
		t.Fail()
	}

	y, x = 1, 3
	(&tg).CheckLeft(y, x)
	if tg[y][x].Visible != false {
		fmt.Println(tg)
		fmt.Printf("tree (%d, %d) should have been false but isn't\n", y, x)
		fmt.Println(tg[y][x].Print())
		t.Fail()
	}

	y, x = 0, 3
	(&tg).CheckRight(y, x)
	if tg[y][x].Visible != true {
		fmt.Println(tg)
		fmt.Printf("tree (%d, %d) should have been true but isn't\n", y, x)
		fmt.Println(tg[y][x].Print())
		t.Fail()
	}

	y, x = 4, 0
	(&tg).CheckRight(y, x)
	if tg[y][x].Visible != true {
		fmt.Println(tg)
		fmt.Printf("Edge tree (%d, %d) should have been true but isn't\n", y, x)
		fmt.Println(tg[y][x].Print())
		t.Fail()
	}

	y, x = 3, 2
	tg[y][x].Visible = false
	(&tg).CheckRight(y, x)
	if tg[y][x].Visible != false {
		fmt.Println(tg)
		fmt.Printf("tree (%d, %d) should have been false but isn't\n", y, x)
		fmt.Println(tg[y][x].Print())
		t.Fail()
	}

	totVis := tg.VisibleTrees()
	if totVis != 21 {
		fmt.Printf("Wrong number of trees. Got: %d Wanted: 21", totVis)
		t.Fail()
	}

}
