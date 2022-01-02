package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Step struct {
    Cuboid
	on bool
}

type Cuboid struct {
  x int
  y int
  z int
  w int
  h int
  d int
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

func subtract(a, b Cuboid) []Cuboid {
  res := []Cuboid
  topHeight := b.y - a.y
  if topHeight > 0 {
    top := Cuboid{
      x: a.x,
      y: a.y,
      z: a.z,
      w: a.w,
      h: topHeight,
      d: a.d,
    }
    res = append(res, top)
  }

  bottomY := b.y + b.h
  bottomHeight := a.h - (bottomY - a.y)
  if bottomHeight > 0 && bottomY < a.y + a.h {
    bottom := Cuboid{
      x: a.x,
      y: bottomY,
      z: a.z,
      w: a.w,
      h: bottomHeight,
      d: a.d,
    }
    res = append(res, bottom)
  }
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
        x1, x2 := parseRange(x)
		y1, y2 := parseRange(y)
		z1, z2 := parseRange(z)
        step.x = x1
        step.y = y1
        step.z = z1
        step.w = x2 - x1
        step.h = y2 - y1
        step.d = z2 - z1
		steps = append(steps, step)
	}

    fmt.Println(steps)
    // cuboids := []Cuboid{}
	return 0
}

func main() {
	os.Exit(run())
}
