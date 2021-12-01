package y2021

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay01(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 1)
	if err != nil {
		t.Fatal(err)
	}

	numbers := ParseNumbers(input)

	t.Run("part_1", func(t *testing.T) {
		up, _, _ := NumberUpOrDown(numbers)
		t.Logf("result='%d'", up)
	})
	t.Run("part_2", func(t *testing.T) {
		threeSums := NumberThreeMeasureSums(numbers)
		up, _, _ := NumberUpOrDown(threeSums)
		t.Logf("result='%d'", up)
	})
}
