package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Config [5][11]byte

type State struct {
	config    Config
	amphipods [][2]int
	cost      int
}

type Item struct {
	value    *State
	priority int
	index    int
}

type PriorityQueue []*Item

const (
	A = 'A'
	B = 'B'
	C = 'C'
	D = 'D'
	W = '#'
	S = '\000'
)

var FinalConfig1 Config = Config{
	[11]byte{S, S, S, S, S, S, S, S, S, S, S},
	[11]byte{W, W, A, W, B, W, C, W, D, W, W},
	[11]byte{W, W, A, W, B, W, C, W, D, W, W},
	[11]byte{W, W, S, W, S, W, S, W, S, W, W},
	[11]byte{W, W, S, W, S, W, S, W, S, W, W},
}
var FinalConfig2 Config = Config{
	[11]byte{S, S, S, S, S, S, S, S, S, S, S},
	[11]byte{W, W, A, W, B, W, C, W, D, W, W},
	[11]byte{W, W, A, W, B, W, C, W, D, W, W},
	[11]byte{W, W, A, W, B, W, C, W, D, W, W},
	[11]byte{W, W, A, W, B, W, C, W, D, W, W},
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value *State, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func amphipodStuckInRoom(amphipodLoc [2]int, state *State) bool {
	for i := amphipodLoc[0] - 1; i > 0; i-- {
		if state.config[i][amphipodLoc[1]] != S {
			return true
		}
	}
	return false
}

func amphipodMultiplier(amphipod byte) int {
	switch amphipod {
	case A:
		return 1
	case B:
		return 10
	case C:
		return 100
	case D:
		return 1000
	default:
		panic(fmt.Errorf("unknown multiplier for amphipod type %c", amphipod))
	}
}

func (s *State) clone() *State {
	state := &State{config: s.config, cost: s.cost}
	state.amphipods = make([][2]int, len(s.amphipods))
	copy(state.amphipods, s.amphipods)
	return state
}

func (s *State) draw() {
	fmt.Println("#############")
	for ri, row := range s.config {
		if ri <= 1 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
		for ci, char := range row {
			if ri > 1 && (ci == 0 || ci+1 == len(row)) {
				fmt.Print(" ")
			} else if char == S {
				fmt.Print(".")
			} else {
				fmt.Printf("%c", char)
			}
		}
		if ri <= 1 {
			fmt.Print("#")
		}
		fmt.Println()
	}
	fmt.Println("-------------")
}

func (s *State) nextStates(p1 bool) []*State {
	states := []*State{}
	for ai, amphipodLoc := range s.amphipods {
		amphipodType := s.config[amphipodLoc[0]][amphipodLoc[1]]
		inRoom := amphipodLoc[0] > 0
		multiplier := amphipodMultiplier(amphipodType)
		targetRoom := -1
		bottom := 2
		if !p1 {
			bottom = 4
		}
		switch amphipodType {
		case A:
			targetRoom = 2
		case B:
			targetRoom = 4
		case C:
			targetRoom = 6
		case D:
			targetRoom = 8
		default:
			panic(fmt.Errorf("unknown amphipod type %c", amphipodType))
		}
		if inRoom {
			if amphipodStuckInRoom(amphipodLoc, s) {
				continue
			}
			if targetRoom == amphipodLoc[1] {
				shouldMove := false
				for i := 1; i <= bottom; i++ {
					occupant := s.config[i][targetRoom]
					if occupant != S && occupant != amphipodType {
						shouldMove = true
						break
					}
				}
				if !shouldMove {
					continue
				}
			}
			for i, pos := range s.config[0] {
				if pos != S {
					continue
				}
				if i%2 == 0 && i != 0 && (i+1) != 11 {
					continue
				}
				loc := [2]int{0, i}
				cost := amphipodLoc[0]
				c1 := loc[1]
				c2 := amphipodLoc[1]
				if c1 > c2 {
					c1, c2 = c2, c1
				}
				cost += c2 - c1
				blocked := false
				for c1 < c2 {
					if s.config[0][c1] != S {
						blocked = true
						break
					}
					c1++
				}
				if blocked {
					continue
				}
				clone := s.clone()
				clone.cost += cost * multiplier
				clone.amphipods[ai] = loc
				clone.config[0][loc[1]] = amphipodType
				clone.config[amphipodLoc[0]][amphipodLoc[1]] = S
				states = append(states, clone)
			}
		}
		roomSpot := [2]int{1, targetRoom}
		occupant := s.config[roomSpot[0]][roomSpot[1]]
		roomAvailable := occupant == S
		if !roomAvailable {
			continue
		}
		foundBottom := false
		for i := 1; i < bottom; i++ {
			occupant = s.config[i+1][roomSpot[1]]
			if occupant != amphipodType && occupant != S {
				roomAvailable = false
				break
			}
			if occupant == S {
				// sanity check
				if foundBottom {
					panic(fmt.Errorf("detected gap in room (%d,%d)!", roomSpot[0], roomSpot[1]))
				}
				roomSpot[0] += 1
			} else {
				// same amphipod type
				foundBottom = true
			}
		}
		if !roomAvailable {
			continue
		}
		cost := amphipodLoc[0] + roomSpot[0]
		start := amphipodLoc[1]
		dest := roomSpot[1]
		dir := 1
		if start > dest {
			dir = -1
			cost += start - dest
		} else {
			cost += dest - start
		}
		if !inRoom {
			// amphipod in the hallway, don't check its position for blockers
			start += dir
		}
		skip := false
		for true {
			if s.config[0][start] != S {
				skip = true
				break
			}
			if start == dest {
				break
			}
			start += dir
		}
		if skip {
			continue
		}
		clone := s.clone()
		clone.cost += cost * multiplier
		clone.amphipods[ai] = roomSpot
		clone.config[roomSpot[0]][roomSpot[1]] = amphipodType
		clone.config[amphipodLoc[0]][amphipodLoc[1]] = S
		states = append(states, clone)
	}
	return states
}

func (s *State) insertAmphipod(loc [2]int, amphipodType byte) {
	s.config[loc[0]][loc[1]] = amphipodType
	s.amphipods = append(s.amphipods, loc)
}

func (s *State) solve(p1 bool) int {
	seen := make(map[Config]int)
	best := 1 << 32
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{value: s, index: 0, priority: s.cost}
	heap.Init(&pq)
	final := FinalConfig1
	if !p1 {
		final = FinalConfig2
	}
	for len(pq) > 0 {
		top := heap.Pop(&pq).(*Item).value
		for _, next := range top.nextStates(p1) {
			if next.config == final && next.cost < best {
				best = next.cost
				continue
			}
			if val, ok := seen[next.config]; !ok || val > next.cost {
				item := &Item{value: next, priority: next.cost}
				seen[next.config] = next.cost
				heap.Push(&pq, item)
			}
		}
	}
	if best == (1 << 32) {
		panic("couldn't find best!")
	}
	return best
}

func (s *State) solve1() int {
	return s.solve(true)
}

func (s *State) solve2() int {
	for i := 2; i <= 8; i += 2 {
		s.config[2][i], s.config[4][i] = S, s.config[2][i]
		updated := false
		for ai, amphipodLoc := range s.amphipods {
			if amphipodLoc == [2]int{2, i} {
				updated = true
				s.amphipods[ai] = [2]int{4, i}
				break
			}
		}
		if !updated {
			panic("inconsistent state")
		}
	}
	s.insertAmphipod([2]int{2, 2}, D)
	s.insertAmphipod([2]int{3, 2}, D)
	s.insertAmphipod([2]int{2, 4}, C)
	s.insertAmphipod([2]int{3, 4}, B)
	s.insertAmphipod([2]int{2, 6}, B)
	s.insertAmphipod([2]int{3, 6}, A)
	s.insertAmphipod([2]int{2, 8}, A)
	s.insertAmphipod([2]int{3, 8}, C)
	return s.solve(false)
}

func run() int {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	grid := strings.Split(strings.TrimSpace(string(data)), "\n")[1:]
	grid = grid[:len(grid)-1]
	if len(grid) != 3 {
		panic(fmt.Errorf("expected grid to have 3 rows, got %d", len(grid)))
	}

	state := &State{}
	for r := 1; r <= 2; r++ {
		for c := 3; c <= 9; c += 2 {
			switch grid[r][c] {
			case A:
				state.config[r][c-1] = A
			case B:
				state.config[r][c-1] = B
			case C:
				state.config[r][c-1] = C
			case D:
				state.config[r][c-1] = D
			default:
				panic(fmt.Errorf("unknown amphipod type at (%d,%d)", r, c))
			}
			state.amphipods = append(state.amphipods, [2]int{r, c - 1})
		}
	}
	for r := 1; r <= 4; r++ {
		for c := 1; c <= 9; c += 2 {
			if c == 1 {
				state.config[r][c-1] = W
			}
			state.config[r][c] = W
			if c == 9 {
				state.config[r][c+1] = W
			}
		}
	}
	fmt.Println(state.solve1())
	fmt.Println(state.solve2())
	return 0
}

func main() {
	os.Exit(run())
}
