package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mod(i, n int64) int64 {
	coefficient := i / n
	if coefficient*n == i {
		return n
	}
	return i - coefficient*n
}

var Moves = map[int64]int64{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

func copyUniverse(p1score, p2score, p1pos, p2pos, depth, universes int64, wins *[2]int64) {
	if p1score >= 21 {
		wins[0] += universes
		return
	}
	if p2score >= 21 {
		wins[1] += universes
		return
	}
	if depth%2 == 0 {
		for move, count := range Moves {
			pos := mod(p1pos+move, 10)
			score := p1score + pos
			copyUniverse(score, p2score, pos, p2pos, depth+1, universes*count, wins)
		}
	} else {
		for move, count := range Moves {
			pos := mod(p2pos+move, 10)
			score := p2score + pos
			copyUniverse(p1score, score, p1pos, pos, depth+1, universes*count, wins)
		}
	}
}

func solveP1(p1Pos, p2Pos int64) int64 {
	p1Score, p2Score, rolls := int64(0), int64(0), int64(0)
	for true {
		rolls++
		p1Pos = mod(p1Pos+mod(rolls, 100), 10)
		rolls++
		p1Pos = mod(p1Pos+mod(rolls, 100), 10)
		rolls++
		p1Pos = mod(p1Pos+mod(rolls, 100), 10)
		p1Score += p1Pos
		if p1Score >= 1000 {
			break
		}

		rolls++
		p2Pos = mod(p2Pos+mod(rolls, 100), 10)
		rolls++
		p2Pos = mod(p2Pos+mod(rolls, 100), 10)
		rolls++
		p2Pos = mod(p2Pos+mod(rolls, 100), 10)
		p2Score += p2Pos
		if p2Score >= 1000 {
			break
		}
	}
	winner := p1Score
	loser := p2Score
	if winner < loser {
		winner, loser = loser, winner
	}

	return loser * rolls
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.SplitN(strings.TrimSpace(string(data)), "\n", 2)
	p1Pos, err := strconv.ParseInt(strings.TrimSpace(strings.SplitN(lines[0], ":", 2)[1]), 10, 64)
	p2Pos, err := strconv.ParseInt(strings.TrimSpace(strings.SplitN(lines[1], ":", 2)[1]), 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(solveP1(p1Pos, p2Pos))

	wins := [2]int64{0, 0}
	copyUniverse(0, 0, p1Pos, p2Pos, 0, 1, &wins)
	winner := wins[0]
	if wins[1] > winner {
		winner = wins[1]
	}
	fmt.Println(winner)
	return 0
}

func main() {
	os.Exit(run())
}
