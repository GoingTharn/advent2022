package problems

import (
	"bytes"
	"log"
	"strings"

	"golang.org/x/exp/slices"
)

type rucksack struct {
	contents string
	first    string
	second   string
	shared   []string
}

func priorityOf(s string) int {
	var priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(priorities, s) + 1
}

func (r rucksack) Split() rucksack {
	compartment_size := len(r.contents) / 2
	r.first = r.contents[:compartment_size]
	r.second = r.contents[compartment_size:]
	return r
}

func (r rucksack) Compare() rucksack {
	var shared []string
	for i := 0; i < len(r.first); i++ {
		working := string(r.first[i])
		if strings.Contains(r.second, working) && !(slices.Contains(shared, working)) {
			shared = append(shared, working)
		}

	}
	r.shared = shared
	return r
}

func (r rucksack) getPriority() (priority int) {
	if len(r.shared) == 0 {
		log.Printf("No shared items! %s", r)
		return 0
	}
	for i := 0; i < len(r.shared); i++ {
		priority += priorityOf(r.shared[i])
	}
	return
}

func compareRucks(rucks []rucksack) (priority int) {
	log.Println(len(rucks))
	for i := 0; i < len(rucks[0].contents); i++ {
		working := string(rucks[0].contents[i])
		if strings.Contains(rucks[1].contents, working) &&
			strings.Contains(rucks[2].contents, working) {
			priority = priorityOf(working)
		}
		if priority != 0 {
			return
		}
	}
	return
}

func Problem3() {
	input, _ := getInput("input/problem3.txt")
	//input := testProblem3Input()

	var rucksacks []rucksack
	compartments := bytes.Split(input, []byte("\n"))
	for i := 0; i < len(compartments); i++ {
		ruck := rucksack{contents: string(compartments[i])}
		ruck = ruck.Split()
		ruck = ruck.Compare()
		rucksacks = append(rucksacks, ruck)
	}

	totalPriority := 0
	for i := 0; i < len(rucksacks); i++ {
		r := rucksacks[i]
		totalPriority += r.getPriority()
	}
	log.Println(totalPriority)
	sharedPriority := 0
	for i := 0; i < len(rucksacks); i = i + 3 {
		sharedPriority += compareRucks(rucksacks[i : i+3])
	}
	log.Println(sharedPriority)
}

func testProblem3Input() []byte {
	return []byte(`ApqgZZSZgcZJqpzBbqTbbLjBDBLhBA
wHptFFsHttHFLMDQDFTbbj
fVfvsstwPHwNwfNGfHWRSnlpClcJzCWCzddSrddg`)
}
