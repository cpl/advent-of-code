package aoc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func Puzzle(year, day int) ([]byte, error) {
	filename := puzzleFilename(year, day)

	var data []byte

	fp, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		if os.IsNotExist(err) {
			return puzzleDownloadAndSave(os.Getenv("AOC_COOKIE"), year, day)
		}

		return nil, fmt.Errorf("cannot open file %q: %w", filename, err)
	}
	defer fp.Close()

	data, err = io.ReadAll(fp)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %q: %w", filename, err)
	}

	return data, nil
}

func PuzzleReader(year, day int) *bufio.Reader {
	data, err := Puzzle(year, day)
	if err != nil {
		panic(err)
	}

	return bufio.NewReader(bytes.NewReader(data))
}

func PuzzleScanner(year, day int) *bufio.Scanner {
	r := PuzzleReader(year, day)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	return scanner
}
