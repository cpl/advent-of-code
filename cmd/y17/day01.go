package main

import (
	"bytes"
	"strconv"
)

func SolveDay01Part01(data []byte) (string, error) {
	data = bytes.TrimSpace(data)
	l := len(data)
	sum := 0

	for idx := range data {
		if data[idx] < '0' || data[idx] > '9' {
			continue
		}

		n0 := data[idx]
		n1 := data[(idx+1)%l]

		if n0 == n1 {
			sum += int(n0 - '0')
		}
	}

	return strconv.FormatInt(int64(sum), 10), nil
}

func SolveDay01Part02(data []byte) (string, error) {
	data = bytes.TrimSpace(data)
	l := len(data)
	sum := 0

	for idx := range data[:l/2] {
		if data[idx] < '0' || data[idx] > '9' {
			continue
		}

		n0 := data[idx]
		n1 := data[(idx+l/2)%l]

		if n0 == n1 {
			sum += int(n0 - '0' + n1 - '0')
		}
	}

	return strconv.FormatInt(int64(sum), 10), nil
}
