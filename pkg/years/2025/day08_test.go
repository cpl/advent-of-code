package y2025_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	aocspace "github.com/cpl/advent-of-code/pkg/aoc-space"
)

type Junction struct {
	Circuit  *Circuit
	Position aocspace.Vec
}

type Distance struct {
	J0, J1 *Junction
	Value  int64
}

type Circuit struct {
	Junctions map[*Junction]struct{}
}

func (circuit *Circuit) Move(to *Circuit) {
	for junction := range circuit.Junctions {
		to.Junctions[junction] = struct{}{}
		junction.Circuit = to
		delete(circuit.Junctions, junction)
	}
}

func (circuit *Circuit) Add(junction *Junction) {
	junction.Circuit = circuit
	circuit.Junctions[junction] = struct{}{}
}

func (circuit *Circuit) Size() int {
	return len(circuit.Junctions)
}

func TestSolve2025Day08(t *testing.T) {
	// aoc.SolveExamplesOnly(t)

	aoc.SolveExample(t, "example", 1, 40, "162,817,812\n57,618,57\n906,360,560\n592,479,940\n352,342,300\n466,668,158\n542,29,236\n431,825,988\n739,650,466\n52,470,668\n216,146,977\n819,987,18\n117,168,530\n805,96,715\n346,949,466\n970,615,88\n941,993,340\n862,61,35\n984,92,344\n425,690,689")
	aoc.SolveExample(t, "example", 2, 25272, "162,817,812\n57,618,57\n906,360,560\n592,479,940\n352,342,300\n466,668,158\n542,29,236\n431,825,988\n739,650,466\n52,470,668\n216,146,977\n819,987,18\n117,168,530\n805,96,715\n346,949,466\n970,615,88\n941,993,340\n862,61,35\n984,92,344\n425,690,689")

	parseJunctions := func(positions []aocspace.Vec) ([]Junction, []Distance) {
		junctions := make([]Junction, len(positions))

		for idx, pos := range positions {
			junctions[idx] = Junction{
				Circuit:  nil,
				Position: pos,
			}
		}

		distances := make([]Distance, 0, len(positions)*len(positions))

		for idx := 0; idx < len(junctions); idx++ {
			junctionI := &junctions[idx]

			for jdx := idx + 1; jdx < len(junctions); jdx++ {
				junctionJ := &junctions[jdx]

				distance := junctionI.Position.DistanceDelta(junctionJ.Position)
				distances = append(distances, Distance{
					J0:    junctionI,
					J1:    junctionJ,
					Value: distance,
				})
			}
		}

		slices.SortFunc(distances, func(a, b Distance) int {
			return cmp.Compare(a.Value, b.Value)
		})

		return junctions, distances
	}

	aoc.Solve(t, "part 1", aocspace.Vec3Parser(','), func(positions []aocspace.Vec) int {
		junctions, distances := parseJunctions(positions)
		circuits := make([]Circuit, 0, len(junctions))

		limit := 10
		if !aoc.IsExample() {
			limit = 1000
		}

		for _, distance := range distances[:limit] {
			j0, j1 := distance.J0, distance.J1

			j0HasCircuit := j0.Circuit != nil
			j1HasCircuit := j1.Circuit != nil

			if j0HasCircuit && j1HasCircuit {
				// same circuit
				if j0.Circuit == j1.Circuit {
					continue
				}

				// move j0 to j1 circuit
				j0.Circuit.Move(j1.Circuit)
			} else if !j0HasCircuit && !j1HasCircuit {
				// create new circuit
				circuits = append(circuits, Circuit{
					Junctions: make(map[*Junction]struct{}, 64),
				})

				// add j0 and j1
				circuit := &circuits[len(circuits)-1]
				circuit.Add(j0)
				circuit.Add(j1)
			} else if j0HasCircuit {
				j0.Circuit.Add(j1)
			} else if j1HasCircuit {
				j1.Circuit.Add(j0)
			} else {
				panic("impossible")
			}
		}

		slices.SortFunc(circuits, func(a, b Circuit) int {
			return b.Size() - a.Size()
		})

		result := 1
		for _, circuit := range circuits[:3] {
			if circuit.Size() == 0 {
				continue
			}

			result *= circuit.Size()
		}

		return result
	})

	aoc.Solve(t, "part 2", aocspace.Vec3Parser(','), func(positions []aocspace.Vec) int {
		junctions, distances := parseJunctions(positions)

		circuits := make([]Circuit, 0, len(junctions))
		circuitLess := make(map[*Junction]struct{}, len(junctions))
		for idx := range junctions {
			circuitLess[&junctions[idx]] = struct{}{}
		}

		for _, distance := range distances {
			j0, j1 := distance.J0, distance.J1

			j0HasCircuit := j0.Circuit != nil
			j1HasCircuit := j1.Circuit != nil

			if j0HasCircuit && j1HasCircuit {
				// same circuit
				if j0.Circuit == j1.Circuit {
					continue
				}

				j0.Circuit.Move(j1.Circuit)
			} else if !j0HasCircuit && !j1HasCircuit {
				// create new circuit
				circuits = append(circuits, Circuit{
					Junctions: make(map[*Junction]struct{}, 64),
				})

				// add j0 and j1
				circuit := &circuits[len(circuits)-1]
				circuit.Add(j0)
				circuit.Add(j1)

				delete(circuitLess, j0)
				delete(circuitLess, j1)
			} else if j0HasCircuit {
				j0.Circuit.Add(j1)
				delete(circuitLess, j1)
			} else if j1HasCircuit {
				j1.Circuit.Add(j0)
				delete(circuitLess, j0)
			} else {
				panic("impossible")
			}

			if len(circuitLess) == 0 {
				ccount := 0
				for _, circuit := range circuits {
					if circuit.Size() > 0 {
						ccount++
					}

					if ccount > 1 {
						break
					}
				}

				if ccount == 1 {
					return int(distance.J0.Position.X * distance.J1.Position.X)
				}
			}
		}

		return 0
	})
}
