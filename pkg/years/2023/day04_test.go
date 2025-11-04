package y2023

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	"github.com/cpl/advent-of-code/pkg/aoc-parse"
)

func TestSolveDay04(t *testing.T) {
	t.Parallel()

	type card struct {
		instances int
		winning   map[int]struct{}
		numbers   map[int]struct{}
	}

	parseCard := func(line string) *card {
		line = line[strings.IndexByte(line, ':')+2:]
		divider := strings.IndexByte(line, '|')
		winningLine := line[:divider-1]
		numbersLine := line[divider+1:]

		winning := strings.Fields(winningLine)
		numbers := strings.Fields(numbersLine)

		c := &card{
			instances: 1,
			winning:   make(map[int]struct{}, len(winning)),
			numbers:   make(map[int]struct{}, len(numbers)),
		}

		for _, v := range winning {
			num, _ := strconv.Atoi(v)
			c.winning[num] = struct{}{}
		}

		for _, v := range numbers {
			num, _ := strconv.Atoi(v)
			c.numbers[num] = struct{}{}
		}

		return c
	}

	score := func(c *card) int {
		score := 1
		for number := range c.numbers {
			if _, ok := c.winning[number]; ok {
				score = score << 1
			}
		}

		return score >> 1
	}

	matches := func(c *card) int {
		matches := 0
		for number := range c.numbers {
			if _, ok := c.winning[number]; ok {
				matches++
			}
		}

		return matches
	}

	part1 := func(cards []*card) int {
		total := 0
		for _, c := range cards {
			total += score(c)
		}

		return total
	}

	part2 := func(cards []*card) int {
		total := 0
		for idx, c := range cards {
			m := matches(c)

			for jdx := idx + 1; jdx < idx+1+m && jdx < len(cards); jdx++ {
				target := cards[jdx]
				target.instances += c.instances
			}

			total += c.instances
		}

		return total
	}

	t.Run("example 1", func(t *testing.T) {
		t.Log(part1(aoc_parse.EachLine(aoc.InputScanner(
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"),
			parseCard)))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(part1(aoc_parse.EachLine(aoc.PuzzleScanner(2023, 4), parseCard)))
	})

	t.Run("example 2", func(t *testing.T) {
		t.Log(part2(aoc_parse.EachLine(aoc.InputScanner(
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"),
			parseCard)))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(part2(aoc_parse.EachLine(aoc.PuzzleScanner(2023, 4), parseCard)))
	})
}
