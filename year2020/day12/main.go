package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"

	"vitess.io/vitess/go/vt/orchestrator/external/golib/math"

	"../../utils"
)

var x, y int
var direction = 'E'
var directionIdx = 1

func main() {
	data, err := utils.GetInput(2020, 12)
	utils.CheckErr(err)

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(string(direction), x, y)

		val, _ := strconv.Atoi(line[1:])
		dir := rune(line[0])

		switch dir {
		case 'N', 'S', 'E', 'W':
			goInDirection(dir, val)
		case 'L', 'R':
			goRotate(dir, val)
		case 'F':
			goInDirection(direction, val)
		}
	}

	fmt.Println(string(direction), x, y)
	fmt.Println(math.AbsInt(x) + math.AbsInt(y))
}

var rotateDef = [...]rune{'N', 'E', 'S', 'W'}

func goRotate(dir rune, val int) {
	switch dir {
	case 'L':
		goRotateLeft(val / 90)
	case 'R':
		goRotateRight(val / 90)
	}
	direction = rotateDef[directionIdx]
}

func goRotateLeft(val int) {
	directionIdx = directionIdx - val
	if directionIdx < 0 {
		directionIdx = 4 + directionIdx
	}
}

func goRotateRight(val int) {
	directionIdx = (directionIdx + val) % 4
}

func goInDirection(dir rune, val int) {
	switch dir {
	case 'N':
		y += val
	case 'S':
		y -= val
	case 'E':
		x += val
	case 'W':
		x -= val
	}
}
