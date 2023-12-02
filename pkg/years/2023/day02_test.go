package y2023

import (
	"bufio"
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay02(t *testing.T) {
	t.Parallel()

	type game struct {
		id   int
		sets [][3]int // 0 = red, 1 = blue, 2 = gree
	}

	// Game 1: 2 blue, 4 green; 7 blue, 1 red, 14 green; 5 blue, 13 green, 1 red; 1 red, 7 blue, 11 green
	parseGames := func(scanner *bufio.Scanner) []game {
		return aoc.ParseLines(scanner, func(line string) game {
			idx0 := strings.IndexRune(line, ' ')
			idx1 := strings.IndexRune(line, ':')
			gameId, _ := strconv.ParseInt(line[idx0+1:idx1], 10, 64)

			sets := strings.Split(line[idx1+1:], ";")
			gameSets := make([][3]int, 0, len(sets))

			for _, set := range sets {
				cubes := strings.Split(set, ",")
				var setCubes [3]int

				for _, cube := range cubes {
					split := strings.Split(strings.TrimSpace(cube), " ")
					count, _ := strconv.Atoi(split[0])

					switch split[1] {
					case "red":
						setCubes[0] = count
					case "green":
						setCubes[1] = count
					case "blue":
						setCubes[2] = count
					}
				}

				gameSets = append(gameSets, setCubes)
			}

			return game{
				id:   int(gameId),
				sets: gameSets,
			}
		})
	}

	part1 := func(games []game) int {
		sum := 0
		for _, g := range games {
			isPossible := true
			for _, s := range g.sets {
				if s[0] > 12 {
					isPossible = false
					break
				}

				if s[1] > 13 {
					isPossible = false
					break
				}

				if s[2] > 14 {
					isPossible = false
					break
				}
			}

			if isPossible {
				sum += g.id
			}
		}

		return sum
	}

	part2 := func(games []game) int {
		sum := 0

		for _, g := range games {
			var maxRed, maxGreen, maxBlue int

			for _, s := range g.sets {
				maxRed = max(maxRed, s[0])
				maxGreen = max(maxGreen, s[1])
				maxBlue = max(maxBlue, s[2])
			}

			power := maxRed * maxGreen * maxBlue
			sum += power
		}

		return sum
	}

	t.Run("example 1", func(t *testing.T) {
		input := aoc.InputScanner(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`)

		t.Log(part1(parseGames(input)))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(part1(parseGames(aoc.PuzzleScanner(2023, 2))))
	})

	t.Run("example 2", func(t *testing.T) {
		input := aoc.InputScanner(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`)

		t.Log(part2(parseGames(input)))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(part2(parseGames(aoc.PuzzleScanner(2023, 2))))
	})
}
