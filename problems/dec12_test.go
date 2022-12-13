package problems

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestBuildArray(t *testing.T) {
	raw := []byte(`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`)

	output, _, _, _ := buildArray(raw)

	tests := []struct {
		input string
		sep   string
		want  string
	}{
		{input: "0/0", sep: "/", want: "S"},
		{input: "0/3", sep: "/", want: "q"},
		{input: "2/4", sep: "/", want: "z"},
		{input: "2/5", sep: "/", want: "E"},
	}

	for _, tc := range tests {
		got := strings.Split(tc.input, tc.sep)
		y, _ := strconv.Atoi(got[0])
		x, _ := strconv.Atoi(got[1])
		if output[y][x].value != tc.want {
			fmt.Printf("Test Failed. Actual: %s Expected: %s", output[y][x].value, tc.want)
			t.Fail()
		}
	}
}

func TestShortestPath(t *testing.T) {
	raw := []byte(`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`)
	dist := getShortest(raw)
	if dist != int64(31) {
		fmt.Printf("Wrong path. Got: %d expected 31", dist)
		t.Fail()
	}

}
