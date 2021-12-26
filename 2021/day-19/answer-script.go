package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
    "math"
)

type Distance struct {
  value int
  toBeacon *Beacon
}

type Beacon struct {
  position [3]int
  distances []int
}

type Scanner struct {
	beacons []*Beacon
}

func abs(v int) int {
  return int(math.Abs(float64(v)))
}

func areNeighbors(s1, s2 *Scanner) bool {
  for _, b1 := range s1.beacons {
    for _, b2 := range s2.beacons {
      count := 0
      for _, d1 := range b1.distances {
        for _, d2 := range b2.distances {
          if d1 == d2 {
            count++
          }
        }
      }
      if count >= 12 {
        return true
      }
    }
  }
  return false
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

func commonCount(a, b []int) int {
  res := 0
  for _, aa := range a {
    for _, bb := range b {
      if aa == bb {
        res++
      }
    }
  }
  return res
}

// 0 -> 1 -> 4 -> 2 -> 3
func findNext(index int, matches map[int]int, done []int, scanners []*Scanner) bool {
  potentials := []int{}
  for i, scanner := range scanners {
    if i == index { continue }
    skip := false
    for _, n := range done {
      if i == n {
        skip = true
        break
      }
    }
    if skip { continue }
    // if _, ok := matches[i]; ok { continue }
    if areNeighbors(scanners[index], scanner) {
      potentials = append(potentials, i)
    }
  }
  fmt.Println(index, potentials, done)
  if len(potentials) == 1 {
    done = append(done, index)
    matches[index] = potentials[0]
    return findNext(potentials[0], matches,done, scanners)
  } else if len(potentials) > 1 {
    for _, p := range potentials {
      newMatches := make(map[int]int)
      for k, v := range matches {
        newMatches[k] = v
      }
      newDone := make([]int, len(done))
      copy(newDone, done)
      newDone = append(newDone, index)
      if findNext(p, newMatches, newDone, scanners) {
        for k, v := range newMatches {
          matches[k] = v
        }
        matches[index] = p
        return true
      }
    }
    return false
  } else {
    if len(done) + 1 == len(scanners) - 1 {
      last := -1
      for i, _ := range scanners {
        skip := false
        if i == index { continue }
        for _, n := range done {
          if n == i {
            skip = true
            break
          }
        }
        if skip { continue }
        last = i
        break
      }
      if last == -1 {
        panic("couldn't find last scanner!")
      }
      matches[index] = last
      return true
    } else {
      return false
    }
  }
}

func run() int {
	data, err := os.ReadFile("input2.txt")
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
            beacon.position = [3]int{x,y,z}
			scanner.beacons = append(scanner.beacons, beacon)
		}
        for i, beacon := range scanner.beacons {
          for j, beacon2 := range scanner.beacons {
            if i == j { continue }
            x := abs(beacon.position[0] - beacon2.position[0])
            y := abs(beacon.position[1] - beacon2.position[1])
            z := abs(beacon.position[2] - beacon2.position[2])
            beacon.distances = append(beacon.distances, x + y + z)
          }
        }
		scanners = append(scanners, scanner)
	}
    left := make([]int, 0, len(scanners))
    for i,_ := range scanners {
      left = append(left, i)
    }
    // matches := make(map[int]int)
    // done := []int{}
    // findNext(0, matches, done, scanners)
    // fmt.Println(matches)
    pos := make([][]int, 0, len(scanners))
    for i, s1 := range scanners {
      gg := []int{}
      for j, s2 := range scanners {
        if j == i { continue }
        if areNeighbors(s1, s2) {
          gg = append(gg, j)
        }
      }
      pos = append(pos, gg)
    }
    for k, v := range pos {
      fmt.Printf("%02d: %v\n", k, v)
    }
    // 0 -> 1 -> 4 -> 2 -> 3
    // sequence := []int{}
    // matchesMap := make(map[int]int)
    // for len(left) > 0 {
    //   neighborFor := left[0]
    //   left = left[1:]
    //   matches := []int{}
    //   for i, scanner := range scanners {
    //     if i == neighborFor { continue }
    //     skip := false
    //     for _, n := range sequence {
    //       if n == i {
    //         skip = true
    //         break
    //       }
    //     }
    //     if skip { continue }
    //     if areNeighbors(scanner, scanners[neighborFor]) {
    //       matches = append(matches, i)
    //     }
    //   }
    //   if len(matches) == 1 {
    //     matchesMap[neighborFor] = matches[0]
    //     sequence = append(sequence, neighborFor)
    //     fmt.Println(neighborFor, sequence)
    //   } else {
    //     left = append(left, neighborFor)
    //     if len(left) == 1 {
    //       break
    //     }
    //   }
    // }
    // fmt.Println(matchesMap)
    // neighbors := []int{0}
    // for len(neighbors) != len(scanners) {
    //   fmt.Println(neighbors)
    //   neighborFor := neighbors[len(neighbors)-1]
    //   done := false
    //   for i, scanner := range scanners {
    //     if i == neighborFor { continue }
    //     skip := false
    //     for _, n := range neighbors {
    //       if i == n {
    //         skip = true
    //         break
    //       }
    //     }
    //     if skip { continue }
    //     for _, b1 := range scanners[neighborFor].beacons {
    //       for _, b2 := range scanner.beacons {
    //         common := 0
    //         for _, d1 := range b1.distances {
    //           for _, d2 := range b2.distances {
    //             if d1.value == d2.value {
    //               common++
    //             }
    //           }
    //         }
    //         if common >= 11 {
    //           neighbors = append(neighbors, i)
    //           done = true
    //           break
    //         }
    //       }
    //       if done { break }
    //     }
    //     if done { break }
    //   }
    // }
    // fmt.Println(neighbors)

    // for i, scanner := range scanners {
      // if i == 0 { continue }
      // matches := make(map[[3]int][3]int)
      // max := -1
      // for i, b1 := range scanners[2].beacons {
      //   for j, b2 := range scanners[4].beacons {
      //     inters := 0
      //     for _, d1 := range b1.distances {
      //       for _, d2 := range b2.distances {
      //         if d1 == d2 {
      //           inters++
      //         }
      //       }
      //     }
      //     if inters > max {
      //       max = inters
      //     }
      //     fmt.Println(inters, i, j)
      //   }
      // }
      // fmt.Println(max)
      // fmt.Println(scanners[2].beacons[0].distances)
      // fmt.Println(scanners[3].beacons[0].distances)
    // }
	return 0
}

func main() {
	os.Exit(run())
}
