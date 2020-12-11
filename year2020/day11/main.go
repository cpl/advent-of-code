package main

import (
	"bufio"
	"bytes"
	"fmt"

	"../../utils"
)

func main() {
	data, err := utils.GetInput(2020, 11)
	utils.CheckErr(err)

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	rows := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		rows = append(rows, []rune(line))
		fmt.Println(line)
	}

	changes := 1
	iter := 1
	for changes != 0 {
		fmt.Println()
		fmt.Println("iteration", iter)
		rows, changes = iterate(rows)
		fmt.Println("occupied", occupied(rows))
		fmt.Println()
		iter++
	}
}

func iterate(rows [][]rune) ([][]rune, int) {
	out := make([][]rune, len(rows))
	width := len(rows[0])
	for idx := range out {
		out[idx] = make([]rune, width)
	}
	changes := 0

	for y, row := range rows {
		for x, char := range row {
			switch char {
			case '.':
				out[y][x] = '.'
			case 'L':
				if occupiedTrace(x, y, rows) == 0 {
					out[y][x] = '#'
					changes++
				} else {
					out[y][x] = 'L'
				}
			case '#':
				if occupiedTrace(x, y, rows) >= 5 {
					out[y][x] = 'L'
					changes++
				} else {
					out[y][x] = '#'
				}
			default:
				panic(char)
			}
			fmt.Print(string(out[y][x]))
		}
		fmt.Println()
	}
	return out, changes
}

func occupied(rows [][]rune) int {
	count := 0
	for _, row := range rows {
		for _, seat := range row {
			if seat == '#' {
				count++
			}
		}
	}
	return count
}

func occupiedTrace(x, y int, rows [][]rune) int {
	return occupiedDirection(x, y, 0, 1, rows) +
		occupiedDirection(x, y, 1, 0, rows) +
		occupiedDirection(x, y, 1, 1, rows) +
		occupiedDirection(x, y, 0, -1, rows) +
		occupiedDirection(x, y, -1, 0, rows) +
		occupiedDirection(x, y, -1, -1, rows) +
		occupiedDirection(x, y, -1, 1, rows) +
		occupiedDirection(x, y, 1, -1, rows)
}

func occupiedDirection(x, y, dx, dy int, rows [][]rune) int {
	mx := len(rows[0])
	my := len(rows)
	sx := false
	sy := false

	for {
		x += dx
		if x < 0 {
			x = 0
			sx = true
		}
		if x >= mx {
			x = mx - 1
			sx = true
		}

		y += dy
		if y < 0 {
			y = 0
			sy = true
		}
		if y >= my {
			y = my - 1
			sy = true
		}

		if dx != 0 && dy != 0 {
			if sx && sy {
				return 0
			}
		}
		if (dx != 0 && sx) || (dy != 0 && sy) {
			return 0
		}

		switch rows[y][x] {
		case '.':
			continue
		case '#':
			return 1
		case 'L':
			return 0
		}
	}

	return 0
}

func occupiedAround(x, y int, rows [][]rune) int {
	count := 0

	rowMid := rows[y]
	if x != 0 {
		if rowMid[x-1] == '#' {
			count++
		}
	}
	if x != len(rowMid)-1 {
		if rowMid[x+1] == '#' {
			count++
		}
	}

	if y != 0 {
		rowTop := rows[y-1]
		if rowTop[x] == '#' {
			count++
		}

		if x != 0 {
			if rowTop[x-1] == '#' {
				count++
			}
		}
		if x != len(rowTop)-1 {
			if rowTop[x+1] == '#' {
				count++
			}
		}
	}

	if y != len(rows)-1 {
		rowBot := rows[y+1]
		if rowBot[x] == '#' {
			count++
		}

		if x != 0 {
			if rowBot[x-1] == '#' {
				count++
			}
		}
		if x != len(rowBot)-1 {
			if rowBot[x+1] == '#' {
				count++
			}
		}
	}

	return count
}
