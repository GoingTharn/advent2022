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
	pm["Problem3"] = problems.Problem3
	pm["Problem4"] = problems.Problem4
	pm["Problem5"] = problems.Problem5
	pm["Problem6"] = problems.Problem6
	pm["Problem7"] = problems.Problem7
	pm["Problem8"] = problems.Problem8
	pm["Problem9"] = problems.Problem9
	pm["Problem10"] = problems.Problem10
	return pm
}

func main() {
	problems := buildProblems()
	problem := flag.String("Problem", "", "Name the problem you wish to run")
	flag.Parse()

	fn := problems[*problem]
	fn()
}
