package y2025_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	aocparse "github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func TestSolve2025Day02(t *testing.T) {
	// aoc.SolveExamplesOnly(t)

	aoc.SolveExample(t, "example", 1, 1227775554, "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124")
	aoc.SolveExample(t, "example", 2, 4174379265, "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124")

	type idRange struct {
		From  int
		FromV string
		To    int
		ToV   string
	}

	parser := aocparse.CommaListEach(func(s string) idRange {
		idx := strings.IndexByte(s, '-')
		if idx == -1 {
			panic(fmt.Sprintf("invalid format: %q", s))
		}

		from, _ := strconv.Atoi(s[:idx])
		to, _ := strconv.Atoi(s[idx+1:])

		return idRange{
			From:  from,
			FromV: s[:idx],
			To:    to,
			ToV:   s[idx+1:],
		}
	})

	isValidId := func(id string) bool {
		if len(id)%2 != 0 {
			return true
		}

		halfIdx := len(id) / 2

		left, right := id[:halfIdx], id[halfIdx:]

		return left != right
	}

	aoc.Solve(t, "part 1", parser, func(ranges []idRange) int {
		sum := 0
		for _, idr := range ranges {
			for n := idr.From; n <= idr.To; n++ {
				if !isValidId(strconv.Itoa(n)) {
					sum += n
				}
			}
		}

		return sum
	})

	isValidComplexId := func(id string) bool {
		if len(id) == 1 {
			return true
		}

		idc := id + id
		at := strings.Index(idc[1:], id)
		at++

		if at < len(id) {
			return false
		}

		return true
	}

	aoc.Solve(t, "part 2", parser, func(ranges []idRange) int {
		sum := 0
		for _, idr := range ranges {
			for n := idr.From; n <= idr.To; n++ {
				if !isValidComplexId(strconv.Itoa(n)) {
					sum += n
				}
			}
		}

		return sum
	})
}
