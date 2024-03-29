package y2022

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay16(t *testing.T) {
	type valve struct {
		name       string
		flow       int
		neighbours map[string]*valve
	}

	parse := func(line string) *valve {
		v := &valve{}
		v.neighbours = make(map[string]*valve)

		_, _ = fmt.Sscanf(line, "Valve %s has flow rate=%d; tunnels lead to valve", &v.name, &v.flow)

		s := line[strings.Index(line, "valve")+5:]
		if s[0] == 's' {
			s = s[2:]
		} else {
			s = s[1:]
		}

		split := strings.Split(s, ", ")
		for _, vName := range split {
			v.neighbours[vName] = nil
		}

		return v
	}

	type cave struct {
		valves map[string]*valve
	}

	interpret := func(valves []*valve) cave {
		known := make(map[string]*valve)

		for _, v := range valves {
			known[v.name] = v
		}

		for _, v := range valves {
			for vName := range v.neighbours {
				v.neighbours[vName] = known[vName]
			}
		}

		return cave{valves: known}
	}

	distanceCache := make(map[string]int)
	distance := func(c cave, from, to string) int {
		key := from + "-" + to
		if d, ok := distanceCache[key]; ok {
			return d
		}

		if from == to {
			return 0
		}

		visited := make(map[string]bool)
		visited[from] = true

		queue := make([]struct {
			v *valve
			d int
		}, 0)

		for _, v := range c.valves[from].neighbours {
			queue = append(queue, struct {
				v *valve
				d int
			}{v: v, d: 1})
		}

		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]

			if v.v.name == to {
				distanceCache[key] = v.d
				return v.d
			}

			for _, n := range v.v.neighbours {
				if !visited[n.name] {
					visited[n.name] = true
					queue = append(queue, struct {
						v *valve
						d int
					}{v: n, d: v.d + 1})
				}
			}
		}

		distanceCache[key] = -1
		return -1
	}

	var nextBest func(cave, *valve, int, int, []string) (int, []string)
	nextBest = func(c cave, current *valve, remaining, flowCum int, open []string) (int, []string) {
		if remaining < 2 {
			return flowCum, open
		}
		if remaining == 2 {
			isOpen := false
			for _, o := range open {
				if o == current.name {
					isOpen = true
					break
				}
			}
			if isOpen {
				return flowCum, open
			}

			return flowCum + current.flow, append(open, current.name)
		}

		calculations := make(map[string]struct {
			score    int
			distance int
		})

		for name, v := range c.valves {
			if name == current.name {
				continue
			}

			if v.flow == 0 {
				continue
			}

			isOpen := false
			for _, o := range open {
				if o == v.name {
					isOpen = true
					break
				}
			}
			if isOpen {
				continue
			}

			d := distance(c, current.name, name)
			if remaining-d-1 <= 0 {
				continue
			}

			score := v.flow * (remaining - d - 1)
			if score <= 0 {
				continue
			}

			calculations[name] = struct {
				score    int
				distance int
			}{score: score, distance: d}
		}

		bestFlowCum := flowCum
		bestOpen := open

		for name, calc := range calculations {
			if calc.score <= 0 {
				continue
			}

			nextBestFlowCum, nextOpen := nextBest(c, c.valves[name], remaining-calc.distance-1, flowCum+calc.score, append(open, name))

			if nextBestFlowCum > bestFlowCum {
				bestFlowCum = nextBestFlowCum
				bestOpen = nextOpen
			}
		}

		return bestFlowCum, bestOpen
	}

	part1 := func(c cave) int {
		bestFlow, _ := nextBest(c, c.valves["AA"], 30, 0, nil)
		return bestFlow
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "Valve AA has flow rate=0; tunnels lead to valves DD, II, BB\nValve BB has flow rate=13; tunnels lead to valves CC, AA\nValve CC has flow rate=2; tunnels lead to valves DD, BB\nValve DD has flow rate=20; tunnels lead to valves CC, AA, EE\nValve EE has flow rate=3; tunnels lead to valves FF, DD\nValve FF has flow rate=0; tunnels lead to valves EE, GG\nValve GG has flow rate=0; tunnels lead to valves FF, HH\nValve HH has flow rate=22; tunnel leads to valve GG\nValve II has flow rate=0; tunnels lead to valves AA, JJ\nValve JJ has flow rate=21; tunnel leads to valve II"
		t.Log(part1(interpret(aoc.ParseLines(aoc.InputScanner(input), parse))))
	})

	t.Run("Custom 1", func(t *testing.T) {
		input := "Valve AA has flow rate=0; tunnels lead to valves BB\nValve BB has flow rate=1; tunnels lead to valves AA"
		t.Log(part1(interpret(aoc.ParseLines(aoc.InputScanner(input), parse))))
	})

	// 1949 too high
	// 1750 too low
	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(interpret(aoc.ParseLines(aoc.PuzzleScanner(2022, 16), parse))))
	})
}
