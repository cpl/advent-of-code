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

var (
	cookie      = os.Getenv("AOC_COOKIE")
	runnerEmail = os.Getenv("AOC_EMAIL")
)

func newRequest(method, u string, body io.Reader) (*http.Request, error) {
	if cookie == "" {
		return nil, fmt.Errorf("cookie is empty")
	}

	if runnerEmail == "" {
		return nil, fmt.Errorf("runner email is empty")
	}

	ctx, cc := context.WithTimeout(context.Background(), time.Second*10)
	defer cc()

	r, _ := http.NewRequestWithContext(ctx, method, u, body)
	r.Header.Set("Cookie", "session="+cookie)
	r.Header.Set("User-Agent", "github.com/cpl/advent-of-code by "+runnerEmail)
	r.Header.Set("Accept", "text/plain")

	return r, nil
}
