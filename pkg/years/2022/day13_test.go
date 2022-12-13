package y2022

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay13(t *testing.T) {
	const (
		tValue = 1
		tList  = 2
	)

	type packet struct {
		t     byte
		value any
	}

	consumeNumber := func(line string) packet {
		num, _ := strconv.Atoi(line)
		return packet{
			t:     tValue,
			value: num,
		}
	}

	var consumeList func(line string, idx *int) packet
	consumeList = func(line string, idx *int) packet {
		p := packet{t: tList}
		v := make([]packet, 0, 10)

		numBuffer := ""

		for {
			c := line[*idx]
			*idx++

			switch c {
			case '[':
				v = append(v, consumeList(line, idx))
			case ']':
				if numBuffer != "" {
					v = append(v, consumeNumber(numBuffer))
					numBuffer = ""
				}

				p.value = v
				return p
			case ',':
				if numBuffer != "" {
					v = append(v, consumeNumber(numBuffer))
					numBuffer = ""
				}
			default:
				numBuffer += string(c)
			}
		}
	}

	consumePacket := func(line string) packet {
		if line == "" {
			return packet{}
		}

		idx := 1

		return consumeList(line, &idx)
	}

	parse := func(line string) packet {
		return consumePacket(line)
	}

	interpret := func(packets []packet) [][2]packet {
		ret := make([][2]packet, 0, len(packets)/3+1)
		var pair [2]packet
		idx := 0

		for _, p := range packets {
			if p.t == 0 {
				idx = 0
				ret = append(ret, pair)
				continue
			}

			pair[idx] = p
			idx++
		}

		ret = append(ret, pair)

		return ret
	}

	var printPacket func(p packet) string
	printPacket = func(p packet) string {
		switch p.t {
		case tValue:
			return fmt.Sprintf("%d", p.value)
		case tList:
			v := p.value.([]packet)

			var b strings.Builder
			b.WriteString("[")
			for idx, sp := range v {
				if idx > 0 {
					b.WriteString(",")
				}
				b.WriteString(printPacket(sp))
			}
			b.WriteString("]")

			return b.String()
		}

		return "<>"
	}

	var compare func(p1, p2 packet) int
	compare = func(p1, p2 packet) int {
		if p1.t == tValue && p2.t == tValue {
			fmt.Println("cmp values", p1.value, p2.value)
			return p2.value.(int) - p1.value.(int)
		}

		if p1.t == tList && p2.t == tList {
			fmt.Println("cmp lists", printPacket(p1), printPacket(p2))

			p2l := len(p2.value.([]packet))
			p1l := len(p1.value.([]packet))

			for idx, p1v := range p1.value.([]packet) {
				if idx >= p2l {
					fmt.Println("p1 longer")
					return -1
				}

				p2v := p2.value.([]packet)[idx]
				cmp := compare(p1v, p2v)
				if cmp == 0 {
					continue
				}

				return cmp
			}

			if p1l == p2l {
				fmt.Println("p1 == p2")
				return 0
			}

			fmt.Println("p1 shorter")
			return 1
		}

		if p1.t == tValue {
			fmt.Println("conv", p1.value, "to", "[", p1.value, "]")
			p1.value = []packet{{t: tValue, value: p1.value}}
			p1.t = tList

			return compare(p1, p2)
		} else {
			fmt.Println("conv", p2.value, "to", "[", p2.value, "]")
			p2.value = []packet{{t: tValue, value: p2.value}}
			p2.t = tList

			return compare(p1, p2)
		}
	}

	part1 := func(pairs [][2]packet) int {
		total := 0

		for idx, pair := range pairs {
			cmp := compare(pair[0], pair[1])
			if cmp > 0 {
				total += idx + 1
			}

			if cmp == 0 {
				fmt.Println(idx)
			}
		}

		return total
	}

	part2 := func(packets []packet) int {
		div0 := consumePacket("[[2]]")
		div1 := consumePacket("[[6]]")

		packets = append(packets,
			div0,
			div1,
		)

		cleanPackets := make([]packet, 0, len(packets))
		for _, p := range packets {
			if p.t == 0 {
				continue
			}

			cleanPackets = append(cleanPackets, p)
		}

		sort.Slice(cleanPackets, func(i, j int) bool {
			return compare(cleanPackets[i], cleanPackets[j]) > 0
		})

		idx0 := 0
		idx1 := 0

		for idx, p := range cleanPackets {
			if compare(p, div0) == 0 {
				idx0 = idx + 1
			} else if compare(p, div1) == 0 {
				idx1 = idx + 1
			}

			if idx0 > 0 && idx1 > 0 {
				break
			}
		}

		return idx0 * idx1
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "[1,1,3,1,1]\n[1,1,5,1,1]\n\n[[1],[2,3,4]]\n[[1],4]\n\n[9]\n[[8,7,6]]\n\n[[4,4],4,4]\n[[4,4],4,4,4]\n\n[7,7,7,7]\n[7,7,7]\n\n[]\n[3]\n\n[[[]]]\n[[]]\n\n[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]"
		t.Log(part1(interpret(aoc.ParseLines(aoc.InputScanner(input), parse))))
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(interpret(aoc.ParseLines(aoc.PuzzleScanner(2022, 13), parse))))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "[1,1,3,1,1]\n[1,1,5,1,1]\n\n[[1],[2,3,4]]\n[[1],4]\n\n[9]\n[[8,7,6]]\n\n[[4,4],4,4]\n[[4,4],4,4,4]\n\n[7,7,7,7]\n[7,7,7]\n\n[]\n[3]\n\n[[[]]]\n[[]]\n\n[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]"
		t.Log(part2(aoc.ParseLines(aoc.InputScanner(input), parse)))
	})

	t.Run("Example 2", func(t *testing.T) {
		t.Log(part2(aoc.ParseLines(aoc.PuzzleScanner(2022, 13), parse)))
	})
}
