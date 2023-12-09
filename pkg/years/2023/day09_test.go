package y2023

import (
	"slices"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay09(t *testing.T) {
	t.Parallel()

	sequence := func(original []int) ([]int, bool) {
		out := make([]int, 0, len(original)-1)

		allZero := true
		for idx := 1; idx < len(original); idx++ {
			result := original[idx] - original[idx-1]
			out = append(out, result)

			if result != 0 {
				allZero = false
			}
		}

		return out, allZero
	}

	sequenceAll := func(original []int) [][]int {
		out := make([][]int, 1, 16)
		out[0] = original

		current := original
		allZero := false

		for !allZero {
			current, allZero = sequence(current)
			out = append(out, current)
		}

		return out
	}

	extrapolate := func(sequences [][]int) int {
		value := 0

		for idx := len(sequences) - 2; idx >= 0; idx-- {
			seq := sequences[idx]
			hist := seq[len(seq)-1]
			value = value + hist
		}

		return value
	}

	t.Run("test 1a", func(t *testing.T) {
		input := []int{0, 3, 6, 9, 12, 15}
		expected := []int{3, 3, 3, 3, 3}
		actual, allZero := sequence(input)

		if !slices.Equal(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}

		if allZero {
			t.Errorf("expected allZero to be false")
		}

		input = expected
		expected = []int{0, 0, 0, 0}

		actual, allZero = sequence(input)
		if !slices.Equal(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}

		if !allZero {
			t.Errorf("expected allZero to be true")
		}
	})

	t.Run("test 1b", func(t *testing.T) {
		input := []int{1, 3, 6, 10, 15, 21}
		expected := [][]int{
			{1, 3, 6, 10, 15, 21},
			{2, 3, 4, 5, 6},
			{1, 1, 1, 1},
			{0, 0, 0},
		}

		actual := sequenceAll(input)
		for idx, expectedItem := range expected {
			if !slices.Equal(actual[idx], expectedItem) {
				t.Errorf("expected %v, got %v", expectedItem, actual[idx])
			}
		}
	})

	t.Run("test 1c", func(t *testing.T) {
		ext := extrapolate([][]int{
			{1, 3, 6, 10, 15, 21},
			{2, 3, 4, 5, 6},
			{1, 1, 1, 1},
			{0, 0, 0},
		})

		if ext != 28 {
			t.Errorf("expected 28, got %d", ext)
		}
	})

	part1 := func(histories [][]int) int {
		total := 0
		for _, history := range histories {
			total += extrapolate(sequenceAll(history))
		}

		return total
	}

	t.Run("example 1", func(t *testing.T) {
		t.Log(part1(aoc.ParseLineNumbers(aoc.InputScanner("0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"))))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(part1(aoc.ParseLineNumbers(aoc.PuzzleScanner(2023, 9))))
	})

	extrapolateBackwards := func(sequences [][]int) int {
		value := 0

		for idx := len(sequences) - 2; idx >= 0; idx-- {
			seq := sequences[idx]
			hist := seq[0]
			value = hist - value
		}

		return value
	}

	part2 := func(histories [][]int) int {
		total := 0
		for _, history := range histories {
			total += extrapolateBackwards(sequenceAll(history))
		}

		return total
	}

	t.Run("example 2", func(t *testing.T) {
		t.Log(part2(aoc.ParseLineNumbers(aoc.InputScanner("0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"))))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(part2(aoc.ParseLineNumbers(aoc.PuzzleScanner(2023, 9))))
	})
}
