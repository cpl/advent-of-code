package y2022

import (
	"fmt"
	"math"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay15(t *testing.T) {
	type vec2 struct {
		x, y int
	}

	type sensor struct {
		pos          vec2
		beacon       vec2
		distToBeacon int
	}

	parse := func(line string) sensor {
		var s sensor

		_, _ = fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&s.pos.x, &s.pos.y, &s.beacon.x, &s.beacon.y)

		s.distToBeacon = int(math.Abs(float64(s.pos.x-s.beacon.x)) + math.Abs(float64(s.pos.y-s.beacon.y)))

		return s
	}

	part1 := func(sensors []sensor, row int) int {
		fill := make(map[int]struct{})
		sub := make(map[int]struct{})

		for _, s := range sensors {
			if s.beacon.y == row {
				sub[s.beacon.x] = struct{}{}
			}

			if s.pos.y-s.distToBeacon > row || row > s.pos.y+s.distToBeacon {
				continue
			}

			distToRow := int(math.Abs(float64(s.pos.y - row)))

			for x := s.pos.x - s.distToBeacon + distToRow; x <= s.pos.x+s.distToBeacon-distToRow; x++ {
				fill[x] = struct{}{}
			}
		}

		for x := range sub {
			delete(fill, x)
		}

		return len(fill)
	}

	tuning := func(p vec2) int {
		return p.x*4000000 + p.y
	}

	overlap := func(r1, r2 vec2) (vec2, bool) {
		if r1.x > r2.x {
			r2, r1 = r1, r2
		}

		if r1.y-r2.x < -1 {
			return vec2{}, false
		}

		if r1.y > r2.y {
			return r1, true
		}

		return vec2{x: r1.x, y: r2.y}, true
	}

	reduce := func(ranges []vec2) []vec2 {
		for len(ranges) > 2 {
			//fmt.Println("reduce", ranges)

			for idx := 0; idx < len(ranges)-1; idx++ {
				r1 := ranges[idx]

				for jdx := idx + 1; jdx < len(ranges); jdx++ {
					r2 := ranges[jdx]

					if r, ok := overlap(r1, r2); ok {
						//fmt.Println("overlap", r1, r2, "=>", r)
						ranges[idx] = r
						if jdx == len(ranges)-1 {
							ranges = ranges[:jdx]
						} else {
							ranges[jdx] = ranges[len(ranges)-1]
							ranges = ranges[:len(ranges)-1]
						}

						break
					}
				}
			}
		}

		if len(ranges) == 2 {
			if r, ok := overlap(ranges[0], ranges[1]); ok {
				return []vec2{r}
			}
		}

		return ranges
	}

	notCovered := func(ranges []vec2, min, max int) int {
		r := reduce(ranges)
		if len(r) == 2 {
			return r[0].y + 1
		}

		return -1
	}

	// ✝ God forgive me for this code ✝
	part2 := func(sensors []sensor, maxSearch int) int {
		filtered := make([]sensor, 0, len(sensors))
		for _, s := range sensors {
			min := vec2{x: s.pos.x - s.distToBeacon, y: s.pos.y - s.distToBeacon}
			max := vec2{x: s.pos.x + s.distToBeacon, y: s.pos.y + s.distToBeacon}

			if min.x < 0 || min.y < 0 || max.x > maxSearch || max.y > maxSearch {
				continue
			}

			filtered = append(filtered, s)
		}

		ranges := make(map[int][]vec2)

		for y := 0; y <= maxSearch; y++ {
			for _, s := range sensors {
				if s.pos.y-s.distToBeacon > y || y > s.pos.y+s.distToBeacon {
					continue
				}

				distToRow := int(math.Abs(float64(s.pos.y - y)))

				from := s.pos.x - s.distToBeacon + distToRow
				if from < 0 {
					from = 0
				}

				to := s.pos.x + s.distToBeacon - distToRow
				if to > maxSearch {
					to = maxSearch
				}

				if from == 0 && to == maxSearch {
					continue
				}

				ranges[y] = append(ranges[y], vec2{x: from, y: to})
			}
		}

		for y, r := range ranges {
			if x := notCovered(r, 0, maxSearch); x != -1 {
				return tuning(vec2{x: x, y: y})
			}
		}

		return -1
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "Sensor at x=2, y=18: closest beacon is at x=-2, y=15\nSensor at x=9, y=16: closest beacon is at x=10, y=16\nSensor at x=13, y=2: closest beacon is at x=15, y=3\nSensor at x=12, y=14: closest beacon is at x=10, y=16\nSensor at x=10, y=20: closest beacon is at x=10, y=16\nSensor at x=14, y=17: closest beacon is at x=10, y=16\nSensor at x=8, y=7: closest beacon is at x=2, y=10\nSensor at x=2, y=0: closest beacon is at x=2, y=10\nSensor at x=0, y=11: closest beacon is at x=2, y=10\nSensor at x=20, y=14: closest beacon is at x=25, y=17\nSensor at x=17, y=20: closest beacon is at x=21, y=22\nSensor at x=16, y=7: closest beacon is at x=15, y=3\nSensor at x=14, y=3: closest beacon is at x=15, y=3\nSensor at x=20, y=1: closest beacon is at x=15, y=3"
		t.Log(part1(aoc.ParseLines(aoc.InputScanner(input), parse), 10))
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(aoc.ParseLines(aoc.PuzzleScanner(2022, 15), parse), 2000000))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "Sensor at x=2, y=18: closest beacon is at x=-2, y=15\nSensor at x=9, y=16: closest beacon is at x=10, y=16\nSensor at x=13, y=2: closest beacon is at x=15, y=3\nSensor at x=12, y=14: closest beacon is at x=10, y=16\nSensor at x=10, y=20: closest beacon is at x=10, y=16\nSensor at x=14, y=17: closest beacon is at x=10, y=16\nSensor at x=8, y=7: closest beacon is at x=2, y=10\nSensor at x=2, y=0: closest beacon is at x=2, y=10\nSensor at x=0, y=11: closest beacon is at x=2, y=10\nSensor at x=20, y=14: closest beacon is at x=25, y=17\nSensor at x=17, y=20: closest beacon is at x=21, y=22\nSensor at x=16, y=7: closest beacon is at x=15, y=3\nSensor at x=14, y=3: closest beacon is at x=15, y=3\nSensor at x=20, y=1: closest beacon is at x=15, y=3"
		t.Log(part2(aoc.ParseLines(aoc.InputScanner(input), parse), 20))
	})

	t.Run("Part 2", func(t *testing.T) {
		t.Log(part2(aoc.ParseLines(aoc.PuzzleScanner(2022, 15), parse), 4000000))
	})
}
