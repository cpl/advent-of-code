package year2017

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay01(t *testing.T) {
	input, err := aoc.MetaGetInput(2017, 1)
	if err != nil {
		t.Fatal(err)
	}

	numbers := ParseDigitList(input)

	t.Run("part_1", func(t *testing.T) {
		t.Logf("solution: %d", SumMatchNext(numbers))
	})
	t.Run("part_2", func(t *testing.T) {
		t.Logf("solution: %d", SumMatchHalfSplitList(numbers))
	})
}
