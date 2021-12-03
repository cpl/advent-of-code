package y2021

import (
	"reflect"
	"strings"
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

func TestParseBinaryListAsStrings(t *testing.T) {
	t.Parallel()

	t.Run("example", func(t *testing.T) {
		input := "00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"
		want := strings.Split(input, "\n")
		binaryList := ParseBinaryListAsStrings([]byte(input))

		if !reflect.DeepEqual(binaryList, want) {
			t.Errorf("got %v, want %v", binaryList, want)
		}
	})
}

func TestCalculateMostCommonBits(t *testing.T) {
	t.Parallel()

	t.Run("example", func(t *testing.T) {
		input := "00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"
		binaryList := ParseBinaryListAsStrings([]byte(input))
		got := CalculateMostCommonBits(binaryList)
		want := "10110"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestFlipBits(t *testing.T) {
	t.Parallel()

	t.Run("simple", func(t *testing.T) {
		in := "111011000011"
		want := "000100111100"
		got := FlipBits(in)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestWalkBinaryList(t *testing.T) {
	t.Parallel()

	t.Run("example", func(t *testing.T) {
		input := "00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"
		got := WalkBinaryList(strings.Split(input, "\n"), true)
		want := "10111"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

		got = WalkBinaryList(strings.Split(input, "\n"), false)
		want = "01010"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
