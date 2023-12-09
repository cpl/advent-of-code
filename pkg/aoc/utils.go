package aoc

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
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

func ParseLineNumber(scan *bufio.Scanner) []int {
	return ParseLines(scan, func(line string) int {
		num, _ := strconv.Atoi(line)
		return num
	})
}

func ParseLineNumbers(scan *bufio.Scanner) [][]int {
	return ParseLines(scan, func(line string) []int {
		fields := strings.Fields(line)
		numbers := make([]int, len(fields))

		for idx, field := range fields {
			numbers[idx], _ = strconv.Atoi(field)
		}

		return numbers
	})
}
