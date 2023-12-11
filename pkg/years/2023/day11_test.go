package y2023

import (
	"math"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay11(t *testing.T) {
	t.Parallel()

	type pos struct {
		x, y int
	}

	parseUniverse := func(input string) ([][]byte, []int, []int) {
		lines := strings.Split(input, "\n")
		out := make([][]byte, 0, len(lines))

		galaxyColCount := make([]int, len(lines[0]))
		expansionX := make([]int, 0, 128)
		expansionY := make([]int, 0, 128)

		for y, line := range lines {
			out = append(out, []byte(line))
			if strings.IndexByte(line, '#') == -1 {
				expansionY = append(expansionY, y)
			} else {
				for idx, c := range line {
					if c == '#' {
						galaxyColCount[idx]++
					}
				}
			}
		}

		for x, gc := range galaxyColCount {
			if gc != 0 {
				continue
			}

			expansionX = append(expansionX, x)
		}

		return out, expansionX, expansionY
	}

	parseGalaxies := func(data [][]byte, expansionX, expansionY []int, expansionFactor int) []pos {
		galaxies := make([]pos, 0, 16)

		for y, row := range data {
			for x, cell := range row {
				if cell == '#' {
					realX, realY := x, y

					for _, ex := range expansionX {
						if x >= ex {
							realX += expansionFactor
						}
					}

					for _, ey := range expansionY {
						if y >= ey {
							realY += expansionFactor
						}
					}

					galaxies = append(galaxies, pos{realX, realY})
				}
			}
		}

		return galaxies
	}

	distance := func(p1, p2 pos) int {
		return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
	}

	part1 := func(galaxies []pos) int {
		sum := 0

		for idx, g1 := range galaxies {
			for _, g2 := range galaxies[idx+1:] {
				sum += distance(g1, g2)
			}
		}

		return sum
	}

	t.Run("example 1", func(t *testing.T) {
		universe, expansionX, expansionY := parseUniverse("...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#.....")
		t.Log(part1(parseGalaxies(universe, expansionX, expansionY, 1)))
	})

	t.Run("part 1", func(t *testing.T) {
		universe, expansionX, expansionY := parseUniverse(aoc.PuzzleString(2023, 11))
		t.Log(part1(parseGalaxies(universe, expansionX, expansionY, 1)))
	})

	t.Run("example 2", func(t *testing.T) {
		universe, expansionX, expansionY := parseUniverse("...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#.....")
		t.Log(part1(parseGalaxies(universe, expansionX, expansionY, 10-1)))
		t.Log(part1(parseGalaxies(universe, expansionX, expansionY, 100-1)))
	})

	t.Run("part 2", func(t *testing.T) {
		universe, expansionX, expansionY := parseUniverse(aoc.PuzzleString(2023, 11))
		t.Log(part1(parseGalaxies(universe, expansionX, expansionY, 1000000-1)))
	})
}
