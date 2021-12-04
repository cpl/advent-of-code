package y2021

import (
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay04(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 4)
	if err != nil {
		t.Fatal(err)
	}

	bingo, err := ParseBingo(input)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("part_1", func(t *testing.T) {
		for _, number := range bingo.Numbers {
			for idx := range bingo.Boards {
				win := bingo.Boards[idx].Mark(number)
				if win {
					score := bingo.Boards[idx].Score()
					t.Logf("board=%d score=%d number=%d solution=%d", idx+1, score, number, score*number)
					return
				}
			}
		}
	})
	t.Run("part_2", func(t *testing.T) {
		lastWinBoardIdx := 0
		lastWinNumberIdx := 0

		for idx := range bingo.Boards {
			for numberIdx, number := range bingo.Numbers {
				if bingo.Boards[idx].IsWinner() {
					break
				}

				win := bingo.Boards[idx].Mark(number)
				if win {
					if numberIdx >= lastWinNumberIdx {
						lastWinNumberIdx = numberIdx
						lastWinBoardIdx = idx
					}

					break
				}
			}
		}

		score := bingo.Boards[lastWinBoardIdx].Score()
		t.Logf("board=%d number=%d score=%d solution=%d",
			lastWinBoardIdx+1, bingo.Numbers[lastWinNumberIdx], score, score*bingo.Numbers[lastWinNumberIdx])
	})
	t.Run("part_2_example", func(t *testing.T) {
		example := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

		bingo, err := ParseBingo([]byte(example))
		if err != nil {
			t.Fatal(err)
		}

		lastWinBoardIdx := 0
		lastWinNumberIdx := 0

		for idx := range bingo.Boards {
			for numberIdx, number := range bingo.Numbers {
				if bingo.Boards[idx].IsWinner() {
					break
				}

				win := bingo.Boards[idx].Mark(number)
				if win {
					if numberIdx >= lastWinNumberIdx {
						lastWinNumberIdx = numberIdx
						lastWinBoardIdx = idx
					}

					break
				}
			}
		}

		score := bingo.Boards[lastWinBoardIdx].Score()
		t.Logf("board=%d number=%d score=%d solution=%d",
			lastWinBoardIdx+1, bingo.Numbers[lastWinNumberIdx], score, score*bingo.Numbers[lastWinNumberIdx])
	})
}
