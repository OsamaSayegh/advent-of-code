package main

import (
	"fmt"
	"os"
	"strings"
)

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the input: %s\n", err)
		return 1
	}
	sections := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	template := sections[0]
	rules := make(map[string]string)
	for _, line := range strings.Split(sections[1], "\n") {
		rules[line[0:2]] = line[len(line)-1:]
	}
	polymers := make(map[string]int)
	counts := make(map[string]int)
	for i := 0; i < len(template); i++ {
		counts[template[i:i+1]] += 1
		if i == len(template)-1 {
			continue
		}
		polymer := template[i : i+2]
		polymers[polymer] += 1
	}
	for s := 1; s <= 40; s++ {
		newPolymers := make(map[string]int)
		for k, v := range polymers {
			counts[rules[k]] += v
			newPolymers[k[:1]+rules[k]] += v
			newPolymers[rules[k]+k[1:2]] += v
		}
		polymers = newPolymers
		if s == 10 || s == 40 {
			max := -1
			min := (1 << 63) - 1
			for _, v := range counts {
				if v > max {
					max = v
				}
				if v < min {
					min = v
				}
			}
			fmt.Println(max - min)
		}
	}
	return 0
}

func main() {
	os.Exit(run())
}
