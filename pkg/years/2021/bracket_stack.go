package y2021

type BracketStack struct {
	stack []rune
}

func (bs *BracketStack) Push(r rune) {
	bs.stack = append(bs.stack, r)
}

func (bs *BracketStack) Pop() rune {
	if len(bs.stack) == 0 {
		return ' '
	}
	r := bs.stack[len(bs.stack)-1]
	bs.stack = bs.stack[:len(bs.stack)-1]
	return r
}

func (bs *BracketStack) IsEmpty() bool {
	return len(bs.stack) == 0
}

var bracketMatch = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var bracketScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func (bs *BracketStack) Parse(input string) (int, rune, rune) {
	runes := []rune(input)

	for idx, r := range runes {
		switch r {
		case '(', '[', '{', '<':
			bs.Push(r)
		case ')', ']', '}', '>':
			if bs.IsEmpty() {
				return idx, ' ', ' '
			}

			poped := bs.Pop()
			if bracketMatch[r] != poped {
				return idx, r, poped
			}
		}
	}

	return -1, ' ', ' '
}

func (bs *BracketStack) RemainingScore() int {
	score := 0
	for idx := len(bs.stack) - 1; idx >= 0; idx-- {
		r := bs.stack[idx]

		score *= 5
		switch r {
		case '(':
			score += 1
		case '[':
			score += 2
		case '{':
			score += 3
		case '<':
			score += 4
		}
	}

	return score
}
