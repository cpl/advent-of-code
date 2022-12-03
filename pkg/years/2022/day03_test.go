package y2022

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay03(t *testing.T) {
	parse := func(line string) [2][]rune {
		l := len(line)
		return [2][]rune{
			[]rune(line[:l/2]),
			[]rune(line[l/2:]),
		}
	}

	priority := func(c rune) int {
		if c >= 'A' && c <= 'Z' {
			return 26 + int(c-'A'+1)
		} else {
			return int(c - 'a' + 1)
		}
	}

	part1 := func(backpacks [][2][]rune) int {
		score := 0

		for _, backpack := range backpacks {
			dupes := make(map[rune]struct{})

			comp1 := backpack[0]
			comp2 := backpack[1]

			for _, c := range comp1 {
				dupes[c] = struct{}{}
			}
			for _, c := range comp2 {
				if _, ok := dupes[c]; ok {
					score += priority(c)
					delete(dupes, c)
				}
			}
		}

		return score
	}

	part2 := func(backpacks [][2][]rune) int {
		score := 0

		findBadge := func(elfs ...[2][]rune) rune {
			dupes := make(map[rune]int)

			for _, elf := range elfs {
				has := make(map[rune]struct{})

				for _, comp := range elf {
					for _, c := range comp {
						has[c] = struct{}{}
					}
				}

				for c := range has {
					dupes[c]++
				}
			}

			for c, n := range dupes {
				if n == len(elfs) {
					return c
				}
			}

			panic("no badge found")
		}

		for idx := 0; idx < len(backpacks); idx += 3 {
			elf1 := backpacks[idx]
			elf2 := backpacks[idx+1]
			elf3 := backpacks[idx+2]

			score += priority(findBadge(elf1, elf2, elf3))
		}

		return score
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"
		t.Log(part1(aoc.ParseLines(aoc.InputScanner(input), parse)))
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(aoc.ParseLines(aoc.PuzzleScanner(2022, 3), parse)))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"
		t.Log(part2(aoc.ParseLines(aoc.InputScanner(input), parse)))
	})

	t.Run("Part 2", func(t *testing.T) {
		t.Log(part2(aoc.ParseLines(aoc.PuzzleScanner(2022, 3), parse)))
	})
}
