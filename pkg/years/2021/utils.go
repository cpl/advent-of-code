package y2021

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
)

func ParseNumbers(input []byte) []int {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanLines)

	var numbers []int
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	return numbers
}

func NumberUpOrDown(numbers []int) (int, int, []int) {
	out := make([]int, len(numbers))
	var up, down int

	for idx := 1; idx < len(numbers); idx++ {
		current := numbers[idx]
		prev := numbers[idx-1]

		if current > prev {
			up++
		} else if current < prev {
			down++
		}

		out[idx] = current - prev
	}

	return up, down, out
}

func NumberThreeMeasureSums(numbers []int) []int {
	var out []int

	for idx := 0; idx < len(numbers)-2; idx++ {
		fmt.Println(numbers[idx], numbers[idx+1], numbers[idx+2])
		sum := numbers[idx] + numbers[idx+1] + numbers[idx+2]
		out = append(out, sum)
	}

	return out
}
