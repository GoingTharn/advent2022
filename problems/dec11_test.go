package problems

import (
	"fmt"
	"sort"
	"testing"
)

func TestParseMonkey(t *testing.T) {
	raw := []byte(`Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`)

	monkeys := parseMonkeys(raw)

	for i := 0; i < 10000; i++ {
		monkeys = Rummage(monkeys)
	}
	var inspectedItems []int
	for _, v := range monkeys {
		inspectedItems = append(inspectedItems, v.itemsInspected)
	}
	sort.Ints(inspectedItems)
	monkeyBiz := inspectedItems[len(inspectedItems)-1] * inspectedItems[len(inspectedItems)-2]
	if monkeyBiz != 2713310158 {
		fmt.Printf("Got wrong MonkeyBiz. Got %v", inspectedItems)
		for k, v := range monkeys {
			fmt.Printf("Monkey: %d\n%s\n\n", k, v)
		}
		t.Fail()
	}

}
