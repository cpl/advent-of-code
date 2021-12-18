package y2021

import (
	"fmt"
	"strings"
	"testing"
)

func TestSnailMath(t *testing.T) {
	t.Parallel()

	inputs := []string{
		"[1,2]",
		"[[1,2],3]",
		"[9,[8,7]]",
		"[[1,9],[8,5]]",
		"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
		"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
		"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
	}

	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			got, numbers := ParseSnailNumbers(input)
			if got.String() != input {
				t.Errorf("got %v, want %v", got, input)
			}

			t.Log(numbers)
		})
	}
}

func TestSnailMath_reduce(t *testing.T) {
	t.Parallel()

	t.Run("explode", func(t *testing.T) {
		tests := map[string]string{
			"[[[[[9,8],1],2],3],4]":                 "[[[[0,9],2],3],4]",
			"[7,[6,[5,[4,[3,2]]]]]":                 "[7,[6,[5,[7,0]]]]",
			"[[6,[5,[4,[3,2]]]],1]":                 "[[6,[5,[7,0]]],3]",
			"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]": "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]":     "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		}

		for input, want := range tests {
			t.Run(input, func(t *testing.T) {
				root, numbers := ParseSnailNumbers(input)
				numbers, _ = SnailNumbersReduce(numbers)
				if root.String() != want {
					t.Errorf("got %v, want %v", root.String(), want)
				}
			})
		}
	})

	t.Run("split", func(t *testing.T) {
		t.Parallel()

		tests := map[string]string{
			"[[[[0,7],4],[15,[0,13]]],[1,1]]":    "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
			"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]": "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
		}

		for input, want := range tests {
			t.Run(input, func(t *testing.T) {
				root, numbers := ParseSnailNumbers(input)
				numbers, _ = SnailNumbersReduce(numbers)
				if root.String() != want {
					t.Errorf("got %v, want %v", root.String(), want)
				}
			})
		}
	})
}

func TestSnailMath_add_reduce(t *testing.T) {
	t.Run("simple example", func(t *testing.T) {
		x, xN := ParseSnailNumbers("[[[[4,3],4],4],[7,[[8,4],9]]]")
		y, yN := ParseSnailNumbers("[1,1]")
		z, zN := SnailNumbersAdd(x, y, xN, yN)

		repeat := true
		for repeat {
			zN, repeat = SnailNumbersReduce(zN)
		}

		want := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
		if z.String() != want {
			t.Errorf("got %v, want %v", z.String(), want)
		}
	})

	t.Run("added 6", func(t *testing.T) {
		want := "[[[[5,0],[7,4]],[5,5]],[6,6]]"

		a, aN := ParseSnailNumbers("[1,1]")
		b, bN := ParseSnailNumbers("[2,2]")
		a, aN = SnailNumbersAdd(a, b, aN, bN)
		b, bN = ParseSnailNumbers("[3,3]")
		a, aN = SnailNumbersAdd(a, b, aN, bN)
		b, bN = ParseSnailNumbers("[4,4]")
		a, aN = SnailNumbersAdd(a, b, aN, bN)
		b, bN = ParseSnailNumbers("[5,5]")
		a, aN = SnailNumbersAdd(a, b, aN, bN)
		b, bN = ParseSnailNumbers("[6,6]")
		a, aN = SnailNumbersAdd(a, b, aN, bN)

		reduce := true
		for reduce {
			aN, reduce = SnailNumbersReduce(aN)
			fmt.Println(a.String())
			fmt.Println(aN)
		}

		if want != a.String() {
			t.Errorf("got %v, want %v", a.String(), want)
		}
	})

	t.Run("complex example", func(t *testing.T) {
		lines := strings.Split(`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`, "\n")

		a, aN := ParseSnailNumbers(lines[0])
		for _, line := range lines[1:] {
			b, bN := ParseSnailNumbers(line)
			a, aN = SnailNumbersAdd(a, b, aN, bN)
			reduce := true
			for reduce {
				aN, reduce = snailNumbersReduce(aN)
			}
			t.Logf(a.String())
		}

		t.Log(a.String())

		want := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"
		if a.String() != want {
			t.Errorf("got %v, want %v", a.String(), want)
		}
	})

	t.Run("complex example step", func(t *testing.T) {
		x, xN := ParseSnailNumbers("[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]")
		y, yN := ParseSnailNumbers("[[[[4,2],2],6],[8,7]]")
		z, zN := SnailNumbersAdd(x, y, xN, yN)

		repeat := true
		for repeat {
			zN, repeat = snailNumbersReduce(zN)
		}

		want := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"
		if z.String() != want {
			t.Errorf("got %v, want %v", z.String(), want)
		}

		t.Log(zN)
	})
}

func TestSnailNumber_Magnitude(t *testing.T) {
	t.Parallel()

	tests := map[string]int{
		"[[1,2],[[3,4],5]]":                                     143,
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]":                     1384,
		"[[[[1,1],[2,2]],[3,3]],[4,4]]":                         445,
		"[[[[3,0],[5,3]],[4,4]],[5,5]]":                         791,
		"[[[[5,0],[7,4]],[5,5]],[6,6]]":                         1137,
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]": 3488,
	}

	for input, ans := range tests {
		t.Run(input, func(t *testing.T) {
			a, _ := ParseSnailNumbers(input)
			if a.Magnitude() != ans {
				t.Errorf("expected ans %d, got %d", ans, a.Magnitude())
			}
		})
	}

}
