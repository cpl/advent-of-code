package aoc_parse

import (
	"bufio"
	"fmt"
)

func EachLine[T any](r *bufio.Scanner, fn func(line string) T) []T {
	var lines []T

	for r.Scan() {
		lines = append(lines, fn(r.Text()))
	}

	return lines
}

func EachColumn[T any](r *bufio.Scanner, strict int, fn func(line string) []T) [][]T {
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
