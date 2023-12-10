package y2023

import (
	"fmt"
	"math"
	"os"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay10(t *testing.T) {
	t.Parallel()

	type pos struct {
		x, y int
	}

	type pipe struct {
		c                    rune
		distance             int
		position             pos
		seenWalk1, seenWalk2 bool
	}

	parsePipe := func(r rune) pipe {
		var c rune

		switch r {
		case '|':
			c = '│'
		case '-':
			c = '─'
		case 'L':
			c = '└'
		case 'J':
			c = '┘'
		case '7':
			c = '┐'
		case 'F':
			c = '┌'
		case '.':
			c = ' '
		case 'S':
			return pipe{
				c:         '#',
				distance:  0,
				seenWalk1: true,
				seenWalk2: true,
			}
		default:
			panic(fmt.Sprintf("unknown pipe: %c", r))
		}

		return pipe{
			c:        c,
			distance: math.MaxInt,
		}
	}

	parseMap := func(input string) ([][]pipe, pos) {
		lines := strings.Split(input, "\n")
		m := make([][]pipe, 0, len(lines))
		startX, startY := -1, -1

		for y, line := range lines {
			row := make([]pipe, 0, len(line))
			for x, r := range line {
				p := parsePipe(r)
				p.position = pos{x, y}

				row = append(row, p)
				if r == 'S' {
					startX = x
					startY = y
				}
			}

			m = append(m, row)
		}

		return m, pos{startX, startY}
	}

	getConnections := func(m [][]pipe, x, y int) (*pipe, *pipe) {
		var pipe1, pipe2 *pipe
		current := m[y][x]

		switch current.c {
		case ' ':
			return nil, nil
		case '#':
			// assume start is not on the edge

			north := m[y-1][x]
			south := m[y+1][x]
			east := m[y][x+1]
			west := m[y][x-1]

			switch north.c {
			case '│', '┐', '┌':
				pipe1 = &north
			}

			switch south.c {
			case '│', '┘', '└':
				if pipe1 != nil {
					return pipe1, &south
				} else {
					pipe1 = &south
				}
			}

			switch east.c {
			case '─', '┐', '┘':
				if pipe1 != nil {
					return pipe1, &east
				} else {
					pipe1 = &east
				}
			}

			switch west.c {
			case '─', '┌', '└':
				if pipe1 != nil {
					return pipe1, &west
				} else {
					pipe1 = &west
				}
			}

			panic("not enough starting position connections")
		case '│':
			if y > 0 {
				pipe1 = &m[y-1][x]
			}

			if y < len(m)-1 {
				pipe2 = &m[y+1][x]
			}
		case '─':
			if x > 0 {
				pipe1 = &m[y][x-1]
			}

			if x < len(m[y])-1 {
				pipe2 = &m[y][x+1]
			}
		case '└':
			if y > 0 {
				pipe1 = &m[y-1][x]
			}

			if x < len(m[y])-1 {
				pipe2 = &m[y][x+1]
			}
		case '┘':
			if x > 0 {
				pipe1 = &m[y][x-1]
			}

			if y > 0 {
				pipe2 = &m[y-1][x]
			}
		case '┐':
			if y < len(m)-1 {
				pipe1 = &m[y+1][x]
			}

			if x > 0 {
				pipe2 = &m[y][x-1]
			}
		case '┌':
			if x < len(m[y])-1 {
				pipe1 = &m[y][x+1]
			}

			if y < len(m)-1 {
				pipe2 = &m[y+1][x]
			}
		default:
			panic(fmt.Sprintf("unknown pipe: %c", current.c))
		}

		return pipe1, pipe2
	}

	var walkAndMeasure func(m [][]pipe, p *pipe, walked, seen int) int
	walkAndMeasure = func(m [][]pipe, p *pipe, walked, seen int) int {
		switch seen {
		case 1:
			if p.seenWalk1 {
				return p.distance
			}
			p.seenWalk1 = true
		case 2:
			if p.seenWalk2 {
				return p.distance
			}
			p.seenWalk2 = true
		}

		if p.distance <= walked {
			return p.distance
		}

		p.distance = walked

		next1, next2 := getConnections(m, p.position.x, p.position.y)

		w1 := walkAndMeasure(m, next1, walked+1, seen)
		w2 := walkAndMeasure(m, next2, walked+1, seen)

		return max(w1, w2)
	}

	part1 := func(m [][]pipe, startPos pos) int {
		p1, p2 := getConnections(m, startPos.x, startPos.y)

		_ = walkAndMeasure(m, p1, 1, 1)
		w2 := walkAndMeasure(m, p2, 1, 2)

		return w2
	}

	part2 := func(m [][]pipe, startPos pos) int {
		p1, p2 := getConnections(m, startPos.x, startPos.y)

		_ = walkAndMeasure(m, p1, 1, 1)
		_ = walkAndMeasure(m, p2, 1, 2)

		fp, _ := os.OpenFile("map.txt", os.O_CREATE|os.O_WRONLY, 0o644)
		defer fp.Close()

		for _, row := range m {
			for _, p := range row {
				if p.distance == math.MaxInt {
					fp.WriteString(" ")
					continue
				}

				fp.WriteString(string(p.c))
			}
			fp.WriteString("\n")
		}

		return 0
	}

	t.Run("example 1a", func(t *testing.T) {
		t.Log(part1(parseMap(".....\n.S-7.\n.|.|.\n.L-J.\n.....")))
	})

	t.Run("example 1b", func(t *testing.T) {
		t.Log(part1(parseMap("...F7.\n..FJ|.\n.SJ.L7\n.|F--J\n.LJ...")))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(part1(parseMap(aoc.PuzzleString(2023, 10))))
	})

	t.Run("example 2a", func(t *testing.T) {
		t.Log(part2(parseMap("FF7FSF7F7F7F7F7F---7\nL|LJ||||||||||||F--J\nFL-7LJSJ||||||LJL-77\nF--JF--7||LJLJ7F7FJ-\nL---JF-JLJ.||-FJLJJ7\n|F|F-JF---7F7-L7L|7|\n|FFJF7L7F-JF7|JL---7\n7-L-JL7||F7|L7F-7F7|\nL.L7LFJ|||||FJL7||LJ\nL7JLJL-JLJLJL--JLJ.L")))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(part2(parseMap(aoc.PuzzleString(2023, 10))))
		// neah, not happening
	})
}
