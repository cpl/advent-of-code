package y2021

import (
	"sort"
	"strings"
)

var mapSegmentToDigit = map[string]string{
	"abcefg":  "0",
	"cf":      "1",
	"acdeg":   "2",
	"acdfg":   "3",
	"bcdf":    "4",
	"abdfg":   "5",
	"abdefg":  "6",
	"acf":     "7",
	"abcdefg": "8",
	"abcdfg":  "9",
}

func Parse7Segments(input []string) *SevenSegments {
	segmentsByLen := map[int][]*SevenSegment{
		2: {},
		3: {},
		4: {},
		7: {},
		5: {},
		6: {},
	}
	output := make([]*SevenSegment, 0, len(input))
	digits := make([]*SevenSegment, len(input))
	for _, segmentStr := range input {
		m := make(map[string]bool)
		for _, r := range segmentStr {
			m[string(r)] = true
		}

		segment := &SevenSegment{
			str:      segmentStr,
			segments: m,
		}

		output = append(output, segment)

		segmentsByLen[len(segmentStr)] = append(segmentsByLen[len(segmentStr)], segment)
		switch len(segmentStr) {
		case 2:
			digits[1] = segment
		case 3:
			digits[7] = segment
		case 4:
			digits[4] = segment
		case 7:
			digits[8] = segment
		}
	}

	return &SevenSegments{
		segments: output,
		byDigit:  digits,
		byLen:    segmentsByLen,
	}
}

func Parse7Segment(str string) *SevenSegment {
	m := make(map[string]bool)
	for _, r := range str {
		m[string(r)] = true
	}
	return &SevenSegment{
		str:      str,
		segments: m,
	}
}

type SevenSegment struct {
	str      string
	segments map[string]bool
}

func (s *SevenSegment) String() string {
	return s.str
}

func (s *SevenSegment) Len() int {
	return len(s.str)
}

func (s *SevenSegment) Diff(other *SevenSegment) map[string]int {
	m := make(map[string]int)
	for r := range s.segments {
		m[r]++
	}
	for r := range other.segments {
		m[r]--
	}
	return m
}

func (s *SevenSegment) Sub(other *SevenSegment) *SevenSegment {
	m := make(map[string]bool)
	for r := range s.segments {
		if other.segments[r] {
			continue
		}

		m[r] = true
	}

	str := ""
	for r := range m {
		str += r
	}

	segment := &SevenSegment{
		str:      str,
		segments: m,
	}

	return segment
}

func (s *SevenSegment) Add(other *SevenSegment) *SevenSegment {
	m := make(map[string]bool)
	for r := range s.segments {
		m[r] = true
	}
	for r := range other.segments {
		m[r] = true
	}

	str := ""
	for r := range m {
		str += r
	}

	return &SevenSegment{
		str:      str,
		segments: m,
	}
}

type SevenSegments struct {
	segments []*SevenSegment
	byLen    map[int][]*SevenSegment
	byDigit  []*SevenSegment
}

func (s *SevenSegments) SolveMapping() map[rune]rune {
	mapReal := make(map[string]*SevenSegment)
	mapReal["a"] = s.byDigit[7].Sub(s.byDigit[1])

	for _, segment := range s.byLen[6] {
		mapReal["g"] = segment.Sub(s.byDigit[4]).Sub(mapReal["a"])
		if mapReal["g"].Len() == 1 {
			s.byDigit[9] = segment
			break
		}
	}

	for _, segment := range s.byLen[5] {
		mapReal["d"] = segment.Sub(s.byDigit[7].Add(mapReal["g"]))
		if mapReal["d"].Len() == 1 {
			s.byDigit[3] = segment
			break
		}
	}

	mapReal["b"] = s.byDigit[9].Sub(s.byDigit[3])
	mapReal["e"] = s.byDigit[8].Sub(s.byDigit[4].Add(mapReal["a"]).Add(mapReal["g"]))

	find2 := mapReal["a"].Add(mapReal["d"]).Add(mapReal["e"]).Add(mapReal["g"])
	for _, segment := range s.byLen[5] {
		mapReal["c"] = segment.Sub(find2)
		if mapReal["c"].Len() == 1 {
			s.byDigit[2] = segment
			break
		}
	}

	mapReal["f"] = s.byDigit[1].Sub(mapReal["c"])

	out := make(map[rune]rune)
	for k, v := range mapReal {
		out[rune(v.String()[0])] = rune(k[0])
	}

	return out
}

func Parse7SegmentOutput(output []string, mapping map[rune]rune) string {
	var build strings.Builder
	for _, str := range output {
		converted := make([]rune, len(str))
		for idx, r := range str {
			converted[idx] = mapping[r]
		}

		sort.Slice(converted, func(i, j int) bool {
			return converted[i] < converted[j]
		})

		build.WriteString(mapSegmentToDigit[string(converted)])
	}
	return build.String()
}
