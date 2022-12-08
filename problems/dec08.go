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
	Height  int
	Visible bool
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
			fmt.Println("Empty line")
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
		for j, tree := range treeList {
			tg.CheckUp(i, j)
			if tree.Visible == true {
				continue
			}
			tg.CheckDown(i, j)
			if tree.Visible == true {
				continue
			}
			tg.CheckLeft(i, j)
			if tree.Visible == true {
				continue
			}
			tg.CheckRight(i, j)
			if tree.Visible == true {
				continue
			}
		}
	}
}

func (tg *TreeGrid) CheckUp(y, x int) {
	candidate := (*tg)[y][x]
	//fmt.Printf("Candidate: %s (%d, %d)\n", candidate.Print(), y, x)
	if y == len((*tg))-1 {
		// edge, always visible
		candidate.Visible = true
		(*tg)[y][x] = candidate
		return
	}
	max := 0
	for i := y - 1; i >= 0; i-- {
		if max < (*tg)[i][x].Height {
			max = (*tg)[i][x].Height
		}
	}
	if candidate.Height > max {
		candidate.Visible = true
		(*tg)[y][x] = candidate
		return
	}
}

func (tg *TreeGrid) CheckDown(y, x int) {
	candidate := (*tg)[y][x]
	//fmt.Printf("Candidate: %s\n", candidate.Print())
	if y == 0 {
		// edge, always visible
		candidate.Visible = true
		(*tg)[y][x] = candidate
		return
	}
	max := 0
	for i := y + 1; i < len((*tg)); i++ {
		if max < (*tg)[i][x].Height {
			max = (*tg)[i][x].Height
		}
	}
	if candidate.Height > max {
		candidate.Visible = true
		(*tg)[y][x] = candidate
		return
	}

}

func (tg *TreeGrid) CheckLeft(y, x int) {
	candidate := (*tg)[y][x]
	//fmt.Printf("Candidate: %s\n", candidate.Print())
	if x == len((*tg)[y])-1 {
		// edge, always visible
		candidate.Visible = true
		(*tg)[y][x] = candidate
		return
	}
	max := 0
	for i := x - 1; i >= 0; i-- {
		//fmt.Printf("i: %d, y:%d check: %s \n", i, y, (*tg)[i][x].Print())
		if max < (*tg)[y][i].Height {
			max = (*tg)[y][i].Height
		}
	}
	if candidate.Height > max {
		candidate.Visible = true
		(*tg)[y][x] = candidate
		return
	}
}

func (tg *TreeGrid) CheckRight(y, x int) {
	candidate := (*tg)[y][x]
	//fmt.Printf("Candidate: %s\n", candidate.Print())
	if x == 0 {
		// edge, always visible
		candidate.Visible = true
		(*tg)[y][x] = candidate
		return
	}
	max := 0
	for i := x + 1; i < len((*tg)[y]); i++ {
		//fmt.Printf("i: %d, x:%d check: %s \n", i, x, (*tg)[y][i].Print())
		if max < (*tg)[y][i].Height {
			max = (*tg)[y][i].Height
		}
	}
	//fmt.Printf("Max: %d  Candidate.Height: %d\n", max, candidate.Height)
	if candidate.Height > max {
		candidate.Visible = true
		(*tg)[y][x] = candidate
		return
	}
}

func (tg TreeGrid) VisibleTrees() (totalVis int) {
	(&tg).CheckVisibility()
	for _, treeList := range tg {
		for _, tree := range treeList {
			if tree.Visible == true {
				totalVis++
			}
		}
	}
	return totalVis
}

func Problem8() {
	raw, _ := getInput("input/problem8.txt")
	tg := BuildTreeGrid(ProblemInput(raw))
	fmt.Println(tg.VisibleTrees())

}
