package problems

import (
	"bytes"
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
	"strings"
)

type sectionRange struct {
	low  int
	high int
}

func Problem4() {

	path := "input/problem4.txt"
	contents, err := getInput(path)
	if err != nil {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		log.Output(2, trace)
		return
	}
	//contents = testInput4()

	assignments := bytes.Split(contents, []byte("\n"))
	containedRanges := 0
	for i := 0; i < len(assignments); i++ {
		assigned := string(assignments[i])
		log.Println(assigned)
		ranges := strings.Split(assigned, ",")
		if len(ranges) < 2 {
			log.Printf("wrong number of split ranges. Input: %s", ranges)
			continue
		}
		isContained, err := compare(ranges)
		if err != nil {
			trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
			log.Output(2, trace)
			return
		}
		if isContained == true {
			containedRanges++
		}
	}
	log.Printf("Contained Ranges: %d", containedRanges)
}

func compare(ranges []string) (isContained bool, err error) {
	first, err := splitRange(ranges[0])
	if err != nil {
		return false, err
	}
	second, err := splitRange(ranges[1])
	if err != nil {
		return false, err
	}
	isContained = first.isContained(second)
	return
}

func (r sectionRange) isContained(s sectionRange) (contains bool) {
	if r.low <= s.low && s.low <= r.high {
		contains = true
	}
	if s.low <= r.low && r.low <= s.high {
		contains = true
	}
	return
}

func splitRange(elem string) (sr sectionRange, err error) {
	items := strings.Split(elem, "-")
	sr.low, err = strconv.Atoi(items[0])
	if err != nil {
		return sr, err
	}
	sr.high, err = strconv.Atoi(items[1])
	if err != nil {
		return sr, err
	}
	return sr, err
}

func testInput4() []byte {
	return []byte(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)

}
