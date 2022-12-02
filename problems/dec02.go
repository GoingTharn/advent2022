package problems

import (
	"bytes"
	"log"
	"strings"
	//"os"
	//"runtime/debug"
	//"sort"
	//"strconv"
)

type game struct {
	me       piece
	opponent piece
}

type piece struct {
	value string
	score int
}

func Problem2() {
	gamesInput, _ := getInput("input/problem2.txt")
	// gamesInput := testInput2()
	games := parseGames(gamesInput)

	running_total := 0
	for i := 0; i < len(games); i++ {
		running_total += games[i].play()
	}
	log.Print(running_total)
}

// A X rock
// B Y paper
// C Z scissors
func (g game) play() (result int) {
	gameMap := make(map[string]int)
	gameMap["A X"] = 3
	gameMap["A Y"] = 1
	gameMap["A Z"] = 2
	gameMap["B X"] = 1
	gameMap["B Y"] = 2
	gameMap["B Z"] = 3
	gameMap["C Z"] = 1
	gameMap["C X"] = 2
	gameMap["C Y"] = 3

	result = gameMap[g.String()]
	result += g.me.score
	log.Printf("Game: %s  Result: %d", g, result)
	return
}

func (g game) String() string {
	this := []string{g.opponent.value, g.me.value}
	return strings.Join(this, " ")
}

func testInput2() []byte {
	return []byte(`B Y
B Y
A X
B Y
A Y`)
}

func createGame(input []byte) game {
	gm := make(map[string]int)
	gm["X"] = 0
	gm["Y"] = 3
	gm["Z"] = 6

	opp, me, found := strings.Cut(string(input), " ")

	piece_val := opp
	piece_score, found := gm[piece_val]
	if found != true {
		piece_score = 0
	}

	oppPiece := piece{score: piece_score, value: piece_val}
	piece_val = me
	piece_score, found = gm[piece_val]
	if found != true {
		piece_score = 0
	}
	myPiece := piece{score: piece_score, value: piece_val}

	return game{me: myPiece, opponent: oppPiece}
}

func parseGames(contents []byte) []game {
	data := bytes.Split(contents, []byte("\n"))
	var output []game
	for i := 0; i < len(data); i++ {
		parsedGame := createGame(data[i])
		output = append(output, parsedGame)
	}
	return output
}
