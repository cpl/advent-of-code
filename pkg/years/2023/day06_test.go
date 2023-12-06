package y2023

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay06(t *testing.T) {
	t.Parallel()

	parse := func(input string, part2 bool) ([]int, []int) {
		lines := strings.Split(input, "\n")

		if part2 {
			timesLine := strings.ReplaceAll(strings.TrimPrefix(lines[0], "Time:"), " ", "")
			distancesLine := strings.ReplaceAll(strings.TrimPrefix(lines[1], "Distance:"), " ", "")

			tValue, _ := strconv.Atoi(timesLine)
			dValue, _ := strconv.Atoi(distancesLine)

			return []int{tValue}, []int{dValue}
		}

		timesLine := strings.Fields(lines[0])
		distancesLine := strings.Fields(lines[1])

		times := make([]int, 0, len(timesLine)-1)
		distances := make([]int, 0, len(distancesLine)-1)

		for idx := 1; idx < len(timesLine); idx++ {
			vTime, _ := strconv.Atoi(timesLine[idx])
			vDistance, _ := strconv.Atoi(distancesLine[idx])

			times = append(times, vTime)
			distances = append(distances, vDistance)
		}

		return times, distances
	}

	raceOptions := func(t, d int) int {
		options := 0
		for power := 0; power < t; power++ {
			if power*(t-power) > d {
				options++
			}
		}

		return options
	}

	part1 := func(times, distances []int) int {
		m := 1

		for idx := range times {
			m = m * raceOptions(times[idx], distances[idx])
		}

		return m
	}

	t.Run("example 1", func(t *testing.T) {
		t.Log(part1(parse("Time:      7  15   30\nDistance:  9  40  200\n", false)))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(part1(parse(aoc.PuzzleString(2023, 6), false)))
	})

	t.Run("example 2", func(t *testing.T) {
		t.Log(part1(parse("Time:      7  15   30\nDistance:  9  40  200\n", true)))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(part1(parse(aoc.PuzzleString(2023, 6), true)))
	})
}
