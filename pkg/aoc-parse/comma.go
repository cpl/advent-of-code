package aoc_parse

import (
	"bufio"
	"strings"
)

func CommaListEach[T any](each func(s string) T) Parser[[]T] {
	// todo stream parse

	return func(scanner *bufio.Scanner) []T {
		scanner.Split(bufio.ScanLines)
		out := make([]T, 0, 64)

		for scanner.Scan() {
			for _, line := range strings.Split(scanner.Text(), ",") {
				out = append(out, each(line))
			}
		}

		return out
	}
}
