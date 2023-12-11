package y2023

import (
	"fmt"
	"math"
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

	convertStartToPipe := func(startPos pos, p1, p2 *pipe) rune {
		startOpenN, startOpenS, startOpenE, startOpenW := false, false, false, false
		if p1.position.x == startPos.x {
			if p1.position.y < startPos.y {
				startOpenN = true
			} else {
				startOpenS = true
			}
		} else {
			if p1.position.x < startPos.x {
				startOpenW = true
			} else {
				startOpenE = true
			}
		}

		if p2.position.x == startPos.x {
			if p2.position.y < startPos.y {
				startOpenN = true
			} else {
				startOpenS = true
			}
		} else {
			if p2.position.x < startPos.x {
				startOpenW = true
			} else {
				startOpenE = true
			}
		}

		if startOpenN && startOpenS {
			return '│'
		} else if startOpenE && startOpenW {
			return '─'
		} else if startOpenN && startOpenE {
			return '└'
		} else if startOpenN && startOpenW {
			return '┘'
		} else if startOpenS && startOpenE {
			return '┌'
		} else if startOpenS && startOpenW {
			return '┐'
		} else {
			panic("not enough starting position connections")
		}
	}

	part2 := func(m [][]pipe, startPos pos) int {
		p1, p2 := getConnections(m, startPos.x, startPos.y)

		_ = walkAndMeasure(m, p1, 1, 1)
		_ = walkAndMeasure(m, p2, 1, 2)

		m[startPos.y][startPos.x].c = convertStartToPipe(startPos, p1, p2)
		countContained := 0

		for _, row := range m {
			contained := false

			for _, p := range row {
				loopPipe := p.seenWalk1 || p.seenWalk2

				if !loopPipe && contained {
					countContained++
					// fmt.Print("#")
				} else {
					if p.c == ' ' {
						p.c = '.'
					}

					// fmt.Print(string(p.c))
				}

				northPipe := p.c == '│' || p.c == '┘' || p.c == '└'

				if loopPipe && northPipe {
					contained = !contained
				}
			}

			// fmt.Println()
		}

		return countContained
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
		t.Log(part2(parseMap("...........\n.S-------7.\n.|F-----7|.\n.||.....||.\n.||.....||.\n.|L-7.F-J|.\n.|..|.|..|.\n.L--J.L--J.\n...........")))
	})

	t.Run("example 2b", func(t *testing.T) {
		t.Log(part2(parseMap(".F----7F7F7F7F-7....\n.|F--7||||||||FJ....\n.||.FJ||||||||L7....\nFJL7L7LJLJ||LJ.L-7..\nL--J.L7...LJS7F-7L7.\n....F-J..F7FJ|L7L7L7\n....L7.F7||L7|.L7L7|\n.....|FJLJ|FJ|F7|.LJ\n....FJL-7.||.||||...\n....L---J.LJ.LJLJ...")))
	})

	t.Run("example 2c", func(t *testing.T) {
		t.Log(part2(parseMap("....................\nFF7FSF7F7F7F7F7F---7\nL|LJ||||||||||||F--J\nFL-7LJLJ||||||LJL-77\nF--JF--7||LJLJ7F7FJ-\nL---JF-JLJ.||-FJLJJ7\n|F|F-JF---7F7-L7L|7|\n|FFJF7L7F-JF7|JL---7\n7-L-JL7||F7|L7F-7F7|\nL.L7LFJ|||||FJL7||LJ\nL7JLJL-JLJLJL--JLJ.L")))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(part2(parseMap(aoc.PuzzleString(2023, 10))))
	})
}
