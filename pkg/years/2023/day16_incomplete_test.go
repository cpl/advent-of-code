package y2023

import (
	"strings"
	"testing"
)

func TestSolveDay16(t *testing.T) {
	t.Parallel()

	parse := func(input string) [][]byte {
		lines := strings.Split(input, "\n")
		out := make([][]byte, len(lines))
		for idx, line := range lines {
			out[idx] = []byte(line)
		}
		return out
	}

	var lightPhysics func(em map[int]struct{}, input [][]byte, sx, sy int, dir rune)
	lightPhysics = func(em map[int]struct{}, input [][]byte, sx, sy int, dir rune) {
		nx, ny := sx, sy

		for {
			switch dir {
			case 'U':
				ny--
			case 'D':
				ny++
			case 'L':
				nx--
			case 'R':
				nx++
			}

			if nx < 0 || nx >= len(input[0]) || ny < 0 || ny >= len(input) {
				return
			}
			pos := ny*len(input[0]) + nx

			nc := input[ny][nx]
			switch nc {
			case '.':
				em[pos] = struct{}{}
			case '-':
				em[pos] = struct{}{}
				if dir == 'U' || dir == 'D' {
					lightPhysics(em, input, nx-1, ny, 'L')
					lightPhysics(em, input, nx+1, ny, 'R')
					return
				}
			case '|':
				em[pos] = struct{}{}
				if dir == 'L' || dir == 'R' {
					lightPhysics(em, input, nx, ny-1, 'U')
					lightPhysics(em, input, nx, ny+1, 'D')
					return
				}
			case '/':
				em[pos] = struct{}{}
				switch dir {
				case 'U':
					nx++
					dir = 'R'
				case 'D':
					nx--
					dir = 'L'
				case 'L':
					ny++
					dir = 'D'
				case 'R':
					ny--
					dir = 'U'
				}
			case '\\':
				em[pos] = struct{}{}
				switch dir {
				case 'U':
					nx--
					dir = 'L'
				case 'D':
					nx++
					dir = 'R'
				case 'L':
					ny--
					dir = 'U'
				case 'R':
					ny++
					dir = 'D'
				}
			}
		}
	}

	part1 := func(contraption [][]byte) int {
		em := make(map[int]struct{})
		lightPhysics(em, contraption, 0, 0, 'R')
		return len(em)
	}

	t.Run("example 1", func(t *testing.T) {
		t.Log(part1(parse(".|...\\....\n|.-.\\.....\n.....|-...\n........|.\n..........\n.........\\\n..../.\\\\..\n.-.-/..|..\n.|....-|.\\\n..//.|....")))
	})
}
