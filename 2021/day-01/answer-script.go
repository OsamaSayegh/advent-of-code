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
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	measurements := make([]int, len(lines))
	for i, l := range lines {
		parsed, err := strconv.Atoi(l)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing integer %s at %d: %s\n", l, i, err)
			return 1
		}
		measurements[i] = parsed
	}

	increases := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			increases++
		}
	}
	fmt.Printf("%d\n", increases)

	windows := make([]int, len(measurements)-2)
	for i := 0; i < len(measurements)-2; i++ {
		windows[i] = measurements[i] + measurements[i+1] + measurements[i+2]
	}

	winIncreases := 0
	for i := 1; i < len(windows); i++ {
		if windows[i] > windows[i-1] {
			winIncreases++
		}
	}
	fmt.Printf("%d\n", winIncreases)
	return 0
}

func main() {
	os.Exit(run())
}
