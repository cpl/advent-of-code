package y2025_test

import (
	"strconv"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	aocparse "github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func TestSolve2025Day01(t *testing.T) {
	// aoc.SolveExamplesOnly(t)

	aoc.SolveExample(t, "example", 1, 3,
		"L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82")

	aoc.SolveExample(t, "example", 2, 23,
		"L1000\nR1000\nR50\nR100\nR2\nL3")

	aoc.SolveExample(t, "example", 2, 6,
		"L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82")

	type move struct {
		direction byte
		value     int
	}

	parser := aocparse.EachLine[move](func(line string) move {
		char := line[0]
		v := line[1:]
		value, _ := strconv.Atoi(v)
		return move{
			direction: char,
			value:     value,
		}
	})

	aoc.Solve(t, "part 1", parser, func(moves []move) int {
		position := 50
		count := 0

		for _, m := range moves {
			mv := m.value % 100
			if m.direction == 'L' {
				mv = -mv
			}
			position += mv

			if position < 0 {
				position = 100 + position
			} else if position >= 100 {
				position = position % 100
			}

			if position == 0 {
				count++
			}
		}

		return count
	})

	aoc.Solve(t, "part 2", parser, func(moves []move) int {
		position := 50
		count := 0

		for _, m := range moves {
			mv := m.value

			if mv >= 100 {
				loops := mv / 100
				mv = mv % 100
				count += loops

				if mv == 0 {
					continue
				}
			}

			isZero := position == 0
			if m.direction == 'L' {
				position -= mv
			} else {
				position += mv
			}

			if position < 0 {
				if !isZero {
					count++
				}

				position = 100 + position
				t.Log("dial is rotated", string(m.direction), m.value, "to position", position, "|", count)
				continue
			} else if position >= 100 {
				if !isZero {
					count++
				}

				position = position % 100
				t.Log("dial is rotated", string(m.direction), m.value, "to position", position, "|", count)
				continue
			}

			t.Log("dial is rotated", string(m.direction), m.value, "to position", position, "|", count)
			if position == 0 {
				count++
			}
		}

		return count // 6663 too high
	})
}
