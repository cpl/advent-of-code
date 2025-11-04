package aoc

import (
	"os"
	"strings"
	"testing"
	"time"

	aoc_parse "github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func Solve[T, K any](t *testing.T, name string, parse aoc_parse.Parser[T], solve func(T) K) {
	t.Helper()

	t.Run(name, func(t *testing.T) {
		t.Helper()

		year, day, part := puzzleTestInfo(t)

		start := time.Now()
		solution := solve(parse(PuzzleScanner(year, day)))
		duration := time.Since(start)

		t.Log("✅", solution)
		t.Log("🕐", duration)

		if part != 0 && os.Getenv("AOC_AUTOSUBMIT") == "true" {
			response, err := Submit(year, day, part, solution)
			if err != nil {
				t.Errorf("failed to submit puzzle: %s", err)
			}

			t.Log("📣", strings.ToUpper(response.String()), response.Emoji())
		}
	})
}

func SolveAutoSubmit(t *testing.T) {
	t.Setenv("AOC_AUTOSUBMIT", "true")
}
