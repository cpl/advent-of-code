package y2021

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay05(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 5)
	if err != nil {
		t.Fatal(err)
	}

	vectors := Parse2Vec2(input)

	t.Run("part_1", func(t *testing.T) {
		mapper := vec2mapper{data: make(map[string]int)}
		for _, vecs := range vectors {
			mapper.Map(vecs)
		}
		t.Logf("solution=%d", mapper.Overlapping())
	})
	t.Run("part_2", func(t *testing.T) {
		mapper := vec2mapper{data: make(map[string]int)}
		for _, vecs := range vectors {
			mapper.Map2(vecs)
		}
		t.Logf("solution=%d", mapper.Overlapping())
	})
}
