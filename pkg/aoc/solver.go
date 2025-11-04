package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	aoc_parse "github.com/cpl/advent-of-code/pkg/aoc-parse"
)

type exampleEntry[K comparable] struct {
	Name     string
	Data     string
	Part     int
	Expected K
}

var examples sync.Map

func SolveExample[K comparable](t *testing.T, name string, part int, expected K, data string) {
	tname := t.Name() + "/" + name
	if part != 0 {
		tname = fmt.Sprintf("%s/part_%d/%s", t.Name(), part, name)
	}

	fmt.Println("adding example", tname)
	examples.Store(tname, exampleEntry[K]{
		Name:     name,
		Data:     data,
		Part:     part,
		Expected: expected,
	})
}

func SolveExamplesOnly(t *testing.T) {
	t.Setenv("AOC_EXAMPLES", "only")
}

func SolveNoExamples(t *testing.T) {
	t.Setenv("AOC_EXAMPLES", "none")
}

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

func solveExamples[T any, K comparable](t *testing.T, parse aoc_parse.Parser[T], solve func(T) K) {
	t.Helper()

	tname := t.Name()
	examples.Range(func(xkey, xvalue any) bool {
		key := xkey.(string)
		if !strings.HasPrefix(key, tname) {
			return true
		}

		example := xvalue.(exampleEntry[K])

		scanner := bufio.NewScanner(strings.NewReader(example.Data))
		testname := "example"
		if example.Name != "example" {
			testname = "example " + example.Name
		}

		t.Run(testname, func(t *testing.T) {
			t.Helper()

			result := solveAny[T, K](t, scanner, parse, solve)
			if result != example.Expected {
				t.Errorf("\n got %v\nwant %v", result, example.Expected)
			}
		})

		return true
	})
}

func SolveAutoSubmit(t *testing.T) {
	t.Setenv("AOC_AUTOSUBMIT", "true")
}
