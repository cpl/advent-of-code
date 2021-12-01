package y2021

import (
	"reflect"
	"testing"
)

func TestParsers(t *testing.T) {
	t.Parallel()

	t.Run("numbers", func(t *testing.T) {
		wantNumbers := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
		input := "199\n200\n208\n210\n200\n207\n240\n269\n260\n263"
		numbers := ParseNumbers([]byte(input))

		if !reflect.DeepEqual(numbers, wantNumbers) {
			t.Errorf("got %v, want %v", numbers, wantNumbers)
		}
	})
}

func TestNumberUpOrDown(t *testing.T) {
	t.Parallel()

	t.Run("example", func(t *testing.T) {
		numbers := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
		wantOut := []int{0, 1, 8, 2, -10, 7, 33, 29, -9, 3}
		up, down, out := NumberUpOrDown(numbers)

		if up != 7 {
			t.Errorf("got %v, want %v", up, 7)
		}
		if down != 2 {
			t.Errorf("got %v, want %v", down, 2)
		}

		if !reflect.DeepEqual(out, wantOut) {
			t.Errorf("got %v, want %v", out, wantOut)
		}

		t.Log(out)
	})
	t.Run("example_3ms", func(t *testing.T) {
		numbers := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
		threeSums := NumberThreeMeasureSums(numbers)
		up, _, _ := NumberUpOrDown(threeSums)

		if up != 5 {
			t.Errorf("got %v, want %v", up, 5)
		}
	})
}

func TestNumberThreeMeasureSums(t *testing.T) {
	t.Parallel()

	t.Run("example", func(t *testing.T) {
		numbers := []int{199, 200, 208, 210, 200, 207}
		wantOut := []int{607, 618, 618, 617}
		out := NumberThreeMeasureSums(numbers)

		if !reflect.DeepEqual(out, wantOut) {
			t.Errorf("got %v, want %v", out, wantOut)
		}
	})
}
