package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
    "math"
)

type Beacon struct {
  x int
  y int
  z int
}

type Scanner struct {
	beacons []*Beacon
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

// 0 -> 1 -> 4 -> 2 -> 3

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	scanners := []*Scanner{}
	for _, beacons := range strings.Split(strings.TrimSpace(string(data)), "\n\n") {
		scanner := &Scanner{}
		for i, position := range strings.Split(beacons, "\n") {
			if i == 0 {
				continue
			}
			x, y, z := parse(position)
            beacon := &Beacon{}
            beacon.x = x
            beacon.y = y
            beacon.z = z
			scanner.beacons = append(scanner.beacons, beacon)
		}
		scanners = append(scanners, scanner)
	}
    fmt.Println(scanners[0].beacons[0])
    fmt.Println(len(scanners))
	return 0
}

func main() {
	os.Exit(run())
}
