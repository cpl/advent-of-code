package aoc

import (
	"bufio"
	"bytes"
	"strings"
)

func InputScanner(input string) *bufio.Scanner {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	return scanner
}

func InputScannerB(input []byte) *bufio.Scanner {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanLines)

	return scanner
}
