package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
