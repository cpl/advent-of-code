package y2021

import (
	"strings"
	"testing"
)

func TestParseCaveGraph(t *testing.T) {
	t.Parallel()

	input := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

	graph := ParseCaveGraph(input)
	if len(graph.nodes) != 6 {
		t.Errorf("expected 6 nodes, got %d", len(graph.nodes))
	}

	if _, ok := graph.nodes["start"].neighbors["A"]; !ok {
		t.Errorf("expected node A to be a neighbor of start")
	}
	if _, ok := graph.nodes["start"].neighbors["b"]; !ok {
		t.Errorf("expected node b to be a neighbor of start")
	}
	if _, ok := graph.nodes["A"].neighbors["c"]; !ok {
		t.Errorf("expected node c to be a neighbor of A")
	}
	if _, ok := graph.nodes["A"].neighbors["b"]; !ok {
		t.Errorf("expected node b to be a neighbor of A")
	}
	if _, ok := graph.nodes["b"].neighbors["d"]; !ok {
		t.Errorf("expected node d to be a neighbor of b")
	}
	if _, ok := graph.nodes["A"].neighbors["end"]; !ok {
		t.Errorf("expected node end to be a neighbor of A")
	}
	if _, ok := graph.nodes["b"].neighbors["end"]; !ok {
		t.Errorf("expected node end to be a neighbor of b")
	}
	if _, ok := graph.nodes["end"].neighbors["b"]; !ok {
		t.Errorf("expected node b to be a neighbor of end")
	}
	if _, ok := graph.nodes["end"].neighbors["A"]; !ok {
		t.Errorf("expected node A to be a neighbor of end")
	}
}

func testCaveGraphWalkBFS(t *testing.T, input string, wantPaths int, useBFS2 bool) {
	graph := ParseCaveGraph(input)

	if useBFS2 {
		graph.WalkBFS2()
	} else {
		graph.WalkBFS()
	}

	if len(graph.paths) != wantPaths {
		for _, path := range graph.paths {
			t.Logf("path: %v", strings.Join(path, " -> "))
		}

		t.Errorf("expected %d paths, got %d", wantPaths, len(graph.paths))
	}
}

func TestCaveGraph_WalkBFS(t *testing.T) {
	t.Parallel()

	t.Run("small", func(t *testing.T) {
		input := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

		testCaveGraphWalkBFS(t, input, 10, false)
	})

	t.Run("medium", func(t *testing.T) {
		input := `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

		testCaveGraphWalkBFS(t, input, 19, false)
	})

	t.Run("large", func(t *testing.T) {
		input := `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

		testCaveGraphWalkBFS(t, input, 226, false)
	})
}

func TestCaveGraph_WalkBFS2(t *testing.T) {
	t.Run("small", func(t *testing.T) {
		input := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

		testCaveGraphWalkBFS(t, input, 36, true)
	})

	t.Run("medium", func(t *testing.T) {
		input := `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

		testCaveGraphWalkBFS(t, input, 103, true)
	})

	t.Run("large", func(t *testing.T) {
		input := `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

		testCaveGraphWalkBFS(t, input, 3509, true)
	})
}
