package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cuboid struct {
	x int
	y int
	z int
	w int
	h int
	d int
}

type Step struct {
	*Cuboid
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intersection(a, b *Cuboid) *Cuboid {
	x1 := max(a.x, b.x)
	y1 := max(a.y, b.y)
	z1 := max(a.z, b.z)
	x2 := min(a.x+a.w, b.x+b.w)
	y2 := min(a.y+a.h, b.y+b.h)
	z2 := min(a.z+a.d, b.z+b.d)
	if x2 > x1 && y2 > y1 && z2 > z1 {
		return &Cuboid{
			x: x1,
			y: y1,
			z: z1,
			w: x2 - x1,
			h: y2 - y1,
			d: z2 - z1,
		}
	} else {
		return nil
	}
}

func subtract(a, b *Cuboid) []*Cuboid {
	if intersection(a, b) == nil {
		dup := *a
		return []*Cuboid{&dup}
	}
	res := []*Cuboid{}
	topCuboidHeight := b.y - a.y
	if topCuboidHeight > 0 {
		top := Cuboid{
			x: a.x,
			y: a.y,
			z: a.z,
			w: a.w,
			h: topCuboidHeight,
			d: a.d,
		}
		res = append(res, &top)
	}

	bottomCuboidY := b.y + b.h
	bottomCuboidHeight := a.h - (bottomCuboidY - a.y)
	if bottomCuboidHeight > 0 && bottomCuboidY < a.y+a.h {
		bottom := Cuboid{
			x: a.x,
			y: bottomCuboidY,
			z: a.z,
			w: a.w,
			h: bottomCuboidHeight,
			d: a.d,
		}
		res = append(res, &bottom)
	}

	topY := max(a.y, b.y)
	bottomY := min(b.y+b.h, a.y+a.h)

	leftRightCuboidHeight := bottomY - topY
	leftCuboidWidth := b.x - a.x
	if leftCuboidWidth > 0 && leftRightCuboidHeight > 0 {
		left := Cuboid{
			x: a.x,
			y: topY,
			z: a.z,
			w: leftCuboidWidth,
			h: leftRightCuboidHeight,
			d: a.d,
		}
		res = append(res, &left)
	}

	rightCuboidX := b.x + b.w
	rightCuboidWidth := a.w - (rightCuboidX - a.x)
	if rightCuboidWidth > 0 && leftRightCuboidHeight > 0 {
		right := Cuboid{
			x: rightCuboidX,
			y: topY,
			z: a.z,
			w: rightCuboidWidth,
			h: leftRightCuboidHeight,
			d: a.d,
		}
		res = append(res, &right)
	}

	frontBackCuboidX := max(a.x, b.x)
	frontBackCuboidWidth := min((a.x+a.w)-frontBackCuboidX, (b.x+b.w)-frontBackCuboidX)
	frontBackCuboidHeight := min((a.y+a.h)-topY, (b.y+b.h)-topY)

	frontCuboidZ := b.z + b.d
	frontCuboidDepth := a.d - (frontCuboidZ - a.z)
	if frontCuboidDepth > 0 {
		front := Cuboid{
			x: frontBackCuboidX,
			y: topY,
			z: frontCuboidZ,
			w: frontBackCuboidWidth,
			h: frontBackCuboidHeight,
			d: frontCuboidDepth,
		}
		res = append(res, &front)
	}

	backCuboidDepth := b.z - a.z
	if backCuboidDepth > 0 {
		back := Cuboid{
			x: frontBackCuboidX,
			y: topY,
			z: a.z,
			w: frontBackCuboidWidth,
			h: frontBackCuboidHeight,
			d: backCuboidDepth,
		}
		res = append(res, &back)
	}
	return res
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
		step.Cuboid = &Cuboid{}
		step.x = x1
		step.y = y1
		step.z = z1
		step.w = x2 - x1 + 1
		step.h = y2 - y1 + 1
		step.d = z2 - z1 + 1
		steps = append(steps, step)
	}

	ranges := []*Cuboid{}
	for _, step := range steps {
		if step.on {
			additionalRanges := []*Cuboid{step.Cuboid}
			for _, r := range ranges {
				updatedAdditionalRanges := []*Cuboid{}
				for _, a := range additionalRanges {
					updatedAdditionalRanges = append(updatedAdditionalRanges, subtract(a, r)...)
				}
				additionalRanges = updatedAdditionalRanges
			}
			ranges = append(ranges, additionalRanges...)
		} else {
			turnOff := []*Cuboid{step.Cuboid}
			for len(turnOff) > 0 {
				off := turnOff[0]
				turnOff = turnOff[1:]
				newRanges := []*Cuboid{}
				for i, on := range ranges {
					if intersection(on, off) != nil {
						turnOff = append(turnOff, subtract(off, on)...)
						newRanges = append(newRanges, subtract(on, off)...)
						newRanges = append(newRanges, ranges[i+1:]...)
						break
					}
					newRanges = append(newRanges, on)
				}
				ranges = newRanges
			}
		}
	}
	allSum := 0
	initSum := 0
	for _, a := range ranges {
		if a.x >= -50 && a.x+a.w <= 50 &&
			a.y >= -50 && a.y+a.h <= 50 &&
			a.z >= -50 && a.z+a.d <= 50 {
			initSum += a.w * a.h * a.d
		}
		allSum += a.w * a.h * a.d
	}
	fmt.Println(initSum)
	fmt.Println(allSum)
	return 0
}

func main() {
	os.Exit(run())
}
