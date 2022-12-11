package problems

import (
	"fmt"
	"strings"
	"testing"
)

func TestClock(t *testing.T) {
	raw := []byte(`noop
addx 3
addx -5
`)
	instructionList := strings.Split(strings.Trim(string(raw), "\n"), "\n")
	c := buildClock(instructionList)

	expectedV := map[int]int{1: 1, 2: 1, 3: 1, 4: 4, 5: 4, 6: -1, 7: -1}
	for i := 1; i < 8; i++ {
		c.Tick()
		if c.time != i {
			fmt.Printf("Time diverged! i: %d c.time %d\n", i, c.time)
			t.Fail()
		}
		fmt.Println(c.results)
		if expectedV[i-1] != c.results[i-1] {
			fmt.Printf("Expected: %d Got:%d at Tick: %d\n", expectedV[i], c.results[i], i)
			t.Fail()
		}
	}

}

func TestBigInput(t *testing.T) {
	raw := []byte(`addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
`)

	instructionList := strings.Split(strings.Trim(string(raw), "\n"), "\n")
	c := buildClock(instructionList)

	for i := 1; i < 225; i++ {
		c.Tick()
	}
	output := c.GetResults()
	if output[20] != 420 {
		fmt.Printf("20 wrong. Got %d instead of 420\n", output[20])
		t.Fail()
	}

	if output[220] != 3960 {
		fmt.Printf("220 wrong. Got %d instead of 3960\n", output[220])
		t.Fail()
	}

	sum := 0
	for _, v := range output {
		sum += v
	}
	if sum != 13140 {
		fmt.Printf("Wrong sum. Got %d wanted 13140", sum)
		t.Fail()
	}

	DrawResults(c)
}
