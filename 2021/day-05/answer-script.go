package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the input: %s\n", err)
		return 1
	}
	lines := [][2][2]int{}
	for _, l := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		line := [2][2]int{}
		for j, p := range strings.SplitN(l, " -> ", 2) {
			point := [2]int{}
			for k, c := range strings.SplitN(p, ",", 2) {
				coordinate, err := strconv.Atoi(c)
				if err != nil {
					fmt.Fprintf(os.Stderr, "could not parse coordinate %s of line %s: %s\n", c, l, err)
					return 1
				}
				point[k] = coordinate
			}
			line[j] = point
		}
		lines = append(lines, line)
	}
	overlaps := make(map[int]int)
	overlapsWithDiagonal := make(map[int]int)
	for _, line := range lines {
		if line[0][0] == line[1][0] {
			common := line[0][0]
			biggest := line[0][1]
			smallest := line[1][1]
			if smallest > biggest {
				biggest, smallest = smallest, biggest
			}
			for v := smallest; v <= biggest; v++ {
				overlaps[common*1000+v] += 1
				overlapsWithDiagonal[common*1000+v] += 1
			}
		} else if line[0][1] == line[1][1] {
			common := line[0][1]
			biggest := line[0][0]
			smallest := line[1][0]
			if smallest > biggest {
				biggest, smallest = smallest, biggest
			}
			for v := smallest; v <= biggest; v++ {
				overlaps[v*1000+common] += 1
				overlapsWithDiagonal[v*1000+common] += 1
			}
		} else {
			x1 := line[0][0]
			y1 := line[0][1]
			x2 := line[1][0]
			y2 := line[1][1]
			yDir := 1
			if y2 < y1 {
				yDir = -1
			}
			xDir := 1
			if x2 < x1 {
				xDir = -1
			}
			x := x1
			y := y1
			for true {
				overlapsWithDiagonal[x*1000+y] += 1
				if x == x2 && y == y2 {
					break
				}
				x += xDir
				y += yDir
			}
		}
	}
	count := 0
	for _, v := range overlaps {
		if v > 1 {
			count++
		}
	}
	fmt.Println(count)

	countWithDiagonal := 0
	for _, v := range overlapsWithDiagonal {
		if v > 1 {
			countWithDiagonal++
		}
	}
	fmt.Println(countWithDiagonal)
	return 0
}

func main() {
	os.Exit(run())
}
