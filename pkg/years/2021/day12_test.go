package y2021

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay12(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 12)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("part_1", func(t *testing.T) {
		graph := ParseCaveGraph(string(input))
		graph.WalkBFS()

		t.Logf("paths=%d", len(graph.paths))
	})
	t.Run("part_2", func(t *testing.T) {
		graph := ParseCaveGraph(string(input))
		graph.WalkBFS2()

		t.Logf("paths=%d", len(graph.paths))
	})
}
