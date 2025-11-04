package aoc

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func submitSuccessFile(year, day, part int) string {
	return filepath.Join("data", fmt.Sprintf("%d_%02d_part%d_submit.txt", year, day, part))
}

func submitSuccessExists(year, day, part int) bool {
	path := submitSuccessFile(year, day, part)

	fp, err := os.OpenFile(path, os.O_RDONLY, 0o600)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		panic(err)
	}
	defer fp.Close()

	stat, err := fp.Stat()
	if err != nil {
		panic(err)
	}

	if stat.Size() != 0 {
		return true
	}

	return false
}

func submitSuccessWrite(year, day, part int, answer any) {
	fp, err := os.OpenFile(submitSuccessFile(year, day, part), os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	_, _ = fp.WriteString(time.Now().String())
	_, _ = fp.WriteString("\n")
	_, _ = fp.WriteString(fmt.Sprintf("%v\n", answer))
}
