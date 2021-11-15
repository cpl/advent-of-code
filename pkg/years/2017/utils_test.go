package year2017

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseDigitList(t *testing.T) {
	type testCase struct {
		input string
		want  []int
	}

	testData := []testCase{
		{"1122", []int{1, 1, 2, 2}},
		{"1111", []int{1, 1, 1, 1}},
		{"1234", []int{1, 2, 3, 4}},
		{"91212129", []int{9, 1, 2, 1, 2, 1, 2, 9}},
	}

	for _, test := range testData {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			output := ParseDigitList([]byte(test.input))
			if !reflect.DeepEqual(output, test.want) {
				t.Errorf("got %v, want %v", output, test.want)
			}
		})
	}
}

func TestSumMatchNext(t *testing.T) {
	type testCase struct {
		input []int
		want  int
	}

	testData := []testCase{
		{[]int{1, 1, 2, 2}, 3},
		{[]int{1, 1, 1, 1}, 4},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{9, 1, 2, 1, 2, 1, 2, 9}, 9},
	}

	for _, test := range testData {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			output := SumMatchNext(test.input)
			if output != test.want {
				t.Errorf("SumMatchNext(%#v) = %#v, want %#v", test.input, output, test.want)
			}
		})
	}
}

func TestSumMatchHalfSplitList(t *testing.T) {
	type testCase struct {
		input []int
		want  int
	}

	testData := []testCase{
		{[]int{1, 2, 1, 2}, 6},
		{[]int{1, 2, 2, 1}, 0},
		{[]int{1, 2, 3, 4, 2, 5}, 4},
		{[]int{1, 2, 3, 1, 2, 3}, 12},
		{[]int{1, 2, 1, 3, 1, 4, 1, 5}, 4},
	}

	for _, test := range testData {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			output := SumMatchHalfSplitList(test.input)
			if output != test.want {
				t.Errorf("SumMatchHalfSplitList(%#v) = %#v, want %#v", test.input, output, test.want)
			}
		})
	}
}

func TestMinMaxInList(t *testing.T) {
	type testCase struct {
		input   []int
		wantMin int
		wantMax int
	}

	testData := []testCase{
		{[]int{5, 1, 9, 5}, 1, 9},
		{[]int{7, 5, 3}, 3, 7},
		{[]int{2, 4, 6, 8}, 2, 8},
	}

	for _, test := range testData {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			outputMin, outputMax := MinMaxInList(test.input)
			if outputMin != test.wantMin || outputMax != test.wantMax {
				t.Errorf("MinMaxInList(%#v) = %#v, %#v, want %#v, %#v",
					test.input, outputMin, outputMax, test.wantMin, test.wantMax)
			}
		})
	}
}

func TestParseNumberListByLine(t *testing.T) {
	type testCase struct {
		input string
		want  [][]int
	}

	testData := []testCase{
		{
			"5 1 9 5\n7 5 3\n2 4 6 8\n",
			[][]int{
				{5, 1, 9, 5},
				{7, 5, 3},
				{2, 4, 6, 8},
			},
		},
		{
			"5 9 2 8\n9 4 7 3\n3 8 6 5",
			[][]int{
				{5, 9, 2, 8},
				{9, 4, 7, 3},
				{3, 8, 6, 5},
			},
		},
	}

	for _, test := range testData {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			output := ParseNumberListByLine(test.input)
			if !reflect.DeepEqual(output, test.want) {
				t.Errorf("got %v, want %v", output, test.want)
			}
		})
	}
}

func TestDivisibleByEachOtherInList(t *testing.T) {
	type testCase struct {
		input    []int
		wantNum0 int
		wantNum1 int
	}

	testData := []testCase{
		{[]int{5, 9, 2, 8}, 8, 2},
		{[]int{9, 4, 7, 3}, 9, 3},
		{[]int{3, 8, 6, 5}, 6, 3},
	}

	for _, test := range testData {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			outputNum0, outputNum1 := DivisibleByEachOtherInList(test.input)
			if outputNum0 != test.wantNum0 || outputNum1 != test.wantNum1 {
				t.Errorf("DivisibleByEachOtherInList(%#v) = %#v, %#v, want %#v, %#v",
					test.input, outputNum0, outputNum1, test.wantNum0, test.wantNum1)
			}
		})
	}
}
