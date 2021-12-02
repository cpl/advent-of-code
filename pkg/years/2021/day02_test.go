package y2021

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay02(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 2)
	if err != nil {
		t.Fatal(err)
	}

	steps := ParseNavigation(input)

	t.Run("part_1", func(t *testing.T) {
		x, depth := CalculateNavigation(steps)
		t.Logf("x=%d, depth=%d, solution=%d", x, depth, x*depth)
	})
	t.Run("part_2", func(t *testing.T) {
		x, depth := CalculateNavigation2(steps)
		t.Logf("x=%d, depth=%d, solution=%d", x, depth, x*depth)
	})
}
