package y2024

import (
	"slices"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	aocparse "github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func TestSolve2024Day01(t *testing.T) {
	// aoc.SolveAutoSubmit(t)
	// aoc.SolveNoExamples(t)
	// aoc.SolveExamplesOnly(t)

	aoc.SolveExample(t, "example", 1,
		11, "3   4\n4   3\n2   5\n1   3\n3   9\n3   3")
	aoc.SolveExample(t, "example", 2,
		31, "3   4\n4   3\n2   5\n1   3\n3   9\n3   3")

	aoc.Solve(t, "part 1", aocparse.NumbersColumn(2), func(lists [][]int) int {
		listA := lists[0]
		listB := lists[1]

		slices.Sort(listA)
		slices.Sort(listB)

		total := 0
		for idx := range listA {
			vA := listA[idx]
			vB := listB[idx]

			distance := vA - vB
			if distance < 0 {
				distance = -distance
			}

			total += distance
		}

		return total
	})

	aoc.Solve(t, "part 2", aocparse.NumbersColumn(2), func(lists [][]int) int {
		countSet := make(map[int]int)
		for _, num := range lists[1] {
			countSet[num]++
		}

		score := 0
		for _, num := range lists[0] {
			score += countSet[num] * num
		}

		return score
	})
}
