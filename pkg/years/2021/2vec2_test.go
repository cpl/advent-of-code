package y2021

import (
	"testing"
)

func TestParse2Vec2(t *testing.T) {
	t.Parallel()

	t.Run("example", func(t *testing.T) {
		input := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

		vectors := Parse2Vec2([]byte(input))
		if len(vectors) != 10 {
			t.Fatalf("expected 10, got %d", len(vectors))
		}

		if vectors[2][0].x != 9 || vectors[2][0].y != 4 || vectors[2][1].x != 3 || vectors[2][1].y != 4 {
			t.Fatalf("expected 9,4->3,4 got %v", vectors[2])
		}
	})

	t.Run("mapping", func(t *testing.T) {
		mapper := vec2mapper{data: make(map[string]int)}

		for _, vecs := range [][2]vec2{
			{vec2{0, 9}, vec2{5, 9}},
			{vec2{8, 0}, vec2{0, 8}},
			{vec2{9, 4}, vec2{3, 4}},
			{vec2{2, 2}, vec2{2, 1}},
			{vec2{7, 0}, vec2{7, 4}},
			{vec2{6, 4}, vec2{2, 0}},
			{vec2{0, 9}, vec2{2, 9}},
			{vec2{3, 4}, vec2{1, 4}},
			{vec2{0, 0}, vec2{8, 8}},
			{vec2{5, 5}, vec2{8, 2}},
		} {
			mapper.Map(vecs)
		}

		mapper.Print(12, 12)
		overlapping := mapper.Overlapping()
		if overlapping != 5 {
			t.Fatalf("expected 5, got %d", overlapping)
		}
	})

	t.Run("mapping2", func(t *testing.T) {
		mapper := vec2mapper{data: make(map[string]int)}

		for _, vecs := range [][2]vec2{
			{vec2{0, 9}, vec2{5, 9}},
			{vec2{8, 0}, vec2{0, 8}},
			{vec2{9, 4}, vec2{3, 4}},
			{vec2{2, 2}, vec2{2, 1}},
			{vec2{7, 0}, vec2{7, 4}},
			{vec2{6, 4}, vec2{2, 0}},
			{vec2{0, 9}, vec2{2, 9}},
			{vec2{3, 4}, vec2{1, 4}},
			{vec2{0, 0}, vec2{8, 8}},
			{vec2{5, 5}, vec2{8, 2}},
		} {
			mapper.Map2(vecs)
		}

		mapper.Print(12, 12)
		overlapping := mapper.Overlapping()
		if overlapping != 12 {
			t.Fatalf("expected 12, got %d", overlapping)
		}
	})
}
