package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func SolveDay02Part01(input []byte) (string, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		min := 1000000
		max := 0

		for _, number := range numbers {
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				return "", err
			}
			if numberInt > max {
				max = numberInt
			}
			if numberInt < min {
				min = numberInt
			}
		}

		fmt.Println("min", min, "max", max, "nums", numbers)

		sum += max - min
	}

	return strconv.Itoa(sum), nil
}
