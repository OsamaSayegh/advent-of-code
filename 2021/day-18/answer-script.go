package main

import (
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	left   interface{}
	right  interface{}
	parent *Pair
}

func copyPair(original *Pair) *Pair {
	dup := &Pair{}

	leftNumber, leftIsNumber := original.left.(int)
	rightNumber, rightIsNumber := original.right.(int)
	leftPair, leftIsPair := original.left.(*Pair)
	rightPair, rightIsPair := original.right.(*Pair)

	if !leftIsNumber && !leftIsPair {
		panic(fmt.Errorf("left child of %v is niether a number nor nested Pair!", original))
	}
	if !rightIsNumber && !rightIsPair {
		panic(fmt.Errorf("right child of %v is niether a number nor nested Pair!", original))
	}

	if leftIsNumber {
		dup.left = leftNumber
	} else {
		left := copyPair(leftPair)
		dup.left = left
		left.parent = dup
	}
	if rightIsNumber {
		dup.right = rightNumber
	} else {
		right := copyPair(rightPair)
		dup.right = right
		right.parent = dup
	}
	return dup
}

func sumPairs(left, right *Pair) *Pair {
	pair := &Pair{left: left, right: right}
	left.parent = pair
	right.parent = pair
	reduce(pair)
	return pair
}

func magnitude(pair *Pair) int {
	leftNumber, leftIsNumber := pair.left.(int)
	rightNumber, rightIsNumber := pair.right.(int)
	leftPair, leftIsPair := pair.left.(*Pair)
	rightPair, rightIsPair := pair.right.(*Pair)

	if !leftIsNumber && !leftIsPair {
		panic(fmt.Errorf("left child of %v is niether a number nor nested Pair!", pair))
	}
	if !rightIsNumber && !rightIsPair {
		panic(fmt.Errorf("right child of %v is niether a number nor nested Pair!", pair))
	}

	left := 0
	if leftIsNumber {
		left = leftNumber
	} else {
		left = magnitude(leftPair)
	}
	right := 0
	if rightIsNumber {
		right = rightNumber
	} else {
		right = magnitude(rightPair)
	}
	return 3*left + 2*right
}

func reduce(pair *Pair) {
	again := explode(pair, 0)
	again = split(pair, pair, 0) || again
	if again {
		reduce(pair)
	}
}

func explode(pair *Pair, depth int) bool {
	leftNumber, leftIsNumber := pair.left.(int)
	rightNumber, rightIsNumber := pair.right.(int)
	leftPair, leftIsPair := pair.left.(*Pair)
	rightPair, rightIsPair := pair.right.(*Pair)

	if !leftIsNumber && !leftIsPair {
		panic(fmt.Errorf("left child of %v is niether a number nor nested Pair!", pair))
	}
	if !rightIsNumber && !rightIsPair {
		panic(fmt.Errorf("right child of %v is niether a number nor nested Pair!", pair))
	}

	if leftIsPair || rightIsPair {
		exploded := 0
		if leftIsPair && explode(leftPair, depth+1) {
			exploded += 1
		}
		if rightIsPair && explode(rightPair, depth+1) {
			exploded += 1
		}
		return exploded > 0
	}

	if depth >= 4 {
		child := pair
		parent := pair.parent
		foundRightMatch := false
		foundLeftMatch := false
		for parent != nil {
			if parent.left == child {
				if !foundRightMatch {
					foundRightMatch = true
					target, ok := parent.right.(int)
					if ok {
						parent.right = target + rightNumber
					} else {
						targetParent := parent.right.(*Pair)
						target, ok := targetParent.left.(int)
						for !ok {
							targetParent = targetParent.left.(*Pair)
							target, ok = targetParent.left.(int)
						}
						targetParent.left = target + rightNumber
					}
				}
			} else {
				if !foundLeftMatch {
					foundLeftMatch = true
					target, ok := parent.left.(int)
					if ok {
						parent.left = target + leftNumber
					} else {
						targetParent := parent.left.(*Pair)
						target, ok := targetParent.right.(int)
						for !ok {
							targetParent = targetParent.right.(*Pair)
							target, ok = targetParent.right.(int)
						}
						targetParent.right = target + leftNumber
					}
				}
			}
			if foundLeftMatch && foundRightMatch {
				break
			}
			child = parent
			parent = parent.parent
		}
		if pair == pair.parent.left {
			pair.parent.left = 0
		} else {
			pair.parent.right = 0
		}
		return true
	} else {
		return false
	}
}

func split(pair *Pair, root *Pair, depth int) bool {
	leftNumber, leftIsNumber := pair.left.(int)
	rightNumber, rightIsNumber := pair.right.(int)
	leftPair, leftIsPair := pair.left.(*Pair)
	rightPair, rightIsPair := pair.right.(*Pair)

	if leftIsNumber && leftNumber >= 10 {
		newLeft := leftNumber / 2
		newRight := newLeft
		if leftNumber%2 != 0 {
			newRight++
		}
		pair.left = &Pair{left: newLeft, right: newRight, parent: pair}
		if depth >= 3 {
			explode(pair.left.(*Pair), depth+1)
		}
		return true
	}
	if leftIsPair {
		if split(leftPair, root, depth+1) {
			return true
		}
	}
	if rightIsNumber && rightNumber >= 10 {
		newLeft := rightNumber / 2
		newRight := newLeft
		if rightNumber%2 != 0 {
			newRight++
		}
		pair.right = &Pair{left: newLeft, right: newRight, parent: pair}
		if depth >= 3 {
			explode(pair.right.(*Pair), depth+1)
		}
		return true
	}
	if rightIsPair {
		if split(rightPair, root, depth+1) {
			return true
		}
	}
	return false
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading the input: %s\n", err)
		return 1
	}
	pairs := []*Pair{}
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		stack := []*Pair{}
		second := false
		for _, char := range line {
			if char == '[' {
				stack = append(stack, &Pair{})
				second = false
			} else if char == ',' {
				second = true
			} else if char == ']' {
				if len(stack) > 1 {
					last := len(stack) - 1
					stack[last].parent = stack[last-1]
					if stack[last-1].left == nil {
						stack[last-1].left = stack[last]
					} else if stack[last-1].right == nil {
						stack[last-1].right = stack[last]
					} else {
						panic(fmt.Errorf("unexpected state when parsing line %s: parent=%v", line, stack[last-1]))
					}
					stack = stack[:last]
				}
			} else {
				val := int(char - 48)
				last := len(stack) - 1
				if second {
					stack[last].right = val
				} else {
					stack[last].left = val
				}
			}
		}
		pairs = append(pairs, stack[0])
	}

	sum := copyPair(pairs[0])
	for i := 1; i < len(pairs); i++ {
		sum = sumPairs(sum, copyPair(pairs[i]))
	}
	fmt.Println(magnitude(sum))

	max := -1
	for i := 0; i < len(pairs); i++ {
		for j := i + 1; j < len(pairs); j++ {
			a := copyPair(pairs[i])
			b := copyPair(pairs[j])
			mag := magnitude(sumPairs(a, b))
			if mag > max {
				max = mag
			}
			a = copyPair(pairs[j])
			b = copyPair(pairs[i])
			mag = magnitude(sumPairs(a, b))
			if mag > max {
				max = mag
			}
		}
	}
	fmt.Println(max)
	return 0
}

func main() {
	os.Exit(run())
}
