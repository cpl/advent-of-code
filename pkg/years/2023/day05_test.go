package y2023

import (
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay05(t *testing.T) {
	t.Parallel()

	type almanacMap struct {
		srcStart int64
		dstStart int64
		length   int64
	}

	type almanac struct {
		seeds []int64
		maps  map[string][]almanacMap
	}

	match := func(alm *almanac, typ string, v int64) int64 {
		maps := alm.maps[typ]

		for _, m := range maps {
			if v >= m.srcStart && v < m.srcStart+m.length {
				dst := v - m.srcStart + m.dstStart
				return dst
			}
		}

		return v
	}

	matchSeedToLocation := func(seed int64, alm *almanac) int64 {
		return match(alm, "humidity-to-location",
			match(alm, "temperature-to-humidity",
				match(alm, "light-to-temperature",
					match(alm, "water-to-light",
						match(alm, "fertilizer-to-water",
							match(alm, "soil-to-fertilizer",
								match(alm, "seed-to-soil", seed)))))))
	}

	parse := func(input string) *almanac {
		lines := strings.Split(input, "\n")

		lineSeeds := strings.Fields(strings.TrimPrefix(lines[0], "seeds: "))
		seeds := make([]int64, len(lineSeeds))
		for idx, s := range lineSeeds {
			seeds[idx], _ = strconv.ParseInt(s, 10, 64)
		}

		maps := make(map[string][]almanacMap)
		typ := ""
		typList := make([]almanacMap, 0)

		for _, line := range lines[2:] {
			if line == "" {
				maps[typ] = typList
				typList = make([]almanacMap, 0)
				typ = ""
				continue
			}

			if typ == "" {
				typ = strings.TrimSuffix(line, " map:")
				continue
			}

			fields := strings.Fields(line)
			srcStart, _ := strconv.ParseInt(fields[1], 10, 64) // SOURCE IS SECOND!
			dstStart, _ := strconv.ParseInt(fields[0], 10, 64) // DEST IS FIRST! FFS!
			length, _ := strconv.ParseInt(fields[2], 10, 64)

			typList = append(typList, almanacMap{
				srcStart: srcStart,
				dstStart: dstStart,
				length:   length,
			})
		}

		return &almanac{
			seeds: seeds,
			maps:  maps,
		}
	}

	part1 := func(alm *almanac) int64 {
		smallest := int64(1<<63 - 1)

		for _, seed := range alm.seeds {
			smallest = min(smallest, matchSeedToLocation(seed, alm))
		}

		return smallest
	}

	seedRangeReduce := func(seedStart, seedLength int64, alm *almanac) []int64 {
		maps := alm.maps["seed-to-soil"]
		ranges := make([]int64, 0)

		for _, m := range maps {
			// if seed start in map range
			if seedStart >= m.srcStart && seedStart < m.srcStart+m.length {
				// if seed end in map range (fully contained by 1 soil range)
				if seedStart+seedLength >= m.srcStart && seedStart+seedLength < m.srcStart+m.length {
					// reduce to 1 seed
					return []int64{seedStart, 1}
				}

				// part 2 "brute force" finished before I could write out reduce logic :)
			}
		}

		return ranges
	}
	_ = seedRangeReduce

	part2 := func(alm *almanac) int64 {
		ans := make([]int64, len(alm.seeds)/2)
		wg := sync.WaitGroup{}
		wg.Add(len(alm.seeds) / 2)

		for idx := 0; idx < len(alm.seeds); idx += 2 {
			go func(idx int) {
				start := alm.seeds[idx]
				l := alm.seeds[idx+1]

				m := int64(1<<63 - 1)
				for seed := start; seed < start+l; seed++ {
					m = min(matchSeedToLocation(seed, alm), m)
				}

				wg.Done()
				ans[idx/2] = m
			}(idx)
		}

		wg.Wait()

		smallest := int64(1<<63 - 1)
		for _, v := range ans {
			smallest = min(smallest, v)
		}

		return smallest
	}

	t.Run("example 1", func(t *testing.T) {
		t.Log(part1(parse("seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4")))
	})

	t.Run("part 1", func(t *testing.T) {
		t.Log(part1(parse(aoc.PuzzleString(2023, 5))))
	})

	t.Run("example 2", func(t *testing.T) {
		t.Log(part2(parse("seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4")))
	})

	t.Run("part 2", func(t *testing.T) {
		t.Log(part2(parse(aoc.PuzzleString(2023, 5))))
	})
}
