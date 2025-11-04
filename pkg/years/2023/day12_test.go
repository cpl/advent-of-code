package y2023

import (
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	"github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func TestSolveDay12(t *testing.T) {
	t.Parallel()

	type datum struct {
		s       string
		damaged []int
	}

	parse := func(line string) datum {
		idx := strings.LastIndexByte(line, ' ')
		s := line[:idx]

		nums := strings.Split(line[idx+1:], ",")
		damaged := make([]int, 0, len(nums))
		for _, num := range nums {
			count, _ := strconv.Atoi(num)
			damaged = append(damaged, count)
		}

		return datum{
			s:       s,
			damaged: damaged,
		}
	}

	isValid := func(s string, damaged []int) bool {
		damagedStr := strings.FieldsFunc(s, func(r rune) bool {
			if r == '.' {
				return true
			}
			return false
		})

		if len(damagedStr) != len(damaged) {
			return false
		}

		for idx, str := range damagedStr {
			if len(str) != damaged[idx] {
				return false
			}
		}

		return true
	}

	permutations := func(s string) []string {
		questionCount := strings.Count(s, "?")
		if questionCount == 0 {
			return []string{s}
		}

		permutations := make([]string, 0, 1<<questionCount)

		for idx := 0; idx < 1<<questionCount; idx++ {
			perm := s
			for i := 0; i < questionCount; i++ {
				if idx&(1<<i) != 0 {
					perm = strings.Replace(perm, "?", "#", 1)
				} else {
					perm = strings.Replace(perm, "?", ".", 1)
				}
			}
			permutations = append(permutations, perm)
		}

		return permutations
	}

	permutationsValid := func(s string, damaged []int) []string {
		perms := permutations(s)
		valid := make([]string, 0, len(perms))

		for _, perm := range perms {
			if isValid(perm, damaged) {
				valid = append(valid, perm)
			}
		}

		return valid
	}

	part1 := func(data []datum) int {
		total := 0
		for _, dat := range data {
			total += len(permutationsValid(dat.s, dat.damaged))
		}
		return total
	}

	unfold := func(d datum) datum {
		d.s = strings.Repeat(d.s+"?", 5)
		d.s = d.s[:len(d.s)-1]

		damaged := make([]int, len(d.damaged)*5)
		for idx := 0; idx < len(d.damaged)*5; idx++ {
			damaged[idx] = d.damaged[idx%len(d.damaged)]
		}
		d.damaged = damaged

		return d
	}

	t.Run("test 1a", func(t *testing.T) {
		data := aoc_parse.EachLine(aoc.InputScanner("#.#.### 1,1,3\n.#...#....###. 1,1,3\n.#.###.#.###### 1,3,1,6\n####.#...#... 4,1,1\n#....######..#####. 1,6,5\n.###.##....# 3,2,1"), parse)
		for idx, dat := range data {
			if !isValid(dat.s, dat.damaged) {
				t.Errorf("expected %d: %s to be valid", idx, dat.s)
			}
		}
	})

	t.Run("test 1b", func(t *testing.T) {
		for _, perm := range permutations("???.###") {
			if isValid(perm, []int{1, 1, 3}) {
				t.Log("valid perm:", perm)
			}
		}
	})

	t.Run("test 1c", func(t *testing.T) {
		if isValid(".##.###", []int{1, 1, 3}) {
			t.Error("expected .##.### to be invalid")
		}
	})

	t.Run("test 1d", func(t *testing.T) {
		if l := len(permutationsValid(".??..??...?##.", []int{1, 1, 3})); l != 4 {
			t.Errorf("expected 4 valid permutations, got %d", l)
		}

		v := permutationsValid("?###????????", []int{3, 2, 1})
		if l := len(v); l != 10 {
			t.Errorf("expected 10 valid permutations, got %d", l)
			for _, perm := range v {
				t.Log(perm)
			}
		}
	})

	t.Run("test 1e", func(t *testing.T) {
		if isValid(".###.##.####", []int{3, 2, 1}) {
			t.Error("expected .###.##.#### to be invalid")
		}

		if isValid(".###.#.##.#.", []int{3, 2, 1}) {
			t.Error("expected .###.#.##.#. to be invalid")
		}
	})

	t.Run("example 1", func(t *testing.T) {
		t.Log(part1(aoc_parse.EachLine(aoc.InputScanner("???.### 1,1,3\n.??..??...?##. 1,1,3\n?#?#?#?#?#?#?#? 1,3,1,6\n????.#...#... 4,1,1\n????.######..#####. 1,6,5\n?###???????? 3,2,1"), parse)))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(part1(aoc_parse.EachLine(aoc.PuzzleScanner(2023, 12), parse)))
	})

	t.Run("test 2a", func(t *testing.T) {
		d := unfold(datum{
			s:       ".#",
			damaged: []int{1},
		})

		if d.s != ".#?.#?.#?.#?.#" {
			t.Errorf("expected .#?.#?.#?.#?.#, got %s", d.s)
		}

		if !slices.Equal(d.damaged, []int{1, 1, 1, 1, 1}) {
			t.Errorf("expected [1,1,1,1,1], got %v", d.damaged)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		// not doing DP for fun right now
	})
}
