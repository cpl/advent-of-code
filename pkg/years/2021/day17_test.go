package y2021

import (
	"fmt"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay17(t *testing.T) {
	input, err := aoc.MetaGetInput(2021, 17)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("shoot_for_fun", func(t *testing.T) {
		vY := -10

		ts := ParseTrickShot(string(input))
		hits, x, y, maxY := ts.Shoot(30, vY, 1000)
		ts.Print(-5, 25, 210, -115)
		fmt.Println(vY, hits, x, y, maxY)
	})

	// 5356 too low (X=19)
	// 5995 HIT (X=18)
	t.Run("part_1", func(t *testing.T) {
		for vY := 0; vY < 5000; vY++ {
			ts := ParseTrickShot(string(input))
			hits, x, y, maxY := ts.Shoot(18, vY, 1000)
			//ts.Print(-5, 25, 210, -115)
			if len(hits) > 0 {
				fmt.Println(vY, hits, x, y, maxY)
			}
		}
	})

	// 1102 too low
	// 3202 well, brute force works I guess
	// math is for suckers (https://youtu.be/sHzdsFiBbFc?t=70)
	t.Run("part_2", func(t *testing.T) {
		count := 0
		for vX := 1; vX < 500; vX++ {
			for vY := -150; vY < 500; vY++ {
				ts := ParseTrickShot(string(input))
				hits, _, _, _ := ts.Shoot(vX, vY, 500)
				//ts.Print(-5, 25, 210, -115)
				if len(hits) > 0 {
					fmt.Println(vX, vY)
					count++
				}
			}
		}
		t.Logf("count=%d", count)
	})
}
