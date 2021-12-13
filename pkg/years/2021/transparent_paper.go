package y2021

import (
	"fmt"
	"strconv"
	"strings"
)

type TransparentPaper struct {
	dotsGrid  [][]*transparentPaperDot
	folds     []transparentPaperFold
	foldIndex int
}

type transparentPaperFold struct {
	cord      int
	direction rune
}

type transparentPaperDot struct {
	x int
	y int

	overlap int
	fold    bool
}

func ParseTransparentPaper(input string) *TransparentPaper {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	tp := &TransparentPaper{}

	dotsMap := make(map[string]bool)
	maxX := 0
	maxY := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		if line[0] == 'f' {
			idx := strings.IndexRune(line, '=')
			cord, _ := strconv.Atoi(line[idx+1:])
			direction := rune(line[idx-1])
			tp.folds = append(tp.folds, transparentPaperFold{cord: cord, direction: direction})

			continue
		}

		split := strings.Split(line, ",")
		dotX, _ := strconv.Atoi(split[0])
		dotY, _ := strconv.Atoi(split[1])

		if dotX > maxX {
			maxX = dotX
		}
		if dotY > maxY {
			maxY = dotY
		}

		dotsMap[line] = true
	}

	dotsGrid := make([][]*transparentPaperDot, maxY+1)
	for y := 0; y <= maxY; y++ {
		dotsGrid[y] = make([]*transparentPaperDot, maxX+1)
		for x := 0; x <= maxX; x++ {
			dotsGrid[y][x] = &transparentPaperDot{
				x: x,
				y: y,
			}

			if dotsMap[fmt.Sprintf("%d,%d", x, y)] {
				dotsGrid[y][x].overlap = 1
			}
		}
	}

	for _, fold := range tp.folds {
		if fold.direction == 'y' {
			row := dotsGrid[fold.cord]
			for idx := range row {
				row[idx].fold = true
			}
		} else {
			for y := 0; y <= maxY; y++ {
				dotsGrid[y][fold.cord].fold = true
			}
		}
	}

	tp.dotsGrid = dotsGrid

	return tp
}

func (tp *TransparentPaper) Print() {
	fold := transparentPaperFold{
		cord:      -1,
		direction: 'x',
	}
	if tp.foldIndex < len(tp.folds) {
		fold = tp.folds[tp.foldIndex]
	}

	for y, dots := range tp.dotsGrid {
		for x, dot := range dots {
			if fold.direction == 'y' {
				if fold.cord == y {
					if dot.overlap > 0 {
						fmt.Print("*")
						continue
					}

					fmt.Print("+")
					continue
				}
			} else {
				if fold.cord == x {
					if dot.overlap > 0 {
						fmt.Print("*")
						continue
					}

					fmt.Print("+")
					continue
				}
			}

			if dot.overlap < 0 {
				fmt.Print("x")
			} else if dot.overlap > 0 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (tp *TransparentPaper) Fold() bool {
	fold := tp.folds[tp.foldIndex]
	tp.foldIndex++

	if fold.direction == 'y' {
		tp.foldY(fold.cord)
	} else {
		tp.foldX(fold.cord)
	}

	return tp.foldIndex < len(tp.folds)
}

func (tp *TransparentPaper) foldX(cord int) {
	for y := 0; y < len(tp.dotsGrid); y++ {
		for x := 0; x < cord; x++ {
			on := tp.dotsGrid[y][x]
			from := tp.dotsGrid[y][len(tp.dotsGrid[y])-1-x]

			on.overlap += from.overlap
			from.overlap = -1
		}

		tp.dotsGrid[y] = tp.dotsGrid[y][:cord]
	}
}

func (tp *TransparentPaper) foldY(cord int) {
	for y := 0; y < cord; y++ {
		foldOn := tp.dotsGrid[y]
		foldFrom := tp.dotsGrid[len(tp.dotsGrid)-1-y]

		for x, dot := range foldFrom {
			foldOn[x].overlap += dot.overlap
			dot.overlap = -1
		}
	}

	tp.dotsGrid = tp.dotsGrid[:cord]
}

func (tp *TransparentPaper) Displayed() int {
	total := 0

	for _, row := range tp.dotsGrid {
		for _, dot := range row {
			if dot.overlap > 0 {
				total++
			}
		}
	}

	return total
}
