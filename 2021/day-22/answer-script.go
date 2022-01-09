package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Step struct {
    *Cuboid
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

func subtract2(a, b Cuboid) {
  xChunks := [3]int{0, 0, 0}
  yChunks := [3]int{0, 0, 0}
  zChunks := [3]int{0, 0, 0}
  xChunks[0] = b.x - a.x
  xChunks[1] = max(a.x, b.x)
  xChunks[1] = min((a.x + a.w) - xChunks[1], (b.x + b.w) - xChunks[1])
  xChunks[2] = (a.x + a.w) - (b.x + b.w)

  yChunks[0] = b.y - a.y
  yChunks[1] = max(a.y, b.y)
  yChunks[1] = min((a.y + a.h) - yChunks[1], (b.y + b.h) - yChunks[1])
  yChunks[2] = (a.y + a.h) - (b.y + b.h)

  zChunks[0] = b.z - a.z
  zChunks[1] = max(a.z, b.z)
  zChunks[1] = min((a.z + a.h) - zChunks[1], (b.z + b.h) - zChunks[1])
  zChunks[2] = (a.z + a.h) - (b.z + b.h)
}

// func intersect(a, b *Cuboid) bool {
//   return ((b.x >= a.x && (a.x + a.w) >= b.x) || ((b.x + b.w) >= a.x && (a.x + a.w) >= (b.x + b.w))) &&
//   ((b.y >= a.y && (a.y + a.h) >= b.y) || ((b.y + b.h) >= a.y && (a.y + a.h) >= (b.y + b.h))) &&
//   ((b.z >= a.z && (a.z + a.d) >= b.z) || ((b.z + b.d) >= a.z && (a.z + a.d) >= (b.z + b.d)))
// }

func intersection(a, b *Cuboid) *Cuboid {
  x1 := max(a.x, b.x)
  y1 := max(a.y, b.y)
  z1 := max(a.z, b.z)
  x2 := min(a.x + a.w, b.x + b.w)
  y2 := min(a.y + a.h, b.y + b.h)
  z2 := min(a.z + a.d, b.z + b.d)
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
  if intersect := intersection(a, b); intersect == nil {
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
  if bottomCuboidHeight > 0 && bottomCuboidY < a.y + a.h {
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

  topY := a.y
  if topY < b.y {
    topY = b.y
  }
  bottomY := b.y + b.h
  if bottomY > a.y + a.h {
    bottomY = a.y + a.h
  }

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

  frontBackCuboidX := a.x
  if frontBackCuboidX < b.x {
    frontBackCuboidX = b.x
  }

  frontBackCuboidWidth := (a.x + a.w) - frontBackCuboidX
  if frontBackCuboidWidth2 := (b.x + b.w) - frontBackCuboidX; frontBackCuboidWidth2 < frontBackCuboidWidth {
    frontBackCuboidWidth = frontBackCuboidWidth2
  }

  frontBackCuboidHeight := (a.y + a.h) - topY
  if frontBackCuboidHeight2 := (b.y + b.h) - topY; frontBackCuboidHeight2 < frontBackCuboidHeight {
    frontBackCuboidHeight = frontBackCuboidHeight2
  }

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

func add(a, b *Cuboid) []*Cuboid {
  intersect := intersection(a, b)
  if intersect == nil {
    dupA := *a
    dupB := *b
    return []*Cuboid{&dupA, &dupB}
  }
  res := []*Cuboid{intersect}
  res = append(res, subtract(a, b)...)
  res = append(res, subtract(b, a)...)
  return res
}

func run() int {
	data, err := os.ReadFile("input2.txt")
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
        turnedOff := []*Cuboid{step.Cuboid}
        updatedRanges := []*Cuboid{}
        stopAt := len(ranges)
        i := 0
        for i < stopAt {
          current := ranges[i]
        }
        for i, r := range ranges {
          updatedTurnedOff := []*Cuboid{}
          for _, t := range turnedOff {
            // fmt.Println(i, t, r)
            // for _, fdx := range subtract(t, r) {
            //   fmt.Println(fdx)
            // }
            // fmt.Println("------------------------------------")
            updatedTurnedOff = append(updatedTurnedOff, subtract(t, r)...)
            updatedRanges = append(updatedRanges, subtract(r, t)...)
          }
          if false && i == 0 {
            fmt.Println("updatedRanges start")
            for _, uu := range updatedRanges {
              fmt.Println(uu)
            }
            fmt.Println("updatedRanges end\n")
            fmt.Println("updatedTurnedOff start")
            for _, uu := range updatedTurnedOff {
              fmt.Println(uu)
            }
            fmt.Println("updatedTurnedOff end\n")
          }
          turnedOff = updatedTurnedOff
        }
        ranges = updatedRanges
      }
    }
      // fmt.Println(len(ranges))
      // sum := 0
      for _, a := range ranges {
        fmt.Println(a)
      //   // sum += a.w * a.h * a.d
      }
      fmt.Println("---------------------------------------")
      // fmt.Println("sum", sum)
    // cuboids := []Cuboid{}
	return 0
}

func main() {
	os.Exit(run())
}
