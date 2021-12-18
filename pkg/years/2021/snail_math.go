package y2021

import (
	"fmt"
	"strconv"
	"strings"
)

type SnailNumber struct {
	left, right *SnailNumber
	parent      *SnailNumber

	isLeft bool
	value  int
	depth  int
}

func (s *SnailNumber) incrementDepth() {
	s.depth++
	if s.left != nil {
		s.left.incrementDepth()
	}
	if s.right != nil {
		s.right.incrementDepth()
	}
}

func (s *SnailNumber) String() string {
	if s.value != -1 {
		return fmt.Sprintf("%d", s.value)
	}

	return fmt.Sprintf("[%s,%s]", s.left.String(), s.right.String())
}

func (s *SnailNumber) StringDepth() string {
	if s.value != -1 {
		return fmt.Sprintf("%d", s.depth)
	}

	return fmt.Sprintf("[%s,%s]", s.left.StringDepth(), s.right.StringDepth())
}

func (s *SnailNumber) Magnitude() int {
	if s.value != -1 {
		return s.value
	}

	return s.left.Magnitude()*3 + s.right.Magnitude()*2
}

func ParseSnailNumbers(input string) (*SnailNumber, []*SnailNumber) {
	idx := 0
	idxMax := len(input)

	num := &SnailNumber{
		value: -1,
		depth: 0,
	}

	numbers := make([]*SnailNumber, 0)

	depth := 0
	current := num

	for idx < idxMax {
		ch := input[idx]
		switch {
		case ch == '[':
			depth++
			current.left = &SnailNumber{
				parent: current,
				isLeft: true,
				value:  -1,
				depth:  depth,
			}
			current = current.left
		case ch == ']':
			depth--
			current = current.parent
		case ch == ',':
			current = current.parent
			current.right = &SnailNumber{
				parent: current,
				isLeft: false,
				value:  -1,
				depth:  depth,
			}
			current = current.right
		case ch >= '0' && ch <= '9':
			end := strings.IndexAny(input[idx:], ",]")
			value, err := strconv.ParseInt(input[idx:idx+end], 10, 64)
			if err != nil {
				panic(err)
			}
			current.value = int(value)
			idx += end - 1

			numbers = append(numbers, current)
		}

		idx++
	}

	return num, numbers
}

func SnailNumbersReduce(numbers []*SnailNumber) ([]*SnailNumber, bool) {
	newNumbers := make([]*SnailNumber, 0, len(numbers)+4)

	idxMax := len(numbers)
	for idx, number := range numbers {
		if number.depth > 4 && number.parent.left.value != -1 && number.parent.right.value != -1 {
			parent := number.parent

			// explode
			fmt.Println("explode", idx, parent)
			if idx != 0 {
				fmt.Println("explode add", numbers[idx-1].value, parent.left.value)
				numbers[idx-1].value += parent.left.value
			}
			if idx < idxMax-2 {
				fmt.Println("explode add", numbers[idx+2].value, parent.right.value)
				numbers[idx+2].value += parent.right.value
			}

			parent.value = 0
			parent.left = nil
			parent.right = nil

			newNumbers = append(newNumbers, parent)
			newNumbers = append(newNumbers, numbers[idx+2:]...)

			return newNumbers, true
		}

		if number.value >= 10 {
			// split
			half := number.value / 2
			rest := number.value - half
			fmt.Println("split", number)

			number.value = -1
			number.left = &SnailNumber{
				parent: number,
				isLeft: true,
				value:  half,
				depth:  number.depth + 1,
			}
			number.right = &SnailNumber{
				parent: number,
				isLeft: false,
				value:  rest,
				depth:  number.depth + 1,
			}

			newNumbers = append(newNumbers, number.left)
			newNumbers = append(newNumbers, number.right)
			newNumbers = append(newNumbers, numbers[idx+1:]...)

			return newNumbers, true
		}

		newNumbers = append(newNumbers, number)
	}

	return newNumbers, false
}

func snailNumbersReduce(numbers []*SnailNumber) ([]*SnailNumber, bool) {
	numbers, reduce := snailNumbersExplode(numbers)
	if reduce {
		return numbers, true
	}
	numbers, reduce = snailNumbersSplit(numbers)

	return numbers, reduce
}

func snailNumbersExplode(numbers []*SnailNumber) ([]*SnailNumber, bool) {
	newNumbers := make([]*SnailNumber, 0, len(numbers)+4)

	idxMax := len(numbers)
	for idx, number := range numbers {
		if number.depth > 4 && number.parent.left.value != -1 && number.parent.right.value != -1 {
			parent := number.parent

			// explode
			if idx != 0 {
				numbers[idx-1].value += parent.left.value
			}
			if idx < idxMax-2 {
				numbers[idx+2].value += parent.right.value
			}

			parent.value = 0
			parent.left = nil
			parent.right = nil

			newNumbers = append(newNumbers, parent)
			newNumbers = append(newNumbers, numbers[idx+2:]...)

			return newNumbers, true
		}

		newNumbers = append(newNumbers, number)
	}

	return newNumbers, false
}

func snailNumbersSplit(numbers []*SnailNumber) ([]*SnailNumber, bool) {
	newNumbers := make([]*SnailNumber, 0, len(numbers)+4)

	for idx, number := range numbers {
		if number.value >= 10 {
			// split
			half := number.value / 2
			rest := number.value - half

			number.value = -1
			number.left = &SnailNumber{
				parent: number,
				isLeft: true,
				value:  half,
				depth:  number.depth + 1,
			}
			number.right = &SnailNumber{
				parent: number,
				isLeft: false,
				value:  rest,
				depth:  number.depth + 1,
			}

			newNumbers = append(newNumbers, number.left)
			newNumbers = append(newNumbers, number.right)
			newNumbers = append(newNumbers, numbers[idx+1:]...)

			return newNumbers, true
		}

		newNumbers = append(newNumbers, number)
	}

	return newNumbers, false
}

func SnailNumbersAdd(x, y *SnailNumber, xNumbers, yNumbers []*SnailNumber) (*SnailNumber, []*SnailNumber) {
	x.isLeft = true
	z := &SnailNumber{
		left:   x,
		right:  y,
		parent: nil,
		isLeft: false,
		value:  -1,
	}
	x.parent = z
	y.parent = z

	x.incrementDepth()
	y.incrementDepth()

	return z, append(xNumbers, yNumbers...)
}
