package y2021

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay11(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 11)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("part_1", func(t *testing.T) {
		grid := ParseOctopusGrid(string(input))
		t.Logf("width=%d height=%d", len(grid.Grid[0]), len(grid.Grid))

		for iter := 0; iter < 100; iter++ {
			grid.ChargeAll()
		}

		t.Logf("flashes=%d", grid.TotalFlashes)
	})
	t.Run("part_2", func(t *testing.T) {
		grid := ParseOctopusGrid(string(input))
		t.Logf("width=%d height=%d", len(grid.Grid[0]), len(grid.Grid))

		flashedAll := len(grid.Grid) * len(grid.Grid[0])
		flashed := 0
		iter := 0
		for ; flashed != flashedAll; iter++ {
			flashed = grid.ChargeAll()
		}

		t.Logf("iter=%d", iter)
	})
}
