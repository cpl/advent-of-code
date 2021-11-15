package main

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay01Part01(t *testing.T) {
	testData := []testCase{
		{
			name:   "1111",
			input:  []byte("1111"),
			output: "4",
		},
		{
			name:   "1122",
			input:  []byte("1122"),
			output: "3",
		},
		{
			name:   "1234",
			input:  []byte("1234"),
			output: "0",
		},
		{
			name:   "91212129",
			input:  []byte("91212129"),
			output: "9",
		},
		{
			name:   "15818755262456378627116779495435596139617",
			input:  []byte("1581875526245637862711679495435596139617"),
			output: "11",
		},
		{
			name:   "0000",
			input:  []byte("0000"),
			output: "0",
		},
	}

	for _, test := range testData {
		t.Run("example_"+test.name, func(t *testing.T) {
			output, err := SolveDay01Part01(test.input)
			if err != nil {
				t.Fatal(err)
			}

			if output != test.output {
				t.Fatalf("expected '%s', got '%s'", test.output, output)
			}
		})
	}

	if t.Failed() {
		return
	}

	t.Run("live", func(t *testing.T) {
		input, err := aoc.MetaGetInput(2017, 1)
		if err != nil {
			t.Fatal(err)
		}

		output, err := SolveDay01Part01(input)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("output: %s", output)
	})
}

func TestSolveDay01Part02(t *testing.T) {
	testData := []testCase{
		{
			name:   "1212",
			input:  []byte("1212"),
			output: "6",
		},
		{
			name:   "123425",
			input:  []byte("123425"),
			output: "4",
		},
		{
			name:   "1221",
			input:  []byte("1221"),
			output: "0",
		},
		{
			name:   "12131415",
			input:  []byte("12131415"),
			output: "4",
		},
		{
			name:   "123123",
			input:  []byte("123123"),
			output: "12",
		},
	}

	for _, test := range testData {
		t.Run("example_"+test.name, func(t *testing.T) {
			output, err := SolveDay01Part02(test.input)
			if err != nil {
				t.Fatal(err)
			}

			if output != test.output {
				t.Fatalf("expected '%s', got '%s'", test.output, output)
			}
		})
	}

	if t.Failed() {
		return
	}

	t.Run("live", func(t *testing.T) {
		input, err := aoc.MetaGetInput(2017, 1)
		if err != nil {
			t.Fatal(err)
		}

		output, err := SolveDay01Part02(input)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("output: %s", output)
	})
}
