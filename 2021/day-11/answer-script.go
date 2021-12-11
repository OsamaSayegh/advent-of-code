package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Size = 10

var Dirs = [...][2]int{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the input: %s\n", err)
		return 1
	}
	octopuses := [Size][Size]int{}
	for i, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		row := [Size]int{}
		for j, level := range strings.Split(line, "") {
			parsed, err := strconv.Atoi(level)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error parsing integer %s at %d: %s\n", level, j, err)
				return 1
			}
			row[j] = parsed
		}
		octopuses[i] = row
	}

	flashes := 0
	k := 1
	for true {
		flashQueue := [][2]int{}
		for r, row := range octopuses {
			for c := range row {
				octopuses[r][c] = (octopuses[r][c] + 1) % 10
				if octopuses[r][c] == 0 {
					flashQueue = append(flashQueue, [2]int{r, c})
					flashes += 1
				}
			}
		}
		for len(flashQueue) > 0 {
			r := flashQueue[0][0]
			c := flashQueue[0][1]
			flashQueue = flashQueue[1:]
			if octopuses[r][c] == 0 {
				for _, dir := range Dirs {
					nr := r + dir[0]
					nc := c + dir[1]
					if nr < 0 || nr >= Size || nc < 0 || nc >= Size {
						continue
					}
					if octopuses[nr][nc] == 0 {
						continue
					}
					octopuses[nr][nc] = (octopuses[nr][nc] + 1) % 10
					if octopuses[nr][nc] == 0 {
						flashQueue = append(flashQueue, [2]int{nr, nc})
						flashes += 1
					}
				}
			}
		}
		if k == 100 {
			fmt.Println(flashes)
		}
		flashAll := true
		for i := 0; i < len(octopuses) && flashAll; i++ {
			row := octopuses[i]
			for _, oct := range row {
				flashAll = flashAll && oct == 0
				if !flashAll {
					break
				}
			}
		}
		if flashAll {
			break
		}
		k++
	}
	fmt.Println(k)
	return 0
}

func main() {
	os.Exit(run())
}
