package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"math"
)

const (
	Regular = iota
	Start
	End
)

var (
	emptyImage    = ebiten.NewImage(3, 3)
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	emptyImage.Fill(color.White)
}

func Link(n1, n2 *Node) {
	if !n1.NeighborsWith(n2) {
		n1.Neighbors = append(n1.Neighbors, n2)
	}

	if !n2.NeighborsWith(n1) {
		n2.Neighbors = append(n2.Neighbors, n1)
	}
}


type Node struct {
	X, Y float64
	Neighbors []*Node
	Status int

	// for A*
	fScore, gScore float64
	prev *Node
	pathPoint bool

	// for graphics
	path vector.Path
}

func (n *Node) Reset() {
	n.fScore = math.Inf(1)
	n.gScore = math.Inf(1)
	n.prev = nil
	n.pathPoint = false
}

func (n *Node) NeighborsWith(other *Node) bool {
	for _, neighbor := range n.Neighbors {
		if other == neighbor {
			return true
		}
	}
	return false
}

func NewNode(x, y float64) *Node {
	return &Node{
		X: x,
		Y: y,
		fScore: math.Inf(1),
		gScore: math.Inf(1),
		prev:   nil,
	}
}

func (n *Node) dist(x, y float64) float64{
	dx, dy := n.X - x, n.Y - y
	return math.Sqrt(dx * dx + dy * dy)
}

func (n *Node) Dist(other *Node) float64 {
	return other.dist(other.X, other.Y)
}

func drawCircle(dst *ebiten.Image, x, y, radius float64, clr color.Color) {
	r, g, b, _ := clr.RGBA()
	var path vector.Path
	path.Arc(float32(x), float32(y),
		float32(radius), 0, 2 * math.Pi, vector.Clockwise)


	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}
	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = float32(r) / float32(0xffff)
		vs[i].ColorG = float32(g) / float32(0xffff)
		vs[i].ColorB = float32(b) / float32(0xffff)
	}
	dst.DrawTriangles(vs, is, emptySubImage, op)
}

func (n *Node) Draw(dst *ebiten.Image) {
	clr := colornames.Purple
	switch n.Status {
	case Regular:
		clr = colornames.Orange
	case Start:
		clr = colornames.White
	case End:
		clr = colornames.Lightgreen
	}

	drawCircle(dst, n.X, n.Y, 10, clr)
}


