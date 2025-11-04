package aoc_parse

import (
	"bufio"
	"strconv"
	"strings"
)

func NumberLine(scan *bufio.Scanner) []int {
	return EachLine(func(line string) int {
		num, _ := strconv.Atoi(line)
		return num
	})(scan)
}

func NumbersLine(scan *bufio.Scanner) [][]int {
	return EachLine(func(line string) []int {
		fields := strings.Fields(line)
		numbers := make([]int, len(fields))

		for idx, field := range fields {
			numbers[idx], _ = strconv.Atoi(field)
		}

		return numbers
	})(scan)
}

func NumbersColumn(columns int) Parser[[][]int] {
	return EachColumn(columns, func(line string) []int {
		fields := strings.Fields(line)
		numbers := make([]int, len(fields))

		for idx, field := range fields {
			numbers[idx], _ = strconv.Atoi(field)
		}

		return numbers
	})
}
