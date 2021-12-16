package y2021

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay16(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 16)
	if err != nil {
		t.Fatal(err)
	}

	packets, err := ParseBits(string(input))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("part_1", func(t *testing.T) {
		for _, packet := range packets {
			t.Logf("sum=%d", packet.VersionSum())
		}
	})
	t.Run("part_2", func(t *testing.T) {
		for _, packet := range packets {
			t.Logf("value=%d", packet.Value)
		}
	})
}
