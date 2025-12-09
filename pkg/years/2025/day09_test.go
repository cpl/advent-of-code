package y2025_test

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	aocspace "github.com/cpl/advent-of-code/pkg/aoc-space"
)

func TestSolve2025Day09(t *testing.T) {
	// aoc.SolveExamplesOnly(t)

	aoc.SolveExample(t, "example", 1, int64(50), "7,1\n11,1\n11,7\n9,7\n9,5\n2,5\n2,3\n7,3")
	aoc.SolveExample(t, "example", 2, int64(24), "7,1\n11,1\n11,7\n9,7\n9,5\n2,5\n2,3\n7,3")

	mhd := func(a, b int64) int64 {
		if a > b {
			return (a - b) + 1
		}
		return (b - a) + 1
	}

	aoc.Solve(t, "part 1", aocspace.Vec2Parser(0), func(points []aocspace.Vec) int64 {
		largestArea := int64(0)

		for idx, p0 := range points {
			for _, p1 := range points[idx:] {
				area := mhd(p0.X, p1.X) * mhd(p0.Y, p1.Y)

				if area > largestArea {
					largestArea = area
				}
			}
		}

		return largestArea
	})

	type Rect struct {
		P0, P1 aocspace.Vec
	}

	aoc.Solve(t, "part 2", aocspace.Vec2Parser(0), func(points []aocspace.Vec) int64 {
		filled := make([]Rect, 0, len(points)+1)
		for idx, p := range points[:len(points)-1] {
			filled = append(filled, Rect{
				P0: p,
				P1: points[idx+1],
			})
		}
		filled = append(filled, Rect{
			P0: points[len(points)-1],
			P1: points[0],
		})

		intersects := func(rect Rect, p0, p1 aocspace.Vec) bool {
			minX := min(p0.X, p1.X) + 1
			maxX := max(p0.X, p1.X) - 1
			minY := min(p0.Y, p1.Y) + 1
			maxY := max(p0.Y, p1.Y) - 1

			fillMinX := min(rect.P0.X, rect.P1.X)
			fillMaxX := max(rect.P0.X, rect.P1.X)
			fillMinY := min(rect.P0.Y, rect.P1.Y)
			fillMaxY := max(rect.P0.Y, rect.P1.Y)

			if fillMaxX < minX || fillMinX > maxX {
				return false
			}
			if fillMaxY < minY || fillMinY > maxY {
				return false
			}

			return true
		}

		largestArea := int64(0)
		for idx, p0 := range points {
			for _, p1 := range points[idx:] {
				area := mhd(p0.X, p1.X) * mhd(p0.Y, p1.Y)

				if area < largestArea {
					continue
				}

				isSafe := true
				for _, fill := range filled {
					if intersects(fill, p0, p1) {
						isSafe = false
						break
					}
				}

				if isSafe {
					largestArea = area
				}
			}
		}

		return largestArea
	})
}
