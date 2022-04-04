package main

import (
	"container/heap"
	"fmt"
	"math"
)


func h(here, goal Pos) float64 {
	return math.Abs(float64(here.X - goal.X)) + math.Abs(float64(here.Y - goal.Y))
}

func g(here, start Pos) (sum float64) {
	for here.prev != nil {
		sum += here.fscore
		here = *here.prev
	}
	return
}

type Pos struct {
	X, Y int
	prev *Pos

	fscore float64
	gScore float64
}

type PosMinHeap []Pos

func (p PosMinHeap) Len() int {
	return len(p)
}

func (p PosMinHeap) Less(i, j int) bool {
	return p[i].fscore < p[j].fscore
}

func (p PosMinHeap) Swap(i, j int) {
	p[i],  p[j] = p[j],  p[i]
}

func (p *PosMinHeap) Push(x interface{}) {
	*p = append(*p, x.(Pos))
}

func (p *PosMinHeap) Pop() interface{} {
	old := *p
	i := len(old) - 1
	ret := old[i]
	*p = old[:i]
	return ret
}


func AStar(start, end Pos, maze Maze) {
	openSet := &PosMinHeap{ start }
	heap.Init(openSet)

	fmt.Println(Pos{X: 10, Y: 0} == Pos{X: 0, Y: 0})

	heap.Push(openSet, Pos{fscore: 10})
	heap.Push(openSet, Pos{fscore: 3})
	heap.Push(openSet, Pos{fscore: 2})

	for i := 0; i < 4; i++ {
		fmt.Println(heap.Pop(openSet))
	}

	//gScore := math.Inf(1)
	//_ = gScore

	for len(*openSet) > 0 {
		current := heap.Pop(openSet).(Pos)
		if current == end {
			break // return path
		}

	}
}