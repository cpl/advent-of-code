package aoc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var MetaAuthCookie = os.Getenv("AUTH_COOKIE")

func init() {
	if MetaAuthCookie == "" {
		data, _ := ioutil.ReadFile(".cookie")
		if data != nil {
			MetaAuthCookie = strings.TrimSpace(string(data))
		}
	}
}

func MetaInputURL(year, day int) string {
	if year == 0 {
		year = time.Now().Year()
	}
	if day == 0 {
		day = time.Now().Day()
	}

	return fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
}

func MetaGetInput(year, day int) ([]byte, error) {
	filename := fmt.Sprintf("./data/year%d/day%02d/input.txt", year, day)
	data, err := ioutil.ReadFile(filename)
	if err == nil {
		return data, nil
	}

	req, err := http.NewRequest(http.MethodGet, MetaInputURL(year, day), nil)
	if err != nil {
		return nil, fmt.Errorf("creating new request, %w", err)
	}

	req.AddCookie(&http.Cookie{
		Name:     "session",
		Value:    MetaAuthCookie,
		Path:     "/",
		Domain:   ".adventofcode.com",
		Secure:   true,
		HttpOnly: true,
	})

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("doing request, %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		dump, _ := httputil.DumpResponse(res, true)
		fmt.Println(string(dump))

		return nil, fmt.Errorf("bad http status code, expected 200, got %d", res.StatusCode)
	}

	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body, %w", err)
	}

	_ = os.MkdirAll(filepath.Dir(filename), 0755)
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return data, err
	}

	return data, nil
}
