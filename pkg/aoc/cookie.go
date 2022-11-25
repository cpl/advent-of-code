package aoc

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func LoadCookie() {
	fp, err := os.Open(filepath.Join("..", "..", "..", ".cookie"))
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	cookie, err := io.ReadAll(fp)
	if err != nil {
		panic(err)
	}

	if err = os.Setenv("AOC_COOKIE", strings.TrimSpace(string(cookie))); err != nil {
		panic(err)
	}
}
