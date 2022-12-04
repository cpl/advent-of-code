package y2022

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay04(t *testing.T) {
	type pair struct {
		start int
		end   int
	}

	parsePair := func(s string) pair {
		var p pair
		_, _ = fmt.Sscanf(s, "%d-%d", &p.start, &p.end)
		return p
	}

	parse := func(line string) [2]pair {

		split := strings.Index(line, ",")

		pair1 := parsePair(line[:split])
		pair2 := parsePair(line[split+1:])

		return [2]pair{pair1, pair2}
	}

	encapsulates := func(p1, p2 pair) bool {
		return p1.start <= p2.start && p1.end >= p2.end || p2.start <= p1.start && p2.end >= p1.end
	}

	overlaps := func(p1, p2 pair) bool {
		return p1.start <= p2.start && p1.end >= p2.start || p2.start <= p1.start && p2.end >= p1.start
	}

	solve := func(groups [][2]pair, fn func(p1, p2 pair) bool) int {
		count := 0

		for _, group := range groups {
			if fn(group[0], group[1]) {
				count++
			}
		}

		return count
	}

	t.Run("Examples 1", func(t *testing.T) {
		input := "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"
		t.Log(solve(aoc.ParseLines(aoc.InputScanner(input), parse), encapsulates))
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(solve(aoc.ParseLines(aoc.PuzzleScanner(2022, 4), parse), encapsulates))
	})

	t.Run("Examples 2", func(t *testing.T) {
		input := "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"
		t.Log(solve(aoc.ParseLines(aoc.InputScanner(input), parse), overlaps))
	})

	t.Run("Part 2", func(t *testing.T) {
		t.Log(solve(aoc.ParseLines(aoc.PuzzleScanner(2022, 4), parse), overlaps))
	})
}
