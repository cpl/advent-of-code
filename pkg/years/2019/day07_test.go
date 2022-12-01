package y2019

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	"github.com/cpl/advent-of-code/pkg/years/2019/intcode"
)

func solveDay07Part1Phases() [][5]int {
	phases := make([][5]int, 0, 120)
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			if b == a {
				continue
			}
			for c := 0; c < 5; c++ {
				if c == a || c == b {
					continue
				}
				for d := 0; d < 5; d++ {
					if d == a || d == b || d == c {
						continue
					}
					for e := 0; e < 5; e++ {
						if e == a || e == b || e == c || e == d {
							continue
						}
						phases = append(phases, [5]int{a, b, c, d, e})
					}
				}
			}
		}
	}

	return phases
}

func solveDay07Part2Phases() [][5]int {
	phases := make([][5]int, 0, 120)
	for a := 5; a < 10; a++ {
		for b := 5; b < 10; b++ {
			if b == a {
				continue
			}
			for c := 5; c < 10; c++ {
				if c == a || c == b {
					continue
				}
				for d := 5; d < 10; d++ {
					if d == a || d == b || d == c {
						continue
					}
					for e := 5; e < 10; e++ {
						if e == a || e == b || e == c || e == d {
							continue
						}
						phases = append(phases, [5]int{a, b, c, d, e})
					}
				}
			}
		}
	}

	return phases
}

func TestSolveDay07(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		program := intcode.Parse(aoc.PuzzleString(2019, 7))

		vm := &intcode.VM{}
		vm.Bootstrap(program)

		phases := solveDay07Part1Phases()
		max := 0

		for _, phase := range phases {
			in := 0
			for _, p := range phase {
				vm.IOWrite(p, in)
				vm.Run()
				in = vm.IOReadBlocking()

				vm.Reset()
			}

			if in > max {
				max = in
			}
		}

		t.Log(max)
	})

	t.Run("Part 2", func(t *testing.T) {
		program := intcode.Parse(aoc.PuzzleString(2019, 7))

		phases := solveDay07Part2Phases()

		ampA := &intcode.VM{}
		ampA.Bootstrap(program)
		ampB := &intcode.VM{}
		ampB.Bootstrap(program)
		ampC := &intcode.VM{}
		ampC.Bootstrap(program)
		ampD := &intcode.VM{}
		ampD.Bootstrap(program)
		ampE := &intcode.VM{}
		ampE.Bootstrap(program)

		max := 0
		for _, phase := range phases {
			// init phases
			ampA.IOWrite(phase[0])
			ampB.IOWrite(phase[1])
			ampC.IOWrite(phase[2])
			ampD.IOWrite(phase[3])
			ampE.IOWrite(phase[4])

			ampA.IOWrite(0)
			exit := false
			for !exit {
				ampA.RunUntilInput()
				ampB.IOWrite(ampA.IOReadBlocking())
				ampB.RunUntilInput()
				ampC.IOWrite(ampB.IOReadBlocking())
				ampC.RunUntilInput()
				ampD.IOWrite(ampC.IOReadBlocking())
				ampD.RunUntilInput()
				ampE.IOWrite(ampD.IOReadBlocking())
				ampE.RunUntilInput()
				v := ampE.IOReadBlocking()

				select {
				case <-ampE.Halted():
					exit = true
					if v > max {
						max = v
					}
				default:
					ampA.IOWrite(v)
				}
			}

			ampA.Reset()
			ampB.Reset()
			ampC.Reset()
			ampD.Reset()
			ampE.Reset()
		}

		t.Log(max)
	})
}
