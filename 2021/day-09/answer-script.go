package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var NESW = [...][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func findBasinSize(grid [][]int, visited map[[2]int]bool, x, y int) int {
	visited[[2]int{x, y}] = true
	count := 0
	for _, dir := range NESW {
		ty := y + dir[0]
		tx := x + dir[1]
		if ty < 0 || ty > len(grid)-1 || tx < 0 || tx > len(grid[ty])-1 {
			continue
		}
		if visited[[2]int{tx, ty}] {
			continue
		}
		val := grid[ty][tx]
		if val == 9 {
			continue
		} else {
			count += findBasinSize(grid, visited, tx, ty) + 1
		}
	}
	return count
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the input: %s\n", err)
		return 1
	}
	grid := [][]int{}
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		row := []int{}
		for _, n := range strings.Split(line, "") {
			parsed, err := strconv.Atoi(n)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error parsing integer %s in line %s: %s\n", n, line, err)
				return 1
			}
			row = append(row, parsed)
		}
		grid = append(grid, row)
	}
	sum := 0
	basins := []int{}
	for y, row := range grid {
		for x, point := range row {
			isLow := true
			for _, dir := range NESW {
				ty := y + dir[0]
				tx := x + dir[1]
				if ty < 0 || ty > len(grid)-1 || tx < 0 || tx > len(grid[ty])-1 {
					continue
				}
				isLow = isLow && grid[ty][tx] > point
			}
			if isLow {
				sum += point + 1
				visited := make(map[[2]int]bool)
				basins = append(basins, findBasinSize(grid, visited, x, y)+1)
			}
		}
	}
	fmt.Println(sum)
	sort.Ints(basins)
	basins = basins[len(basins)-3:]
	fmt.Println(basins[0] * basins[1] * basins[2])
	return 0
}

func main() {
	os.Exit(run())
}
