package problems

import (
	"fmt"
	"strings"
)

type buff []byte

const BOUND int = 14

func Problem6() {
	raw := testInput6()
	raw, _ = getInput("input/problem6.txt")

	buffer := strings.Split(string(raw), "")
	for i := range buffer {
		start := checkStartOfPacket(buffer, i)
		if start == true {
			fmt.Printf("Packet Start: %d\n", i)
			break
		}
	}

}

func checkStartOfPacket(buffer []string, i int) bool {
	if i < BOUND {
		return false
	}
	charSet := make(map[string]int)

	for j := BOUND; j > 0; j-- {
		charSet[buffer[i-j]] = 0
	}

	fmt.Printf("%d:", i)
	fmt.Println(charSet)
	if len(charSet) == BOUND {
		return true
	}
	return false
}

func testInput6() buff {
	return []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
}
