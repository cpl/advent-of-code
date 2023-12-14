package y2023

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay13(t *testing.T) {
	t.Parallel()

	parse := func(input string) [][][]rune {
		patterns := strings.Split(input, "\n\n")
		out := make([][][]rune, len(patterns))

		for idx := range patterns {
			lines := strings.Split(patterns[idx], "\n")
			for lineIdx := range lines {
				out[idx] = append(out[idx], []rune(lines[lineIdx]))
			}
		}

		return out
	}

	rotate := func(pattern [][]rune) [][]rune {
		out := make([][]rune, len(pattern[0]))
		for idx := range out {
			out[idx] = make([]rune, len(pattern))
		}

		for y := range pattern {
			for x := range pattern[y] {
				out[x][len(pattern)-y-1] = pattern[y][x]
			}
		}

		return out
	}

	printPattern := func(pattern [][]rune) {
		for y := range pattern {
			for x := range pattern[y] {
				fmt.Printf("%c", pattern[y][x])
			}
			fmt.Println()
		}
	}

	equal := func(row1, row2 []rune) bool {
		return slices.Equal(row1, row2)
	}

	checkMirroring := func(side1, side2 [][]rune) bool {
		maxLen := min(len(side1), len(side2))

		for y := 0; y < maxLen; y++ {
			if !equal(side1[len(side1)-y-1], side2[y]) {
				return false
			}
		}

		return true
	}

	var findMirrorPoint func(pattern [][]rune, rotated bool, skip int, skipRot bool) (int, bool)
	findMirrorPoint = func(pattern [][]rune, rotated bool, skip int, skipRot bool) (int, bool) {
		for y := 0; y < len(pattern)-1; y++ {
			if equal(pattern[y], pattern[y+1]) {
				if checkMirroring(pattern[:y], pattern[y+2:]) {
					if y+1 == skip && rotated == skipRot {
						continue
					}

					return y + 1, rotated
				}
			}
		}

		if rotated {
			return -1, false
		}

		return findMirrorPoint(rotate(pattern), true, skip, skipRot)
	}

	_, _, _ = parse, rotate, printPattern

	part1 := func(patterns [][][]rune) int {
		total := 0
		for _, pattern := range patterns {
			point, rotated := findMirrorPoint(pattern, false, -1, false)
			if rotated {
				total += point
			} else {
				total += point * 100
			}
		}

		return total
	}

	flip := func(r rune) rune {
		if r == '#' {
			return '.'
		}

		return '#'
	}

	smudge := func(pattern [][]rune) [][][]rune {
		height := len(pattern)
		width := len(pattern[0])

		out := make([][][]rune, height*width)

		for idx := 0; idx < len(out); idx++ {
			out[idx] = make([][]rune, height)
			for y := range out[idx] {
				out[idx][y] = make([]rune, width)
			}

			for y := range pattern {
				for x := range pattern[y] {
					out[idx][y][x] = pattern[y][x]
				}
			}

			out[idx][idx/width][idx%width] = flip(out[idx][idx/width][idx%width])
		}

		return out
	}

	part2 := func(patterns [][][]rune) int {
		total := 0
		for _, pattern := range patterns {
			originalPoint, originalRotated := findMirrorPoint(pattern, false, -1, false)

			for _, smudged := range smudge(pattern) {
				point, rotated := findMirrorPoint(smudged, false, originalPoint, originalRotated)
				if point == -1 {
					continue
				}

				if rotated {
					total += point
				} else {
					total += point * 100
				}

				break
			}
		}

		return total
	}

	t.Run("example 1", func(t *testing.T) {
		patterns := parse("#.##..##.\n..#.##.#.\n##......#\n##......#\n..#.##.#.\n..##..##.\n#.#.##.#.\n\n#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#")
		if len(patterns) != 2 {
			t.Fatalf("expected 2 patterns, got %d", len(patterns))
		}

		t.Log(part1(patterns))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(part1(parse(aoc.PuzzleString(2023, 13))))
	})

	t.Run("example 2", func(t *testing.T) {
		patterns := parse("#.##..##.\n..#.##.#.\n##......#\n##......#\n..#.##.#.\n..##..##.\n#.#.##.#.\n\n#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#")
		if len(patterns) != 2 {
			t.Fatalf("expected 2 patterns, got %d", len(patterns))
		}

		t.Log(part2(patterns))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(part2(parse(aoc.PuzzleString(2023, 13))))
	})
}
