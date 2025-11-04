package y2022

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	"github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func TestSolveDay10(t *testing.T) {
	type instruction struct {
		op    byte
		value int
	}

	const OpNOOP = 0
	const OpADDX = 1

	parse := func(line string) instruction {
		if line == "noop" {
			return instruction{}
		}

		value, _ := strconv.Atoi(line[strings.Index(line, " ")+1:])
		return instruction{
			op:    OpADDX,
			value: value,
		}
	}

	execute := func(instructions []instruction, interrupt func(c, r int)) int {
		cycle := 1
		register := 1

		tick := func(n int) {
			for n > 0 {
				interrupt(cycle, register)
				cycle++
				n--
			}
		}

		for _, ins := range instructions {
			switch ins.op {
			case OpNOOP:
				tick(1)
			case OpADDX:
				tick(2)
				register += ins.value
			}
		}

		return register * cycle
	}

	part1 := func(instructions []instruction) int {
		var total int

		execute(instructions, func(c, r int) {
			switch c {
			case 20:
				total += r * c
			case 60:
				total += r * c
			case 100:
				total += r * c
			case 140:
				total += r * c
			case 180:
				total += r * c
			case 220:
				total += r * c
			}
		})

		return total
	}

	part2 := func(instructions []instruction) [6 * 40]int {
		screen := [6 * 40]int{}
		pos := 0

		execute(instructions, func(c, r int) {
			tPos := pos % 40
			if r == tPos || r-1 == tPos || r+1 == tPos {
				screen[pos] = 1
			}

			pos++
		})

		return screen
	}

	printScreen := func(screen [6 * 40]int, width int) {
		for idx := 0; idx < len(screen); idx++ {
			if idx%width == 0 {
				fmt.Println()
			}

			if screen[idx] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("█")
			}
		}

		fmt.Println()
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "addx 15\naddx -11\naddx 6\naddx -3\naddx 5\naddx -1\naddx -8\naddx 13\naddx 4\nnoop\naddx -1\naddx 5\naddx -1\naddx 5\naddx -1\naddx 5\naddx -1\naddx 5\naddx -1\naddx -35\naddx 1\naddx 24\naddx -19\naddx 1\naddx 16\naddx -11\nnoop\nnoop\naddx 21\naddx -15\nnoop\nnoop\naddx -3\naddx 9\naddx 1\naddx -3\naddx 8\naddx 1\naddx 5\nnoop\nnoop\nnoop\nnoop\nnoop\naddx -36\nnoop\naddx 1\naddx 7\nnoop\nnoop\nnoop\naddx 2\naddx 6\nnoop\nnoop\nnoop\nnoop\nnoop\naddx 1\nnoop\nnoop\naddx 7\naddx 1\nnoop\naddx -13\naddx 13\naddx 7\nnoop\naddx 1\naddx -33\nnoop\nnoop\nnoop\naddx 2\nnoop\nnoop\nnoop\naddx 8\nnoop\naddx -1\naddx 2\naddx 1\nnoop\naddx 17\naddx -9\naddx 1\naddx 1\naddx -3\naddx 11\nnoop\nnoop\naddx 1\nnoop\naddx 1\nnoop\nnoop\naddx -13\naddx -19\naddx 1\naddx 3\naddx 26\naddx -30\naddx 12\naddx -1\naddx 3\naddx 1\nnoop\nnoop\nnoop\naddx -9\naddx 18\naddx 1\naddx 2\nnoop\nnoop\naddx 9\nnoop\nnoop\nnoop\naddx -1\naddx 2\naddx -37\naddx 1\naddx 3\nnoop\naddx 15\naddx -21\naddx 22\naddx -6\naddx 1\nnoop\naddx 2\naddx 1\nnoop\naddx -10\nnoop\nnoop\naddx 20\naddx 1\naddx 2\naddx 2\naddx -6\naddx -11\nnoop\nnoop\nnoop"
		t.Log(part1(aoc_parse.EachLine(aoc.InputScanner(input), parse)))
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(aoc_parse.EachLine(aoc.PuzzleScanner(2022, 10), parse)))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "addx 15\naddx -11\naddx 6\naddx -3\naddx 5\naddx -1\naddx -8\naddx 13\naddx 4\nnoop\naddx -1\naddx 5\naddx -1\naddx 5\naddx -1\naddx 5\naddx -1\naddx 5\naddx -1\naddx -35\naddx 1\naddx 24\naddx -19\naddx 1\naddx 16\naddx -11\nnoop\nnoop\naddx 21\naddx -15\nnoop\nnoop\naddx -3\naddx 9\naddx 1\naddx -3\naddx 8\naddx 1\naddx 5\nnoop\nnoop\nnoop\nnoop\nnoop\naddx -36\nnoop\naddx 1\naddx 7\nnoop\nnoop\nnoop\naddx 2\naddx 6\nnoop\nnoop\nnoop\nnoop\nnoop\naddx 1\nnoop\nnoop\naddx 7\naddx 1\nnoop\naddx -13\naddx 13\naddx 7\nnoop\naddx 1\naddx -33\nnoop\nnoop\nnoop\naddx 2\nnoop\nnoop\nnoop\naddx 8\nnoop\naddx -1\naddx 2\naddx 1\nnoop\naddx 17\naddx -9\naddx 1\naddx 1\naddx -3\naddx 11\nnoop\nnoop\naddx 1\nnoop\naddx 1\nnoop\nnoop\naddx -13\naddx -19\naddx 1\naddx 3\naddx 26\naddx -30\naddx 12\naddx -1\naddx 3\naddx 1\nnoop\nnoop\nnoop\naddx -9\naddx 18\naddx 1\naddx 2\nnoop\nnoop\naddx 9\nnoop\nnoop\nnoop\naddx -1\naddx 2\naddx -37\naddx 1\naddx 3\nnoop\naddx 15\naddx -21\naddx 22\naddx -6\naddx 1\nnoop\naddx 2\naddx 1\nnoop\naddx -10\nnoop\nnoop\naddx 20\naddx 1\naddx 2\naddx 2\naddx -6\naddx -11\nnoop\nnoop\nnoop"
		printScreen(part2(aoc_parse.EachLine(aoc.InputScanner(input), parse)), 40)
	})

	t.Run("Part 2", func(t *testing.T) {
		printScreen(part2(aoc_parse.EachLine(aoc.PuzzleScanner(2022, 10), parse)), 40)
	})
}
