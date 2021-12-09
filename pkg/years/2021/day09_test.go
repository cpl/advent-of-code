package y2021

import (
	"sort"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay09(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 9)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("part_1", func(t *testing.T) {
		lows := HeightmapLows(ParseHeightmap(string(input)))
		risk := 0
		for _, low := range lows {
			risk += low + 1
		}

		t.Logf("solution=%d (lows=%v)", risk, lows)
	})
	t.Run("part_2", func(t *testing.T) {
		hm := ParseHeightmap2(string(input))
		hm.SeekBasins()

		sort.Slice(hm.basins, func(i, j int) bool {
			return hm.basins[i].count > hm.basins[j].count
		})

		total := hm.basins[0].count * hm.basins[1].count * hm.basins[2].count
		t.Logf("solution=%d", total)
	})
}
