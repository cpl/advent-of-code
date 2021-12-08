package y2021

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay08(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 8)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("part_1", func(t *testing.T) {
		scanner := bufio.NewScanner(bytes.NewReader(input))
		scanner.Split(bufio.ScanLines)

		total := 0
		for scanner.Scan() {
			line := scanner.Text()
			idx := strings.IndexRune(line, '|')
			line = line[idx+1:]
			segments := strings.Split(line, " ")

			for _, segment := range segments {
				switch len(segment) {
				case 2, 3, 4, 7:
					total++
				}
			}
		}

		t.Logf("solution=%d", total)
	})
	t.Run("part_2", func(t *testing.T) {
		scanner := bufio.NewScanner(bytes.NewReader(input))
		scanner.Split(bufio.ScanLines)

		total := 0
		for scanner.Scan() {
			line := scanner.Text()
			splitIdx := strings.IndexRune(line, '|')
			segments := Parse7Segments(strings.Split(line[:splitIdx-1], " "))
			mapping := segments.SolveMapping()
			numberStr := Parse7SegmentOutput(strings.Split(line[splitIdx+2:], " "), mapping)
			numberInt, _ := strconv.Atoi(numberStr)
			total += numberInt
		}

		t.Logf("solution=%d", total)
	})
}
