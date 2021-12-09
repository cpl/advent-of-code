package y2021

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestParseHeightmap(t *testing.T) {
	t.Parallel()

	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	hm := ParseHeightmap(input)
	want := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	if !reflect.DeepEqual(hm, want) {
		t.Errorf("got %v, want %v", hm, want)
	}
}

func TestHeightmapLows(t *testing.T) {
	t.Parallel()

	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	lows := HeightmapLows(ParseHeightmap(input))
	want := []int{1, 0, 5, 5}

	if !reflect.DeepEqual(lows, want) {
		t.Errorf("got %v, want %v", lows, want)
	}
}

func TestHeightmap2(t *testing.T) {
	t.Parallel()

	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	hm := ParseHeightmap2(input)
	hm.Print()
	hm.SeekBasins()

	fmt.Println("--------------------------")
	hm.Print()

	fmt.Println("--------------------------")

	sort.Slice(hm.basins, func(i, j int) bool {
		return hm.basins[i].count > hm.basins[j].count
	})
	for _, basin := range hm.basins {
		t.Logf("%s=%d", string(rune('a'+basin.id)), basin.count)
	}

	total := hm.basins[0].count * hm.basins[1].count * hm.basins[2].count
	t.Logf("total=%d", total)

	if total != 1134 {
		t.Errorf("got %d, want %d", total, 1134)
	}
}
