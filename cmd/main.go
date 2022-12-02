package main

import (
	"com.advent2022.goingtharn/problems"
	"flag"
	//"fmt"
	//"log"
)

func buildProblems() map[string]func() {
	pm := make(map[string]func())
	pm["Problem1"] = problems.Problem1
	pm["Problem2"] = problems.Problem2
	return pm
}

func main() {
	problems := buildProblems()
	problem := flag.String("Problem", "", "Name the problem you wish to run")
	flag.Parse()

	fn := problems[*problem]
	fn()
}
