package aoc

import (
	"bufio"
	"fmt"
	"strings"
	"sync"
	"testing"

	aoc_parse "github.com/cpl/advent-of-code/pkg/aoc-parse"
)

var exampleMode bool

func IsExample() bool {
	return exampleMode
}

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

	// fmt.Println("adding example", tname)
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

func solveExamples[T any, K comparable](t *testing.T, parse aoc_parse.Parser[T], solve func(T) K) {
	exampleMode = true
	defer func() {
		exampleMode = false
	}()

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
