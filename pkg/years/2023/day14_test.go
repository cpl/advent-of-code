package y2023

import (
	"encoding/hex"
	"hash/crc32"
	"slices"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay14(t *testing.T) {
	t.Parallel()

	parsePlatform := func(input string) [][]byte {
		lines := strings.Split(input, "\n")
		out := make([][]byte, len(lines))
		for idx := range lines {
			out[idx] = []byte(lines[idx])
		}
		return out
	}

	physics := func(platform [][]byte, x, y int, direction rune) {
		object := platform[y][x]
		if object == '.' || object == '#' {
			return
		}

		switch direction {
		case 'N':
			if y == 0 {
				return
			}
			platform[y][x] = '.'

			for dy := y - 1; dy >= 0; dy-- {
				if platform[dy][x] != '.' {
					platform[dy+1][x] = object
					return
				}
			}

			platform[0][x] = object
		case 'S':
			if y == len(platform)-1 {
				return
			}
			platform[y][x] = '.'

			for dy := y + 1; dy < len(platform); dy++ {
				if platform[dy][x] != '.' {
					platform[dy-1][x] = object
					return
				}
			}

			platform[len(platform)-1][x] = object
		case 'W':
			if x == 0 {
				return
			}
			platform[y][x] = '.'

			for dx := x - 1; dx >= 0; dx-- {
				if platform[y][dx] != '.' {
					platform[y][dx+1] = object
					return
				}
			}

			platform[y][0] = object
		case 'E':
			if x == len(platform[y])-1 {
				return
			}

			platform[y][x] = '.'

			for dx := x + 1; dx < len(platform[y]); dx++ {
				if platform[y][dx] != '.' {
					platform[y][dx-1] = object
					return
				}
			}

			platform[y][len(platform[y])-1] = object
		}
	}

	tilt := func(platform [][]byte, direction rune) [][]byte {
		switch direction {
		case 'N':
			for y := 0; y < len(platform); y++ {
				for x := 0; x < len(platform[y]); x++ {
					physics(platform, x, y, direction)
				}
			}
		case 'S':
			for y := len(platform) - 1; y >= 0; y-- {
				for x := 0; x < len(platform[y]); x++ {
					physics(platform, x, y, direction)
				}
			}
		case 'W':
			for y := 0; y < len(platform); y++ {
				for x := 0; x < len(platform[y]); x++ {
					physics(platform, x, y, direction)
				}
			}
		case 'E':
			for y := 0; y < len(platform); y++ {
				for x := len(platform[y]) - 1; x >= 0; x-- {
					physics(platform, x, y, direction)
				}
			}
		}

		return platform
	}

	tiltCycle := func(platform [][]byte) [][]byte {
		return tilt(tilt(tilt(tilt(platform, 'N'), 'W'), 'S'), 'E')
	}

	hash := func(platform [][]byte) string {
		h := crc32.New(crc32.IEEETable)
		for _, row := range platform {
			_, _ = h.Write(row)
		}

		return hex.EncodeToString(h.Sum(nil))
	}

	findCycle := func(seen []string) int {
		for idx := len(seen) - 1; idx >= 0; idx-- {
			for jdx := idx - 1; jdx >= 0; jdx-- {
				if seen[idx] == seen[jdx] {
					return idx - jdx
				}
			}
		}

		return -1
	}

	tiltCycleFind := func(platform [][]byte, cycles int) [][]byte {
		seen := make([]string, 0, 1000)

		for cycle := 0; cycle < 1000; cycle++ {
			h := hash(platform)
			seen = append(seen, h)
			tiltCycle(platform)
		}

		cycleSize := findCycle(seen)
		cycleRemaining := cycles - 1000

		for cycle := 0; cycle < cycleRemaining%cycleSize; cycle++ {
			tiltCycle(platform)
		}

		return platform
	}

	loadScore := func(platform [][]byte) int {
		total := 0

		l := len(platform)
		for y, row := range platform {
			for _, object := range row {
				if object == 'O' {
					total += l - y
				}
			}
		}

		return total
	}

	t.Run("example 1", func(t *testing.T) {
		original := parsePlatform("O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#....")
		tilted := parsePlatform("O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#....")
		tilt(tilted, 'N')

		expected := parsePlatform("OOOO.#.O..\nOO..#....#\nOO..O##..O\nO..#.OO...\n........#.\n..#....#.#\n..O..#.O.O\n..O.......\n#....###..\n#....#....")

		for idx := range expected {
			if !slices.Equal(expected[idx], tilted[idx]) {
				t.Errorf("expected %s | got %s | original %s", string(expected[idx]), string(tilted[idx]), string(original[idx]))
			}
		}

		t.Log(loadScore(tilted))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(loadScore(tilt(parsePlatform(aoc.PuzzleString(2023, 14)), 'N')))
	})

	t.Run("example 2", func(t *testing.T) {
		platform := parsePlatform("O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#....")
		tiltCycleFind(platform, 1000000000)
		t.Log(loadScore(platform))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(loadScore(tiltCycleFind(parsePlatform(aoc.PuzzleString(2023, 14)), 1000000000)))
	})
}
