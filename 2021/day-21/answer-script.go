package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mod(i, n int) int {
	coefficient := i / n
	if coefficient*n == i {
		return n
	}
	return i - coefficient*n
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.SplitN(strings.TrimSpace(string(data)), "\n", 2)
	player1Pos, err := strconv.Atoi(strings.TrimSpace(strings.SplitN(lines[0], ":", 2)[1]))
	player2Pos, err := strconv.Atoi(strings.TrimSpace(strings.SplitN(lines[1], ":", 2)[1]))
	if err != nil {
		panic(err)
	}

	player1Score, player2Score, rolls := 0, 0, 0
	for true {
		rolls++
		player1Pos = mod(player1Pos+mod(rolls, 100), 10)
		rolls++
		player1Pos = mod(player1Pos+mod(rolls, 100), 10)
		rolls++
		player1Pos = mod(player1Pos+mod(rolls, 100), 10)
		player1Score += player1Pos
		if player1Score >= 1000 {
			break
		}

		rolls++
		player2Pos = mod(player2Pos+mod(rolls, 100), 10)
		rolls++
		player2Pos = mod(player2Pos+mod(rolls, 100), 10)
		rolls++
		player2Pos = mod(player2Pos+mod(rolls, 100), 10)
		player2Score += player2Pos
		if player2Score >= 1000 {
			break
		}
	}
	winner := player1Score
	loser := player2Score
	if winner < loser {
		winner, loser = loser, winner
	}

	fmt.Println(loser * rolls)
	return 0
}

func main() {
	os.Exit(run())
}
