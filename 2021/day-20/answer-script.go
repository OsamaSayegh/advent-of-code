package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Size    = 100
	Padding = 1
)

var Dirs = [...][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 0},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.SplitN(strings.TrimSpace(string(data)), "\n\n", 2)
	algorithm := input[0]
	if len(algorithm) != 512 {
		panic(fmt.Errorf("expected algorithm to be 512 chars, but got %d", len(algorithm)))
	}
	image := make(map[[2]int]bool)
	for r, row := range strings.Split(input[1], "\n") {
		for c, bit := range row {
			image[[2]int{r, c}] = bit == '#'
		}
	}
	iterations := 50
	edges := [2]int{0 - Padding, (Size - 1) + Padding}
	missingIsLit := false
	for iterations > 0 {
		enhanced := make(map[[2]int]bool)
		for r := edges[0]; r <= edges[1]; r++ {
			for c := edges[0]; c <= edges[1]; c++ {
				index := 0
				for _, dir := range Dirs {
					index <<= 1
					nr := r + dir[0]
					nc := c + dir[1]
					if v, exists := image[[2]int{nr, nc}]; v || (!exists && missingIsLit) {
						index |= 1
					}
				}
				enhanced[[2]int{r, c}] = algorithm[index] == '#'
			}
		}
		edges[0] -= Padding
		edges[1] += Padding
		if algorithm[0] == '#' {
			missingIsLit = !missingIsLit
		}
		iterations--
		image = enhanced
		if iterations == 48 || iterations == 0 {
			count := 0
			for _, lit := range image {
				if lit {
					count++
				}
			}
			fmt.Println(count)
		}
	}
	return 0
}

func main() {
	os.Exit(run())
}
