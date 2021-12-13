package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	X = 0
	Y = 1
)

func visualize(dots map[[2]int]int) {
	maxX := -1
	maxY := -1
	for dot := range dots {
		if dot[0] > maxX {
			maxX = dot[0]
		}
		if dot[1] > maxY {
			maxY = dot[1]
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if dots[[2]int{x, y}] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println("")
	}
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the input: %s\n", err)
		return 1
	}
	input := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	dots := make(map[[2]int]int)
	instructions := [][2]int{}
	for _, line := range strings.Split(input[0], "\n") {
		dot := [2]int{}
		xy := strings.Split(line, ",")
		x, err := strconv.Atoi(xy[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing integer x %s at line: %s\n", xy[0], line, err)
			return 1
		}
		y, err := strconv.Atoi(xy[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing integer y %s at line: %s\n", xy[1], line, err)
			return 1
		}
		dot[0], dot[1] = x, y
		dots[dot] += 1
	}
	for _, line := range strings.Split(input[1], "\n") {
		split := strings.Split(line, "=")
		val, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing integer for instruction %s at line: %s\n", split[1], line, err)
			return 1
		}
		instruction := [2]int{}
		instruction[0] = val
		if strings.Contains(split[0], "x") {
			instruction[1] = X
		} else {
			instruction[1] = Y
		}
		instructions = append(instructions, instruction)
	}
	p1Done := false
	for _, instruction := range instructions {
		fold := instruction[0]
		isX := instruction[1] == X
		newDots := make(map[[2]int]int)
		for dot := range dots {
			x, y := dot[0], dot[1]
			if isX {
				if x < fold {
					newDots[dot] += 1
					continue
				}
				x2 := x - (x-fold)*2
				newDots[[2]int{x2, y}] += 1
			} else {
				if y < fold {
					newDots[dot] += 1
					continue
				}
				y2 := y - (y-fold)*2
				newDots[[2]int{x, y2}] += 1
			}
		}
		dots = newDots
		if !p1Done {
			p1Done = true
			fmt.Println(len(dots))
		}
	}
	visualize(dots)
	return 0
}

func main() {
	os.Exit(run())
}
