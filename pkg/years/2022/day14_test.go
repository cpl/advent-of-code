package y2022

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	"github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func TestSolveDay14(t *testing.T) {
	type vec2 struct {
		x, y int
	}

	type simulation struct {
		grid [][]byte
		// origin     vec2
		dimensions vec2
		floorY     int

		minP vec2
		maxP vec2
	}

	parse := func(line string) []vec2 {
		points := strings.Split(line, " -> ")
		ret := make([]vec2, 0, len(points))
		for _, point := range points {
			var v vec2
			_, _ = fmt.Sscanf(point, "%d,%d", &v.x, &v.y)
			ret = append(ret, v)
		}
		return ret
	}

	simInBounds := func(sim *simulation, v vec2) bool {
		// return v.x >= sim.origin.x && v.x < sim.origin.x+sim.dimensions.x &&
		//	v.y >= sim.origin.y && v.y < sim.origin.y+sim.dimensions.y

		if sim.floorY == 0 {
			return v.x >= sim.minP.x && v.x <= sim.maxP.x && v.y >= sim.minP.y && v.y <= sim.maxP.y
		}

		return v.x >= 0 && v.x < 10000 && v.y >= 0 && v.y < 10000
	}

	simValue := func(sim *simulation, v vec2) byte {
		if sim.floorY != 0 && sim.floorY == v.y {
			return '='
		}

		if !simInBounds(sim, v) {
			return 'x'
		}

		// ret := sim.grid[v.y-sim.origin.y][v.x-sim.origin.x]
		ret := sim.grid[v.y][v.x]
		if ret == 0 {
			return '.'
		}

		return ret
	}

	simValueSet := func(sim *simulation, v vec2, value byte) {
		// sim.grid[v.y-sim.origin.y][v.x-sim.origin.x] = value
		sim.grid[v.y][v.x] = value
	}

	simPrint := func(sim *simulation) *simulation {
		for y := -1; y < sim.dimensions.y+1; y++ {
			for x := sim.minP.x - 10; x < sim.dimensions.x+10; x++ {
				fmt.Printf("%c", simValue(sim, vec2{x: x, y: y}))
			}
			fmt.Println()
		}

		return sim
	}

	simIsSolid := func(sim *simulation, v vec2) bool {
		value := simValue(sim, v)
		return value == '#' || value == 'o' || value == '='
	}

	interpret := func(walls [][]vec2) *simulation {
		minX, minY := math.MaxInt, 0
		maxX, maxY := 500, 0

		for _, wall := range walls {
			for _, point := range wall {
				if point.x < minX {
					minX = point.x
				}
				if point.y < minY {
					minY = point.y
				}
				if point.x > maxX {
					maxX = point.x
				}
				if point.y > maxY {
					maxY = point.y
				}
			}
		}

		// origin := vec2{x: minX, y: minY}
		dimensions := vec2{x: maxX, y: maxY}
		grid := make([][]byte, 10000)

		for y := range grid {
			grid[y] = make([]byte, 10000)
		}

		sim := &simulation{
			grid: grid,
			// origin:     origin,
			dimensions: dimensions,
			minP: vec2{
				x: minX,
				y: minY,
			},
			maxP: vec2{
				x: maxX,
				y: maxY,
			},
		}

		for _, wall := range walls {
			from := wall[0]
			for _, point := range wall[1:] {
				if from.x == point.x {
					minY := from.y
					maxY := point.y
					if point.y < minY {
						minY = point.y
						maxY = from.y
					}

					for y := minY; y <= maxY; y++ {
						simValueSet(sim, vec2{x: from.x, y: y}, '#')
					}
				} else {
					minX := from.x
					maxX := point.x
					if point.x < minX {
						minX = point.x
						maxX = from.x
					}

					for x := minX; x <= maxX; x++ {
						simValueSet(sim, vec2{x: x, y: from.y}, '#')
					}
				}

				from = point
			}
		}

		simValueSet(sim, vec2{x: 500, y: 0}, '+')

		return sim
	}

	simSand := func(sim *simulation) byte {
		pos := vec2{x: 500, y: 0}

		for {
			if !simInBounds(sim, pos) {
				return 'x'
			}

			below := vec2{x: pos.x, y: pos.y + 1}
			if !simIsSolid(sim, below) {
				pos = below
				continue
			}

			belowLeft := vec2{x: pos.x - 1, y: pos.y + 1}
			if !simIsSolid(sim, belowLeft) {
				pos = belowLeft
				continue
			}

			belowRight := vec2{x: pos.x + 1, y: pos.y + 1}
			if !simIsSolid(sim, belowRight) {
				pos = belowRight
				continue
			}

			if pos.x == 500 && pos.y == 0 {
				return '+'
			}

			simValueSet(sim, pos, 'o')
			return '#'
		}
	}

	part1 := func(sim *simulation) int {
		count := 0
		for simSand(sim) != 'x' {
			count++
		}
		return count
	}

	part2 := func(sim *simulation) int {
		sim.floorY = sim.dimensions.y + 2
		sim.dimensions.y = sim.floorY

		count := 0
		for simSand(sim) != '+' {
			count++
		}
		return count + 1
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9"
		t.Log(part1(simPrint(interpret(aoc_parse.EachLine(aoc.InputScanner(input), parse)))))
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(interpret(aoc_parse.EachLine(aoc.PuzzleScanner(2022, 14), parse))))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9"
		sim := interpret(aoc_parse.EachLine(aoc.InputScanner(input), parse))
		t.Log(part2(sim))
		simPrint(sim)
	})

	t.Run("Part 2", func(t *testing.T) {
		sim := interpret(aoc_parse.EachLine(aoc.PuzzleScanner(2022, 14), parse))
		t.Log(part2(sim))
		simPrint(sim)
	})
}
