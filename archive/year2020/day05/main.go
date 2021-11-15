package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"

	"../../utils"
)

const (
	AirplaneRows = 128
	AirplaneCols = 8
)

func binarySearchChar(char rune, min, max int) (int, int) {
	switch char {
	case 'F', 'L': // lower
		return min, (min + max) / 2
	case 'B', 'R': // upper
		return (min+max)/2 + 1, max
	default:
		return min, max
	}
}

func binarySearch(str string, min, max int) int {
	for _, char := range str {
		min, max = binarySearchChar(char, min, max)
	}

	return min
}

func extractSeat(boardingPass string) (int, int) {
	rowString := boardingPass[:7]
	colString := boardingPass[7:]

	return binarySearch(rowString, 0, AirplaneRows-1), binarySearch(colString, 0, AirplaneCols-1)
}

func seatID(row, col int) int {
	return row*8 + col
}

func main() {
	data, err := utils.GetInput(2020, 05)
	utils.CheckErr(err)

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	var seats [AirplaneRows][AirplaneCols]byte

	maxSeatID := 0
	for scanner.Scan() {
		line := scanner.Text()
		row, col := extractSeat(line)

		sid := seatID(row, col)
		if sid > maxSeatID {
			maxSeatID = sid
		}

		seats[row][col] = 1
	}

	log.Println("max seat ID", maxSeatID)

	// solve part2 visually
	// ###. ####  73
	// my seat :)
	for rowNum, row := range seats {
		for idx, seat := range row {
			if idx == 4 {
				fmt.Print(" ")
			}

			if seat == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Printf(" %3d", rowNum)

		fmt.Println()
	}
}
