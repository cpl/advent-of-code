package y2022

import (
	"sort"
	"strconv"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func solveDay01Part1(calories []int) int {
	var sum int
	var sums []int
	var max int

	for _, c := range calories {
		if c != 0 {
			sum += c
		} else {
			sums = append(sums, sum)
			if sum > max {
				max = sum
			}
			sum = 0
		}
	}

	return max
}

func solveDay01Part2(calories []int) int {
	var sum int
	var sums []int

	for _, c := range calories {
		if c != 0 {
			sum += c
		} else {
			sums = append(sums, sum)
			sum = 0
		}
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	return sums[0] + sums[1] + sums[2]
}

func TestSolveDay01(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		calories := aoc.ParseLines(aoc.PuzzleScanner(2022, 1), func(line string) int {
			if line == "" {
				return 0
			}

			num, _ := strconv.Atoi(line)
			return num
		})

		t.Log(solveDay01Part1(calories))
	})

	t.Run("Part 2", func(t *testing.T) {
		calories := aoc.ParseLines(aoc.PuzzleScanner(2022, 1), func(line string) int {
			if line == "" {
				return 0
			}

			num, _ := strconv.Atoi(line)
			return num
		})

		t.Log(solveDay01Part2(calories))
	})
}
