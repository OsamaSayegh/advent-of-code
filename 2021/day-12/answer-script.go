package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Start = "start"
	End   = "end"
)

func isSmall(cave string) bool {
	if cave == Start || cave == End {
		return false
	}
	return (cave[0] & (1 << 5)) != 0
}

func solvePart1(caves map[string][]string) {
	paths := [][]string{}
	queue := [][]string{{Start}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastCave := path[len(path)-1]
		for _, nextCave := range caves[lastCave] {
			if nextCave == Start {
				continue
			}
			seenSmall := false
			loop := false
			for i := 0; i < len(path)-1; i++ {
				if path[i] == lastCave && path[i+1] == nextCave {
					loop = true
					break
				}
				if isSmall(nextCave) && path[i] == nextCave {
					seenSmall = true
					break
				}
			}
			if loop || seenSmall {
				continue
			}
			pathDup := make([]string, len(path))
			copy(pathDup, path)
			pathDup = append(pathDup, nextCave)
			if nextCave == End {
				paths = append(paths, pathDup)
			} else {
				queue = append(queue, pathDup)
			}
		}
	}
	fmt.Println(len(paths))
}

func solvePart2(caves map[string][]string) {
	paths := [][]string{}
	queue := [][]string{{Start}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastCave := path[len(path)-1]
		for _, nextCave := range caves[lastCave] {
			if nextCave == Start {
				continue
			}
			visitedMap := make(map[string]int)
			visitedMap[nextCave] += 1
			twiceCount := 0
			invalid := false
			for i := 0; i < len(path); i++ {
				if isSmall(path[i]) {
					visitedMap[path[i]] += 1
					count := visitedMap[path[i]]
					if count > 2 {
						invalid = true
						break
					}
					if count == 2 {
						twiceCount += 1
						if twiceCount > 1 {
							invalid = true
							break
						}
					}
				}
			}
			if invalid {
				continue
			}
			pathDup := make([]string, len(path))
			copy(pathDup, path)
			pathDup = append(pathDup, nextCave)
			if nextCave == End {
				paths = append(paths, pathDup)
			} else {
				queue = append(queue, pathDup)
			}
		}
	}
	fmt.Println(len(paths))
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the input: %s\n", err)
		return 1
	}
	caves := make(map[string][]string)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		path := strings.Split(line, "-")
		start := path[0]
		end := path[1]
		caves[start] = append(caves[start], end)
		caves[end] = append(caves[end], start)
	}
	solvePart1(caves)
	solvePart2(caves)
	return 0
}

func main() {
	os.Exit(run())
}
