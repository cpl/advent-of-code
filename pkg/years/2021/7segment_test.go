package y2021

import (
	"bufio"
	"strings"
	"testing"
)

//  aaaa
// b    c
// b    c
//  dddd
// e    f
// e    f
//  gggg

// 1 = _ _ c _ _ f _ # 2 x
// 7 = a _ c _ _ f _ # 3 x
// 4 = _ b c d _ f _ # 4 x
// 8 = a b c d e f g # 7 x

// 2 = a _ c d e _ g # 5
// 3 = a _ c d _ f g # 5 x
// 5 = a b _ d _ f g # 5

// 6 = a b _ d e f g # 6
// 0 = a b c _ e f g # 6
// 9 = a b c d _ f g # 6 x

func Test7Segment(t *testing.T) {
	t.Parallel()

	t.Run("example", func(t *testing.T) {
		input := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"

		segments := Parse7Segments(strings.Split(input, " "))
		mapping := segments.SolveMapping()
		want := map[rune]rune{
			'd': 'a',
			'e': 'b',
			'a': 'c',
			'f': 'd',
			'g': 'e',
			'b': 'f',
			'c': 'g',
		}

		for k, v := range want {
			if mapping[k] != v {
				t.Errorf("%s -> %s, want %s", string(k), string(mapping[k]), string(v))
			}
		}

		output := strings.Split("cdfeb fcadb cdfeb cdbaf", " ")
		solution := Parse7SegmentOutput(output, mapping)

		if solution != "5353" {
			t.Errorf("%s, want %s", solution, "5353")
		}
	})

	t.Run("example long", func(t *testing.T) {
		example := `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

		scanner := bufio.NewScanner(strings.NewReader(example))
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			splitIdx := strings.IndexRune(line, '|')
			segments := Parse7Segments(strings.Split(line[:splitIdx-1], " "))
			mapping := segments.SolveMapping()
			solution := Parse7SegmentOutput(strings.Split(line[splitIdx+2:], " "), mapping)
			t.Logf(solution)
		}
	})
}
