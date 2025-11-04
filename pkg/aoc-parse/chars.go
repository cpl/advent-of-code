package aoc_parse

import (
	"bufio"
	"fmt"
	"io"
)

func EachChar[T any](r *bufio.Reader, fn func(char rune) T) []T {
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
