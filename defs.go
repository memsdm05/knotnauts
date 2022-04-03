package main

import (
	"strings"
)

type Wall uint
const (
	N Wall = 1 << iota
	S
	E
	W
)

func (w Wall) Has(other Wall) bool {
	return w & other != 0
}

var Opposite = map[Wall]Wall {
	N: S,
	S: N,
	E: W,
	W: E,
}

func (w Wall) Delta(x, y int) (int, int) {
	switch w {
	case N:
		y--
	case S:
		y++
	case E:
		x++
	case W:
		x--
	}

	return x, y
}

type Maze [][]Wall

func NewMaze(width, height int) Maze {
	m := make(Maze, height)
	for i := 0; i < height; i++ {
		m[i] = make([]Wall, width)
	}
	return m
}

func (m Maze) Width() int {
	return len(m[0])
}

func (m Maze) Height() int {
	return len(m)
}

func (m Maze) WithinBounds(x, y int) bool {
	return x >= 0 && x < m.Width() && y >= 0 && y < m.Height()
}

func ifWrite(cond bool, sb *strings.Builder) {
	if cond {
		sb.WriteString(" ")
	} else {
		sb.WriteString("_")
	}
}

func (m Maze) String() string {
	var sb strings.Builder
	sb.WriteString("  " + strings.Repeat("_", (maze.Width() * 2) - 1) + "\n")
	for i := 0; i < m.Height(); i++ {
		sb.WriteString("|")
		for j := 0; j < m.Width(); j++ {
			ifWrite(m[i][j].Has(S), &sb)
			if m[i][j].Has(E) {
				ifWrite((m[i][j] | m[i][j+1]).Has(S), &sb)
			} else {
				sb.WriteString("|")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

