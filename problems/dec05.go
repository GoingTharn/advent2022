package problems

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
	"strings"
)

type crates struct {
	stacks []stack
}

type stack []string

type moves []string

type move struct {
	from   int
	to     int
	amount int
}

func (s *stack) pop() (val string, err error) {
	i := len((*s))
	val = (*s)[i-1]
	*s = (*s)[:i-1]
	return val, err
}

func (c crates) String() string {
	var elements []string
	builder := strings.Builder{}
	for i := 0; i < len(c.stacks); i++ {
		builder.WriteString(fmt.Sprintf(" %d ", i+1))

		if i != len(c.stacks)-1 {
			builder.WriteString(" ")
		}
	}
	elements = append(elements, builder.String())
	builder.Reset()

	maxLen := c.maxLength()
	for depth := 0; depth < maxLen; depth++ {
		for width := 0; width < len(c.stacks); width++ {
			//fmt.Printf("Depth: %d Stack: %s Width: %d\n", depth, c.stacks[width], len(c.stacks[width]))
			if len(c.stacks[width]) <= depth {
				builder.WriteString("   ")
			} else {
				builder.WriteString(fmt.Sprintf("%s", c.stacks[width][depth]))
			}

			if width < len(c.stacks)-1 {
				builder.WriteString(" ")
			}
		}
		elements = append(elements, builder.String())
		builder.Reset()
	}
	for i := len(elements) - 1; i >= 0; i-- {
		builder.WriteString(elements[i])
		builder.WriteString("\n")
	}
	return builder.String()
}

func (c crates) maxLength() (length int) {

	for i := 0; i < len(c.stacks); i++ {
		if len(c.stacks[i]) > length {
			length = len(c.stacks[i])
		}
	}
	return
}

func testInput5() []byte {

	return []byte(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`)
}

func Problem5() {
	path := "input/problem5.txt"
	content, err := getInput(path)
	if err != nil {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		log.Output(2, trace)
		return
	}

	//content = testInput5()
	c, mv := parseInput5(content)

	fmt.Println(c)

	for i := 0; i < len(mv); i++ {
		moveCrates(mv[i], &c)
	}
	fmt.Println(c)

	for i := 0; i < len(c.stacks); i++ {
		fmt.Print(string(c.stacks[i][len(c.stacks[i])-1][1]))
	}
}

func moveCrates(r string, c *crates) {
	var m move
	m.Build(r)
	c.Execute(m)
}

func (m *move) Build(raw string) (err error) {
	splits := strings.Split(raw, " ")
	if len(splits) < 6 {
		log.Println("String not long enough!")
		return errors.New(fmt.Sprintf("Split string not long enough for moves: %s", splits))
	}
	m.amount, err = strconv.Atoi(splits[1])
	if err != nil {
		log.Printf("Error!: %s", err)
		return
	}

	m.from, err = strconv.Atoi(splits[3])
	if err != nil {
		log.Printf("Error!: %s", err)
		return
	}

	m.to, err = strconv.Atoi(splits[5])
	if err != nil {
		log.Printf("Error!: %s", err)
		return
	}
	m.from--
	m.to--
	return
}

func (m move) String() string {
	return fmt.Sprintf("Amount: %d, From: %d, To: %d", m.amount, m.from, m.to)
}

func (c *crates) Execute(m move) (err error) {
	fmt.Println(m)
	var temp []string
	for i := 0; i < m.amount; i++ {
		working := c.stacks[m.from]
		box, err := working.pop()
		if err != nil {
			return err
		}
		fmt.Printf("Box: '%s'\n", box)
		c.stacks[m.from] = working
		temp = append(temp, box)
	}
	fmt.Printf("To Append: %s\n", temp)
	for i := len(temp) - 1; i >= 0; i-- {
		c.stacks[m.to] = append(c.stacks[m.to], temp[i])
	}
	fmt.Printf("Stacks: %s\n", c.stacks)
	fmt.Printf("Crates: \n%s", c)
	return
}

func parseInput5(contents []byte) (c crates, m moves) {

	cut := bytes.Split(contents, []byte("\n\n"))
	rawMoves := cut[1]
	m = strings.Split(string(rawMoves), "\n")

	rawCrates := cut[0]
	c = buildCrates(string(rawCrates))
	return c, m
}

func buildCrates(raw string) (c crates) {
	contents := strings.Split(raw, "\n")
	if len(contents) < 2 {
		log.Printf("Contents too small! raw: %s", contents)
		return
	}

	// ensure there's enough array space in c.stacks
	stackLabels := contents[len(contents)-1]
	stacks := strings.Split(strings.TrimSpace(stackLabels), " ")
	stackCount, err := strconv.Atoi(stacks[len(stacks)-1])
	if err != nil {
		log.Println(err)
		return
	}
	c.stacks = make([]stack, stackCount)

	contents = contents[:len(contents)-1]
	// get just the stacks here, in reverse to get the bottom at [0] index
	for i := len(contents) - 1; i >= 0; i-- {
		stackIdx := 0
		for lineIdx := 0; lineIdx < len(contents[i]); lineIdx = lineIdx + 4 {
			working := contents[i][lineIdx : lineIdx+3]
			if working == "   " {
				stackIdx++
				continue
			} else {
				c.stacks[stackIdx] = append(c.stacks[stackIdx], working)
				stackIdx++
			}
		}
	}
	return c
}
