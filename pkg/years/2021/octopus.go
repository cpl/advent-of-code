package y2021

import "strings"

type Octopus struct {
	Charge  int
	Flashed bool
}

type OctopusGrid struct {
	Grid         [][]Octopus
	TotalFlashes int
}

func ParseOctopusGrid(input string) *OctopusGrid {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	grid := &OctopusGrid{
		Grid: make([][]Octopus, len(lines)),
	}

	for y, line := range lines {
		grid.Grid[y] = make([]Octopus, len(line))
		for x, char := range line {
			grid.Grid[y][x] = Octopus{
				Charge:  int(char - '0'),
				Flashed: false,
			}
		}
	}

	return grid
}

func (g *OctopusGrid) ChargeAll() int {
	for y, row := range g.Grid {
		for x := range row {
			g.Charge(x, y)
		}
	}

	flashed := 0
	for y, row := range g.Grid {
		for x := range row {
			if g.Grid[y][x].Flashed {
				flashed++
				g.Grid[y][x].Flashed = false
			}
		}
	}

	return flashed
}

func (g *OctopusGrid) Charge(x, y int) {
	if x < 0 || y < 0 || x >= len(g.Grid[0]) || y >= len(g.Grid) {
		return
	}
	if g.Grid[y][x].Flashed {
		return
	}

	g.Grid[y][x].Charge++

	if g.Grid[y][x].Charge > 9 {
		g.Flash(x, y)
	}
}

func (g *OctopusGrid) Flash(x, y int) {
	if x < 0 || y < 0 || x >= len(g.Grid[0]) || y >= len(g.Grid) {
		return
	}
	if g.Grid[y][x].Flashed {
		return
	}

	g.Grid[y][x].Flashed = true
	g.TotalFlashes++
	g.Grid[y][x].Charge = 0

	g.Charge(x-1, y)
	g.Charge(x+1, y)
	g.Charge(x, y-1)
	g.Charge(x, y+1)
	g.Charge(x-1, y-1)
	g.Charge(x+1, y-1)
	g.Charge(x-1, y+1)
	g.Charge(x+1, y+1)
}

func (g *OctopusGrid) String() string {
	var builder strings.Builder

	for _, row := range g.Grid {
		for _, octopus := range row {
			builder.WriteString(string(rune(octopus.Charge + '0')))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
