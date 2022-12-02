package problems

import (
	"bytes"
	"fmt"
	"log"
	"runtime/debug"
	"sort"
	"strconv"
)

type elf struct {
	calories int
}

func Problem1() {

	path := "input/problem1.txt"
	content, err := getInput(path)
	if err != nil {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		log.Output(2, trace)
		return
	}
	//content = testInput()
	elves := parseInput(content)
	sort.Slice(elves, func(i, j int) bool { return elves[i].calories > elves[j].calories })
	fmt.Println("Top 3 calories:")
	total := 0
	for i := 0; i < 3; i++ {
		total += elves[i].calories
	}
	fmt.Println(total)

}

func testInput() []byte {
	test_input := []byte("100\n200\n300\n\n100\n\n900")
	return test_input
}

func parseInput(contents []byte) (elves []elf) {
	data := bytes.Split(contents, []byte("\n"))
	var working elf
	working = elf{0}

	for i := 0; i < len(data); i++ {
		cal_string := string(data[i])
		calories, err := strconv.Atoi(cal_string)
		if err != nil {
			elves = append(elves, working)
			fmt.Println("New elf!!")
			working = elf{0}
			continue
		}

		working.calories += calories
		fmt.Print("Calories:")
		fmt.Println(working.calories)
		if i == len(data)-1 {
			fmt.Println("Last elf!!")
			elves = append(elves, working)

		}
	}

	return elves
}
