package intcode

import (
	"strconv"
	"strings"
)

func Parse(s string) []int {
	strs := strings.Split(s, ",")
	program := make([]int, len(strs))

	for i, str := range strs {
		program[i], _ = strconv.Atoi(str)
	}

	return program
}
