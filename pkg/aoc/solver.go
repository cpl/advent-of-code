package aoc

import (
	"bufio"
	"os"
	"strings"
	"testing"
	"time"

	aoc_parse "github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func Solve[T any, K comparable](t *testing.T, name string, parse aoc_parse.Parser[T], solve func(T) K) {
	t.Helper()

	t.Run(name, func(t *testing.T) {
		t.Helper()

		examplesMode := os.Getenv("AOC_EXAMPLES")
		if examplesMode != "none" {
			solveExamples(t, parse, solve)
		}
		if examplesMode == "only" {
			t.Skip("examples only")
		}

		year, day, part := puzzleTestInfo(t)
		scanner := PuzzleScanner(year, day)
		solution := solveAny(t, scanner, parse, solve)

		if part != 0 && os.Getenv("AOC_AUTOSUBMIT") == "true" {
			response, err := Submit(year, day, part, solution)
			if err != nil {
				t.Errorf("failed to submit puzzle: %s", err)
			}

			t.Log("📣", strings.ToUpper(response.String()), response.Emoji())
		}
	})
}

func solveAny[T, K any](t *testing.T, scanner *bufio.Scanner, parse aoc_parse.Parser[T], solve func(T) K) K {
	t.Helper()

	start := time.Now()
	solution := solve(parse(scanner))
	duration := time.Since(start)

	t.Log("✅", solution)
	t.Log("🕐", duration)

	return solution
}

func SolveAutoSubmit(t *testing.T) {
	t.Setenv("AOC_AUTOSUBMIT", "true")
}
