package y2023

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay15(t *testing.T) {
	t.Parallel()

	holidayHash := func(b []byte) byte {
		v := uint64(0)
		for _, c := range b {
			v += uint64(c)
			v *= 17
			v = v % 256
		}

		return byte(v)
	}

	part1 := func(input string) int {
		total := 0

		for _, s := range strings.Split(input, ",") {
			total += int(holidayHash([]byte(s)))
		}

		return total
	}

	type hmv struct {
		s   string
		num int
	}

	part2 := func(input string) int {
		sequence := strings.Split(input, ",")
		var hm [256][]*hmv

		for _, s := range sequence {
			if idx := strings.IndexByte(s, '='); idx != -1 {
				n := s[idx+1:]
				s = s[:idx]
				h := holidayHash([]byte(s))
				num, _ := strconv.Atoi(n)

				update := false
				for vdx, v := range hm[h] {
					if v.s == s {
						hm[h][vdx].num = num
						update = true
						break
					}
				}

				if update {
					continue
				}

				hm[h] = append(hm[h], &hmv{s, num})
				continue
			}

			if idx := strings.IndexByte(s, '-'); idx != -1 {
				s = s[:idx]
				h := holidayHash([]byte(s))

				values := hm[h]
				if len(values) == 0 {
					continue
				}

				for vdx, v := range values {
					if v.s == s {
						values = append(values[:vdx], values[vdx+1:]...)
						break
					}
				}

				hm[h] = values
			}
		}

		total := 0
		for idx, values := range hm {
			for vdx, v := range values {
				total += (idx + 1) * (vdx + 1) * v.num
			}
		}

		return total
	}

	t.Run("example 1a", func(t *testing.T) {
		t.Log(holidayHash([]byte("HASH")))
	})

	t.Run("example 1b", func(t *testing.T) {
		t.Log(part1("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(part1(aoc.PuzzleString(2023, 15)))
	})

	t.Run("example 2", func(t *testing.T) {
		t.Log(part2("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(part2(aoc.PuzzleString(2023, 15)))
	})
}
