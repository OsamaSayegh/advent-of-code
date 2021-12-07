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
	positions := []int{}
	for i, p := range strings.Split(strings.TrimSpace(string(data)), ",") {
		pos, err := strconv.Atoi(p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing integer %s at %d: %s\n", p, i, err)
			return 1
		}
		positions = append(positions, pos)
	}

	sum := 0
	for _, p := range positions {
		sum += p
	}

	average := sum / len(positions)
	p1Ans := calculateCost(average, positions)
	{
		oneHigher := calculateCost(average+1, positions)
		oneLower := calculateCost(average-1, positions)
		if oneHigher < p1Ans || oneLower < p1Ans {
			dir := 1
			p1Ans = oneHigher
			if oneLower < p1Ans {
				dir = -1
				p1Ans = oneLower
			}
			i := dir
			for true {
				newCost := calculateCost(average+i, positions)
				if newCost > p1Ans {
					break
				} else {
					p1Ans = newCost
					i += dir
				}
			}
		}
	}
	fmt.Println(p1Ans)
	fmt.Println(calculateCostP2(average, positions))
	return 0
}

func calculateCost(pos int, positions []int) int {
	sum := 0
	for _, p := range positions {
		val := pos - p
		if val < 0 {
			val *= -1
		}
		sum += val
	}
	return sum
}

func calculateCostP2(pos int, positions []int) int {
	sum := 0
	for _, p := range positions {
		val := pos - p
		if val < 0 {
			val *= -1
		}
		sum += (val * (val + 1)) / 2
	}
	return sum
}

func main() {
	os.Exit(run())
}
