package main

import (
	"github.com/fogleman/delaunay"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
	"math/rand"
	"time"
)

type MamaMiaGame struct {
	nodeset []*Node
	start, goal *Node
	selected *Node
}

func randomNode() *Node {
	w, h := ebiten.WindowSize()
	return NewNode(rand.Float64() * float64(w), rand.Float64() * float64(h))
}

func (m *MamaMiaGame) RecalculateDelaunay() {
	vertexset := make([]delaunay.Point, len(m.nodeset))
	for i, node := range m.nodeset {
		node.Neighbors = nil
		vertexset[i].X = node.X
		vertexset[i].Y = node.Y
	}

	triangulation, err := delaunay.Triangulate(vertexset)
	if err != nil {
		panic(err)
	}

	ts := triangulation.Triangles
	hs := triangulation.Halfedges
	for i, h := range hs {
		if i > h {
			n1 := m.nodeset[ts[i]]
			n2 := m.nodeset[ts[nextHE(i)]]
			Link(n1, n2)
		}
	}
}

func (m MamaMiaGame) find(x, y float64) *Node {
	for _, node := range m.nodeset {
		if dist := node.dist(x, y); dist < 10 {
			return node
		}
	}
	return nil
}

func nextHE(e int) int {
	if e % 3 == 2 {
		return e - 2
	}
	return e + 1
}

func (m *MamaMiaGame) Init(n int)  {
	m.start = randomNode()
	m.start.Status = Start
	m.start.pathPoint = true

	for i := 0; i < 10; i++ {
		rand.Float64() // stop spawning rig
	}
	m.goal = randomNode()
	m.goal.Status = End

	// build nodes
	m.nodeset = []*Node{ m.start, m.goal }
	for i := 0; i < n - 2; i++ {
		m.nodeset = append(m.nodeset, randomNode())
	}

	// trianglate
	m.recalc()
}

func (m *MamaMiaGame) recalc()  {
	m.RecalculateDelaunay()
	for _, node := range m.nodeset {
		node.Reset()
	}
	m.start.pathPoint = true
	AStar(m.start, m.goal)
}

func (m *MamaMiaGame) Update() error {
	x, y := ebiten.CursorPosition()
	mx, my := float64(x), float64(y)
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if node := m.find(mx, my); node != nil {
			m.selected = node
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && m.selected != nil {
		m.recalc()
		m.selected.X = mx
		m.selected.Y = my
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		m.selected = nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		m.nodeset = append(m.nodeset, NewNode(mx, my))
		m.recalc()
	}

	if ebiten.IsKeyPressed(ebiten.KeyDelete) || ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		if node := m.find(mx, my); node != nil && len(m.nodeset) > 3{
			if node == m.start || node == m.goal {
				goto nope // lol
			}

			for i := 0; i < len(m.nodeset); i++ {
				if m.nodeset[i] == node {
					m.nodeset[i] = m.nodeset[len(m.nodeset) - 1]
					m.nodeset = m.nodeset[:len(m.nodeset) - 1]
					m.recalc()
					break
				}
			}

			nope:
		}
	}

	return nil
}

func (m MamaMiaGame) Draw(screen *ebiten.Image) {
	for _, node := range m.nodeset {
		for _, neighbor := range node.Neighbors {
			clr := colornames.White
			if node.pathPoint && neighbor.pathPoint {
				clr = colornames.Hotpink
			}
			ebitenutil.DrawLine(screen,
				node.X, node.Y, neighbor.X, neighbor.Y, clr)
		}

		node.Draw(screen)
	}
}

func (m MamaMiaGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main()  {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(500, 500)
	ebiten.SetWindowResizable(true)
	game := new(MamaMiaGame)
	game.Init(20)
	ebiten.RunGame(game)


	//path := AStar(Pos{}, Pos{}, maze)
	//fmt.Println(path)
}

