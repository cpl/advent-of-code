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

func TestNavigation(t *testing.T) {
	t.Parallel()

	t.Run("example_parse", func(t *testing.T) {
		input := "forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2"
		stepsWant := []navigationStep{
			{
				direction: "forward",
				units:     5,
			},
			{
				direction: "down",
				units:     5,
			},
			{
				direction: "forward",
				units:     8,
			},
			{
				direction: "up",
				units:     3,
			},
			{
				direction: "down",
				units:     8,
			},
			{
				direction: "forward",
				units:     2,
			},
		}
		steps := ParseNavigation([]byte(input))

		if !reflect.DeepEqual(steps, stepsWant) {
			t.Errorf("got %v, want %v", steps, stepsWant)
		}
	})
	t.Run("example_calculate", func(t *testing.T) {
		steps := []navigationStep{
			{
				direction: "forward",
				units:     5,
			},
			{
				direction: "down",
				units:     5,
			},
			{
				direction: "forward",
				units:     8,
			},
			{
				direction: "up",
				units:     3,
			},
			{
				direction: "down",
				units:     8,
			},
			{
				direction: "forward",
				units:     2,
			},
		}
		x, depth := CalculateNavigation(steps)
		if x != 15 {
			t.Errorf("got %v, want %v", x, 15)
		}
		if depth != 10 {
			t.Errorf("got %v, want %v", depth, 10)
		}
	})
	t.Run("example_calculate_p2", func(t *testing.T) {
		steps := []navigationStep{
			{
				direction: "forward",
				units:     5,
			},
			{
				direction: "down",
				units:     5,
			},
			{
				direction: "forward",
				units:     8,
			},
			{
				direction: "up",
				units:     3,
			},
			{
				direction: "down",
				units:     8,
			},
			{
				direction: "forward",
				units:     2,
			},
		}
		x, depth := CalculateNavigation2(steps)
		if x != 15 {
			t.Errorf("got %v, want %v", x, 15)
		}
		if depth != 60 {
			t.Errorf("got %v, want %v", depth, 60)
		}
	})
}
