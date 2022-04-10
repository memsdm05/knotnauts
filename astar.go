package main

import (
	"container/heap"
	"math"
)


func h(from, to Pos) float64 {
	return math.Abs(float64(from.X - to.X)) + math.Abs(float64(from.Y - to.Y))
}

type node struct {
	Pos

	fScore float64
	gScore float64
	prev *node
}

type NodeMinHeap []*node

func (n NodeMinHeap) has(node *node) bool {
	for _, check := range n {
		if check == node {
			return true
		}
	}
	return false
}

func (n NodeMinHeap) Len() int {
	return len(n)
}

func (n NodeMinHeap) Less(i, j int) bool {
	return n[i].fScore < n[j].fScore
}

func (n NodeMinHeap) Swap(i, j int) {
	n[i],  n[j] = n[j],  n[i]
}

func (n *NodeMinHeap) Push(x interface{}) {
	*n = append(*n, x.(*node))
}

func (n *NodeMinHeap) Pop() interface{} {
	old := *n
	i := len(old) - 1
	ret := old[i]
	*n = old[:i]
	return ret
}

var graph = map[Pos]*node{}

func push(n *node) *node {
	if n == nil {
		return nil
	}
	graph[n.Pos] = n
	return n
}

func maybeGet(pos Pos, prev *node) *node {
	if n, ok := graph[pos]; ok {
		return n
	} else {
		return push(newNode(pos, prev))
	}
}

func newNode(pos Pos, prev *node) *node {
	return &node{
		Pos:    pos,
		fScore: math.Inf(1),
		gScore: math.Inf(1),
		prev:   prev,
	}
}

func AStar(start, end Pos, maze Maze) (path []Pos) {
	// init start node
	startNode := push(newNode(start, nil))
	startNode.gScore = 0
	startNode.fScore = h(startNode.Pos, end)

	// init heap
	openSet := &NodeMinHeap{ startNode }
	heap.Init(openSet)

	for len(*openSet) > 0 {
		current := heap.Pop(openSet).(*node)
		if current.Pos == end {
			for current.prev != nil {
				path = append(path, current.Pos)
				current = current.prev
			}
			return
		}

		tile := maze[current.Y][current.X]
		for _, wall := range EveryWall {
			if tile.Has(wall) {
				continue
			}
			npos := wall.Delta(current.Pos)
			neighbor := maybeGet(npos, current)

			tgScore := current.gScore + 1 // delta is always 1
			if tgScore < neighbor.gScore {
				neighbor.prev = current
				neighbor.gScore = tgScore
				neighbor.fScore = tgScore + h(npos, end)

				if !openSet.has(neighbor) {
					openSet.Push(neighbor)
				}
			}
		}
	}

	return
}