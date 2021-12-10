package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var ScoresPart1 = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var ScoresPart2 = map[byte]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var OpeningToClosing = map[byte]byte{
	'{': '}',
	'[': ']',
	'(': ')',
	'<': '>',
}

var ClosingToOpening = map[byte]byte{
	'}': '{',
	']': '[',
	')': '(',
	'>': '<',
}

func isOpening(char byte) bool {
	return char == '{' || char == '[' || char == '(' || char == '<'
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the input: %s\n", err)
		return 1
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	score := 0
	scores := []int{}
	for _, line := range lines {
		lineScore := 0
		pending := []byte{}
		corrupt := false
		for i := 0; i < len(line); i++ {
			cur := line[i]
			if isOpening(cur) {
				pending = append(pending, cur)
			} else if len(pending) > 0 {
				expected := OpeningToClosing[pending[len(pending)-1]]
				if cur != expected {
					score += ScoresPart1[cur]
					corrupt = true
					break
				}
				pending = pending[:len(pending)-1]
			}
		}
		if !corrupt {
			for i := len(pending) - 1; i >= 0; i-- {
				lineScore *= 5
				lineScore += ScoresPart2[OpeningToClosing[pending[i]]]
			}
			scores = append(scores, lineScore)
		}
	}
	fmt.Println(score)
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
	return 0
}

func main() {
	os.Exit(run())
}
