package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Step struct {
	x1 int
	x2 int
	y1 int
	y2 int
	z1 int
	z2 int
	on bool
}

func parseRange(coords string) (int, int) {
	split := strings.SplitN(coords, "..", 2)
	v1, err := strconv.Atoi(split[0])
	v2, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	return v1, v2
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	steps := []Step{}
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		step := Step{}
		if strings.HasPrefix(line, "on ") {
			line = strings.Replace(line, "on ", "", 1)
			step.on = true
		} else if strings.HasPrefix(line, "off ") {
			line = strings.Replace(line, "off ", "", 1)
			step.on = false
		} else {
			panic(fmt.Errorf("unknown on/off prefix for line %s", line))
		}
		coords := strings.SplitN(line, ",", 3)
		x := strings.Replace(coords[0], "x=", "", 1)
		y := strings.Replace(coords[1], "y=", "", 1)
		z := strings.Replace(coords[2], "z=", "", 1)
		step.x1, step.x2 = parseRange(x)
		step.y1, step.y2 = parseRange(y)
		step.z1, step.z2 = parseRange(z)
		steps = append(steps, step)
	}
	lit := make(map[[3]int]bool)
	fmt.Println("steps", len(steps))
	for i, step := range steps {
		fmt.Println(i)
		for z := step.z1; z <= step.z2; z++ {
			// if z < -50 || z > 50 { continue }
			for y := step.y1; y <= step.y2; y++ {
				// if y < -50 || y > 50 { continue }
				for x := step.x1; x <= step.x2; x++ {
					// if x < -50 || x > 50 { continue }
					loc := [3]int{x, y, z}
					if step.on {
						lit[loc] = true
					} else {
						delete(lit, loc)
					}
				}
			}
		}
	}
	fmt.Println(len(lit))
	fmt.Println(steps)
	return 0
}

func main() {
	os.Exit(run())
}
