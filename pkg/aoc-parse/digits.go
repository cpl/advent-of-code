package aoc_parse

import (
	"bufio"
	"fmt"
)

func DigitsLine(scan *bufio.Scanner) [][]uint8 {
	scan.Split(bufio.ScanLines)

	return EachLine(func(line string) []uint8 {
		digits := make([]uint8, 0, len(line))
		for _, ch := range line {
			if ch < '0' || ch > '9' {
				panic(fmt.Sprintf("invalid char '%s'", string(ch)))
			}

			digits = append(digits, uint8(ch-'0'))
		}

		return digits
	})(scan)
}
