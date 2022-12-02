package y2022

import (
	"fmt"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay02(t *testing.T) {
	type rockPaperScissorsRound struct {
		ElfPlay     string
		CounterPlay string
	}

	parse := func(line string) rockPaperScissorsRound {
		var r rockPaperScissorsRound
		_, _ = fmt.Sscanf(line, "%s %s", &r.ElfPlay, &r.CounterPlay)
		return r
	}

	part1 := func(rounds []rockPaperScissorsRound) int {
		score := 0
		for _, round := range rounds {
			switch round.CounterPlay {
			case "X": // rock
				score += 1
				switch round.ElfPlay {
				case "A": // rock
					score += 3
				case "B": // paper
				case "C": // scissors
					score += 6
				}
			case "Y": // paper
				score += 2
				switch round.ElfPlay {
				case "A": // rock
					score += 6
				case "B": // paper
					score += 3
				case "C": // scissors
				}
			case "Z": // scissors
				score += 3
				switch round.ElfPlay {
				case "A": // rock
				case "B": // paper
					score += 6
				case "C": // scissors
					score += 3
				}
			}
		}

		return score
	}

	part2 := func(rounds []rockPaperScissorsRound) int {
		counterPlays := map[string]map[string]string{
			"A": { // rock
				"X": "C",
				"Y": "A",
				"Z": "B",
			},
			"B": { // paper
				"X": "A",
				"Y": "B",
				"Z": "C",
			},
			"C": { // scissors
				"X": "B",
				"Y": "C",
				"Z": "A",
			},
		}

		score := 0
		for _, round := range rounds {
			elfPlay := round.ElfPlay
			pick := counterPlays[elfPlay][round.CounterPlay]
			switch pick {
			case "A": // rock
				score += 1
			case "B": // paper
				score += 2
			case "C": // scissors
				score += 3
			}

			switch round.CounterPlay {
			case "X": // lose
			case "Y": // draw
				score += 3
			case "Z": // win
				score += 6
			}
		}

		return score
	}

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(aoc.ParseLines(aoc.PuzzleScanner(2022, 2), parse)))
	})

	t.Run("Examples 2", func(t *testing.T) {
		t.Log(part2(aoc.ParseLines(aoc.InputScanner("A Y\nB X\nC Z"), parse)))
	})

	t.Run("Part 2", func(t *testing.T) {
		t.Log(part2(aoc.ParseLines(aoc.PuzzleScanner(2022, 2), parse)))
	})
}
