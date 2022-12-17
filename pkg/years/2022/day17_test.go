package y2022

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay17(t *testing.T) {
	type vec2 struct {
		x, y int
	}

	vec2add := func(a, b vec2) vec2 {
		return vec2{a.x + b.x, a.y + b.y}
	}

	type rock struct {
		points []vec2
	}

	rocks := []rock{
		{
			points: []vec2{
				{0, 0},
				{1, 0},
				{2, 0},
				{3, 0},
			},
		},
		{
			points: []vec2{
				{1, 0},
				{0, 1},
				{1, 1},
				{2, 1},
				{1, 2},
			},
		},
		{
			points: []vec2{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
				{2, 2},
			},
		},
		{
			points: []vec2{
				{0, 0},
				{0, 1},
				{0, 2},
				{0, 3},
			},
		},
		{
			points: []vec2{
				{1, 0},
				{0, 0},
				{0, 1},
				{1, 1},
			},
		},
	}

	type tower struct {
		maxHeight int
		grid      [][7]byte
	}

	printTower := func(t *tower) {
		for y := t.maxHeight + 5; y >= 0; y-- {
			for x := -1; x < len(t.grid[y])+1; x++ {
				if x == -1 {
					if y%10 == 0 {
						fmt.Printf(" %4d |", y)
					} else {
						fmt.Printf("      |")
					}
					continue
				} else if x == len(t.grid[y]) {
					fmt.Printf("|")
					continue
				}

				if v := t.grid[y][x]; v == 0 {
					if y == t.maxHeight {
						fmt.Print("-")
					} else {
						fmt.Print(".")
					}
				} else {
					fmt.Print(strconv.Itoa(int(v)))
				}
			}
			fmt.Println()
		}
		fmt.Println("      +-------+")
	}

	printRockFall := func(r *rock, p vec2, maxHeight int) {
		pointMap := make(map[vec2]struct{})
		for _, point := range r.points {
			pointMap[vec2add(point, p)] = struct{}{}
		}

		for y := maxHeight + 5; y >= 0; y-- {
			for x := -1; x < 8; x++ {
				if x == -1 || x == 7 {
					fmt.Print("|")
					continue
				}

				if _, ok := pointMap[vec2{x, y}]; ok {
					fmt.Print("@")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println("+-------+")
		fmt.Println()
	}
	_ = printRockFall

	physics := func(tow *tower, r *rock, pos vec2, jet rune) (vec2, rune) {
		//fmt.Println("jet", string(jet))
		//printRockFall(r, pos, tow.maxHeight)

		tPos := vec2{}
		switch jet {
		case '<':
			tPos = vec2add(pos, vec2{-1, 0})
		case '>':
			tPos = vec2add(pos, vec2{1, 0})
		}

		status := '.'
		for _, p := range r.points {
			p = vec2add(p, tPos)

			if p.x < 0 || p.x >= len(tow.grid[p.y]) {
				status = '|'
				break
			}

			if tow.grid[p.y][p.x] != 0 {
				status = '#'
				break
			}
		}

		if status == '|' || status == '#' {
			tPos = pos
		} else {
			pos = tPos
		}

		tPos = vec2add(tPos, vec2{0, -1})
		status = '.'
		for _, p := range r.points {
			p = vec2add(p, tPos)

			if p.y < 0 {
				status = '-'
				break
			}

			if tow.grid[p.y][p.x] != 0 {
				status = '#'
				break
			}
		}

		if status == '#' || status == '-' {
			tPos = pos
		}

		return tPos, status
	}

	rockIdx := byte(1)
	commitRock := func(t *tower, r *rock, pos vec2) {
		maxY := t.maxHeight

		for _, p := range r.points {
			p = vec2add(p, pos)
			t.grid[p.y][p.x] = rockIdx

			if p.y+1 > maxY {
				maxY = p.y + 1
			}
		}

		rockIdx++
		if rockIdx >= 10 {
			rockIdx = 1
		}

		t.maxHeight = maxY
	}

	simulate := func(tow *tower, stream []rune, start, iterations, tick int) int {
		for iter := start; iter < start+iterations; iter++ {
			r := rocks[iter%len(rocks)]
			rPos := vec2{2, tow.maxHeight + 3}

			status := ' '
			for {
				jet := stream[tick%len(stream)]
				tick++

				rPos, status = physics(tow, &r, rPos, jet)
				if status == '.' {
					continue
				} else if status == '#' || status == '-' {
					commitRock(tow, &r, rPos)
					break
				}
			}
		}

		return tick
	}

	part1 := func(stream []rune, iterations int) int {
		tow := tower{
			maxHeight: 0,
			grid:      make([][7]byte, 10000),
		}

		simulate(&tow, stream, 0, iterations, 0)
		printTower(&tow)

		return tow.maxHeight
	}

	indexPatternEq := func(a, b [][7]byte) bool {
		for y := range a {
			for x, v := range a[y] {
				if v == 0 {
					if b[y][x] != 0 {
						return false
					}

					continue
				}

				if b[y][x] == 0 {
					return false
				}
			}
		}

		return true
	}

	indexPattern := func(grid [][7]byte, pattern [][7]byte) int {
		for towY := 0; towY < len(grid)-len(pattern); towY++ {
			if indexPatternEq(grid[towY:towY+len(pattern)], pattern) {
				return towY
			}
		}

		return -1
	}

	parsePattern := func(pattern string) [][7]byte {
		lines := strings.Split(pattern, "\n")

		l := len(lines) - 1
		p := make([][7]byte, l)

		for y, line := range lines[1:] {
			x := 0
			for _, c := range line {
				switch c {
				case '.':
					p[l-y-1][x] = 0
					x++
				case '#', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					p[l-y-1][x] = 1
					x++
				default:
					continue
				}
			}
		}

		return p
	}

	part2 := func(stream []rune, pattern [][7]byte) int64 {
		tow := tower{
			maxHeight: 0,
			grid:      make([][7]byte, 10000),
		}

		tick := 0
		iter := 0

		// find first occurence of pattern
		pIdx1 := -1
		iter1 := 0
		height1 := 0
		for {
			pIdx := indexPattern(tow.grid, pattern)
			if pIdx != -1 {
				pIdx1 = pIdx
				iter1 = iter
				height1 = tow.maxHeight
				break
			}

			tick = simulate(&tow, stream, iter, 1, tick)
			iter += 1
		}

		// find recurring pattern height and iteration
		iterLast := iter1
		from := pIdx1
		delta := -1
		height := tow.maxHeight
		for {
			pIdx := indexPattern(tow.grid[from+1:], pattern)
			if pIdx != -1 {
				if delta == iter-iterLast {
					break
				}
				height = tow.maxHeight - height
				delta = iter - iterLast
				from += pIdx + 1
				iterLast = iter
			}

			tick = simulate(&tow, stream, iter, 1, tick)
			iter += 1
		}

		target := 1000000000000

		// number of cycles (excluding first occurence)
		cycles := (target - iter1) / delta

		// remaining iterations
		rem := (target - iter1) % delta

		// height of tower after first cycle + repeating cycles
		heightSoFar := int64(cycles)*int64(height) + int64(height1)

		// simulate remaining iterations
		currentHeight := tow.maxHeight
		simulate(&tow, stream, iter, rem, tick)
		heightSoFar += int64(tow.maxHeight - currentHeight)

		return heightSoFar
	}

	t.Run("Example 1", func(t *testing.T) {
		input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
		t.Log(part1([]rune(input), 2022))
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1([]rune(aoc.PuzzleString(2022, 17)), 2022))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
		pattern := parsePattern(`
|.455...|
|.455...|
|.4.3...|
|.4.3.2.|
|.333222|
|.11112.|`)

		t.Log(part2([]rune(input), pattern))
	})

	t.Run("Part 2", func(t *testing.T) {
		pattern := parsePattern(`
|.66....|
|.66....|
|..5....|
|..5....|
|..5.4..|
|..5.4..|
|.3444..|
|3332222|
|.311...|
|..11...|`)
		t.Log(part2([]rune(aoc.PuzzleString(2022, 17)), pattern))
	})
}
