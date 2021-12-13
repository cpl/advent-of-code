package y2021

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay13(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 13)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("part_1", func(t *testing.T) {
		tp := ParseTransparentPaper(string(input))
		tp.Fold()

		t.Logf("dots=%d", tp.Displayed())
	})
	t.Run("part_2", func(t *testing.T) {
		tp := ParseTransparentPaper(string(input))
		for tp.Fold() {
		}

		t.Logf("dots=%d", tp.Displayed())
		tp.Print()
	})
}
