package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func puzzleFilename(year, day int) string {
	return filepath.Join("data", fmt.Sprintf("%d_%02d.txt", year, day))
}

func puzzleDownload(year, day int) ([]byte, error) {
	r, err := newRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		return nil, fmt.Errorf("creating puzzle download request: %w", err)
	}

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

func puzzleDownloadAndSave(year, day int) ([]byte, error) {
	data, err := puzzleDownload(year, day)
	if err != nil {
		return nil, err
	}

	filename := puzzleFilename(year, day)

	if err = os.MkdirAll(filepath.Dir(filename), 0o755); err != nil {
		return nil, fmt.Errorf("cannot create puzzle directory: %w", err)
	}

	fp, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, fmt.Errorf("cannot create puzzle file: %w", err)
	}
	defer fp.Close()

	if _, err = fp.Write(data); err != nil {
		return nil, fmt.Errorf("cannot write puzzle file: %w", err)
	}

	return data, nil
}

func puzzleTestInfo(t *testing.T) (year, day, part int) {
	name := t.Name()
	name, after, _ := strings.Cut(name, "/")

	_, err := fmt.Sscanf(name, "TestSolve%dDay%02d", &year, &day)
	if err != nil {
		panic(err)
	}

	if after != "" {
		_, _ = fmt.Sscanf(after, "part_%d", &part)
	}

	return year, day, part
}
