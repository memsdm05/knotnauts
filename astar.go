package main

import (
	"container/heap"
)

type NodeMinHeap []*Node

func (n NodeMinHeap) has(node *Node) bool {
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
	*n = append(*n, x.(*Node))
}

func (n *NodeMinHeap) Pop() interface{} {
	old := *n
	i := len(old) - 1
	ret := old[i]
	*n = old[:i]
	return ret
}

func AStar(start, end *Node) {
	h := func(node *Node) float64 {
		return node.Dist(end)
	}

	// init start node
	start.gScore = 0
	start.fScore = h(start)

	// init heap
	openSet := &NodeMinHeap{ start }
	heap.Init(openSet)

	for len(*openSet) > 0 {
		current := heap.Pop(openSet).(*Node)
		if current == end {
			for current.prev != nil {
				current.pathPoint = true
				current = current.prev
			}
			return
		}

		for _, neighbor := range current.Neighbors {

			tgScore := current.gScore + current.Dist(neighbor)
			if tgScore < neighbor.gScore {
				neighbor.prev = current
				neighbor.gScore = tgScore
				neighbor.fScore = tgScore + h(neighbor)

				if !openSet.has(neighbor) {
					openSet.Push(neighbor)
				}
			}
		}
	}

	return
}