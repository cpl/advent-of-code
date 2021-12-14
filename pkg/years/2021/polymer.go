package y2021

import (
	"math"
	"strings"
)

type Polymer struct {
	str   []rune
	m     map[string]rune
	pairs map[string]int
}

func ParsePolymer(input string) *Polymer {
	p := &Polymer{
		m: make(map[string]rune),
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	p.str = []rune(lines[0])

	pairs := make(map[string]int)
	for idx := 0; idx < len(lines[0])-1; idx++ {
		pairs[string([]rune{p.str[idx], p.str[idx+1]})]++
	}
	p.pairs = pairs

	for _, line := range lines[2:] {
		parts := strings.Split(strings.TrimSpace(line), " -> ")
		p.m[parts[0]] = rune(parts[1][0])
	}

	return p
}

func (poly *Polymer) Apply() []rune {
	out := make([]rune, 0, len(poly.str)*2)

	for idx := 0; idx < len(poly.str)-1; idx++ {
		out = append(out, poly.str[idx])
		if n, ok := poly.m[string(poly.str[idx:idx+2])]; ok {
			out = append(out, n)
		}
	}
	out = append(out, poly.str[len(poly.str)-1])

	return out
}

func (poly *Polymer) Elements() (rune, int, rune, int) {
	minCount := math.MaxInt
	minRune := ' '
	maxCount := 0
	maxRune := ' '

	m := make(map[rune]int)
	for _, r := range poly.str {
		m[r]++
	}

	for r, count := range m {
		if count < minCount {
			minCount = count
			minRune = r
		}
		if count > maxCount {
			maxCount = count
			maxRune = r
		}
	}

	return minRune, minCount, maxRune, maxCount
}

// NNCB (NN, NC, CB)
// NCNBCHB (NC, CN, NB, BC, CH, HB)

func (poly *Polymer) ApplyPairs() {
	newPairs := make(map[string]int)
	for pair, count := range poly.pairs {
		l, r := rune(pair[0]), rune(pair[1])
		n := poly.m[pair]

		newPairs[string([]rune{l, n})] += count
		newPairs[string([]rune{n, r})] += count
	}
	poly.pairs = newPairs
}

func (poly *Polymer) ElementsPairs() (rune, int, rune, int) {
	m := make(map[rune]int)

	minCount := math.MaxInt
	minRune := ' '
	maxCount := 0
	maxRune := ' '

	for pair, count := range poly.pairs {
		l, r := rune(pair[0]), rune(pair[1])
		m[l] += count
		m[r] += count
	}

	for r, count := range m {
		if count < minCount {
			minCount = count
			minRune = r
		}
		if count > maxCount {
			maxCount = count
			maxRune = r
		}
	}

	if minCount%2 != 0 {
		minCount++
	}
	if maxCount%2 != 0 {
		maxCount++
	}

	return minRune, minCount / 2, maxRune, maxCount / 2
}
