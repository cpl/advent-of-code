package aoc

import (
	"bufio"
	"fmt"
	"io"
)

func ParseLines[T any](r *bufio.Scanner, fn func(line string) T) []T {
	var lines []T

	for r.Scan() {
		lines = append(lines, fn(r.Text()))
	}

	return lines
}

func ParseChars[T any](r *bufio.Reader, fn func(char rune) T) []T {
	var chars []T

	for {
		char, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(fmt.Errorf("cannot read rune: %w", err))
		}

		chars = append(chars, fn(char))
	}

	return chars
}
