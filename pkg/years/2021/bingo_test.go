package y2021

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestParseBingoBoard(t *testing.T) {
	t.Parallel()

	t.Run("example", func(t *testing.T) {
		input := `22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19`

		board := ParseBingoBoard(strings.Split(input, "\n"))
		want := BingoBoard{
			Numbers: [5][5]int{
				{22, 13, 17, 11, 0},
				{8, 2, 23, 4, 24},
				{21, 9, 14, 16, 7},
				{6, 10, 3, 18, 5},
				{1, 12, 20, 15, 19},
			},
		}

		if !reflect.DeepEqual(board, want) {
			t.Errorf("got %v, want %v", board, want)
		}
	})
}

func TestParseBingo(t *testing.T) {
	t.Parallel()

	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

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

	t.Run("example", func(t *testing.T) {
		bingo, err := ParseBingo([]byte(input))
		if err != nil {
			t.Fatal(err)
		}

		wantNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
		if !reflect.DeepEqual(bingo.Numbers, wantNumbers) {
			t.Errorf("got %v, want %v", bingo.Numbers, wantNumbers)
		}

		wantBoards := []BingoBoard{
			{Numbers: [5][5]int{
				{22, 13, 17, 11, 0},
				{8, 2, 23, 4, 24},
				{21, 9, 14, 16, 7},
				{6, 10, 3, 18, 5},
				{1, 12, 20, 15, 19},
			}},
			{Numbers: [5][5]int{
				{3, 15, 0, 2, 22},
				{9, 18, 13, 17, 5},
				{19, 8, 7, 25, 23},
				{20, 11, 10, 24, 4},
				{14, 21, 16, 12, 6},
			}},
			{Numbers: [5][5]int{
				{14, 21, 17, 24, 4},
				{10, 16, 15, 9, 19},
				{18, 8, 23, 26, 20},
				{22, 11, 13, 6, 5},
				{2, 0, 12, 3, 7},
			}},
		}

		if len(bingo.Boards) != len(wantBoards) {
			t.Errorf("got %v, want %v", bingo.Boards, wantBoards)
		}

		for idx, board := range wantBoards {
			if !reflect.DeepEqual(bingo.Boards[idx], board) {
				t.Errorf("got %v, want %v", bingo.Boards[idx], board)
			}
		}
	})
}

func testSetupExampleBing(t *testing.T) *Bingo {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

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

	bingo, err := ParseBingo([]byte(input))
	if err != nil {
		t.Fatal(err)
	}

	return bingo
}

func TestPrintBoards(t *testing.T) {
	t.Parallel()

	bingo := testSetupExampleBing(t)

	PrintBoards(bingo.Boards)

	for _, num := range bingo.Numbers {
		for idx := range bingo.Boards {
			_ = bingo.Boards[idx].Mark(num)
		}
	}

	fmt.Println()
	PrintBoards(bingo.Boards)
}

func TestBingoBoard_Mark_and_Score(t *testing.T) {
	t.Parallel()
	bingo := testSetupExampleBing(t)

	for _, num := range bingo.Numbers {
		for idx := range bingo.Boards {
			win := bingo.Boards[idx].Mark(num)
			if win {
				if num != 24 {
					t.Fatalf("expected 24, got %d", num)
				}
				if idx != 2 {
					t.Fatalf("expected 2, got %d", idx)
				}
				if score := bingo.Boards[idx].Score(); score != 188 {
					t.Fatalf("expected 188, got %d", score)
				}

				return
			}
		}
	}
}
