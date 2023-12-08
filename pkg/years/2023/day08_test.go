package y2023

import (
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	aoc_math "github.com/cpl/advent-of-code/pkg/aoc-math"
)

func TestSolveDay08(t *testing.T) {
	t.Parallel()

	type node struct {
		name  string
		left  string
		right string
	}

	parse := func(input string) (string, map[string]*node) {
		lines := strings.Split(input, "\n")

		lookup := make(map[string]*node)
		for _, line := range lines[2:] {
			lookup[line[:3]] = &node{
				name:  line[:3],
				left:  line[7:10],
				right: line[12:15],
			}
		}

		return lines[0], lookup
	}

	part1 := func(instructions string, lookup map[string]*node) int {
		current := lookup["AAA"]

		l := len(instructions)
		for steps := 0; steps < l*1000; steps++ {
			c := instructions[steps%l]

			switch c {
			case 'L':
				current = lookup[current.left]
			case 'R':
				current = lookup[current.right]
			}

			if current.name == "ZZZ" {
				return steps + 1
			}
		}

		return -1
	}

	part2 := func(instructions string, lookup map[string]*node) int64 {
		starts := make([]*node, 0, 16)
		startsEnds := make([]int64, 0, 16)

		for k, v := range lookup {
			if k[2] == 'A' {
				starts = append(starts, v)
				startsEnds = append(startsEnds, 0)
			}
		}

		step := func(c byte, target *node) (*node, bool) {
			var next *node

			switch c {
			case 'L':
				next = lookup[target.left]
			case 'R':
				next = lookup[target.right]
			}

			return next, next.name[2] == 'Z'
		}

		l := len(instructions)
		doneCount := 0

		for steps := 0; steps < l*1000; steps++ {
			c := instructions[steps%l]

			for idx, start := range starts {
				var done bool

				starts[idx], done = step(c, start)
				if done && startsEnds[idx] == 0 {
					startsEnds[idx] = int64(steps + 1)
					doneCount++
				}
			}

			if doneCount == len(starts) {
				break
			}
		}

		return aoc_math.LCM(startsEnds...)
	}

	t.Run("example 1a", func(t *testing.T) {
		t.Log(part1(parse("RL\n\nAAA = (BBB, CCC)\nBBB = (DDD, EEE)\nCCC = (ZZZ, GGG)\nDDD = (DDD, DDD)\nEEE = (EEE, EEE)\nGGG = (GGG, GGG)\nZZZ = (ZZZ, ZZZ)")))
	})

	t.Run("example 1b", func(t *testing.T) {
		t.Log(part1(parse("LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)")))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(part1(parse(aoc.PuzzleString(2023, 8))))
	})

	t.Run("example 2", func(t *testing.T) {
		t.Log(part2(parse("LR\n\n11A = (11B, XXX)\n11B = (XXX, 11Z)\n11Z = (11B, XXX)\n22A = (22B, XXX)\n22B = (22C, 22C)\n22C = (22Z, 22Z)\n22Z = (22B, 22B)\nXXX = (XXX, XXX)")))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(part2(parse(aoc.PuzzleString(2023, 8))))
	})
}
