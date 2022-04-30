package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const MATCH_THRESHOLD = 12

var Rotations = [24][3][3]int{
	{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	},
	{
		{0, 0, 1},
		{0, 1, 0},
		{-1, 0, 0},
	},
	{
		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, -1},
	},
	{
		{0, 0, -1},
		{0, 1, 0},
		{1, 0, 0},
	},
	{
		{0, -1, 0},
		{1, 0, 0},
		{0, 0, 1},
	},
	{
		{0, 0, 1},
		{1, 0, 0},
		{0, 1, 0},
	},
	{
		{0, 1, 0},
		{1, 0, 0},
		{0, 0, -1},
	},
	{
		{0, 0, -1},
		{1, 0, 0},
		{0, -1, 0},
	},
	{
		{0, 1, 0},
		{-1, 0, 0},
		{0, 0, 1},
	},
	{
		{0, 0, 1},
		{-1, 0, 0},
		{0, -1, 0},
	},
	{
		{0, -1, 0},
		{-1, 0, 0},
		{0, 0, -1},
	},
	{
		{0, 0, -1},
		{-1, 0, 0},
		{0, 1, 0},
	},
	{
		{1, 0, 0},
		{0, 0, -1},
		{0, 1, 0},
	},
	{
		{0, 1, 0},
		{0, 0, -1},
		{-1, 0, 0},
	},
	{
		{-1, 0, 0},
		{0, 0, -1},
		{0, -1, 0},
	},
	{
		{0, -1, 0},
		{0, 0, -1},
		{1, 0, 0},
	},
	{
		{1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	},
	{
		{0, 0, -1},
		{0, -1, 0},
		{-1, 0, 0},
	},
	{
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, 1},
	},
	{
		{0, 0, 1},
		{0, -1, 0},
		{1, 0, 0},
	},
	{
		{1, 0, 0},
		{0, 0, 1},
		{0, -1, 0},
	},
	{
		{0, -1, 0},
		{0, 0, 1},
		{-1, 0, 0},
	},
	{
		{-1, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
	},
	{
		{0, 1, 0},
		{0, 0, 1},
		{1, 0, 0},
	},
}

type Beacon struct {
	position     [3]int
	orientations [][3]int
}

type Scanner struct {
	beacons   []*Beacon
	distances []int
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

func rotate(pos [3]int, rotation [3][3]int) [3]int {
	return [3]int{
		pos[0]*rotation[0][0] + pos[1]*rotation[0][1] + pos[2]*rotation[0][2],
		pos[0]*rotation[1][0] + pos[1]*rotation[1][1] + pos[2]*rotation[1][2],
		pos[0]*rotation[2][0] + pos[1]*rotation[2][1] + pos[2]*rotation[2][2],
	}
}

func manhattan(p1 [3]int, p2 [3]int) int {
	x := abs(p1[0] - p2[0])
	y := abs(p1[1] - p2[1])
	z := abs(p1[2] - p2[2])
	return x + y + z
}

func run() int {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	scanners := []*Scanner{}
	for i, beacons := range strings.Split(strings.TrimSpace(string(data)), "\n\n") {
		scanner := &Scanner{}
		for j, position := range strings.Split(beacons, "\n") {
			if j == 0 {
				continue
			}
			beacon := &Beacon{}
			x, y, z := parse(position)
			beacon.position = [3]int{x, y, z}
			if i != 0 {
				for _, rotation := range Rotations {
					beacon.orientations = append(beacon.orientations, rotate(beacon.position, rotation))
				}
			}
			scanner.beacons = append(scanner.beacons, beacon)
		}
		scanners = append(scanners, scanner)
	}
	for _, s := range scanners {
		distances := []int{}
		for i1, b1 := range s.beacons {
			if i1+1 == len(s.beacons) {
				continue
			}
			for _, b2 := range s.beacons[i1+1:] {
				distances = append(distances, manhattan(b1.position, b2.position))
			}
		}
		s.distances = distances
	}
	composite := scanners[0]
	scanners[0] = nil
	scanners = scanners[1:]
	scannersDistances := [][3]int{}
	for len(scanners) > 0 {
		lastLen := len(scanners)
		for si, scanner := range scanners {
			for _, cb := range composite.beacons {
				for ri := range Rotations {
					for _, b1 := range scanner.beacons {
						offset := [3]int{
							b1.orientations[ri][0] - cb.position[0],
							b1.orientations[ri][1] - cb.position[1],
							b1.orientations[ri][2] - cb.position[2],
						}
						transformed := [][3]int{}
						for _, b2 := range scanner.beacons {
							transformed = append(
								transformed,
								[3]int{
									b2.orientations[ri][0] - offset[0],
									b2.orientations[ri][1] - offset[1],
									b2.orientations[ri][2] - offset[2],
								},
							)
						}
						matchCount := 0
						duplicates := make(map[int]bool)
						for ti, t := range transformed {
							for _, cb2 := range composite.beacons {
								if t == cb2.position {
									matchCount++
									duplicates[ti] = true
								}
							}
						}
						if matchCount >= MATCH_THRESHOLD {
							scannersDistances = append(scannersDistances, offset)
							scanners[si] = scanners[len(scanners)-1]
							scanners = scanners[:len(scanners)-1]
							newBeacons := []*Beacon{}
							for ti, t := range transformed {
								if duplicates[ti] {
									continue
								}
								beacon := &Beacon{}
								beacon.position = t
								newBeacons = append(newBeacons, beacon)
							}
							for nbi, nb1 := range newBeacons {
								for _, cb2 := range composite.beacons {
									composite.distances = append(
										composite.distances,
										manhattan(nb1.position, cb2.position),
									)
								}
								if nbi+1 == len(newBeacons) {
									continue
								}
								for _, nb2 := range newBeacons[nbi+1:] {
									composite.distances = append(
										composite.distances,
										manhattan(nb1.position, nb2.position),
									)
								}
							}
							composite.beacons = append(
								composite.beacons,
								newBeacons...,
							)
							goto findNextScanner
						}
					}
				}
			}
		}
		if len(scanners) == lastLen {
			panic(fmt.Errorf("infinite loop detected! lastLen=%d.", lastLen))
		}
	findNextScanner:
		continue
	}
	fmt.Println(len(composite.beacons))
	largest := 0
	for _, d1 := range scannersDistances {
		for _, d2 := range scannersDistances {
			distance := manhattan(d1, d2)
			if distance > largest {
				largest = distance
			}
		}
	}
	fmt.Println(largest)
	return 0
}

func main() {
	os.Exit(run())
}
