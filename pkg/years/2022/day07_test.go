package y2022

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay07(t *testing.T) {
	type file struct {
		name     string
		parent   *file
		children []*file
		size     int
	}

	parse := func(lines []string) *file {
		root := &file{name: "/"}
		current := root

		for _, line := range lines[1:] {
			if line[0] == '$' {
				split := strings.Split(strings.TrimSpace(line[1:]), " ")
				switch split[0] {
				case "ls":
				case "cd":
					switch split[1] {
					case "..":
						current = current.parent
					case "/":
						current = root
					default:
						for _, child := range current.children {
							if child.name == split[1] {
								current = child
								break
							}
						}
					}
				default:
					panic(fmt.Errorf("unknown command %q", split[0]))
				}

				continue
			}

			split := strings.Split(strings.TrimSpace(line), " ")
			switch split[0] {
			case "dir":
				current.children = append(current.children, &file{
					name:   split[1],
					parent: current,
				})
			default:
				size, _ := strconv.Atoi(split[0])
				current.children = append(current.children, &file{
					name:   split[1],
					parent: current,
					size:   size,
				})
			}
		}

		return root
	}

	var totalSize func(start *file) int
	totalSize = func(start *file) int {
		if start.size > 0 {
			return start.size
		}

		var total int
		for _, child := range start.children {
			total += totalSize(child)
		}

		start.size = total

		return total
	}

	isDir := func(f *file) bool {
		return len(f.children) > 0
	}

	var atMostSum func(start *file, atMost int) int
	atMostSum = func(start *file, atMost int) int {
		total := 0
		if start.size < atMost && isDir(start) {
			total += start.size
		}

		for _, child := range start.children {
			total += atMostSum(child, atMost)
		}

		return total
	}

	space := func(start *file) (unused, used, target int) {
		available := 70000000
		needed := 30000000

		used = start.size
		unused = available - used
		target = needed - unused

		return
	}

	part1 := func(start *file) int {
		return atMostSum(start, 100000)
	}

	var freeTargets func(start *file, target int) []*file
	freeTargets = func(start *file, target int) []*file {
		targets := make([]*file, 0, 10)

		for _, child := range start.children {
			if !isDir(child) {
				continue
			}

			if child.size < target {
				continue
			}

			if child.size >= target {
				targets = append(targets, child)

				more := freeTargets(child, target)
				if len(more) == 0 {
					continue
				}
				targets = append(targets, more...)
			}
		}

		return targets
	}

	part2 := func(start *file) int {
		_, _, target := space(start)
		targets := freeTargets(start, target)
		min := targets[0].size

		for _, dir := range targets {
			if dir.size < min {
				min = dir.size
			}
		}

		return min
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k"
		root := parse(strings.Split(input, "\n"))
		totalSize(root)
		if root.size != 48381165 {
			t.Fatalf("Expected 48381165, got %d", root.size)
		}

		t.Log(part1(root))
	})

	t.Run("Part 1", func(t *testing.T) {
		root := parse(aoc.PuzzleStringSliceNewline(2022, 7))
		totalSize(root)
		t.Log(part1(root))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k"
		root := parse(strings.Split(input, "\n"))
		totalSize(root)
		t.Log(part2(root))
	})

	t.Run("Part 2", func(t *testing.T) {
		root := parse(aoc.PuzzleStringSliceNewline(2022, 7))
		totalSize(root)
		t.Log(part2(root))
	})
}
