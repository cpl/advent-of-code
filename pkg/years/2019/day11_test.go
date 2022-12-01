package y2019

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	"github.com/cpl/advent-of-code/pkg/years/2019/intcode"
)

func TestSolveDay11(t *testing.T) {
	coord := func(x, y int) string {
		return strconv.Itoa(x) + "," + strconv.Itoa(y)
	}

	paintColor := func(i int) string {
		if i == 0 {
			return "black"
		}
		return "white"
	}

	t.Run("Part 1", func(t *testing.T) {
		program := intcode.Parse(aoc.PuzzleString(2019, 11))
		vm := &intcode.VM{}
		vm.Bootstrap(program)

		hull := make(map[string]string)
		x, y := 0, 0
		dir := 0

		vm.IOWrite(0)

		halted := vm.Halted()
		for exit := false; !exit; {
			select {
			case <-halted:
				exit = true
			default:
				vm.RunUntilInput()
				paint := vm.IOReadBlocking()
				turn := vm.IOReadBlocking()

				hull[coord(x, y)] = paintColor(paint)

				if turn == 0 {
					dir = (dir + 3) % 4
				} else {
					dir = (dir + 1) % 4
				}

				switch dir {
				case 0:
					y++
				case 1:
					x++
				case 2:
					y--
				case 3:
					x--
				}

				color, ok := hull[coord(x, y)]
				if !ok {
					vm.IOWrite(0)
				} else {
					switch color {
					case "white":
						vm.IOWrite(1)
					case "black":
						vm.IOWrite(0)
					}
				}
			}
		}

		t.Log(len(hull))
	})

	t.Run("Part 2", func(t *testing.T) {
		program := intcode.Parse(aoc.PuzzleString(2019, 11))
		vm := &intcode.VM{}
		vm.Bootstrap(program)

		hull := make(map[string]string)
		x, y := 0, 0
		dir := 0

		vm.IOWrite(1)

		halted := vm.Halted()
		for exit := false; !exit; {
			select {
			case <-halted:
				exit = true
			default:
				vm.RunUntilInput()
				paint := vm.IOReadBlocking()
				turn := vm.IOReadBlocking()

				hull[coord(x, y)] = paintColor(paint)

				if turn == 0 {
					dir = (dir + 3) % 4
				} else {
					dir = (dir + 1) % 4
				}

				switch dir {
				case 0:
					y++
				case 1:
					x++
				case 2:
					y--
				case 3:
					x--
				}

				color, ok := hull[coord(x, y)]
				if !ok {
					vm.IOWrite(0)
				} else {
					switch color {
					case "white":
						vm.IOWrite(1)
					case "black":
						vm.IOWrite(0)
					}
				}
			}
		}

		t.Log(len(hull))

		for y = 2; y > -8; y-- {
			for x = -5; x < 50; x++ {
				color, ok := hull[coord(x, y)]
				if !ok {
					fmt.Print(".")
				} else {
					switch color {
					case "white":
						fmt.Print("█")
					case "black":
						fmt.Print(",")
					}
				}
			}

			fmt.Println()
		}
	})
}
