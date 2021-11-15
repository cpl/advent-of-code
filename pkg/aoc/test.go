package aoc

import (
	"testing"
)

func TestLiveData(t *testing.T, year, day int) []byte {
	input, err := MetaGetInput(year, day)
	if err != nil {
		t.Fatal(err)
	}
	return input
}
