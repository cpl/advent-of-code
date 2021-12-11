package y2021

import (
	"reflect"
	"testing"
)

func TestParseOctopusGrid(t *testing.T) {
	input := `123456
789000
123455`

	grid := ParseOctopusGrid(input)
	want := &OctopusGrid{
		Grid: [][]Octopus{
			{Octopus{Charge: 1}, Octopus{Charge: 2}, Octopus{Charge: 3}, Octopus{Charge: 4}, Octopus{Charge: 5}, Octopus{Charge: 6}},
			{Octopus{Charge: 7}, Octopus{Charge: 8}, Octopus{Charge: 9}, Octopus{Charge: 0}, Octopus{Charge: 0}, Octopus{Charge: 0}},
			{Octopus{Charge: 1}, Octopus{Charge: 2}, Octopus{Charge: 3}, Octopus{Charge: 4}, Octopus{Charge: 5}, Octopus{Charge: 5}},
		},
		TotalFlashes: 0,
	}

	if !reflect.DeepEqual(grid, want) {
		t.Errorf("got %v, want %v", grid, want)
	}
}

func TestOctopusGrid_ChargeAll(t *testing.T) {
	input := "000\n000\n000"
	grid := ParseOctopusGrid(input)
	grid.ChargeAll()

	for _, row := range grid.Grid {
		for _, octopus := range row {
			if octopus.Charge != 1 {
				t.Errorf("got %v, want %v", octopus.Charge, 1)
			}
		}
	}
}

func TestOctopusGrid_Flash(t *testing.T) {
	input := `11111
19991
19191
19991
11111`

	grid := ParseOctopusGrid(input)

	t.Run("step 1", func(t *testing.T) {
		grid.ChargeAll()
		want := "34543\n40004\n50005\n40004\n34543\n"
		if got := grid.String(); got != want {
			t.Errorf("got '%v', want '%v'", got, want)
		}
	})

	t.Run("step 2", func(t *testing.T) {
		grid.ChargeAll()
		want := "45654\n51115\n61116\n51115\n45654\n"
		if got := grid.String(); got != want {
			t.Errorf("got '%v', want '%v'", got, want)
		}
	})
}

func TestOctopusGrid_Example(t *testing.T) {
	input := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	t.Run("step 10", func(t *testing.T) {
		grid := ParseOctopusGrid(input)
		for iter := 0; iter < 10; iter++ {
			grid.ChargeAll()
		}

		want := 204
		if got := grid.TotalFlashes; got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("step 100", func(t *testing.T) {
		grid := ParseOctopusGrid(input)
		for iter := 0; iter < 100; iter++ {
			grid.ChargeAll()
		}

		want := 1656
		if got := grid.TotalFlashes; got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("all flash", func(t *testing.T) {
		grid := ParseOctopusGrid(input)
		expected := len(grid.Grid) * len(grid.Grid[0])
		iter := 0
		expectedIter := 194
		for expected != grid.ChargeAll() {
			iter++
			if iter > 200 {
				t.Errorf("went over max number of iterations")
			}
		}
		if iter != expectedIter {
			t.Errorf("got %v, want %v", iter, expectedIter)
		}
	})
}
