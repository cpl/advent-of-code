package y2021

import (
	"sort"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay10(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 10)
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	t.Run("part_1", func(t *testing.T) {
		score := 0
		for _, line := range lines {
			bs := &BracketStack{}
			idx, r, _ := bs.Parse(line)
			if idx != -1 {
				score += bracketScore[r]
			}
		}

		t.Logf("score=%d", score)
	})
	t.Run("part_2", func(t *testing.T) {
		scores := make([]int, 0, len(lines))

		for _, line := range lines {
			bs := &BracketStack{}
			idx, _, _ := bs.Parse(line)
			if idx != -1 {
				continue
			}

			if !bs.IsEmpty() {
				scores = append(scores, bs.RemainingScore())
			}
		}

		sort.Ints(scores)

		t.Logf("score=%d", scores[len(scores)/2])
	})
}
