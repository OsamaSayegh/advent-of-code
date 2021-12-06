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
	fish := []int{}
	for i, f := range strings.Split(strings.TrimSpace(string(data)), ",") {
		parsed, err := strconv.Atoi(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing fish integer %s at %d: %s\n", f, i, err)
			return 1
		}
		fish = append(fish, parsed)
	}
	buckets := [9]int{}
	for _, f := range fish {
		buckets[f] += 1
	}
	for day := 1; day <= 256; day++ {
		newBuckets := [9]int{}
		newBuckets[7] = buckets[8]
		newBuckets[6] = buckets[7] + buckets[0]
		newBuckets[5] = buckets[6]
		newBuckets[4] = buckets[5]
		newBuckets[3] = buckets[4]
		newBuckets[2] = buckets[3]
		newBuckets[1] = buckets[2]
		newBuckets[0] = buckets[1]
		newBuckets[8] = buckets[0]
		buckets = newBuckets
		if day == 80 || day == 256 {
			sum := 0
			for _, d := range buckets {
				sum += d
			}
			fmt.Println(sum)
		}
	}
	return 0
}

func main() {
	os.Exit(run())
}
