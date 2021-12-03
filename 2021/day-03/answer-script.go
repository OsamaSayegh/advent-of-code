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
	report := make([]int64, len(lines))
	for i, l := range lines {
		parsed, err := strconv.ParseInt(l, 2, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing integer %s at %d: %s\n", l, i, err)
			return 1
		}
		report[i] = parsed
	}
	gamma := int64(0)
	bitSize := len(lines[0])
	for i := int64(1 << (bitSize - 1)); i > 0; i >>= 1 {
		ones := 0
		for _, val := range report {
			if i&val != 0 {
				ones++
			}
		}
		if len(report)-ones < ones {
			gamma |= i
		}
	}

	epsilon := (^gamma) & ((1 << bitSize) - 1)
	fmt.Println(gamma * epsilon)

	oxy := make([]int64, len(lines))
	co2 := make([]int64, len(lines))
	copy(oxy, report)
	copy(co2, report)
	for i := int64(1 << (bitSize - 1)); i > 0; i >>= 1 {
		if len(oxy) > 1 {
			oxy_ones := []int64{}
			oxy_zeros := []int64{}
			for _, val := range oxy {
				if i&val != 0 {
					oxy_ones = append(oxy_ones, val)
				} else {
					oxy_zeros = append(oxy_zeros, val)
				}
			}
			if len(oxy_ones) >= len(oxy_zeros) {
				oxy = oxy_ones
			} else {
				oxy = oxy_zeros
			}
		}

		if len(co2) > 1 {
			co2_ones := []int64{}
			co2_zeros := []int64{}
			for _, val := range co2 {
				if i&val != 0 {
					co2_ones = append(co2_ones, val)
				} else {
					co2_zeros = append(co2_zeros, val)
				}
			}
			if len(co2_zeros) <= len(co2_ones) {
				co2 = co2_zeros
			} else {
				co2 = co2_ones
			}
		}

		if len(co2) == 1 && len(oxy) == 1 {
			break
		}
	}
	lifeSupport := oxy[0] * co2[0]
	fmt.Println(lifeSupport)
	return 0
}

func main() {
	os.Exit(run())
}
