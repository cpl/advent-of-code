package y2022

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
)

func TestSolveDay11(t *testing.T) {
	type monkey struct {
		inspected int
		items     []int
		operation func(int) int
		test      func(int) int
	}

	parseItems := func(line string) []int {
		line = strings.TrimSpace(line)
		line = strings.TrimPrefix(line, "Starting items: ")

		nums := strings.Split(line, ", ")
		items := make([]int, len(nums))
		for i, num := range nums {
			items[i], _ = strconv.Atoi(num)
		}

		return items
	}

	parseOperation := func(line string) func(int) int {
		var a string
		var b string
		var op string

		_, _ = fmt.Sscanf(strings.TrimSpace(line), "Operation: new = %s %s %s", &a, &op, &b)

		if a == "old" && b == "old" {
			switch op {
			case "+":
				return func(old int) int { return old + old }
			case "-":
				return func(old int) int { return old - old }
			case "*":
				return func(old int) int { return old * old }
			case "/":
				return func(old int) int { return old / old }
			}
		}

		bNum, _ := strconv.Atoi(b)
		switch op {
		case "+":
			return func(old int) int { return old + bNum }
		case "-":
			return func(old int) int { return old - bNum }
		case "*":
			return func(old int) int { return old * bNum }
		case "/":
			return func(old int) int { return old / bNum }
		}

		panic("invalid operation")
	}

	parseTest := func(lines []string) func(int) int {
		var div int
		var retTrue int
		var retFalse int

		_, _ = fmt.Sscanf(strings.TrimSpace(lines[0]), "Test: divisible by %d", &div)
		_, _ = fmt.Sscanf(strings.TrimSpace(lines[1]), "If true: throw to monkey %d", &retTrue)
		_, _ = fmt.Sscanf(strings.TrimSpace(lines[2]), "If false: throw to monkey %d", &retFalse)

		return func(v int) int {
			if v%div == 0 {
				return retTrue
			}
			return retFalse
		}
	}

	parse := func(lines []string) []*monkey {
		monkeys := make([]*monkey, 0, (len(lines)+1)/6)

		for idx := 0; idx < len(lines); idx += 7 {
			mo := &monkey{
				items:     parseItems(lines[idx+1]),
				operation: parseOperation(lines[idx+2]),
				test:      parseTest(lines[idx+3 : idx+6]),
			}

			monkeys = append(monkeys, mo)
		}

		return monkeys
	}

	round := func(monkeys []*monkey, relief bool) {
		for _, mo := range monkeys {
			for _, item := range mo.items {
				item = mo.operation(item)

				if relief {
					item = item / 3
				}
				target := mo.test(item)
				monkeys[target].items = append(monkeys[target].items, item)
			}

			mo.inspected += len(mo.items)
			mo.items = nil
		}
	}

	part1 := func(monkeys []*monkey) int {
		for roundIdx := 0; roundIdx < 20; roundIdx++ {
			round(monkeys, true)
		}

		var top1, top2 int

		for _, mo := range monkeys {
			if mo.inspected > top1 {
				top2 = top1
				top1 = mo.inspected
			} else if mo.inspected > top2 {
				top2 = mo.inspected
			}
		}

		return top1 * top2
	}

	part2 := func(monkeys []*monkey) int {
		for roundIdx := 0; roundIdx < 10000; roundIdx++ {
			round(monkeys, false)
		}

		var top1, top2 int

		for _, mo := range monkeys {
			if mo.inspected > top1 {
				top2 = top1
				top1 = mo.inspected
			} else if mo.inspected > top2 {
				top2 = mo.inspected
			}
		}

		return top1 * top2
	}

	t.Run("Example 1", func(t *testing.T) {
		input := "Monkey 0:\n  Starting items: 79, 98\n  Operation: new = old * 19\n  Test: divisible by 23\n    If true: throw to monkey 2\n    If false: throw to monkey 3\n\nMonkey 1:\n  Starting items: 54, 65, 75, 74\n  Operation: new = old + 6\n  Test: divisible by 19\n    If true: throw to monkey 2\n    If false: throw to monkey 0\n\nMonkey 2:\n  Starting items: 79, 60, 97\n  Operation: new = old * old\n  Test: divisible by 13\n    If true: throw to monkey 1\n    If false: throw to monkey 3\n\nMonkey 3:\n  Starting items: 74\n  Operation: new = old + 3\n  Test: divisible by 17\n    If true: throw to monkey 0\n    If false: throw to monkey 1"
		lines := strings.Split(input, "\n")
		t.Log(part1(parse(lines)))
	})

	t.Run("Part 1", func(t *testing.T) {
		t.Log(part1(parse(strings.Split(aoc.PuzzleString(2022, 11), "\n"))))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := "Monkey 0:\n  Starting items: 79, 98\n  Operation: new = old * 19\n  Test: divisible by 23\n    If true: throw to monkey 2\n    If false: throw to monkey 3\n\nMonkey 1:\n  Starting items: 54, 65, 75, 74\n  Operation: new = old + 6\n  Test: divisible by 19\n    If true: throw to monkey 2\n    If false: throw to monkey 0\n\nMonkey 2:\n  Starting items: 79, 60, 97\n  Operation: new = old * old\n  Test: divisible by 13\n    If true: throw to monkey 1\n    If false: throw to monkey 3\n\nMonkey 3:\n  Starting items: 74\n  Operation: new = old + 3\n  Test: divisible by 17\n    If true: throw to monkey 0\n    If false: throw to monkey 1"
		lines := strings.Split(input, "\n")
		t.Log(part2(parse(lines)))
	})

	t.Run("Part 2", func(t *testing.T) {
		t.Log(part2(parse(strings.Split(aoc.PuzzleString(2022, 11), "\n"))))
	})
}
