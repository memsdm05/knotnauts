package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
	"strings"
)

type Pos struct {
	X, Y int
}

func randWalls() []Wall {
	walls := []Wall{ N, S, E, W }
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

var EveryWall = []Wall{ N, S, E, W }

func (w Wall) Has(other Wall) bool {
	return w & other != 0
}

//func (w Wall) Neighbors(at Pos, walls Wall) []Pos {
//	var neighbors []Pos
//	for _, currentWall := range []Wall{ N, S, E, W } {
//		if !walls.Has(currentWall) {
//			continue
//		}
//
//		neighbor := Pos{}
//		neighbor = currentWall.Delta(at)
//		neighbors = append(neighbors, neighbor)
//	}
//
//	return neighbors
//}

type Piece int
const(
	Empty = iota
	Path
	Goal
)

var Opposite = map[Wall]Wall {
	N: S,
	S: N,
	E: W,
	W: E,
}

func (w Wall) Delta(p Pos) Pos {
	switch w {
	case N:
		p.Y--
	case S:
		p.Y++
	case E:
		p.X++
	case W:
		p.X--
	}
	return p
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
	m.carve(Pos{})
	return m
}

func (m Maze) Width() int {
	return len(m[0])
}

func (m Maze) Height() int {
	return len(m)
}

func (m Maze) WithinBounds(pos Pos) bool {
	return pos.X >= 0 && pos.X < m.Width() && pos.Y >= 0 && pos.Y < m.Height()
}

func (m *Maze) carve(cpos Pos) {
	maze := *m
	cx, cy := cpos.X, cpos.Y
	for _, wall := range randWalls() {
		npos := wall.Delta(cpos)
		nx, ny := npos.X, npos.Y
		if !(maze.WithinBounds(npos) && maze[ny][nx].Wall == 0) {
			continue
		}

		maze[cy][cx].Wall |= wall
		maze[ny][nx].Wall |= Opposite[wall]
		maze.carve(npos)
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

func (m Maze) Draw(dst *ebiten.Image) {
	width, height := dst.Size()
	dw, dh := float64(width / m.Width()), float64(height / m.Height())
	for i := 0; i < m.Height(); i++ {
		for j := 0; j < m.Width(); j++ {
			tile := m[i][j]
			cx, cy := float64(j) * dw, float64(i) * dh

			// draw tiles
			var c color.Color
			switch tile.Piece {
			case Empty:
				c = colornames.White
			case Path:
				c = colornames.Red
			case Goal:
				c = colornames.Green
			}

			_ = c
			//if (i + j) % 2 == 0 {
			//	c = colornames.Purple
			//}
			//ebitenutil.DrawRect(dst, cx, cy, dw, dh, c)

			//fmt.Printf("%b\n", tile.Wall)

			if tile.Has(S) {
				ebitenutil.DrawLine(dst, cx, cy + dh, cx + dw, cy + dh, colornames.White)
			}

			if tile.Has(E) {
				ebitenutil.DrawLine(dst, cx + dw, cy, cx + dw, cy + dh, colornames.White)
			}




			// draw walls
			//for _, wall := range EveryWall {
			//	if !tile.Has(wall) { continue }
			//	//t := wall.Delta(Pos{j, i})
			//	//x, y := float64(t.X) * dw, float64(t.Y) * dh
			//
			//	// OH GOD
			//	black := colornames.Black
			//	switch wall {
			//	case N:
			//		ebitenutil.DrawLine(dst, cx, cy, cx + dw, cy, colornames.Green)
			//	case S:
			//		ebitenutil.DrawLine(dst, cx, cy + dh, cx + dw, cy + dh, colornames.Black)
			//	case E:
			//		ebitenutil.DrawLine(dst, cx + dw, cy, cx + dw, cy + dh, black)
			//	case W:
			//		ebitenutil.DrawLine(dst, cx, cy, cx, cy + dh, black)
			//	}
			//
			//}
		}
	}
}

