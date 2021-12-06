package y2021

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay06(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 6)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("part_1", func(t *testing.T) {
		bank := ParseLanternFishBank(string(input))
		bank.AgeDays(80)
		t.Logf("solution=%d", bank.Total())
	})
	t.Run("part_2", func(t *testing.T) {
		bank := ParseLanternFishBank(string(input))
		bank.AgeDays(256)
		t.Logf("solution=%d", bank.Total())
	})
}
