package aoc_space

import (
	"bufio"
	"strconv"
	"strings"

	aocparse "github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func Vec2Parser(split byte) aocparse.Parser[[]Vec] {
	if split == 0 {
		split = ','
	}

	return func(scanner *bufio.Scanner) []Vec {
		out := make([]Vec, 0, 1024)

		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			vecs := strings.Split(scanner.Text(), string(split))
			if len(vecs) != 2 {
				panic("invalid input")
			}

			x, _ := strconv.ParseInt(vecs[0], 10, 64)
			y, _ := strconv.ParseInt(vecs[1], 10, 64)

			out = append(out, Vec{
				X: x,
				Y: y,
			})
		}

		return out
	}
}

func Vec3Parser(split byte) aocparse.Parser[[]Vec] {
	if split == 0 {
		split = ','
	}

	return func(scanner *bufio.Scanner) []Vec {
		out := make([]Vec, 0, 1024)

		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			vecs := strings.Split(scanner.Text(), string(split))
			if len(vecs) != 3 {
				panic("invalid input")
			}

			x, _ := strconv.ParseInt(vecs[0], 10, 64)
			y, _ := strconv.ParseInt(vecs[1], 10, 64)
			z, _ := strconv.ParseInt(vecs[2], 10, 64)

			out = append(out, Vec{
				X: x,
				Y: y,
				Z: z,
			})
		}

		return out
	}
}
