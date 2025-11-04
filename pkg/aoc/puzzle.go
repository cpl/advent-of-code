package aoc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Puzzle(year, day int) ([]byte, error) {
	filename := puzzleFilename(year, day)

	var data []byte

	fp, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		if os.IsNotExist(err) {
			return puzzleDownloadAndSave(year, day)
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

func PuzzleString(year, day int) string {
	data, err := Puzzle(year, day)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(data))
}

func PuzzleStringSliceNewline(year, day int) []string {
	s := PuzzleString(year, day)

	return strings.Split(s, "\n")
}

func PuzzleStringRaw(year, day int) string {
	data, err := Puzzle(year, day)
	if err != nil {
		panic(err)
	}

	return string(data)
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

func PuzzleScannerTest(t *testing.T) *bufio.Scanner {
	year, day, _ := puzzleTestInfo(t)
	return PuzzleScanner(year, day)
}
