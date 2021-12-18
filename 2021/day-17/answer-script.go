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
	input := strings.Split(strings.TrimSpace(strings.ReplaceAll(string(data), "target area:", "")), ", ")
	input[0] = strings.ReplaceAll(input[0], "x=", "")
	input[1] = strings.ReplaceAll(input[1], "y=", "")

	xStart, err := strconv.Atoi(strings.SplitN(input[0], "..", 2)[0])
	xEnd, err := strconv.Atoi(strings.SplitN(input[0], "..", 2)[1])

	yStart, err := strconv.Atoi(strings.SplitN(input[1], "..", 2)[0])
	yEnd, err := strconv.Atoi(strings.SplitN(input[1], "..", 2)[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't parse input %v: %s\n", input, err)
		return 1
	}

	if xStart > xEnd {
		xStart, xEnd = xEnd, xStart
	}
	if yStart > yEnd {
		yStart, yEnd = yEnd, yStart
	}

	xvs := []int{}
	for xv := 0; xv <= xEnd; xv++ {
		maxDistance := xv * (xv + 1) / 2
		if maxDistance < xStart {
			continue
		}
		if xv*2-1 > xEnd && !(xv >= xStart && xv <= xEnd) {
			continue
		}
		xvs = append(xvs, xv)
	}

	velocities := [][2]int{}
	highest := -1
	for _, xv := range xvs {
		x := xv
		step := 1
		for x < xStart {
			x += xv - step
			step++
		}
		i := 0
		for x >= xStart && x <= xEnd {
			for y := yStart; y <= yEnd; y++ {
				yv := (y + ((step * (step - 1)) / 2))
				if yv%step != 0 {
					continue
				}
				yv /= step
				velocity := [2]int{xv, yv}
				unique := true
				for _, v := range velocities {
					if v == velocity {
						unique = false
						break
					}
				}
				if unique {
					velocities = append(velocities, velocity)
					max := (yv * (yv + 1)) / 2
					if max > highest {
						highest = max
					}
				}
			}
			if xv-step > 0 {
				x += xv - step
			}
			step++
			i++
			if i > 1000 {
				break
			}
		}
	}
	fmt.Println(highest)
	fmt.Println(len(velocities))
	return 0
}

func main() {
	os.Exit(run())
}
