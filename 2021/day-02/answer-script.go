package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FORWARD = "forward"
	UP      = "up"
	DOWN    = "down"
)

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the input: %s\n", err)
		return 1
	}
	input := strings.TrimSpace(string(data))
	commands := strings.Split(input, "\n")

	depth := 0
	horizPos := 0

	aim := 0
	depth2 := 0
	horizPos2 := 0
	for _, command := range commands {
		split := strings.SplitN(command, " ", 2)
		dir := split[0]
		change, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to parse integer from %s\n", command)
			return 1
		}
		if dir == FORWARD {
			horizPos += change
			horizPos2 += change
			depth2 += aim * change
		} else if dir == DOWN {
			depth += change
			aim += change
		} else if dir == UP {
			depth -= change
			aim -= change
		} else {
			fmt.Fprintf(os.Stderr, "unknown direction %s from command %s\n", dir, command)
			return 1
		}
	}
	fmt.Println(depth * horizPos)
	fmt.Println(depth2 * horizPos2)
	return 0
}

func main() {
	os.Exit(run())
}
