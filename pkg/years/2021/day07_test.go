package y2021

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay07(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 7)
	if err != nil {
		t.Fatal(err)
	}

	crabs := ParseCrabs(string(input))

	t.Run("part_1", func(t *testing.T) {
		position, fuel := crabs.Align()
		t.Logf("position=%d fuel=%d", position, fuel)
	})
	t.Run("part_2", func(t *testing.T) {
		position, fuel := crabs.AlignLinear()
		t.Logf("position=%d fuel=%d", position, fuel)
	})
}
