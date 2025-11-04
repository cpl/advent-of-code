package aoc_parse

import (
	"bufio"
	"strconv"
	"strings"
)

func NumberLine(scan *bufio.Scanner) []int {
	return EachLine(scan, func(line string) int {
		num, _ := strconv.Atoi(line)
		return num
	})
}

func NumbersLine(scan *bufio.Scanner) [][]int {
	return EachLine(scan, func(line string) []int {
		fields := strings.Fields(line)
		numbers := make([]int, len(fields))

		for idx, field := range fields {
			numbers[idx], _ = strconv.Atoi(field)
		}

		return numbers
	})
}

func NumbersColumn(scan *bufio.Scanner, columns int) [][]int {
	return EachColumn(scan, columns, func(line string) []int {
		fields := strings.Fields(line)
		numbers := make([]int, len(fields))

		for idx, field := range fields {
			numbers[idx], _ = strconv.Atoi(field)
		}

		return numbers
	})
}
