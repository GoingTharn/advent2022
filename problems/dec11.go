package problems

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items          []int
	operation      string
	magnitude      int
	test_magnitude int
	trueMonkey     int // place in map
	falseMonkey    int
	itemsInspected int
	worryReducer   int
}

func (m Monkey) String() string {
	return fmt.Sprintf("items: %v\noperation: %s magnitude: %d\n divisible by %d\ntrueMonkey: %d falseMonkey: %d\n",
		m.items, m.operation, m.magnitude, m.test_magnitude, m.trueMonkey, m.falseMonkey)
}

func divide(x, y int) int {
	return x % y
}

func multiply(x, y int) int {
	return x * y
}

func square(x int) int {
	return x * x
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func (m *Monkey) Throw(monkeys map[int]Monkey) map[int]Monkey {
	var item int
	item, m.items = popInt(m.items)

	switch m.operation {
	case "+":
		item = add(item, m.magnitude)
	case "-":
		item = subtract(item, m.magnitude)
	case "/":
		item = divide(item, m.magnitude)
	case "**":
		item = square(item)
	case "*":
		item = multiply(item, m.magnitude)
	}

	item = item % m.worryReducer
	testVar := item%m.test_magnitude == 0
	if testVar {
		working := monkeys[m.trueMonkey]
		working.items = append(working.items, item)
		monkeys[m.trueMonkey] = working
	} else {
		working := monkeys[m.falseMonkey]
		working.items = append(working.items, item)
		monkeys[m.falseMonkey] = working
	}
	return monkeys
}

func popInt(items []int) (item int, poppedItems []int) {
	if len(items) < 1 {
		return
	}
	item = items[0]
	if len(items) < 2 {
		poppedItems = make([]int, 0)
		return
	}
	poppedItems = items[1:]
	return
}

func parseMonkeys(raw []byte) map[int]Monkey {
	monkeys := make(map[int]Monkey)
	monkeysRaw := strings.Split(string(raw), "Monkey ")
	worryReducer := make([]int, 0)

	cleanMonkeyRaw := make([]string, 0)
	for _, line := range monkeysRaw {
		if len(line) > 1 {
			cleanMonkeyRaw = append(cleanMonkeyRaw, line)
		}
	}
	for i := range cleanMonkeyRaw {
		monkeys[i] = Monkey{}
	}
	for i, mRaw := range cleanMonkeyRaw {
		workingMonkey := monkeys[i]
		commands := strings.Split(mRaw, "\n")

		for i := range commands {
			commands[i] = strings.ReplaceAll(commands[i], " ", "")
		}
		rawItems := commands[1]
		rawItems = strings.ReplaceAll(rawItems, "Startingitems:", "")
		itemStrings := strings.Split(rawItems, ",")
		items := make([]int, 0)
		for _, i := range itemStrings {
			item, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println("ERROR CONVERTING ITEM TO INT")
			}
			items = append(items, item)
		}

		workingMonkey.items = items
		rawOperation := commands[2]
		rawOperation = strings.ReplaceAll(rawOperation, "Operation:new=old", "")
		workingMonkey.operation = string(rawOperation[0])
		magnitude, err := strconv.Atoi(rawOperation[1:])
		if err != nil {
			if rawOperation[1:] == "old" {
				workingMonkey.operation = "**"
			} else {
				fmt.Printf("ERROR CONVERTING MAGNITUDE TO INT. %s\n", rawOperation)
			}
		}
		workingMonkey.magnitude = magnitude

		rawTest := commands[3]
		rawTest = strings.ReplaceAll(rawTest, "Test:divisibleby", "")
		testMag, err := strconv.Atoi(rawTest)
		worryReducer = append(worryReducer, testMag)
		if err != nil {
			fmt.Println("ERROR CONVERTING TEST MAGNITUDE TO INT")
		}
		workingMonkey.test_magnitude = testMag

		ifTrueRaw := commands[4][len(commands[4])-1]
		ifTrue, err := strconv.Atoi(string(ifTrueRaw))
		if err != nil {
			fmt.Println("ERROR CONVERTING IFTRUE TO INT")
		}
		workingMonkey.trueMonkey = ifTrue

		ifFalseRaw := commands[5][len(commands[5])-1]
		ifFalse, err := strconv.Atoi(string(ifFalseRaw))
		if err != nil {
			fmt.Println("ERROR CONVERTING IFTRUE TO INT")
		}
		workingMonkey.falseMonkey = ifFalse

		monkeys[i] = workingMonkey
	}
	worryMod := 1
	for _, val := range worryReducer {
		worryMod = worryMod * val
	}
	for i, monkey := range monkeys {
		monkey.worryReducer = worryMod
		monkeys[i] = monkey
	}
	return monkeys
}

func Rummage(monkeys map[int]Monkey) map[int]Monkey {
	for i := 0; i < len(monkeys); i++ {
		working := monkeys[i]
		for range working.items {
			working.itemsInspected = working.itemsInspected + 1
			monkeys = working.Throw(monkeys)
		}
		monkeys[i] = working
	}
	return monkeys
}

func Problem11() {
	raw, _ := getInput("input/problem11.txt")
	monkeys := parseMonkeys(raw)
	for i := 0; i < 10000; i++ {
		monkeys = Rummage(monkeys)
	}

	inspectedItems := make([]int, 1)
	for _, v := range monkeys {
		inspectedItems = append(inspectedItems, v.itemsInspected)
	}
	sort.Ints(inspectedItems)
	monkeyBiz := inspectedItems[len(inspectedItems)-1] * inspectedItems[len(inspectedItems)-2]
	fmt.Printf("MonkeyBiz: %d", monkeyBiz)

}
