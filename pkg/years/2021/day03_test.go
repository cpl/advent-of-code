package y2021

import (
	"strconv"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay03(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 3)
	if err != nil {
		t.Fatal(err)
	}

	binary := ParseBinaryListAsStrings(input)

	t.Run("part_1", func(t *testing.T) {
		gamma := CalculateMostCommonBits(binary)
		epsilon := FlipBits(gamma)

		gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
		epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

		t.Logf("gamma=%d epsilon=%d solution=%d", gammaInt, epsilonInt, gammaInt*epsilonInt)
	})
	t.Run("part_2", func(t *testing.T) {
		o2rating := WalkBinaryList(binary, true)
		co2rating := WalkBinaryList(binary, false)

		o2ratingInt, _ := strconv.ParseInt(o2rating, 2, 64)
		co2ratingInt, _ := strconv.ParseInt(co2rating, 2, 64)

		t.Logf("o2rating=%d co2rating=%d solution=%d", o2ratingInt, co2ratingInt, o2ratingInt*co2ratingInt)
	})
}
