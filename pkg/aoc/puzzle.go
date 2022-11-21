package aoc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var httpClient = &http.Client{
	Transport: &http.Transport{},
	Timeout:   time.Second * 5,
}

func puzzleFilename(year, day int) string {
	return filepath.Join("data", fmt.Sprintf("%d_%02d.txt", year, day))
}

func puzzleDownload(cookie string, year, day int) ([]byte, error) {
	if cookie == "" {
		return nil, fmt.Errorf("cookie is empty")
	}

	r, _ := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	r.Header.Set("Cookie", "session="+cookie)
	r.Header.Set("User-Agent", "Santa's Little Helper")
	r.Header.Set("Accept", "text/plain")

	resp, err := httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("cannot download puzzle: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("cannot download puzzle: %s | %s", resp.Status, string(body))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot download puzzle: %w", err)
	}

	return data, nil
}

func puzzleDownloadAndSave(cookie string, year, day int) ([]byte, error) {
	data, err := puzzleDownload(cookie, year, day)
	if err != nil {
		return nil, err
	}

	filename := puzzleFilename(year, day)

	if err = os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return nil, fmt.Errorf("cannot create puzzle directory: %w", err)
	}

	fp, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("cannot create puzzle file: %w", err)
	}
	defer fp.Close()

	if _, err = fp.Write(data); err != nil {
		return nil, fmt.Errorf("cannot write puzzle file: %w", err)
	}

	return data, nil
}

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

func PuzzleReader(year, day int) (io.Reader, error) {
	data, err := Puzzle(year, day)
	if err != nil {
		return nil, err
	}

	return io.NopCloser(bytes.NewReader(data)), nil
}

func PuzzleLineScanner(year, day int) (*bufio.Scanner, error) {
	r, err := PuzzleReader(year, day)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	return scanner, nil
}

func PuzzleLinesFloats(year, day int) ([]float64, error) {
	scanner, err := PuzzleLineScanner(year, day)
	if err != nil {
		return nil, err
	}

	var lines []float64

	for scanner.Scan() {
		var value float64
		line := scanner.Text()

		value, err = strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse line %q : %w", line, err)
		}

		lines = append(lines, value)
	}

	return lines, nil
}

func PuzzleLinesInts(year, day int) ([]int, error) {
	scanner, err := PuzzleLineScanner(year, day)
	if err != nil {
		return nil, err
	}

	var lines []int

	for scanner.Scan() {
		var value int
		line := scanner.Text()

		value, err = strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("cannot parse line %q : %w", line, err)
		}

		lines = append(lines, value)
	}

	return lines, nil
}

func PuzzleLinesStrings(year, day int) ([]string, error) {
	scanner, err := PuzzleLineScanner(year, day)
	if err != nil {
		return nil, err
	}

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
