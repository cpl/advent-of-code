package y2022

import (
	"math"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay12(t *testing.T) {
	type pos struct {
		x, y int
	}

	type heightmap struct {
		grid   [][]rune
		start  pos
		end    pos
		starts []pos
	}

	parse := func(line string) []rune {
		return []rune(line)
	}

	interpret := func(grid [][]rune) *heightmap {
		hm := &heightmap{grid: grid}

		for y, row := range grid {
			for x, c := range row {
				switch c {
				case 'S':
					hm.start.x, hm.start.y = x, y
				case 'E':
					hm.end.x, hm.end.y = x, y
				case 'a':
					hm.starts = append(hm.starts, pos{x, y})
				}
			}
		}

		return hm
	}

	canTravel := func(from, to rune) bool {
		if to == ' ' {
			return false // out of bounds
		}

		if from == 'S' {
			return to == 'a'
		}

		if to == 'E' { // E == z !!!!!!!!!
			return from == 'y' || from == 'z'
		}

		return to-from <= 1
	}

	height := func(hm *heightmap, p pos) rune {
		if p.x < 0 || p.y < 0 || p.y >= len(hm.grid) || p.x >= len(hm.grid[p.y]) {
			return ' '
		}

		return hm.grid[p.y][p.x]
	}

	neighbors := func(hm *heightmap, p pos) []pos {
		return []pos{
			{p.x + 1, p.y},
			{p.x, p.y + 1},
			{p.x - 1, p.y},
			{p.x, p.y - 1},
		}
	}

	type posWithDist struct {
		pos
		dist int
	}

	findPathDist := func(start pos, hm *heightmap) int {
		var current posWithDist
		visited := make(map[pos]bool)

		pq2 := make([]posWithDist, 0, 10000)
		pq2 = append(pq2, posWithDist{start, 0})

		for len(pq2) > 0 {
			current, pq2 = pq2[0], pq2[1:]

			if current.pos == hm.end {
				return current.dist
			}

			if visited[current.pos] {
				continue
			}
			visited[current.pos] = true

			for _, next := range neighbors(hm, current.pos) {
				if canTravel(height(hm, current.pos), height(hm, next)) {
					pq2 = append(pq2, posWithDist{next, current.dist + 1})
				}
			}
		}

		return math.MaxInt
	}

	part1 := func(hm *heightmap) int {
		return findPathDist(hm.start, hm)
	}

	part2 := func(hm *heightmap) int {
		shortest := math.MaxInt

		for _, start := range hm.starts {
			if dist := findPathDist(start, hm); dist < shortest {
				shortest = dist
			}
		}

		return shortest
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi"
		t.Log(part1(interpret(aoc.ParseLines(aoc.InputScanner(input), parse))))
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(interpret(aoc.ParseLines(aoc.PuzzleScanner(2022, 12), parse))))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi"
		t.Log(part2(interpret(aoc.ParseLines(aoc.InputScanner(input), parse))))
	})

	t.Run("Part 2", func(t *testing.T) {
		t.Log(part2(interpret(aoc.ParseLines(aoc.PuzzleScanner(2022, 12), parse))))
	})
}
