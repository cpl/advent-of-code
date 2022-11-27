package aoc

import (
	"bufio"
	"strings"
)

func InputScanner(input string) *bufio.Scanner {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	return scanner
}
