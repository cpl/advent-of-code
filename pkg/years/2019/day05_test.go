package y2019

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	"github.com/cpl/advent-of-code/pkg/years/2019/intcode"
)

func TestSolveDay05(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		program := intcode.Parse(aoc.PuzzleString(2019, 5))

		vm := &intcode.VM{}
		vm.Bootstrap(program)

		vm.IOWrite(1)
		vm.Run()

		t.Log(vm.IORead())
	})

	t.Run("Part 2", func(t *testing.T) {
		program := intcode.Parse(aoc.PuzzleString(2019, 5))

		vm := &intcode.VM{}
		vm.Bootstrap(program)

		vm.IOWrite(5)
		vm.Run()

		t.Log(vm.IORead())
	})
}
