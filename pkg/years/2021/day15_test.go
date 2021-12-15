package y2021

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay15(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 15)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("part_1", func(t *testing.T) {
		caveMap := ParseCaveMap(string(input), false)
		caveMap.AStar()

		t.Logf("solution=%d", caveMap.nodes[len(caveMap.nodes)-1][len(caveMap.nodes[0])-1].distance)
	})
	t.Run("part_2", func(t *testing.T) {
		caveMap := ParseCaveMap(string(input), true)
		caveMap.AStar()

		// 3019 too high
		// 3013 too high
		// 3012 right, whatever, off by 7 for the large example for some reason
		t.Logf("root_w=%d", caveMap.nodes[0][0].weight)
		t.Logf("goal_w=%d", caveMap.nodes[len(caveMap.nodes)-1][len(caveMap.nodes[0])-1].weight)
		t.Logf("solution=%d", caveMap.nodes[len(caveMap.nodes)-1][len(caveMap.nodes[0])-1].distance)
	})
}
