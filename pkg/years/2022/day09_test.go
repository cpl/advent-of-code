package y2022

import (
	"fmt"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay09(t *testing.T) {
	type movement struct {
		direction rune
		distance  int
	}

	parse := func(line string) (m movement) {
		_, _ = fmt.Sscanf(line, "%c %d", &m.direction, &m.distance)
		return
	}

	type pos struct {
		x, y int
	}

	adjacent := func(p0, p1 pos) bool {
		return (p0.x == p1.x && (p0.y == p1.y+1 || p0.y == p1.y-1)) ||
			(p0.y == p1.y && (p0.x == p1.x+1 || p0.x == p1.x-1)) ||
			(p0.x == p1.x+1 && p0.y == p1.y+1) ||
			(p0.x == p1.x+1 && p0.y == p1.y-1) ||
			(p0.x == p1.x-1 && p0.y == p1.y+1) ||
			(p0.x == p1.x-1 && p0.y == p1.y-1)
	}

	move := func(p pos, m movement, postMove func(dp pos)) pos {
		for m.distance > 0 {
			switch m.direction {
			case 'U':
				p.y++
			case 'D':
				p.y--
			case 'L':
				p.x--
			case 'R':
				p.x++
			case 'q':
				p.x--
				p.y++
			case 'e':
				p.x++
				p.y++
			case 'z':
				p.x--
				p.y--
			case 'c':
				p.x++
				p.y--
			}

			m.distance--
			if postMove != nil {
				postMove(p)
			}
		}

		return p
	}

	moveCloserTo := func(from, target pos) (m movement) {
		m.distance = 1

		xd := ' '
		yd := ' '

		if from.x < target.x {
			xd = 'R'
		} else if from.x > target.x {
			xd = 'L'
		}

		if from.y < target.y {
			yd = 'U'
		} else if from.y > target.y {
			yd = 'D'
		}

		fd := ' '
		if xd != ' ' && yd != ' ' {
			if yd == 'U' {
				fd = 'q'
			} else {
				fd = 'z'
			}

			if fd == 'q' {
				if xd == 'R' {
					fd = 'e'
				}
			} else {
				if xd == 'R' {
					fd = 'c'
				}
			}
		} else {
			if xd != ' ' {
				fd = xd
			} else {
				fd = yd
			}
		}

		m.direction = fd

		return
	}

	part1 := func(movements []movement) map[string]int {
		headPos := pos{0, 0}
		tailPos := pos{0, 0}

		visited := make(map[string]int, 100)
		visited[fmt.Sprintf("%d,%d", 0, 0)] = 1

		for _, m := range movements {
			headPos = move(headPos, m, func(dp pos) {
				if adjacent(dp, tailPos) {
					return
				}

				mdt := moveCloserTo(tailPos, dp)

				tailPos = move(tailPos, mdt, func(dt pos) {
					visited[fmt.Sprintf("%d,%d", dt.x, dt.y)]++
				})
			})
		}

		return visited
	}

	part2 := func(movements []movement) map[string]int {
		headPos := pos{0, 0}
		midkPos := [8]pos{}
		tailPos := pos{0, 0}

		visited := make(map[string]int, 100)
		visited[fmt.Sprintf("%d,%d", 0, 0)] = 1

		moveMiddle := func(dp pos) {
			if !adjacent(midkPos[0], dp) {
				dm := moveCloserTo(midkPos[0], dp)
				midkPos[0] = move(midkPos[0], dm, nil)
			}

			for idx := 1; idx < len(midkPos); idx++ {
				if adjacent(midkPos[idx], midkPos[idx-1]) {
					continue
				}

				dm := moveCloserTo(midkPos[idx], midkPos[idx-1])
				midkPos[idx] = move(midkPos[idx], dm, nil)
			}
		}

		moveTail := func() {
			if adjacent(tailPos, midkPos[7]) {
				return
			}

			mdt := moveCloserTo(tailPos, midkPos[7])

			tailPos = move(tailPos, mdt, func(dt pos) {
				visited[fmt.Sprintf("%d,%d", dt.x, dt.y)]++
			})
		}

		for _, m := range movements {
			headPos = move(headPos, m, func(dp pos) {
				moveMiddle(dp)
				moveTail()
			})
		}

		return visited
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2"
		t.Log(part1(aoc.ParseLines(aoc.InputScanner(input), parse)))
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(aoc.ParseLines(aoc.PuzzleScanner(2022, 9), parse)))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2"
		t.Log(len(part2(aoc.ParseLines(aoc.InputScanner(input), parse))))
		input = "R 5\nU 8\nL 8\nD 3\nR 17\nD 10\nL 25\nU 20"
		visited := part2(aoc.ParseLines(aoc.InputScanner(input), parse))
		t.Log(len(visited))
		for y := 10; y > -10; y-- {
			for x := -20; x < 20; x++ {
				if x == 0 && y == 0 {
					fmt.Print("s")
					continue
				}

				if _, ok := visited[fmt.Sprintf("%d,%d", x, y)]; !ok {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
			}
			fmt.Println()
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		t.Log(len(part2(aoc.ParseLines(aoc.PuzzleScanner(2022, 9), parse))))
	})
}
