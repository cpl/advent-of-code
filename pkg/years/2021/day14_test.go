package y2021

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay14(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 14)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("part_1", func(t *testing.T) {
		poly := ParsePolymer(string(input))

		for iter := 0; iter < 10; iter++ {
			out := poly.Apply()
			poly.str = out
		}

		minRune, minCount, maxRune, maxCount := poly.Elements()

		t.Logf("min: %d %s", minCount, string(minRune))
		t.Logf("max: %d %s", maxCount, string(maxRune))
		t.Logf("solution=%d", maxCount-minCount)
	})
	t.Run("part_1_pairs", func(t *testing.T) {
		poly := ParsePolymer(string(input))

		for iter := 0; iter < 10; iter++ {
			poly.ApplyPairs()
		}

		minRune, minCount, maxRune, maxCount := poly.ElementsPairs()

		t.Logf("min: %d %s", minCount, string(minRune))
		t.Logf("max: %d %s", maxCount, string(maxRune))
		t.Logf("solution=%d", maxCount-minCount)
	})
	t.Run("part_2", func(t *testing.T) {
		poly := ParsePolymer(string(input))

		for iter := 0; iter < 40; iter++ {
			poly.ApplyPairs()
		}

		minRune, minCount, maxRune, maxCount := poly.ElementsPairs()

		t.Logf("min: %d %s", minCount, string(minRune))
		t.Logf("max: %d %s", maxCount, string(maxRune))
		t.Logf("solution=%d", maxCount-minCount)
	})
}
