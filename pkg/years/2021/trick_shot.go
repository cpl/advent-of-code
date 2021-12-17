package y2021

import (
	"fmt"
	"strings"
)

type TrickShot uint8

func (t TrickShot) String() string {
	switch t {
	case TrickShotEmpty:
		return "."
	case TrickShotShip:
		return "@"
	case TrickShotTarget:
		return "X"
	case TrickShotVisited:
		return "#"
	case TrickShotHit:
		return "+"
	}

	return "?"
}

const (
	TrickShotEmpty TrickShot = iota
	TrickShotShip
	TrickShotTarget
	TrickShotVisited
	TrickShotHit
)

type TrickShotMap struct {
	data         map[string]TrickShot
	x, y, tX, tY int
}

func (ts *TrickShotMap) Print(sX, sY, tX, tY int) {
	for y := sY; y != tY; y-- {
		for x := sX; x != tX; x++ {
			fmt.Printf("%s", ts.data[fmt.Sprintf("%d,%d", x, y)].String())
		}
		fmt.Println()
	}
}

func (ts *TrickShotMap) Shoot(vX, vY, i int) ([]int, int, int, int) {
	x := 0
	y := 0
	maxY := 0

	hits := make([]int, 0)
	for iter := 0; iter < i; iter++ {
		x += vX
		switch {
		case vX > 0:
			vX--
		case vX < 0:
			vX++
		}

		y += vY
		vY--

		if y > maxY {
			maxY = y
		}

		str := fmt.Sprintf("%d,%d", x, y)
		if tt := ts.data[str]; tt == TrickShotTarget {
			ts.data[str] = TrickShotHit
			hits = append(hits, iter)
		} else {
			ts.data[str] = TrickShotVisited
		}
	}

	return hits, x, y, maxY
}

func ParseTrickShot(input string) *TrickShotMap {
	var sX, sY, tX, tY int

	input = strings.TrimSpace(input)
	_, _ = fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d",
		&sX, &tX, &sY, &tY)

	dx := 1
	if tX < sX {
		dx = -1
	}
	dy := 1
	if tY < sY {
		dy = -1
	}

	data := make(map[string]TrickShot)
	for y := sY; y != tY+dy; y += dy {
		for x := sX; x != tX+dx; x += dx {
			data[fmt.Sprintf("%d,%d", x, y)] = TrickShotTarget
		}
	}

	data["0,0"] = TrickShotShip

	return &TrickShotMap{
		data: data,
		x:    sX,
		y:    sY,
		tX:   tX,
		tY:   tY,
	}
}

func TrickShotX(target int) int {
	return (target * (target + 1)) / 2
}
