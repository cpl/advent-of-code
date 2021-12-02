package y2021

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
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

type navigationStep struct {
	direction string
	units     int
}

func ParseNavigation(input []byte) []navigationStep {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanLines)

	var steps []navigationStep
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		direction := split[0]
		units, _ := strconv.Atoi(split[1])
		steps = append(steps, navigationStep{direction, units})
	}

	return steps
}

func CalculateNavigation(steps []navigationStep) (int, int) {
	x, depth := 0, 0

	for _, step := range steps {
		switch step.direction {
		case "forward":
			x += step.units
		case "up":
			depth -= step.units
		case "down":
			depth += step.units
		default:
			panic("bad direction: " + step.direction)
		}
	}

	return x, depth
}

func CalculateNavigation2(steps []navigationStep) (int, int) {
	x, depth, aim := 0, 0, 0

	_ = aim
	for _, step := range steps {
		switch step.direction {
		case "up":
			aim -= step.units
		case "down":
			aim += step.units
		case "forward":
			x += step.units
			depth += step.units * aim
		default:
			panic("bad direction: " + step.direction)
		}
	}

	return x, depth
}
