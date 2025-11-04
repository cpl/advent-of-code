package y2022

import (
	"fmt"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	"github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func TestSolveDay18(t *testing.T) {
	type cube struct {
		x, y, z int
	}

	parse := func(line string) cube {
		var c cube
		_, _ = fmt.Sscanf(line, "%d,%d,%d", &c.x, &c.y, &c.z)
		return c
	}

	sides := func(c cube) []cube {
		return []cube{
			{c.x - 1, c.y, c.z},
			{c.x + 1, c.y, c.z},
			{c.x, c.y - 1, c.z},
			{c.x, c.y + 1, c.z},
			{c.x, c.y, c.z - 1},
			{c.x, c.y, c.z + 1},
		}
	}

	part1 := func(cubes []cube) int {
		exposed := make(map[cube]int)
		cubeMap := make(map[cube]struct{})

		for _, c := range cubes {
			cubeMap[c] = struct{}{}

			for _, s := range sides(c) {
				exposed[s]++
			}
		}

		for e := range exposed {
			if _, ok := cubeMap[e]; ok {
				delete(exposed, e)
			}
		}

		sum := 0
		for _, v := range exposed {
			sum += v
		}

		return sum
	}

	part2 := func(cubes []cube) int {
		exposed := make(map[cube]int)
		cubeMap := make(map[cube]struct{})

		for _, c := range cubes {
			cubeMap[c] = struct{}{}

			for _, s := range sides(c) {
				exposed[s]++
			}
		}

		for e := range exposed {
			if _, ok := cubeMap[e]; ok {
				delete(exposed, e)
			}
		}

		exposedSurface := make(map[cube]rune)
		var isSurfaceExposed func(c cube, visited map[cube]struct{}) bool
		isSurfaceExposed = func(c cube, visited map[cube]struct{}) bool {
			if _, ok := visited[c]; ok {
				return false
			}
			visited[c] = struct{}{}

			if c.x < 0 || c.y < 0 || c.z < 0 || c.x > 20 || c.y > 20 || c.z > 20 {
				exposedSurface[c] = ' '
				return true
			}

			if _, ok := cubeMap[c]; ok {
				exposedSurface[c] = '#'
				return false
			}

			if v, ok := exposedSurface[c]; ok {
				if v == ' ' {
					return true
				}
			}

			for _, s := range sides(c) {
				if isSurfaceExposed(s, visited) {
					return true
				}
			}

			exposedSurface[c] = 'x'

			return false
		}

		for e := range exposed {
			if !isSurfaceExposed(e, make(map[cube]struct{})) {
				delete(exposed, e)
			}
		}

		sum := 0
		for _, v := range exposed {
			sum += v
		}

		return sum
	}

	t.Run("Example 1 - short", func(t *testing.T) {
		input := "1,1,1\n2,1,1"
		t.Log(part1(aoc_parse.EachLine(aoc.InputScanner(input), parse)))
	})

	t.Run("Example 1", func(t *testing.T) {
		input := "2,2,2\n1,2,2\n3,2,2\n2,1,2\n2,3,2\n2,2,1\n2,2,3\n2,2,4\n2,2,6\n1,2,5\n3,2,5\n2,1,5\n2,3,5"
		t.Log(part1(aoc_parse.EachLine(aoc.InputScanner(input), parse)))
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(aoc_parse.EachLine(aoc.PuzzleScanner(2022, 18), parse)))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "2,2,2\n1,2,2\n3,2,2\n2,1,2\n2,3,2\n2,2,1\n2,2,3\n2,2,4\n2,2,6\n1,2,5\n3,2,5\n2,1,5\n2,3,5"
		t.Log(part2(aoc_parse.EachLine(aoc.InputScanner(input), parse)))
	})

	t.Run("Part 2", func(t *testing.T) {
		t.Log(part2(aoc_parse.EachLine(aoc.PuzzleScanner(2022, 18), parse)))
	})
}
