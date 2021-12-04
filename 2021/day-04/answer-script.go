package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	BoardSize = 5
)

func calculateId(board, row, column int) int {
	return board*100 + row*10 + column
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the input: %s\n", err)
		return 1
	}
	blocks := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	draws := []int{}
	for i, n := range strings.Split(blocks[0], ",") {
		parsed, err := strconv.Atoi(n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing drawing integer %s at %d: %s\n", n, i, err)
			return 1
		}
		draws = append(draws, parsed)
	}

	boards := [][BoardSize][BoardSize]int{}
	for _, boardStr := range blocks[1:] {
		board := [BoardSize][BoardSize]int{}
		for j, line := range strings.Split(boardStr, "\n") {
			row := [BoardSize]int{}
			line = strings.ReplaceAll(strings.TrimSpace(line), "  ", " ")
			for k, numStr := range strings.Split(line, " ") {
				numStr = strings.TrimSpace(numStr)
				if numStr == "" {
					continue
				}
				parsed, err := strconv.Atoi(numStr)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error parsing board integer %s at %d: %s\n", numStr, k, err)
					return 1
				}
				row[k] = parsed
			}
			board[j] = row
		}
		boards = append(boards, board)
	}

	firstAndLast := [2]int{-1, -1}
	finishedBoards := make(map[int]bool)
	winnerLastCalled := -1
	lastCalled := -1
	marks := make(map[int]bool)
	for _, draw := range draws {
		for boardInd, board := range boards {
			if finishedBoards[boardInd] {
				continue
			}
			for i := 0; i < BoardSize && !finishedBoards[boardInd]; i++ {
				for j := 0; j < BoardSize; j++ {
					cellId := calculateId(boardInd, i, j)
					if draw != board[i][j] || marks[cellId] {
						continue
					}
					marks[cellId] = true
					rowsDone := true
					colsDone := true
					for c := 0; c < BoardSize; c++ {
						rowId := calculateId(boardInd, i, c)
						rowsDone = rowsDone && marks[rowId]
						colId := calculateId(boardInd, c, j)
						colsDone = colsDone && marks[colId]
					}
					if rowsDone || colsDone {
						finishedBoards[boardInd] = true
						if firstAndLast[0] == -1 {
							firstAndLast[0] = boardInd
							winnerLastCalled = draw
						} else if len(finishedBoards) == len(boards) {
							firstAndLast[1] = boardInd
							lastCalled = draw
							goto done
						}
						break
					}
				}
			}
		}
	}

	fmt.Fprintf(os.Stderr, "could not find winner!\n")
	return 1

done:
	for b, boardInd := range firstAndLast {
		board := boards[boardInd]
		sum := 0
		for i := 0; i < BoardSize; i++ {
			for j := 0; j < BoardSize; j++ {
				cellId := calculateId(boardInd, i, j)
				if !marks[cellId] {
					sum += board[i][j]
				}
			}
		}
		last := winnerLastCalled
		if b == 1 {
			last = lastCalled
		}
		fmt.Println(last * sum)
	}
	return 0
}

func main() {
	os.Exit(run())
}
