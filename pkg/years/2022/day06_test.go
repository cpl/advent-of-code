package y2022

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay06(t *testing.T) {
	uniq := func(s string) bool {
		var seen [26]bool

		for _, c := range s {
			if seen[c-'a'] {
				return false
			}
			seen[c-'a'] = true
		}

		return true
	}

	analyzeStream := func(stream string, window int) int {
		for pos := 0; pos < len(stream)-window+1; pos++ {
			if uniq(stream[pos : pos+window]) {
				return pos + window
			}
		}

		return -1
	}

	part1 := func(stream string) int {
		return analyzeStream(stream, 4)
	}

	part2 := func(stream string) int {
		return analyzeStream(stream, 14)
	}

	t.Run("Example 1", func(t *testing.T) {
		tests := []struct {
			input    string
			expected int
		}{
			{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", expected: 5},
			{input: "nppdvjthqldpwncqszvftbrmjlhg", expected: 6},
			{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", expected: 10},
			{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", expected: 11},
		}

		for _, tt := range tests {
			actual := part1(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}
		}
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(aoc.PuzzleString(2022, 6)))
	})

	t.Run("Part 2", func(t *testing.T) {
		t.Log(part2(aoc.PuzzleString(2022, 6)))
	})
}
