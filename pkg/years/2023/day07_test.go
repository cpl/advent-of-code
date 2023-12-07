package y2023

import (
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay07(t *testing.T) {
	t.Parallel()

	replacer := strings.NewReplacer(
		"T", ":",
		"J", ";",
		"Q", "<",
		"K", "=",
		"A", ">",
	)

	type round struct {
		hand      string
		rank      int64 // 0=high, 1=1p, 2=2p, 3=3k, 4=full, 5=4k, 6=5k
		rankJoker int64
		bid       int64
	}

	parse := func(line string) *round {
		idx := strings.IndexByte(line, ' ')
		s := line[:idx]
		bid, _ := strconv.ParseInt(line[idx+1:], 10, 64)

		m := make(map[byte]int)
		for _, c := range []byte(s) {
			m[c]++
		}

		values := make([]int, 0, len(m))
		for _, v := range m {
			values = append(values, v)
		}

		rank := int64(0)
		switch len(m) {
		case 5:
			rank = 0 // high card
		case 4:
			rank = 1 // one pair
		case 1:
			rank = 6 // five of a kind
		case 2:
			if values[0] == 3 || values[1] == 3 {
				rank = 4 // full house
			} else {
				rank = 5 // four of a kind
			}
		case 3:
			if values[0] == 2 || values[1] == 2 || values[2] == 2 {
				rank = 2 // two pair
			} else {
				rank = 3 // three of a kind
			}
		}

		jokerRank := int64(0)
		jokerCount := m['J']
		if jokerCount == 0 || rank == 6 {
			return &round{
				hand:      replacer.Replace(s),
				rank:      rank,
				rankJoker: rank,
				bid:       bid,
			}
		}

		switch rank {
		case 5, 4: // four of a kind, full house
			jokerRank = 6 // five of a kind
		case 3: // three of a kind
			jokerRank = 5 // four of a kind
		case 2:
			if jokerCount == 2 {
				jokerRank = 5 // four of a kind
			} else {
				jokerRank = 4 // full house
			}
		case 1:
			jokerRank = 3 // three of a kind
		case 0:
			jokerRank = 1 // one pair
		}

		return &round{
			hand:      replacer.Replace(s),
			rank:      rank,
			rankJoker: jokerRank,
			bid:       bid,
		}
	}

	part1 := func(rounds []*round) int64 {
		slices.SortFunc(rounds, func(a, b *round) int {
			if a.rank == b.rank {
				return strings.Compare(a.hand, b.hand)
			}

			return int(a.rank - b.rank)
		})

		total := int64(0)

		for idx, r := range rounds {
			total += r.bid * (int64(idx) + 1)
		}

		return total
	}

	part2 := func(rounds []*round) int64 {
		slices.SortFunc(rounds, func(a, b *round) int {
			if a.rankJoker == b.rankJoker {
				a.hand = strings.ReplaceAll(a.hand, ";", "0")
				b.hand = strings.ReplaceAll(b.hand, ";", "0")

				return strings.Compare(a.hand, b.hand)
			}

			return int(a.rankJoker - b.rankJoker)
		})

		total := int64(0)

		for idx, r := range rounds {
			total += r.bid * (int64(idx) + 1)
		}

		return total
	}

	t.Run("example 1", func(t *testing.T) {
		t.Log(part1(aoc.ParseLines(aoc.InputScanner("32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"), parse)))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(part1(aoc.ParseLines(aoc.PuzzleScanner(2023, 7), parse)))
	})

	t.Run("example 2", func(t *testing.T) {
		t.Log(part2(aoc.ParseLines(aoc.InputScanner("32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"), parse)))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(part2(aoc.ParseLines(aoc.PuzzleScanner(2023, 7), parse)))
	})
}
