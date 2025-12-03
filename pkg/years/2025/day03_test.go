package y2025_test

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	aocparse "github.com/cpl/advent-of-code/pkg/aoc-parse"
	"go.sdls.io/pkg"
)

func TestSolve2025Day03(t *testing.T) {
	aoc.SolveExample(t, "example", 1, 357, "987654321111111\n811111111111119\n234234234234278\n818181911112111")
	aoc.SolveExample(t, "example", 2, "3121910778619", "987654321111111\n811111111111119\n234234234234278\n818181911112111")

	parser := aocparse.DigitsLine

	maxJoltBank2 := func(bank []uint8) int {
		var d0, d1 uint8

		l := len(bank)
		for idx, jolt := range bank {
			if jolt > d0 && idx <= l-2 {
				d0 = jolt
				d1 = 0
				continue
			}

			if jolt > d1 {
				d1 = jolt
			}
		}

		return int(d0*10) + int(d1)
	}

	aoc.Solve(t, "part 1", parser, func(banks [][]uint8) int {
		total := 0
		for _, bank := range banks {
			total += maxJoltBank2(bank)
		}

		return total
	})

	maxJoltBank12 := func(bank []uint8) pkg.UInt128 {
		var digits [12]int

		l := len(bank)
		for idx, jolt := range bank {
			startPos := 12 - min(l-idx, 12)

			for didx, d := range digits[startPos:] {
				if int(jolt) <= d {
					continue
				}

				digits[startPos+didx] = int(jolt)
				if startPos+didx == 12 {
					break
				}

				for ridx := startPos + didx + 1; ridx < len(digits); ridx++ {
					digits[ridx] = 0
				}

				break
			}
		}

		total := pkg.UInt128{}

		for _, digit := range digits {
			total = total.Mul64(10)
			total = total.Add64(uint64(digit))
		}

		return total
	}

	aoc.Solve(t, "part 2", parser, func(banks [][]uint8) string {
		total := pkg.UInt128{}
		for _, bank := range banks {
			total = total.Add(maxJoltBank12(bank))
		}

		return total.String()
	})
}
