package y2021

import (
	"testing"
)

func TestParseTransparentPaper(t *testing.T) {
	t.Parallel()

	input := `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

	tp := ParseTransparentPaper(input)

	tp.Fold()
	if tp.Displayed() != 17 {
		t.Errorf("expected 17, got %d", tp.Displayed())
	}

	tp.Fold()
	if tp.Displayed() != 16 {
		t.Errorf("expected 16, got %d", tp.Displayed())
	}

	tp.Print()
}
