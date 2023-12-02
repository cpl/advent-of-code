package aoc

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var httpClient = &http.Client{
	Transport: &http.Transport{},
	Timeout:   time.Second * 5,
}

func newRequest(method, u string, body io.Reader) (*http.Request, error) {
	cookie := os.Getenv("AOC_COOKIE")
	if cookie == "" {
		return nil, fmt.Errorf("cookie is empty")
	}

	runnerEmail := os.Getenv("AOC_EMAIL")
	if runnerEmail == "" {
		return nil, fmt.Errorf("runner email is empty")
	}

	r, _ := http.NewRequestWithContext(context.Background(), method, u, body)
	r.Header.Set("Cookie", "session="+cookie)
	r.Header.Set("User-Agent", "github.com/cpl/advent-of-code by "+runnerEmail)
	r.Header.Set("Accept", "text/plain")

	return r, nil
}
