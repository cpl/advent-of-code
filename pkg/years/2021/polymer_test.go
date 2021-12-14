package y2021

import (
	"testing"
)

func TestPolymer(t *testing.T) {
	t.Parallel()

	input := `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

	poly := ParsePolymer(input)
	if string(poly.str) != "NNCB" {
		t.Errorf("expected NNCB, got %s", string(poly.str))
	}

	if poly.m["CH"] != 'B' {
		t.Errorf("expected 'B', got '%s'", string(poly.m["CH"]))
	}

	out := poly.Apply()
	if string(out) != "NCNBCHB" {
		t.Errorf("expected 'NCNBCHB', got '%s'", string(out))
	}

	poly.str = out
	out = poly.Apply()
	if string(out) != "NBCCNBBBCBHCB" {
		t.Errorf("expected 'NBCCNBBBCBHCB', got '%s'", string(out))
	}

	poly.str = out
	out = poly.Apply()
	if string(out) != "NBBBCNCCNBBNBNBBCHBHHBCHB" {
		t.Errorf("expected 'NBBBCNCCNBBNBNBBCHBHHBCHB', got '%s'", string(out))
	}

	poly.str = out
	out = poly.Apply()
	if string(out) != "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB" {
		t.Errorf("expected 'NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB', got '%s'", string(out))
	}
}

func TestPolymer_10steps(t *testing.T) {
	t.Parallel()

	input := `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

	poly := ParsePolymer(input)
	for idx := 0; idx < 10; idx++ {
		out := poly.Apply()
		poly.str = out
	}

	minRune, minCount, maxRune, maxCount := poly.Elements()

	t.Logf("min: %d %s", minCount, string(minRune))
	t.Logf("max: %d %s", maxCount, string(maxRune))

	if minCount != 161 || minRune != 'H' {
		t.Errorf("expected 161 H, got %d %s", minCount, string(minRune))
	}
	if maxCount != 1749 || maxRune != 'B' {
		t.Errorf("expected 1749 B, got %d %s", maxCount, string(maxRune))
	}
}

func TestPolymer_40steps(t *testing.T) {
	t.Parallel()

	input := `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

	poly := ParsePolymer(input)
	t.Logf("%+v", poly.pairs)

	if _, ok := poly.pairs["NN"]; !ok {
		t.Errorf("expected 'NN' to be a pair")
	}
	if _, ok := poly.pairs["NC"]; !ok {
		t.Errorf("expected 'NC' to be a pair")
	}
	if _, ok := poly.pairs["CB"]; !ok {
		t.Errorf("expected 'CB' to be a pair")
	}

	for idx := 0; idx < 10; idx++ {
		poly.ApplyPairs()
	}
	t.Logf("%+v", poly.pairs)

	minRune, minCount, maxRune, maxCount := poly.ElementsPairs()

	t.Logf("min: %d %s", minCount, string(minRune))
	t.Logf("max: %d %s", maxCount, string(maxRune))

	if minCount != 161 || minRune != 'H' {
		t.Errorf("expected 161 H, got %d %s", minCount, string(minRune))
	}
	if maxCount != 1749 || maxRune != 'B' {
		t.Errorf("expected 1749 B, got %d %s", maxCount, string(maxRune))
	}

	//for idx := 0; idx < 30; idx++ {
	//	poly.ApplyPairs()
	//}

	//minRune, minCount, maxRune, maxCount = poly.ElementsPairs()
	//
	//t.Logf("min: %d %s", minCount, string(minRune))
	//t.Logf("max: %d %s", maxCount, string(maxRune))
	//
	//if minCount != 3849876073 || minRune != 'H' {
	//	t.Errorf("expected 161 H, got %d %s", minCount, string(minRune))
	//}
	//if maxCount != 2192039569602 || maxRune != 'B' {
	//	t.Errorf("expected 1749 B, got %d %s", maxCount, string(maxRune))
	//}
}
