package y2021

import (
	"math"
	"strings"
)

type CaveMap struct {
	w, h  int
	nodes [][]*caveMapNode
}

type caveMapNode struct {
	distance int
	weight   int

	children []*caveMapNode
}

func ParseCaveMap(input string, largeMode bool) *CaveMap {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	h := len(lines)
	w := len(lines[0])

	nw := w
	nh := h
	if largeMode {
		h = h * 5
		w = w * 5
	}

	caveMap := &CaveMap{
		nodes: make([][]*caveMapNode, h),
		w:     w,
		h:     h,
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			weight := int(lines[y%nh][x%nw] - '0')
			if y >= len(lines) || x >= len(lines[0]) {
				extra := 0
				extra += y / len(lines)
				extra += x / len(lines[0])

				weight += extra
				if weight >= 10 {
					weight = weight%10 + 1
				}
			}

			caveMap.nodes[y] = append(caveMap.nodes[y], &caveMapNode{
				distance: math.MaxInt,
				weight:   weight,
			})
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x < w-1 {
				caveMap.nodes[y][x].children = append(caveMap.nodes[y][x].children, caveMap.nodes[y][x+1])
			}
			if y < h-1 {
				caveMap.nodes[y][x].children = append(caveMap.nodes[y][x].children, caveMap.nodes[y+1][x])
			}
		}
	}

	caveMap.nodes[0][0].distance = 0

	return caveMap
}

func (c *CaveMap) String() string {
	var builder strings.Builder
	builder.Grow(len(c.nodes) * len(c.nodes[0]))

	for _, row := range c.nodes {
		for _, node := range row {
			builder.WriteString(string(rune(node.weight) + '0'))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func (c *CaveMap) AStar() {
	root := c.nodes[0][0]
	goal := c.nodes[c.h-1][c.w-1]

	queue := []*caveMapNode{root}
	cameFrom := map[*caveMapNode]*caveMapNode{}
	cost := map[*caveMapNode]int{}

	cameFrom[root] = nil
	cost[root] = 0

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current == goal {
			break
		}

		for _, child := range current.children {
			newCost := cost[current] + child.weight
			if _, ok := cost[child]; !ok || newCost < cost[child] {
				cost[child] = newCost
				child.distance = newCost
				cameFrom[child] = current
				queue = append(queue, child)
			}
		}
	}
}
