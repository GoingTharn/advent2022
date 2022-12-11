package problems

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type TreeGrid [][]Tree

type ProblemInput []byte

type Tree struct {
	Height      int
	Visible     bool
	ScenicScore int
}

func (t Tree) String() string {
	return fmt.Sprintf("%d", t.Height)
}

func (t Tree) Print() string {
	return fmt.Sprintf("Tree height: %d vis: %t", t.Height, t.Visible)
}

func (tg TreeGrid) String() string {
	builder := strings.Builder{}
	for _, treeSlice := range tg {
		for j, tree := range treeSlice {
			builder.WriteString(fmt.Sprint(tree))
			if j < len(treeSlice)-1 {
				builder.WriteString(", ")
			} else {
				builder.WriteString("\n")
			}
		}
	}
	return builder.String()
}

func BuildTreeGrid(p ProblemInput) (tg TreeGrid) {

	var trees []Tree
	lines := strings.Split(string(p), "\n")
	tg = make(TreeGrid, len(lines)-1)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		trees = []Tree{}
		for _, item := range strings.Split(line, "") {
			height, err := strconv.Atoi(item)
			if err != nil {
				log.Printf("Height failed to convert on item: %s", item)
				return tg
			}
			trees = append(trees, Tree{Height: height, Visible: false})
		}
		tg[i] = trees
	}
	return tg
}

func (tg *TreeGrid) CheckVisibility() {
	for i, treeList := range *tg {
		for j := range treeList {
			up := tg.CheckUp(i, j)
			down := tg.CheckDown(i, j)
			left := tg.CheckLeft(i, j)
			right := tg.CheckRight(i, j)
			(*tg)[i][j].ScenicScore = up * down * left * right
		}
	}
}

func (tg *TreeGrid) CheckUp(y, x int) int {
	candidate := (*tg)[y][x]
	//fmt.Printf("Candidate: %s (%d, %d)\n", candidate.Print(), y, x)
	if y == len((*tg))-1 {
		// edge, always visible
		candidate.Visible = true
		(*tg)[y][x] = candidate
		return 0
	}
	for i := y - 1; i >= 0; i-- {
		if candidate.Height <= (*tg)[i][x].Height {
			return y - i
		}
	}
	return y
}

func (tg *TreeGrid) CheckDown(y, x int) int {
	candidate := (*tg)[y][x]
	//fmt.Printf("Candidate: %s\n", candidate.Print())
	if y == 0 {
		// edge, always visible
		candidate.Visible = true
		(*tg)[y][x] = candidate
		return 0
	}
	for i := y + 1; i < len((*tg)); i++ {
		if candidate.Height <= (*tg)[i][x].Height {
			return i - y
		}
	}
	return len((*tg)) - 1 - y
}

func (tg *TreeGrid) CheckLeft(y, x int) int {
	candidate := (*tg)[y][x]
	//fmt.Printf("Candidate: %s\n", candidate.Print())
	if x == len((*tg)[y])-1 {
		// edge, always visible
		candidate.Visible = true
		(*tg)[y][x] = candidate
		return 0
	}
	for i := x - 1; i >= 0; i-- {
		//fmt.Printf("i: %d, y:%d check: %s \n", i, y, (*tg)[i][x].Print())
		if candidate.Height <= (*tg)[y][i].Height {
			return x - i
		}
	}
	return x
}

func (tg *TreeGrid) CheckRight(y, x int) int {
	candidate := (*tg)[y][x]
	//fmt.Printf("Candidate: %s\n", candidate.Print())
	if x == 0 {
		// edge, always visible
		candidate.Visible = true
		(*tg)[y][x] = candidate
		return 0
	}
	for i := x + 1; i < len((*tg)[y]); i++ {
		if candidate.Height <= (*tg)[y][i].Height {
			return i - x
		}
	}
	return len((*tg)[y]) - 1 - x
}

func (tg TreeGrid) MaxScenicView() int {
	tg.CheckVisibility()
	maxVis := 0
	for _, line := range tg {
		for _, tree := range line {
			if maxVis < tree.ScenicScore {
				maxVis = tree.ScenicScore
			}
		}
	}
	return maxVis
}

func Problem8() {
	raw, _ := getInput("input/problem8.txt")
	tg := BuildTreeGrid(ProblemInput(raw))
	fmt.Println(tg.MaxScenicView())

}
