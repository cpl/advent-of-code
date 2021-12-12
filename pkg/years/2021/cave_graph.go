package y2021

import (
	"strings"
)

type CaveGraph struct {
	nodes map[string]*CaveNode
	paths [][]string
}

func (graph *CaveGraph) WalkBFS() {
	start := graph.nodes["start"]
	start.walkBFS(0, []string{"start"})
}

func (graph *CaveGraph) WalkBFS2() {
	start := graph.nodes["start"]
	start.walkBFS2(0, []string{"start"}, false)
}

type CaveNode struct {
	graph     *CaveGraph
	name      string
	neighbors map[string]*CaveNode
	walked    int
}

func (node *CaveNode) walkBFS(depth int, path []string) {
	if node.name == "end" {
		node.graph.paths = append(node.graph.paths, path)
		return
	}
	node.walked++

	for _, n := range node.neighbors {
		// do not walk back to start
		if n.name == "start" {
			continue
		}

		// do not walk small nodes twice
		if !n.IsBig() && n.walked > 0 {
			continue
		}

		n.walkBFS(depth+1, append(path, n.name))
	}

	node.walked--
}

func (node *CaveNode) walkBFS2(depth int, path []string, walkedSmallTwice bool) {
	if node.name == "end" {
		node.graph.paths = append(node.graph.paths, path)
		return
	}
	node.walked++

	for _, n := range node.neighbors {
		// do not walk back to start
		if n.name == "start" {
			continue
		}

		if !n.IsBig() {
			if n.walked > 0 {
				if walkedSmallTwice {
					continue
				}

				if n.walked >= 2 {
					continue
				}

				n.walkBFS2(depth+1, append(path, n.name), true)
			} else {
				n.walkBFS2(depth+1, append(path, n.name), walkedSmallTwice)
			}
		} else {
			n.walkBFS2(depth+1, append(path, n.name), walkedSmallTwice)
		}
	}

	node.walked--
}

func (node *CaveNode) IsBig() bool {
	if node.name == "start" || node.name == "end" {
		return true
	}

	for _, r := range node.name {
		if r < 'A' || r > 'Z' {
			return false
		}
	}

	return true
}

func ParseCaveGraph(input string) *CaveGraph {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	graph := &CaveGraph{
		nodes: make(map[string]*CaveNode),
	}

	for _, line := range lines {
		split := strings.Split(line, "-")
		from := split[0]
		to := split[1]

		fromNode, ok := graph.nodes[from]
		if !ok {
			fromNode = &CaveNode{
				name:      from,
				neighbors: make(map[string]*CaveNode),
				graph:     graph,
			}
			graph.nodes[from] = fromNode
		}

		toNode, ok := graph.nodes[to]
		if !ok {
			toNode = &CaveNode{
				name:      to,
				neighbors: make(map[string]*CaveNode),
				graph:     graph,
			}
			graph.nodes[to] = toNode
		}

		fromNode.neighbors[to] = toNode
		toNode.neighbors[from] = fromNode
	}

	return graph
}
