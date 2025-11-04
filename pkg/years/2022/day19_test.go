package y2022

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	"github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func TestSolveDay19(t *testing.T) {
	type recipe map[string]int
	type blueprint struct {
		id      int
		recipes map[string]recipe
	}

	parseRecipe := func(s string) recipe {
		r := make(recipe)
		for {
			var n int
			var name string

			_, _ = fmt.Sscanf(s, "%d %s", &n, &name)
			r[name] = n

			idx := strings.Index(s, "and ")
			if idx == -1 {
				break
			}

			s = s[idx+4:]
		}

		return r
	}

	parse := func(line string) blueprint {
		bp := blueprint{
			recipes: make(map[string]recipe),
		}

		_, _ = fmt.Sscanf(line, "Blueprint %d:", &bp.id)
		line = line[strings.Index(line, ":")+2:]

		for _, r := range strings.Split(line, ".") {
			r = strings.TrimSpace(r)
			if r == "" {
				continue
			}

			var rType string
			_, _ = fmt.Sscanf(r, "Each %s robot", &rType)
			bp.recipes[rType] = parseRecipe(r[strings.Index(r, "costs ")+6:])
		}

		return bp
	}

	blueprintGeodeProduction := func(bp blueprint) int {
		geodeRecipe := bp.recipes["geode"]
		geodeCosts := make(map[string]int)

		for r, need := range geodeRecipe {
			if r == "ore" {
				geodeCosts[r] = need
				continue
			}

			_ = bp.recipes[r]
		}

		return -1
	}

	_ = blueprintGeodeProduction

	part1 := func(blueprints []blueprint) int {
		return -1
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.\nBlueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian."
		t.Log(part1(aoc_parse.EachLine(aoc.InputScanner(input), parse)))
	})
}
