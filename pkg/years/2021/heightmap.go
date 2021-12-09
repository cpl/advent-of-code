package y2021

import (
	"fmt"
	"strings"
)

type HeightmapBasin struct {
	id    int
	count int
}

type HeightmapPoint struct {
	x, y   int
	height int

	basin *HeightmapBasin
}

type Heightmap struct {
	points [][]*HeightmapPoint
	basins []*HeightmapBasin
}

func (hm *Heightmap) Print() {
	for _, row := range hm.points {
		for _, p := range row {
			if p.basin == nil {
				if p.height == 9 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(string(rune('a' + p.basin.id)))
			}
		}
		fmt.Println()
	}
}

func (hm *Heightmap) SeekBasins() {
	for y := 0; y < len(hm.points); y++ {
		for x := 0; x < len(hm.points[y]); x++ {
			hm.seekBasin(x, y)
		}
	}
}

func (hm *Heightmap) seekBasin(x, y int) {
	if y < 0 || y >= len(hm.points) {
		return
	}
	if x < 0 || x >= len(hm.points[y]) {
		return
	}

	point := hm.points[y][x]
	if point.height == 9 || point.basin != nil {
		return
	}

	basinT := hm.seekBasinNeighbours(x, y-1)
	basinR := hm.seekBasinNeighbours(x+1, y)
	basinB := hm.seekBasinNeighbours(x, y+1)
	basinL := hm.seekBasinNeighbours(x-1, y)
	switch {
	case basinT != nil:
		point.basin = basinT
		basinT.count++
	case basinR != nil:
		point.basin = basinR
		basinR.count++
	case basinB != nil:
		point.basin = basinB
		basinB.count++
	case basinL != nil:
		point.basin = basinL
		basinL.count++
	default:
		point.basin = &HeightmapBasin{
			id:    len(hm.basins),
			count: 1,
		}
		hm.basins = append(hm.basins, point.basin)
	}

	hm.seekBasin(x+1, y)
	hm.seekBasin(x-1, y)
	hm.seekBasin(x, y+1)
	hm.seekBasin(x, y-1)
}

func (hm *Heightmap) seekBasinNeighbours(x, y int) *HeightmapBasin {
	if y < 0 || y >= len(hm.points) {
		return nil
	}
	if x < 0 || x >= len(hm.points[y]) {
		return nil
	}

	point := hm.points[y][x]
	if point.basin == nil {
		return nil
	}

	return point.basin
}

func ParseHeightmap(input string) [][]int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	hm := make([][]int, len(lines))
	for idx, line := range lines {
		hm[idx] = make([]int, len(line))
		for jdx, c := range line {
			hm[idx][jdx] = int(c - '0')
		}
	}
	return hm
}

func ParseHeightmap2(input string) *Heightmap {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	hm := &Heightmap{
		points: make([][]*HeightmapPoint, len(lines)),
		basins: nil,
	}

	for y, line := range lines {
		hm.points[y] = make([]*HeightmapPoint, len(line))
		for x, c := range line {
			height := int(c - '0')
			hm.points[y][x] = &HeightmapPoint{
				x:      x,
				y:      y,
				height: height,
			}
		}
	}

	return hm
}

func HeightmapLows(hm [][]int) []int {
	lows := make([]int, 0)

	for y := 0; y < len(hm); y++ {
		for x := 0; x < len(hm[y]); x++ {
			v := hm[y][x]
			if y > 0 && hm[y-1][x] <= v {
				continue
			}
			if y < len(hm)-1 && hm[y+1][x] <= v {
				continue
			}
			if x > 0 && hm[y][x-1] <= v {
				continue
			}
			if x < len(hm[y])-1 && hm[y][x+1] <= v {
				continue
			}

			lows = append(lows, v)
		}
	}

	return lows
}
