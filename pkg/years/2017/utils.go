package year2017

import (
	"bytes"
	"strconv"
	"strings"
)

func ParseDigitList(list []byte) []int {
	list = bytes.TrimSpace(list)
	numbers := make([]int, len(list))

	for idx, b := range list {
		numbers[idx] = int(b - '0')
	}

	return numbers
}

func SumMatchNext(list []int) int {
	sum := 0
	l := len(list)

	for idx, num := range list {
		if num == list[(idx+1)%l] {
			sum += num
		}
	}

	return sum
}

func SumMatchHalfSplitList(list []int) int {
	sum := 0
	l := len(list)

	for idx, num := range list {
		if num == list[(idx+l/2)%l] {
			sum += num
		}
	}

	return sum
}

func MinMaxInList(list []int) (int, int) {
	min := list[0]
	max := list[0]

	for _, num := range list {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	return min, max
}

func ParseNumberListByLine(input string) [][]int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	out := make([][]int, len(lines))

	for idx, line := range lines {
		numbers := strings.Fields(line)
		out[idx] = make([]int, len(numbers))

		for numIdx, num := range numbers {
			out[idx][numIdx], _ = strconv.Atoi(num)
		}
	}

	return out
}

func DivisibleByEachOtherInList(list []int) (int, int) {
	for idx, num := range list {
		for _, otherNum := range list[idx+1:] {
			if num%otherNum == 0 {
				return num, otherNum
			}

			if otherNum%num == 0 {
				return otherNum, num
			}
		}
	}

	return 0, 0
}
