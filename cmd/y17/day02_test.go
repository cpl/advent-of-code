package main

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay02Part01(t *testing.T) {
	testData := []testCase{
		{
			name:   "example",
			input:  []byte("5 1 9 5\n7 5 3\n2 4 6 8"),
			output: "18",
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			result, err := SolveDay02Part01(test.input)
			if err != nil {
				t.Fatal(err)
			}

			if result != test.output {
				t.Errorf("SolveDay02Part01(%s) = %s, want %s", test.name, result, test.output)
			}
		})
	}

	if t.Failed() {
		return
	}

	t.Run("live", func(t *testing.T) {
		solution, err := SolveDay02Part01(aoc.TestLiveData(t, 2017, 2))
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("solution: %s", solution)
	})
}
