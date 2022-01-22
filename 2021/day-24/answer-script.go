package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	op   string
	reg1 string
	reg2 string
	val  int
}

type MagicNumbers struct {
	a int
	b int
	c int
}

const (
	Inp = "inp"
	Add = "add"
	Mul = "mul"
	Div = "div"
	Mod = "mod"
	Eql = "eql"
)

func solve(magic [14]MagicNumbers, start, increment int, check func(int) bool) int {
	digits := [14]int{}
	z := 0
	cancelled := []int{}
	for chunk := 0; chunk < 14; chunk++ {
		a := magic[chunk].a
		b := magic[chunk].b
		c := magic[chunk].c
		if digits[chunk] == 0 {
			digits[chunk] = start
		}
		w := digits[chunk]
		if a == 26 {
			w = z%26 + b
			jumpBack := cancelled[len(cancelled)-1]
			cancelled = cancelled[:len(cancelled)-1]
			if check(w) {
				digits[jumpBack] += increment
				z = 0
				chunk = -1
				continue
			}
			digits[chunk] = w
		} else {
			cancelled = append(cancelled, chunk)
		}
		if (z%26)+b == w {
			z /= a
		} else {
			z /= a
			z *= 26
			z += w + c
		}
	}
	number := 0
	for _, d := range digits {
		number *= 10
		number += d
	}
	return number
}

func run() int {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	instructions := []Instruction{}
	for _, inst := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		instruction := Instruction{}
		parts := strings.SplitN(inst, " ", 3)
		if parts[0] == Inp {
			instruction.op = Inp
		} else if parts[0] == Add {
			instruction.op = Add
		} else if parts[0] == Mul {
			instruction.op = Mul
		} else if parts[0] == Div {
			instruction.op = Div
		} else if parts[0] == Mod {
			instruction.op = Mod
		} else if parts[0] == Eql {
			instruction.op = Eql
		} else {
			panic(fmt.Errorf("error: unknown op %s", parts[0]))
		}
		reg1 := parts[1]
		if reg1 != "x" && reg1 != "y" && reg1 != "z" && reg1 != "w" {
			panic(fmt.Errorf("error: unknown reg %c", reg1))
		}
		instruction.reg1 = reg1
		if instruction.op != Inp {
			reg2 := parts[2]
			if reg2 == "x" || reg2 == "y" || reg2 == "z" || reg2 == "w" {
				instruction.reg2 = reg2
			} else {
				val, err := strconv.Atoi(parts[2])
				if err != nil {
					panic(err)
				}
				instruction.val = val
			}
		}
		instructions = append(instructions, instruction)
	}
	magic := [14]MagicNumbers{}
	for chunk := 0; chunk < 14; chunk++ {
		i := chunk * 18
		magic[chunk] = MagicNumbers{
			a: instructions[i+4].val,
			b: instructions[i+5].val,
			c: instructions[i+15].val,
		}
	}
	fmt.Println(solve(magic, 9, -1, func(w int) bool { return w > 9 }))
	fmt.Println(solve(magic, 1, 1, func(w int) bool { return w < 1 }))
	return 0
}

func main() {
	os.Exit(run())
}
