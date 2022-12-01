package y2019

import (
	"github.com/cpl/advent-of-code/pkg/years/2019/intcode"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay02(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		program := intcode.Parse(aoc.PuzzleString(2019, 2))

		vm := &intcode.VM{}
		vm.Bootstrap(program)

		vm.Memset(1, 12)
		vm.Memset(2, 2)

		vm.Run()
		t.Log(vm.Memget(0))
	})

	t.Run("Part 2", func(t *testing.T) {
		program := intcode.Parse(aoc.PuzzleString(2019, 2))

		vm := &intcode.VM{}
		vm.Bootstrap(program)

		for noun := 0; noun < 100; noun++ {
			for verb := 0; verb < 100; verb++ {
				vm.Reset()

				vm.Memset(1, noun)
				vm.Memset(2, verb)
				vm.Run()

				if vm.Memget(0) == 19690720 {
					t.Log(100*noun + verb)
					return
				}
			}
		}

		t.Fail()
	})
}
