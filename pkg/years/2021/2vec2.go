package y2021

import (
	"bufio"
	"bytes"
	"fmt"
)

type vec2 struct {
	x, y int
}

func Parse2Vec2(input []byte) [][2]vec2 {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanLines)

	out := make([][2]vec2, 0)

	for scanner.Scan() {
		var vec [2]vec2
		line := scanner.Text()

		_, _ = fmt.Sscanf(line, "%d,%d -> %d,%d",
			&vec[0].x, &vec[0].y,
			&vec[1].x, &vec[1].y)

		out = append(out, vec)
	}

	return out
}

type vec2mapper struct {
	data map[string]int
}

func (vm *vec2mapper) Map(vecs [2]vec2) {
	x0, y0 := vecs[0].x, vecs[0].y
	x1, y1 := vecs[1].x, vecs[1].y

	if x0 != x1 && y0 != y1 {
		return
	}

	xStart, xEnd := x0, x1
	if x0 > x1 {
		xStart = x1
		xEnd = x0
	}

	yStart, yEnd := y0, y1
	if y0 > y1 {
		yStart = y1
		yEnd = y0
	}

	for x := xStart; x <= xEnd; x++ {
		for y := yStart; y <= yEnd; y++ {
			vm.data[fmt.Sprintf("%d,%d", x, y)]++
		}
	}
}

func (vm *vec2mapper) Map2(vecs [2]vec2) {
	x0, y0 := vecs[0].x, vecs[0].y
	x1, y1 := vecs[1].x, vecs[1].y

	if x0 == x1 || y0 == y1 {
		vm.Map(vecs)
		return
	}

	mulX := 1
	if x0 > x1 {
		mulX = -1
	}
	mulY := 1
	if y0 > y1 {
		mulY = -1
	}

	steps := x0 - x1
	if steps < 0 {
		steps = -steps
	}

	for step := 0; step <= steps; step++ {
		x := x0 + mulX*step
		y := y0 + mulY*step
		vm.data[fmt.Sprintf("%d,%d", x, y)]++
	}
}

func (vm *vec2mapper) Overlapping() int {
	sum := 0

	for _, v := range vm.data {
		if v > 1 {
			sum++
		}
	}

	return sum
}

func (vm *vec2mapper) Print(maxX, maxY int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if v := vm.data[fmt.Sprintf("%d,%d", x, y)]; v > 0 {
				fmt.Printf("%d", v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
