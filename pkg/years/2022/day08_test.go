package y2022

import (
	"fmt"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

// 🌲🌲🎄

func TestSolveDay08(t *testing.T) {
	parse := func(line string) []int {
		out := make([]int, len(line))
		for idx, c := range line {
			out[idx] = int(c - '0')
		}
		return out
	}

	part1 := func(grid [][]int) int {
		total := 0

		maxX := len(grid[0]) - 1
		maxY := len(grid) - 1
		for y, row := range grid {
			for x, _ := range row {
				if x == 0 || y == 0 || x == maxX || y == maxY {
					total++
				}
			}
		}

		seenMap := make(map[string]int)
		see := func(x, y, start int) int {
			tree := grid[y][x]
			if tree > start {
				loc := fmt.Sprintf("%d,%d", x, y)
				start = tree

				if _, ok := seenMap[loc]; !ok {
					seenMap[loc] = tree
					total++
				}
			}
			return start
		}

		for y := 1; y < maxY; y++ {
			// left most tree
			start := grid[y][0]

			// scan left to right
			for x := 1; x < maxX; x++ {
				start = see(x, y, start)
			}

			// right most tree
			start = grid[y][maxX]

			// scan right to left
			for x := maxX - 1; x > 0; x-- {
				start = see(x, y, start)
			}
		}

		for x := 1; x < maxX; x++ {
			// top most tree
			start := grid[0][x]

			// scan top to bottom
			for y := 1; y < maxY; y++ {
				start = see(x, y, start)
			}

			// bottom most tree
			start = grid[maxY][x]

			// scan bottom to top
			for y := maxY - 1; y > 0; y-- {
				start = see(x, y, start)
			}
		}

		return total
	}

	viewingDistanceDirection := func(me int, trees []int) int {
		if len(trees) == 0 {
			return 0
		}

		count := 0
		for _, tree := range trees {
			if tree < me {
				count++
			} else {
				return count + 1
			}
		}

		return count
	}

	reverse := func(in []int) []int {
		out := make([]int, len(in))
		for idx := range in {
			out[idx] = in[len(in)-idx-1]
		}
		return out
	}

	treesOnX := func(grid [][]int, x int) []int {
		out := make([]int, 0, len(grid))
		for _, row := range grid {
			out = append(out, row[x])
		}
		return out
	}

	viewingDistance := func(x, y int, grid [][]int) int {
		me := grid[y][x]

		up := viewingDistanceDirection(me, reverse(treesOnX(grid, x)[:y]))
		if up == 0 {
			return 0
		}

		down := viewingDistanceDirection(me, treesOnX(grid, x)[y+1:])
		if down == 0 {
			return 0
		}

		left := viewingDistanceDirection(me, reverse(grid[y][:x]))
		if left == 0 {
			return 0
		}

		right := viewingDistanceDirection(me, grid[y][x+1:])
		if right == 0 {
			return 0
		}

		return left * right * up * down
	}

	part2 := func(grid [][]int) int {
		max := 0

		for y, row := range grid {
			for x := range row {
				v := viewingDistance(x, y, grid)
				if v > max {
					max = v
				}
			}
		}

		return max
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "30373\n25512\n65332\n33549\n35390"
		t.Log(part1(aoc.ParseLines(aoc.InputScanner(input), parse)))
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(aoc.ParseLines(aoc.PuzzleScanner(2022, 8), parse)))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "30373\n25512\n65332\n33549\n35390"
		t.Log(part2(aoc.ParseLines(aoc.InputScanner(input), parse)))
	})

	t.Run("Part 2", func(t *testing.T) {
		t.Log(part2(aoc.ParseLines(aoc.PuzzleScanner(2022, 8), parse)))
	})
}
