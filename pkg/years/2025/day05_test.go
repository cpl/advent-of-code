package y2025_test

import (
	"bufio"
	"sort"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	"go.sdls.io/pkg"
)

func TestSolve2025Day05(t *testing.T) {
	aoc.SolveExample(t, "example", 1, 3, "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32")
	aoc.SolveExample(t, "example", 2, "14", "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32")
	aoc.SolveExample(t, "1", 2, "100", "1-100\n1-100\n2-99\n5-5\n5-99\n2-3\n")
	aoc.SolveExample(t, "2", 2, "200", "1-100\n100-200")
	aoc.SolveExample(t, "3", 2, "200", "1-100\n2-180\n179-200")
	aoc.SolveExample(t, "4", 2, "222", "1-100\n200-300\n150-170")
	aoc.SolveExample(t, "5", 2, "100", "2-5\n10-19\n80-90\n12-82\n1-100")

	type freshRange struct {
		from pkg.UInt128
		to   pkg.UInt128
	}

	type input struct {
		freshRanges []freshRange
		minimum     pkg.UInt128
		ingredients []pkg.UInt128
	}

	parse := func(scanner *bufio.Scanner) input {
		scanner.Split(bufio.ScanLines)

		in := input{
			freshRanges: make([]freshRange, 0, 1024),
			ingredients: make([]pkg.UInt128, 0, 1024),
		}

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}

			from, to, found := strings.Cut(line, "-")
			if !found {
				panic("invalid input")
			}

			fromU128, err := pkg.UInt128FromString(from)
			if err != nil {
				panic(err)
			}

			toU128, err := pkg.UInt128FromString(to)
			if err != nil {
				panic(err)
			}

			if fromU128.Compare(toU128) > 0 {
				panic("invalid input")
			}

			if in.minimum.Compare(fromU128) > 0 {
				in.minimum = fromU128
			}

			in.freshRanges = append(in.freshRanges, freshRange{
				from: fromU128,
				to:   toU128,
			})
		}

		for scanner.Scan() {
			line := scanner.Text()
			u128, err := pkg.UInt128FromString(line)
			if err != nil {
				panic(err)
			}

			in.ingredients = append(in.ingredients, u128)
		}

		return in
	}

	isFresh := func(v pkg.UInt128, ranges []freshRange) bool {
		for _, r := range ranges {
			if v.Compare(r.from) >= 0 && v.Compare(r.to) <= 0 {
				return true
			}
		}

		return false
	}

	aoc.Solve(t, "part 1", parse, func(in input) int {
		count := 0
		for _, ingredient := range in.ingredients {
			if isFresh(ingredient, in.freshRanges) {
				count++
			}
		}

		return count
	})

	type freshIntRange struct {
		from uint64
		to   uint64
	}

	rangesToIntRanges := func(in input) []freshIntRange {
		intRanges := make([]freshIntRange, 0, len(in.freshRanges))

		for _, r := range in.freshRanges {
			_, from := r.from.Sub(in.minimum).UInt64()
			_, to := r.to.Sub(in.minimum).UInt64()
			ri := freshIntRange{
				from: from,
				to:   to,
			}

			intRanges = append(intRanges, ri)
		}

		sort.Slice(intRanges, func(i, j int) bool {
			return intRanges[i].from < intRanges[j].from
		})

		return intRanges
	}

	rangesOverlap := func(ranges []freshIntRange) []freshIntRange {
		out := make([]freshIntRange, 0, len(ranges))
		for idx := range ranges {
			from := ranges[idx].from
			to := ranges[idx].to

			if len(out) > 0 && out[len(out)-1].to >= to {
				continue
			}

			for jdx := idx + 1; jdx < len(ranges); jdx++ {
				if ranges[jdx].from <= to {
					to = max(to, ranges[jdx].to)
				}
			}

			out = append(out, freshIntRange{
				from: from,
				to:   to,
			})
		}

		return out
	}

	aoc.Solve(t, "part 2", parse, func(in input) string {
		ranges := rangesToIntRanges(in)
		merged := rangesOverlap(ranges)
		size := pkg.UInt128{}
		for _, rng := range merged {
			size = size.Add64(rng.to - rng.from + 1)
		}

		return size.String()
	})
}
