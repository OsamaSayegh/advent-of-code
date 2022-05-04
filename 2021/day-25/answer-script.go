package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type Grid map[[2]int]byte

const (
	EAST  = '>'
	SOUTH = 'v'
)

func run() int {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	grid := make(Grid)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	height := len(lines)
	width := len(lines[0])
	for r, line := range lines {
		for c := 0; c < len(line); c++ {
			if line[c] == EAST || line[c] == SOUTH {
				grid[[2]int{r, c}] = line[c]
			}
		}
	}
	i := 0
	for true {
		i += 1
		eastMoves := make(Grid)
		for coord, dir := range grid {
			if dir != EAST {
				eastMoves[coord] = dir
				continue
			}
			newCoord := [2]int{coord[0], (coord[1] + 1) % width}
			if _, ok := grid[newCoord]; !ok {
				eastMoves[newCoord] = dir
			} else {
				eastMoves[coord] = dir
			}
		}
		southMoves := make(Grid)
		for coord, dir := range eastMoves {
			if dir != SOUTH {
				southMoves[coord] = dir
				continue
			}
			newCoord := [2]int{(coord[0] + 1) % height, coord[1]}
			if _, ok := eastMoves[newCoord]; !ok {
				southMoves[newCoord] = dir
			} else {
				southMoves[coord] = dir
			}
		}
		if reflect.DeepEqual(grid, southMoves) {
			break
		} else {
			grid = southMoves
		}
	}
	fmt.Println(i)
	return 0
}

func main() {
	os.Exit(run())
}
