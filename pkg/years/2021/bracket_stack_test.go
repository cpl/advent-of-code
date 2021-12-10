package y2021

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestBracketStack_Parse(t *testing.T) {
	t.Parallel()

	t.Run("single", func(t *testing.T) {
		bs := &BracketStack{}
		idx, r, p := bs.Parse("{([(<{}[<>[]}>{[]{[(<()>")
		if idx != 12 {
			t.Errorf("expected -1, got %d", idx)
		}
		if r != '}' {
			t.Errorf("expected '}' got %c", r)
		}
		if p != '[' {
			t.Errorf("expected '[' got %c", p)
		}
	})

	t.Run("example", func(t *testing.T) {
		input := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

		inputs := strings.Split(input, "\n")
		score := 0
		for _, line := range inputs {
			bs := &BracketStack{}
			idx, r, _ := bs.Parse(line)
			if idx != -1 {
				t.Logf("bad bracket: %c", r)
				score += bracketScore[r]
			}
		}

		t.Logf("score=%d", score)
		if score != 26397 {
			t.Errorf("expected 26397, got %d", score)
		}
	})
}

func TestBracketStack_Parse_withIncomplete(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

		inputs := strings.Split(input, "\n")
		scores := make([]int, 0, len(inputs))
		want := []int{288957, 5566, 1480781, 995444, 294}

		for _, line := range inputs {
			bs := &BracketStack{}
			idx, _, _ := bs.Parse(line)
			if idx != -1 {
				continue
			}

			if !bs.IsEmpty() {
				scores = append(scores, bs.RemainingScore())
			}
		}

		t.Logf("scores=%v", scores)
		if !reflect.DeepEqual(scores, want) {
			t.Errorf("expected %v, got %v", want, scores)
		}

		sort.Ints(scores)
		score := scores[len(scores)/2]
		t.Logf("middle score=%d", score)
		if score != 288957 {
			t.Errorf("expected 288957, got %d", score)
		}
	})
}
