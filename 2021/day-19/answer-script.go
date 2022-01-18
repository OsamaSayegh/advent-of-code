package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
    "math"
)

var Rotations = [...][3]int {
  // multiples of 90-degress for the x, y, z axes
  {0, 0, 0},
  {0, 1, 0},
  {0, 2, 0},
  {0, 3, 0},

  {0, 0, 1},
  {0, 1, 1},
  {0, 2, 1},
  {0, 3, 1},

  {0, 0, 2},
  {0, 1, 2},
  {0, 2, 2},
  {0, 3, 2},

  {0, 0, 3},
  {0, 1, 3},
  {0, 2, 3},
  {0, 3, 3},

  {1, 0, 0},
  {1, 1, 0},
  {1, 2, 0},
  {1, 3, 0},

  {3, 0, 0},
  {3, 1, 0},
  {3, 2, 0},
  {3, 3, 0},
}

type Beacon struct {
  position [3]int
}

type Scanner struct {
	beacons []*Beacon
    id int
}

func abs(v int) int {
  return int(math.Abs(float64(v)))
}

func parse(xyz string) (int, int, int) {
	split := strings.SplitN(xyz, ",", 3)
	x, err := strconv.Atoi(split[0])
	y, err := strconv.Atoi(split[1])
	z, err := strconv.Atoi(split[2])
	if err != nil {
		panic(err)
	}
	return x, y, z
}

func rz(pos [3]int, steps int) [3]int {
  for steps != 0 {
    pos[0], pos[1], pos[2] = -pos[1], pos[0], pos[2]
    steps -= 1
  }
  return pos
}

func rx(pos [3]int, steps int) [3]int {
  for steps != 0 {
    pos[0], pos[1], pos[2] = pos[0], pos[2], -pos[1]
    steps -= 1
  }
  return pos
}

func ry(pos [3]int, steps int) [3]int {
  for steps != 0 {
    pos[0], pos[1], pos[2] = -pos[2], pos[1], pos[0]
    steps -= 1
  }
  return pos
}

// 0 -> 1 -> 4 -> 2 -> 3
// scanner 1 @ 68,-1246,-43
// scanner 4 @ -20,-1133,1061
// scanner 2 @ 1105,-1205,1229
// scanner 3 @ -92,-2380,-20

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	scanners := []*Scanner{}
	for i, beacons := range strings.Split(strings.TrimSpace(string(data)), "\n\n") {
		scanner := &Scanner{}
        scanner.id = i
		for i, position := range strings.Split(beacons, "\n") {
			if i == 0 {
				continue
			}
            beacon := &Beacon{}
			x, y, z := parse(position)
            beacon.position = [3]int{x, y, z}
			scanner.beacons = append(scanner.beacons, beacon)
		}
		scanners = append(scanners, scanner)
	}
    // sorted := []*Scanner{scanners[0]}
    // for len(sorted) != len(scanners) {
    //   last := sorted[len(sorted)-1]
    //   for _, scanner := range scanners {
    //     skip := false
    //     for _, done := range sorted {
    //       if done.id == scanner.id {
    //         skip = true
    //         break
    //       }
    //     }
    //     if skip {
    //       continue
    //     }
    //     for _, rotation := range Rotations {
    //       rotated := &Scanner{}
    //       for _, b := range scanner.beacons {
    //         pos := b.position
    //         x, y, z := rotation[0], rotation[1], rotation[2]
    //         if z != 0 {
    //           pos = rz(pos, z)
    //         }
    //         if x != 0 {
    //           pos = rx(pos, x)
    //         }
    //         if y != 0 {
    //           pos = ry(pos, y)
    //         }
    //         rotated.beacons = append(rotated.beacons, &Beacon{pos})
    //       }
    //     }
    //   }
    // }
    arr := [3]int{1,2,3}
    fmt.Println(ry(rx(arr, 2), 1))
    fmt.Println(ry(rz(arr, 2), 3))
	return 0
}

func main() {
	os.Exit(run())
}
