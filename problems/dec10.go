package problems

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Clock struct {
	time             int
	instructionsList []string
	results          map[int]int
	crt              *CRT
}

type CRT struct {
	X       int
	V       int
	fn_tick int
}

func (c CRT) String() string {
	return fmt.Sprintf("CRT(%p): X: %d V: %d fn_tick: %d", &c, c.X, c.V, c.fn_tick)
}

func (c Clock) String() string {
	return fmt.Sprintf("Clock(%p): time: %d instructionsList: %s", &c, c.time, c.instructionsList)
}

func (c *CRT) execute() {
	c.V = c.V + c.X
	c.X = 0
}

func (crt *CRT) Tick(t int) (need_fn bool) {
	if crt.fn_tick <= t {
		crt.execute()
		need_fn = true
	}
	return need_fn
}

func (c *Clock) Tick() {
	if c.time != 0 {
		c.results[c.time] = (*c.crt).V
	}
	need_fn := c.crt.Tick(c.time)
	if need_fn == true {
		err := addFn(c)
		if err != nil {
			fmt.Println(err)
		}
	}
	c.time = c.time + 1
}

func (c Clock) GetResults() (out map[int]int) {
	out = make(map[int]int)
	var interesting []int = []int{20, 60, 100, 140, 180, 220}
	for _, tick := range interesting {
		out[tick] = c.results[tick] * tick
	}
	return out
}

func addFn(c *Clock) (err error) {
	instList := c.instructionsList
	if len(instList) < 1 {
		err = errors.New("InstructionList empty")
		return
	}
	instruction := instList[0]
	if len(instList) > 1 {
		c.instructionsList = c.instructionsList[1:]
	} else {
		c.instructionsList = []string{}
	}

	parsed := strings.Split(instruction, " ")
	if len(parsed[0]) < 1 {
		return
	}
	var x, tick_time int
	if parsed[0] == "noop" {
		tick_time = c.time + 1
		c.crt.X = x
		c.crt.fn_tick = tick_time
	} else if parsed[0] == "addx" {
		x, err := strconv.Atoi(parsed[1])
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return err
		}
		tick_time = c.time + 2
		c.crt.X = x
		c.crt.fn_tick = tick_time
		return nil
	}
	return
}

func buildClock(instructionsList []string) Clock {
	var crt CRT = CRT{V: 1, fn_tick: 0}
	var c Clock = Clock{crt: &crt, time: 0, results: make(map[int]int), instructionsList: instructionsList}
	addFn(&c)
	return c
}

func DrawResults(c Clock) {
	display := strings.Builder{}
	builder := strings.Builder{}
	for i := 1; i < 260; i++ {
		pixel := i % 40
		sprite := []int{pixel - 2, pixel - 1, pixel}
		if slices.Contains(sprite, c.results[i]) == true {
			builder.WriteString("#")
		} else {
			builder.WriteString(".")
		}
		if pixel == 0 {
			display.WriteString(builder.String())
			display.WriteString("\n")
			builder.Reset()
		}
	}
	fmt.Println(display.String())
}

func Problem10() {
	raw, _ := getInput("input/problem10.txt")
	instructionList := strings.Split(strings.Trim(string(raw), "\n"), "\n")
	c := buildClock(instructionList)

	for i := 1; i < 265; i++ {
		c.Tick()
	}

	DrawResults(c)
}
