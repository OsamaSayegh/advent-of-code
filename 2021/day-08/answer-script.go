package main

import (
	"fmt"
	"os"
	"strings"
)

func commonCount(str, substr string) int {
	count := 0
	for _, a := range strings.Split(str, "") {
		for _, b := range strings.Split(substr, "") {
			if a == b {
				count++
			}
		}
	}
	return count
}

func isEquivalent(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	return commonCount(a, b) == len(a)
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the input: %s\n", err)
		return 1
	}
	entries := make([]map[[10]string][4]string, 0)
	for _, l := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		patterns := [10]string{}
		values := [4]string{}
		for j, seg := range strings.SplitN(l, " | ", 2) {
			for k, val := range strings.Split(seg, " ") {
				if j == 0 {
					patterns[k] = val
				} else if j == 1 {
					values[k] = val
				} else {
					fmt.Fprintf(os.Stderr, "unexpected number of segments at line: %s\n", l)
					return 1
				}
			}
		}
		entry := map[[10]string][4]string{
			patterns: values,
		}
		entries = append(entries, entry)
	}

	count := 0
	countP2 := 0
	for _, entry := range entries {
		for patterns, values := range entry {
			numbers := [10]string{}
			zeroSixNine := []string{}
			twoThreeFive := []string{}

			for _, pattern := range patterns {
				length := len(pattern)
				if length == 2 {
					numbers[1] = pattern
				} else if length == 4 {
					numbers[4] = pattern
				} else if length == 3 {
					numbers[7] = pattern
				} else if length == 7 {
					numbers[8] = pattern
				} else if length == 6 {
					zeroSixNine = append(zeroSixNine, pattern)
				} else if length == 5 {
					twoThreeFive = append(twoThreeFive, pattern)
				}
			}

			for _, potential := range zeroSixNine {
				if commonCount(potential, numbers[1]) == 1 {
					numbers[6] = potential
				}
				if commonCount(potential, numbers[4]) == len(numbers[4]) {
					numbers[9] = potential
				}
			}

			for _, potential := range zeroSixNine {
				if potential != numbers[6] && potential != numbers[9] {
					numbers[0] = potential
				}
			}

			for _, potential := range twoThreeFive {
				if commonCount(potential, numbers[7]) == len(numbers[7]) {
					numbers[3] = potential
				} else if commonCount(numbers[6], potential) == len(potential) {
					numbers[5] = potential
				} else {
					numbers[2] = potential
				}
			}

			for _, n := range numbers {
				if len(n) == 0 {
					fmt.Fprintf(os.Stderr, "couldn't find all numbers for pattern %v\n", patterns)
					return 1
				}
			}
			entryValue := 0
			for _, val := range values {
				entryValue *= 10
				for i, n := range numbers {
					if isEquivalent(val, n) {
						entryValue += i
					}
				}
				size := len(val)
				if size == 2 || size == 4 || size == 3 || size == 7 {
					count += 1
				}
			}
			countP2 += entryValue
		}
	}
	fmt.Println(count)
	fmt.Println(countP2)
	return 0
}

func main() {
	os.Exit(run())
}
