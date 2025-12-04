package y2025_test

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	aocspace "github.com/cpl/advent-of-code/pkg/aoc-space"
)

func TestSolve2025Day04(t *testing.T) {
	aoc.SolveExample(t, "example", 1, 13, "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@.")
	aoc.SolveExample(t, "example", 2, 43, "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@.")

	aoc.Solve(t, "part 1", aocspace.Grid2Parser[byte](aocspace.Grid2CHF), func(grid *aocspace.Grid2[byte]) int {
		canRemove := 0

		grid.Iterate(func(point aocspace.Point[byte]) {
			isPaper := point.Value == '@'
			if !isPaper {
				return
			}

			neighbours := grid.GetNeighbours(point.Position)
			neighbourPaper := 0
			for _, neighbour := range neighbours {
				if neighbour.Value == '@' {
					neighbourPaper++
				}
			}

			if neighbourPaper < 4 {
				canRemove++
			}
		})

		return canRemove
	})

	aoc.Solve(t, "part 2", aocspace.Grid2Parser[byte](aocspace.Grid2CHF), func(grid *aocspace.Grid2[byte]) int {
		canRemove := 0
		toRemove := make([]aocspace.Vec, 0, 8)

		for {
			grid.Iterate(func(point aocspace.Point[byte]) {
				isPaper := point.Value == '@'
				if !isPaper {
					return
				}

				neighbours := grid.GetNeighbours(point.Position)
				neighbourPaper := 0
				for _, neighbour := range neighbours {
					if neighbour.Value == '@' {
						neighbourPaper++
					}
				}

				if neighbourPaper < 4 {
					toRemove = append(toRemove, point.Position)
					canRemove++
				}
			})

			if len(toRemove) == 0 {
				return canRemove
			}

			for _, pos := range toRemove {
				grid.Set('.', pos)
			}

			toRemove = toRemove[:0]
		}
	})
}
