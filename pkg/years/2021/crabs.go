package y2021

import (
	"math"
	"strconv"
	"strings"
)

type CrabSubmarines struct {
	positions []int
}

func ParseCrabs(input string) *CrabSubmarines {
	input = strings.TrimSpace(input)
	numbers := strings.Split(input, ",")
	numbersInt := make([]int, len(numbers))
	max := 0
	for idx, number := range numbers {
		numberInt, _ := strconv.Atoi(number)
		numbersInt[idx] = numberInt
		if numberInt > max {
			max = numberInt
		}
	}

	positions := make([]int, max+1)
	for _, number := range numbersInt {
		positions[number]++
	}

	return &CrabSubmarines{
		positions: positions,
	}
}

func (c *CrabSubmarines) Align() (int, int) {
	minFuel, minPos := math.MaxInt, 0
	for position := range c.positions {
		fuel := c.AlignAt(position)
		if fuel < minFuel {
			minFuel = fuel
			minPos = position
		}
	}
	return minPos, minFuel
}

func (c *CrabSubmarines) AlignAt(x int) int {
	fuel := 0

	for position, crabs := range c.positions {
		if position < x {
			fuel += (x - position) * crabs
		} else if position > x {
			fuel += (position - x) * crabs
		}
	}

	return fuel
}

func (c *CrabSubmarines) AlignLinearAt(x int) int {
	fuel := 0

	for position, crabs := range c.positions {
		var moves int
		if position < x {
			moves = x - position
		} else if position > x {
			moves = position - x
		}

		for move := 1; move <= moves; move++ {
			fuel += move * crabs
		}
	}

	return fuel
}

func (c *CrabSubmarines) AlignLinear() (int, int) {
	minFuel, minPos := math.MaxInt, 0
	for position := range c.positions {
		fuel := c.AlignLinearAt(position)
		if fuel < minFuel {
			minFuel = fuel
			minPos = position
		}
	}
	return minPos, minFuel
}
