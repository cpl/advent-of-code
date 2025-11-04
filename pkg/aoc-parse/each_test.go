package aoc_parse_test

import (
	"bufio"
	"reflect"
	"strings"
	"testing"

	aoc_parse "github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func TestEachColumn(t *testing.T) {
	t.Parallel()

	r := bufio.NewScanner(strings.NewReader("\n\n\n\n\n\n\n\n\n\n\n\n"))

	columns := aoc_parse.EachColumn[int](r, 2, func(line string) []int {
		return []int{1, 2}
	})

	want := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	}

	if !reflect.DeepEqual(columns, want) {
		t.Errorf("EachColumn: got %v, want %v", columns, want)
	}
}
