package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"strings"
)

func randWalls() []Wall {
	walls := []Wall{N, S, E, W}
	rand.Shuffle(4, func(i, j int) {
		walls[i], walls[j] = walls[j], walls[i]
	})
	return walls
}

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

func (w Wall) Neighbors(at Pos, walls Wall) []Pos {
	var neighbors []Pos
	for _, currentWall := range []Wall{ N, S, E, W } {
		if !walls.Has(currentWall) {
			continue
		}

		neighbor := Pos{}
		neighbor.X, neighbor.Y = currentWall.Delta(at.X, at.Y)
		neighbors = append(neighbors, neighbor)
	}

	return neighbors
}

type Piece int
const(
	Empty = iota
	Path
)

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

type Tile struct {
	Wall
	Piece
}

type Maze [][]Tile

func NewMaze(width, height int) Maze {
	m := make(Maze, height)
	for i := 0; i < height; i++ {
		m[i] = make([]Tile, width)
	}
	m.carve(0, 0)
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

func (m *Maze) carve(cx, cy int) {
	maze := *m
	for _, wall := range randWalls() {
		nx, ny := wall.Delta(cx, cy)
		if !(maze.WithinBounds(nx, ny) && maze[ny][nx].Wall == 0) {
			continue
		}

		maze[cy][cx].Wall |= wall
		maze[ny][nx].Wall |= Opposite[wall]
		maze.carve(nx, ny)
	}
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
	sb.WriteString("  " + strings.Repeat("_", (m.Width() * 2) - 1) + "\n")
	for i := 0; i < m.Height(); i++ {
		sb.WriteString("|")
		for j := 0; j < m.Width(); j++ {
			ifWrite(m[i][j].Has(S), &sb)
			if m[i][j].Has(E) {
				ifWrite((m[i][j].Wall | m[i][j+1].Wall).Has(S), &sb)
			} else {
				sb.WriteString("|")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (m *Maze) Draw(dst *ebiten.Image, sidelen int) {
	w, h := dst.Size()
	_, _ = w, h
}

