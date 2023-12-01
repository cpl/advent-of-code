package y2023

import (
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay01(t *testing.T) {
	solvePart1 := func(lines []string) int {
		values := make([]int, 0, len(lines))

		for _, line := range lines {
			var c1, c2 rune

			for _, c := range line {
				if c >= '0' && c <= '9' {
					if c1 == 0 {
						c1 = c
					}

					c2 = c
				}
			}

			v := (c1-'0')*10 + (c2 - '0')
			values = append(values, int(v))
		}

		sum := 0
		for _, v := range values {
			sum += v
		}

		return sum
	}

	digitsSet := map[string]int{
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	part2Value := func(s string) (int, int) {
		for k, v := range digitsSet {
			if strings.HasPrefix(s, k) {
				return v, len(k)
			}
		}

		return -1, 0
	}

	part2LineToValue := func(line string) int {
		c1, c2 := -1, -1

		ptr := 0
		for ptr < len(line) {
			v, skip := part2Value(line[ptr:])
			if v == -1 {
				ptr++
				continue
			}
			ptr += max(skip-1, 1)

			if c1 == -1 {
				c1 = v
			}

			c2 = v
		}

		return c1*10 + c2
	}

	solvePart2 := func(lines []string) int {
		values := make([]int, 0, len(lines))

		for _, line := range lines {
			values = append(values, part2LineToValue(line))
		}

		sum := 0
		for _, v := range values {
			sum += v
		}

		return sum
	}

	t.Run("example 1", func(t *testing.T) {
		t.Log(solvePart1([]string{
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet",
		}))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(solvePart1(aoc.PuzzleStringSliceNewline(2023, 1)))
	})

	t.Run("example 2", func(t *testing.T) {
		t.Log(solvePart2([]string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
		}))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(solvePart2(aoc.PuzzleStringSliceNewline(2023, 1)))
	})
}
