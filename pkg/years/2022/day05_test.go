package y2022

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay05(t *testing.T) {
	extract := func(s string) (containers, moves []string) {
		idx := strings.Index(s, "move")
		containers = strings.Split(s[:idx], "\n")
		containers = containers[:len(containers)-2]
		moves = strings.Split(s[idx:], "\n")
		return
	}

	parseYard := func(containers []string) [][]string {
		yard := make([][]string, 10)

		for cIdx := len(containers) - 2; cIdx >= 0; cIdx-- {
			container := containers[cIdx]
			idx := 0
			pos := 1

			for pos < len(container) {
				symbol := string(container[pos])
				if symbol != " " {
					yard[idx] = append(yard[idx], symbol)
				}

				pos += 4
				idx += 1
			}
		}

		return yard
	}

	makeMoves := func(moves []string, fn func(count, from, to int)) {
		for _, s := range moves {
			if s == "" {
				continue
			}

			var count, from, to int
			_, _ = fmt.Sscanf(s, "move %d from %d to %d", &count, &from, &to)
			fn(count, from, to)
		}
	}

	topCrates := func(yard [][]string) string {
		var builder strings.Builder
		for _, col := range yard {
			if len(col) == 0 {
				continue
			}
			builder.WriteString(col[len(col)-1])
		}

		return builder.String()
	}

	part1 := func(input string) string {
		containers, moves := extract(input)
		yard := parseYard(containers)
		makeMoves(moves, func(count, from, to int) {
			for count > 0 {
				yard[to-1] = append(yard[to-1], yard[from-1][len(yard[from-1])-1])
				yard[from-1] = yard[from-1][:len(yard[from-1])-1]
				count--
			}
		})

		return topCrates(yard)
	}

	part2 := func(input string) string {
		containers, moves := extract(input)
		yard := parseYard(containers)
		makeMoves(moves, func(count, from, to int) {
			fCol := yard[from-1]
			yard[to-1] = append(yard[to-1], fCol[len(fCol)-count:]...)
			yard[from-1] = fCol[:len(fCol)-count]
		})

		return topCrates(yard)
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2"
		t.Log(part1(input))
	})

	t.Run("Part 1", func(t *testing.T) {
		input := aoc.PuzzleStringRaw(2022, 5)
		t.Log(part1(input))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2"
		t.Log(part2(input))
	})

	t.Run("Part 2", func(t *testing.T) {
		input := aoc.PuzzleStringRaw(2022, 5)
		t.Log(part2(input))
	})
}
