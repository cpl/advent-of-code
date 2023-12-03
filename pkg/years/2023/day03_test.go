package y2023

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

// 00000
// 00X00
// 00000

func TestSolveDay03(t *testing.T) {
	t.Parallel()

	isDigit := func(c byte) bool {
		return c >= '0' && c <= '9'
	}

	isSymbol := func(c byte) bool {
		if c == '.' {
			return false
		}

		return true
	}

	type landmark struct {
		isSymbol bool
		isPart   bool

		value int
	}

	parseEngine := func(input string) (map[int]*landmark, int) {
		engine := make(map[int]*landmark)
		width := 0
		ptr := 0

		accStart := -1
		acc := 0

		for _, b := range []byte(input) {
			if b == '\n' {
				if width == 0 {
					width = ptr
				}

				if accStart != -1 {
					part := &landmark{
						isPart: true,
						value:  acc,
					}

					for start := accStart; start < ptr; start++ {
						engine[start] = part
					}

					acc = 0
					accStart = -1
				}

				continue
			}

			ptr++

			if isDigit(b) {
				if accStart == -1 {
					accStart = ptr
				}

				digit := int(b - '0')
				acc = acc*10 + digit
				continue
			} else if accStart != -1 {
				part := &landmark{
					isPart: true,
					value:  acc,
				}

				for start := accStart; start < ptr; start++ {
					engine[start] = part
				}

				acc = 0
				accStart = -1
			}

			if isSymbol(b) {
				engine[ptr] = &landmark{
					isSymbol: true,
					value:    int(b),
				}
			}
		}

		return engine, width
	}

	part1 := func(engine map[int]*landmark, width int) int {
		sum := 0

		getNeighbouringPartValue := func(x, y int) int {
			idx := y*width + x + 1
			part, ok := engine[idx]
			if !ok {
				return 0
			}

			if !part.isPart {
				return 0
			}

			value := part.value
			part.value = 0

			return value
		}

		addNeighbouringParts := func(x, y int) {
			sum += getNeighbouringPartValue(x+1, y)
			sum += getNeighbouringPartValue(x+1, y+1)
			sum += getNeighbouringPartValue(x+1, y-1)
			sum += getNeighbouringPartValue(x-1, y)
			sum += getNeighbouringPartValue(x, y+1)
			sum += getNeighbouringPartValue(x-1, y-1)
			sum += getNeighbouringPartValue(x-1, y+1)
			sum += getNeighbouringPartValue(x, y-1)
		}

		for idx, v := range engine {
			if v.isSymbol {
				addNeighbouringParts(idx%width-1, idx/width)
			}
		}

		return sum
	}

	part2 := func(engine map[int]*landmark, width int) int {
		sum := 0

		getNeighbour := func(x, y int) *landmark {
			part, ok := engine[y*width+x+1]
			if !ok {
				return nil
			}

			if !part.isPart {
				return nil
			}

			return part
		}

		getNeighbours := func(x, y int) map[*landmark]struct{} {
			neighbours := make(map[*landmark]struct{}, 9)

			neighbours[getNeighbour(x+1, y)] = struct{}{}
			neighbours[getNeighbour(x+1, y+1)] = struct{}{}
			neighbours[getNeighbour(x+1, y-1)] = struct{}{}
			neighbours[getNeighbour(x-1, y)] = struct{}{}
			neighbours[getNeighbour(x, y+1)] = struct{}{}
			neighbours[getNeighbour(x-1, y-1)] = struct{}{}
			neighbours[getNeighbour(x-1, y+1)] = struct{}{}
			neighbours[getNeighbour(x, y-1)] = struct{}{}

			delete(neighbours, nil)

			return neighbours
		}

		gearRatio := func(x, y int) int {
			neighbours := getNeighbours(x, y)
			if len(neighbours) != 2 {
				return 0
			}

			product := 1
			for n := range neighbours {
				product = product * n.value
			}

			return product
		}

		for idx, v := range engine {
			if v.isSymbol && rune(v.value) == '*' {
				sum += gearRatio(idx%width-1, idx/width)
			}
		}

		return sum
	}

	t.Run("example 1", func(t *testing.T) {
		t.Log(part1(parseEngine("467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..")))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(part1(parseEngine(aoc.PuzzleString(2023, 3))))
	})

	t.Run("example 2", func(t *testing.T) {
		t.Log(part2(parseEngine("467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..")))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(part2(parseEngine(aoc.PuzzleString(2023, 3))))
	})
}
