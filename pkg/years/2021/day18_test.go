package y2021

import (
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay18(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 18)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("part_1", func(t *testing.T) {
		lines := strings.Split(strings.TrimSpace(string(input)), "\n")
		a, aN := ParseSnailNumbers(lines[0])
		for _, line := range lines[1:] {
			b, bN := ParseSnailNumbers(line)
			a, aN = SnailNumbersAdd(a, b, aN, bN)
			reduce := true
			for reduce {
				aN, reduce = snailNumbersReduce(aN)
			}
		}

		t.Logf("magnitude=%d", a.Magnitude())
	})

	t.Run("part_2", func(t *testing.T) {
		maxMagnitude := 0
		lines := strings.Split(strings.TrimSpace(string(input)), "\n")
		for idx, lineI := range lines {
			for jdx, lineJ := range lines {
				if idx == jdx {
					continue
				}
				a, aN := ParseSnailNumbers(lineI)
				b, bN := ParseSnailNumbers(lineJ)
				a, aN = SnailNumbersAdd(a, b, aN, bN)
				reduce := true
				for reduce {
					aN, reduce = snailNumbersReduce(aN)
				}

				if a.Magnitude() > maxMagnitude {
					maxMagnitude = a.Magnitude()
				}
			}
		}

		t.Logf("maxMagnitude=%d", maxMagnitude)
	})
}
