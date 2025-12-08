package y2025_test

import (
	"bufio"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	space "github.com/cpl/advent-of-code/pkg/aoc-space"
	"go.sdls.io/pkg"
)

func TestSolve2025Day07(t *testing.T) {
	// aoc.SolveExamplesOnly(t)

	aoc.SolveExample(t, "example", 1, 21, ".......S.......\n...............\n.......^.......\n...............\n......^.^......\n...............\n.....^.^.^.....\n...............\n....^.^...^....\n...............\n...^.^...^.^...\n...............\n..^...^.....^..\n...............\n.^.^.^.^.^...^.\n...............")
	aoc.SolveExample(t, "example", 2, 40, ".......S.......\n...............\n.......^.......\n...............\n......^.^......\n...............\n.....^.^.^.....\n...............\n....^.^...^....\n...............\n...^.^...^.^...\n...............\n..^...^.....^..\n...............\n.^.^.^.^.^...^.\n...............")

	type Splitters = pkg.Set[space.Vec]

	type Input = struct {
		MaxY      int64
		Start     space.Vec
		Splitters Splitters
	}

	parser := func(scanner *bufio.Scanner) Input {
		input := Input{
			Splitters: pkg.NewSetSize[space.Vec](1_000_000),
		}

		y := 0
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Bytes()
			for x, ch := range line {
				switch ch {
				case '.':
					continue
				case '^':
					input.Splitters.Add(space.Vec2(x, y))
				case 'S':
					input.Start = space.Vec2(x, y)
				default:
					panic("invalid character in line")
				}
			}
			y++
		}

		input.MaxY = int64(y)

		return input
	}

	findSplitter := func(beam space.Vec, input *Input) (space.Vec, bool) {
		for y := beam.Y; y <= input.MaxY; y++ {
			check := space.Vec{
				X: beam.X,
				Y: y,
			}

			_, isSplitter := input.Splitters[check]
			if isSplitter {
				return check, true
			}
		}

		return space.Vec{}, false
	}

	aoc.Solve(t, "part 1", parser, func(input Input) int {
		beams := make([]space.Vec, 0, 1024)
		results := pkg.NewSet[space.Vec]()
		hitSplitters := pkg.NewSet[space.Vec]()

		beams = append(beams, input.Start)

		// splits := 0
		for len(beams) > 0 {
			// fmt.Println("have beams", len(beams))

			for _, beam := range beams {
				splitter, willSplit := findSplitter(beam, &input)
				if !willSplit {
					continue
				}

				if hitSplitters.Contains(splitter) {
					continue
				}
				hitSplitters.Add(splitter)

				// splits++
				results.Add(space.Vec{X: splitter.X + 1, Y: splitter.Y})
				results.Add(space.Vec{X: splitter.X - 1, Y: splitter.Y})
			}

			beams = results.ToSlice()
			results = results.Clear()
		}

		return hitSplitters.Len()
	})

	aoc.Solve(t, "part 2", parser, func(input Input) int {
		var quantumBeam func(beam space.Vec) int

		knownSolutions := make(map[space.Vec]int)

		quantumBeam = func(beam space.Vec) int {
			splitter, found := findSplitter(beam, &input)
			if !found {
				return 1
			}

			if solutions, known := knownSolutions[splitter]; known {
				return solutions
			}

			left := quantumBeam(space.Vec{X: splitter.X - 1, Y: splitter.Y})
			right := quantumBeam(space.Vec{X: splitter.X + 1, Y: splitter.Y})

			knownSolutions[splitter] = left + right

			return left + right
		}

		return quantumBeam(input.Start)
	})
}
