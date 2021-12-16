package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Size     = 100
	FullSize = Size * 5
)

var Dirs = [...][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func solve(size int, costProvider func(int, int) int) int {
	graph := make(map[int]int)
	graph[0] = 0
	queue := []int{0}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		r := current / 1000
		c := current % 1000
		for _, dir := range Dirs {
			nr := r + dir[0]
			nc := c + dir[1]
			if nr < 0 || nr >= size || nc < 0 || nc >= size {
				continue
			}
			neighbor := nr*1000 + nc
			distance := graph[current] + costProvider(nr, nc)
			if _, exists := graph[neighbor]; !exists || graph[neighbor] > distance {
				graph[neighbor] = distance
				queue = append(queue, neighbor)
			}
		}
	}
	return graph[(size-1)*1000+(size-1)]
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the input: %s\n", err)
		return 1
	}
	cavern := [Size][Size]int{}
	for r, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		row := [Size]int{}
		for c, risk := range strings.Split(line, "") {
			parsed, err := strconv.Atoi(risk)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error parsing integer %s at %d: %s\n", risk, c, err)
				return 1
			}
			row[c] = parsed
		}
		cavern[r] = row
	}
	fmt.Println(solve(Size, func(nr, nc int) int { return cavern[nr][nc] }))
	fullCavern := [FullSize][FullSize]int{}
	for r, row := range cavern {
		for c, val := range row {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					newVal := val + i + j
					if newVal >= 10 {
						newVal %= 10
						newVal += 1
					}
					fullCavern[r+Size*i][c+Size*j] = newVal
				}
			}
		}
	}
	fmt.Println(solve(FullSize, func(nr, nc int) int { return fullCavern[nr][nc] }))
	return 0
}

func main() {
	os.Exit(run())
}
