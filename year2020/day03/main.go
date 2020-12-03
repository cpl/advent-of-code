package main

import (
	"bufio"
	"bytes"
	"fmt"

	"../../utils"
)

const (
	Tree = '#'
	Free = '.'
)

func main() {
	data, err := utils.GetInput(2020, 03)
	utils.CheckErr(err)

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	grid := make([]string, 0, 1000)

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	a := hitTrees(1, 1, grid)
	b := hitTrees(3, 1, grid)
	c := hitTrees(5, 1, grid)
	d := hitTrees(7, 1, grid)
	e := hitTrees(1, 2, grid)

	fmt.Println(a, b, c, d, e)
	fmt.Println(a * b * c * d * e)
}

func hitTrees(right, down int, grid []string) int {
	count := 0
	x := 0
	y := 0
	maxX := len(grid[0])
	maxY := len(grid)

	for y < maxY-down {
		x = (x + right) % maxX
		y = y + down
		if grid[y][x] == Tree {
			count++
		}
	}

	return count
}
