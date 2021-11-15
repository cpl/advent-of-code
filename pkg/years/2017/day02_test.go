package year2017

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay02(t *testing.T) {
	input, err := aoc.MetaGetInput(2017, 2)
	if err != nil {
		t.Fatal(err)
	}

	numbers := ParseNumberListByLine(string(input))

	t.Run("part_01", func(t *testing.T) {
		sum := 0
		for _, line := range numbers {
			min, max := MinMaxInList(line)
			sum += max - min
		}
		t.Logf("solution: %d", sum)
	})

	t.Run("part_02", func(t *testing.T) {
		sum := 0
		for _, line := range numbers {
			num0, num1 := DivisibleByEachOtherInList(line)
			sum += num0 / num1
		}
		t.Logf("solution: %d", sum)
	})
}
