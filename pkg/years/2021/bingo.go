package y2021

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Bingo struct {
	Boards  []BingoBoard
	Numbers []int
}

type pos struct {
	row, col int
}

type BingoBoard struct {
	Numbers [5][5]int
	Marked  [5][5]bool

	lookup map[int]pos
	rowSum [5]int
	colSum [5]int
	won    bool
}

func (bb *BingoBoard) IsWinner() bool {
	return bb.won
}

func (bb *BingoBoard) Mark(num int) bool {
	if p, ok := bb.lookup[num]; ok {
		bb.Marked[p.row][p.col] = true
		bb.rowSum[p.row]++
		bb.colSum[p.col]++

		if bb.rowSum[p.row] == 5 || bb.colSum[p.col] == 5 {
			bb.won = true
			return true
		}
	}

	return false
}

func (bb *BingoBoard) Score() int {
	sum := 0
	for row := range bb.Marked {
		for col, marked := range bb.Marked[row] {
			if !marked {
				sum += bb.Numbers[row][col]
			}
		}
	}

	return sum
}

func ParseBingo(input []byte) (*Bingo, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanLines)

	if !scanner.Scan() {
		return nil, scanner.Err()
	}

	numbersLine := scanner.Text()
	numbersStrs := strings.Split(numbersLine, ",")
	numbersInts := make([]int, len(numbersStrs))
	for idx, s := range numbersStrs {
		numbersInts[idx], _ = strconv.Atoi(s)
	}

	bingo := &Bingo{
		Numbers: numbersInts,
	}

	boardLines := make([]string, 0, 5)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else {
			boardLines = append(boardLines, line)
		}
		if len(boardLines) == 5 {
			bingo.Boards = append(bingo.Boards, ParseBingoBoard(boardLines))
			boardLines = boardLines[:0]
		}
	}

	return bingo, nil
}

func ParseBingoBoard(lines []string) BingoBoard {
	board := BingoBoard{
		lookup: make(map[int]pos),
	}

	for row, line := range lines {
		for col, s := range strings.Fields(line) {
			num, _ := strconv.Atoi(s)
			board.Numbers[row][col] = num
			board.lookup[num] = pos{row, col}
		}
	}
	return board
}

func PrintBoards(boards []BingoBoard) {
	for row := 0; row < 5; row++ {
		for _, board := range boards {
			for col, num := range board.Numbers[row] {
				sym := " "
				if board.Marked[row][col] {
					sym = "X"
				}

				fmt.Printf("[%2d %s]", num, sym)
			}
			fmt.Printf("  ")
		}
		fmt.Println()
	}
}
