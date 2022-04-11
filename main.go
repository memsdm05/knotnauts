package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"math"
	"math/rand"
	"time"
)

type MamaMiaGame struct {
	nodeset []*Node
	start, goal *Node
}

func randomNode() *Node {
	w, h := ebiten.WindowSize()
	return NewNode(rand.Float64() * float64(w), rand.Float64() * float64(h))
}

func (m MamaMiaGame) find(x, y float64) *Node {
	var (
		minDist = math.Inf(1)
		minNode *Node
	)
	for _, node := range m.nodeset {
		if dist := node.dist(x, y); dist < minDist {
			minDist = dist
			minNode = node
		}
	}
	return minNode
}

func (m *MamaMiaGame) Init(n int)  {
	m.start = randomNode()
	m.start.Status = Start
	m.start.pathPoint = true

	m.goal = randomNode()
	m.goal.Status = End

	m.nodeset = []*Node{ m.start, m.goal }

	for i := 0; i < n - 2; i++ {
		m.nodeset = append(m.nodeset, randomNode())
	}

	for _, node := range m.nodeset {
		for i := 0; i < 2; i++ {
			rando := m.nodeset[rand.Intn(len(m.nodeset))]
			if !node.NeighborsWith(rando) {
				Link(node, rando)
			}
		}
	}

	AStar(m.start, m.goal)
}

func (m MamaMiaGame) Update() error {
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

	game := new(MamaMiaGame)
	game.Init(20)

	fmt.Println(game.nodeset[0])
	ebiten.RunGame(game)


	//path := AStar(Pos{}, Pos{}, maze)
	//fmt.Println(path)
}

