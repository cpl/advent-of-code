package aoc_parse

import (
	"bufio"
	"fmt"
)

func EachLine[T any](fn func(line string) T) Parser[[]T] {
	return func(r *bufio.Scanner) []T {
		var lines []T

		for r.Scan() {
			lines = append(lines, fn(r.Text()))
		}

		return lines
	}
}

func EachColumn[T any](strict int, fn func(line string) []T) Parser[[][]T] {
	return func(r *bufio.Scanner) [][]T {
		if strict <= 0 {
			panic("strict must be greater than zero")
		}

		valueCount := 0
		columns := make([][]T, strict)

		for r.Scan() {
			line := r.Text()
			values := fn(line)

			if len(values) != strict {
				panic(fmt.Errorf("invalid number of values for column %d: %q (%v)", valueCount, line, values))
			}

			for cidx := range values {
				columns[cidx] = append(columns[cidx], values[cidx])
			}

			valueCount++
		}

		return columns
	}
}
