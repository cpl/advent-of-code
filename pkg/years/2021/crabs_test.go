package y2021

import (
	"reflect"
	"testing"
)

func TestParseCrabs(t *testing.T) {
	t.Parallel()

	crabs := ParseCrabs("16,1,2,0,4,2,7,1,2,14")
	want := []int{1, 2, 3, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1}

	if !reflect.DeepEqual(crabs.positions, want) {
		t.Errorf("got %v, want %v", crabs.positions, want)
	}
}

func TestCrabSubmarines_AlignConstant(t *testing.T) {
	t.Parallel()

	crabs := &CrabSubmarines{
		positions: []int{1, 2, 3, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1},
	}

	position, fuel := crabs.Align()
	if fuel != 37 {
		t.Errorf("got %d, want %d", fuel, 37)
	}
	if position != 2 {
		t.Errorf("got %d, want %d", position, 2)
	}
}

func TestCrabSubmarines_AlignAtConstant(t *testing.T) {
	t.Parallel()

	crabs := &CrabSubmarines{
		positions: []int{1, 2, 3, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1},
	}

	fuel := crabs.AlignAt(2)
	if fuel != 37 {
		t.Errorf("got %d, want %d", fuel, 37)
	}
}

func TestCrabSubmarines_AlignAtLinear(t *testing.T) {
	t.Parallel()

	crabs := &CrabSubmarines{
		positions: []int{1, 2, 3, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1},
	}

	fuel := crabs.AlignLinearAt(5)
	if fuel != 168 {
		t.Errorf("got %d, want %d", fuel, 168)
	}

	fuel = crabs.AlignLinearAt(2)
	if fuel != 206 {
		t.Errorf("got %d, want %d", fuel, 206)
	}
}

func TestCrabSubmarines_AlignLinear(t *testing.T) {
	t.Parallel()

	crabs := &CrabSubmarines{
		positions: []int{1, 2, 3, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1},
	}

	position, fuel := crabs.AlignLinear()
	if fuel != 168 {
		t.Errorf("got %d, want %d", fuel, 168)
	}
	if position != 5 {
		t.Errorf("got %d, want %d", position, 5)
	}
}
