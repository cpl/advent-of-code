package y2019

import (
	"fmt"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	"github.com/cpl/advent-of-code/pkg/years/2019/intcode"
)

func TestSolveDay13(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		program := intcode.Parse(aoc.PuzzleString(2019, 13))
		vm := &intcode.VM{}
		vm.Bootstrap(program)

		vm.Run()
		output := vm.IORead()

		screen := make(map[string]int)
		for idx := 0; idx < len(output); idx += 3 {
			x := output[idx]
			y := output[idx+1]
			tile := output[idx+2]

			screen[fmt.Sprintf("%d,%d", x, y)] = tile
		}

		var count int
		for _, tile := range screen {
			if tile == 2 {
				count++
			}
		}

		t.Log(count)
	})

	t.Run("Part 2", func(t *testing.T) {
		program := intcode.Parse(aoc.PuzzleString(2019, 13))
		vm := &intcode.VM{}
		vm.Bootstrap(program)

		vm.Memset(0, 2)
		vm.Run()
		output := vm.IORead()

		screen := make(map[string]int)
		for idx := 0; idx < len(output); idx += 3 {
			x := output[idx]
			y := output[idx+1]
			tile := output[idx+2]

			screen[fmt.Sprintf("%d,%d", x, y)] = tile
		}

		var count int
		for _, tile := range screen {
			if tile == 2 {
				count++
			}
		}

		t.Log(count)
	})
}
