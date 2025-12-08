package y2025_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	aocspace "github.com/cpl/advent-of-code/pkg/aoc-space"
)

type Distance struct {
	J0, J1 *Junction
	Value  int64
}

type Junction struct {
	Circuit     int
	Position    aocspace.Vec
	Connections map[*Junction]struct{}
}

func TestSolve2025Day08(t *testing.T) {
	// aoc.SolveExamplesOnly(t)

	aoc.SolveExample(t, "example", 1, 40, "162,817,812\n57,618,57\n906,360,560\n592,479,940\n352,342,300\n466,668,158\n542,29,236\n431,825,988\n739,650,466\n52,470,668\n216,146,977\n819,987,18\n117,168,530\n805,96,715\n346,949,466\n970,615,88\n941,993,340\n862,61,35\n984,92,344\n425,690,689")

	aoc.Solve(t, "part 1", aocspace.Vec3Parser(','), func(positions []aocspace.Vec) int {
		junctions := make([]Junction, len(positions))

		for idx, pos := range positions {
			junctions[idx] = Junction{
				Circuit:     0,
				Position:    pos,
				Connections: make(map[*Junction]struct{}, 1024),
			}
		}

		distances := make([]Distance, 0, len(positions)*len(positions))

		for idx := 0; idx < len(junctions); idx++ {
			junctionI := &junctions[idx]

			for jdx := idx + 1; jdx < len(junctions); jdx++ {
				// for jdx := 0; jdx < len(junctions); jdx++ {
				if jdx == idx {
					continue
				}

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
			return int(a.Value - b.Value)
		})

		//for _, d := range distances {
		//	fmt.Println(d.J0.Position, d.J1.Position, d.Value)
		//}

		// After **making** the ten shortest connections
		connections := 0
		connectionsLimit := 10
		if !aoc.IsExample() {
			connectionsLimit = 1000
		}

		circuitId := 1
		circuitReMap := make(map[int]int, 1024)

		for _, distance := range distances {
			if connections >= connectionsLimit {
				break
			}

			j0, j1 := distance.J0, distance.J1

			// already connected
			if _, connected := j0.Connections[j1]; connected {
				// fmt.Println("already connected", j0.Position, "<->", j1.Position)
				continue
			}
			if _, connected := j1.Connections[j0]; connected {
				// fmt.Println("X already connected", j0.Position, "<->", j1.Position)
				continue
			}

			// already in same circuit
			if j1.Circuit != 0 && j0.Circuit == j1.Circuit {
				// fmt.Println("same circuit", j0.Position, j1.Position, "c", j0.Circuit)
				connections++
				continue
			}

			// different circuits?
			if j0.Circuit != 0 && j1.Circuit != 0 && j0.Circuit != j1.Circuit {
				// fmt.Println("different circuits means what?")
				// THEY MERGE !! ffs ...

				connections++
				circuitReMap[min(j0.Circuit, j1.Circuit)] = max(j0.Circuit, j1.Circuit)
				continue
			}

			// j0 part of circuit
			if j0.Circuit != 0 {
				connections++

				j1.Circuit = j0.Circuit

				j0.Connections[j1] = struct{}{}
				j1.Connections[j0] = struct{}{}

				//fmt.Println("joining",
				//	j1.Position,
				//	"to",
				//	j0.Position,
				//	"to circuit",
				//	j0.Circuit,
				//)

				continue
			}

			// j1 part of circuit
			if j1.Circuit != 0 {
				connections++

				j0.Circuit = j1.Circuit

				j0.Connections[j1] = struct{}{}
				j1.Connections[j0] = struct{}{}

				//fmt.Println("joining",
				//	j0.Position,
				//	"to",
				//	j1.Position,
				//	"to circuit",
				//	j1.Circuit,
				//)

				continue
			}

			// connect
			//fmt.Println("connecting",
			//	j0.Position,
			//	j1.Position,
			//	"to circuit",
			//	circuitId,
			//)

			j0.Connections[j1] = struct{}{}
			j1.Connections[j0] = struct{}{}
			j0.Circuit = circuitId
			j1.Circuit = circuitId
			circuitId++
			connections++
		}

		circuitsSizes := make([]int, circuitId+1)
		for _, junction := range junctions {
			circuitsSizes[junction.Circuit]++
		}

		// Multiplying together the sizes of the three largest circuits
		// -- god... is this a reading & attention exercise?

		fmt.Println(circuitsSizes)
		fmt.Println(circuitReMap)

		// re-map
		for id, count := range circuitsSizes {
			idReplacement, replace := circuitReMap[id]
			if !replace {
				continue
			}

			circuitsSizes[id] = 0
			circuitsSizes[idReplacement] += count
		}

		circuitsSizes[0] = 0 // ignore 1 junction circuits
		slices.Sort(circuitsSizes)

		fmt.Println(circuitsSizes)

		result := 1
		for _, size := range circuitsSizes[len(circuitsSizes)-3:] {
			if size == 0 {
				continue
			}

			result = result * size
		}

		return result // 7920 too low
	})
}
