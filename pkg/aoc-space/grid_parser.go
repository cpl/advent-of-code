package aoc_space

import (
	"bufio"
	"fmt"

	aocparse "github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func Grid2CHF(ch byte) byte {
	return ch
}

func Grid2Parser[T any](chf func(ch byte) T) aocparse.Parser[*Grid2[T]] {
	return func(scanner *bufio.Scanner) *Grid2[T] {
		lines := make([]string, 0, 64)

		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		grid := NewGrid[T](&Space{
			Size: Vec{
				X: int64(len(lines[0])),
				Y: int64(len(lines)),
			},
		})

		for y, line := range lines {
			for x, ch := range []byte(line) {
				if !grid.Set(chf(ch), Vec2(x, y)) {
					panic(fmt.Sprintf("failed to add to grid %d,%d = %s", x, y, string(ch)))
				}
			}
		}

		return grid
	}
}
