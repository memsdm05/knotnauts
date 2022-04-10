package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

const(
	Width = 10
	Height = 10
)

type MazeAstarGame struct {
	maze Maze
}

func (m MazeAstarGame) Update() error {
	return nil
}

func (m MazeAstarGame) Draw(screen *ebiten.Image) {
	m.maze.Draw(screen)
}

func (m MazeAstarGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return m.maze.Width() * 15, m.maze.Height() * 15
}

func main()  {
	//rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowResizable(true) // this stupid
	ebiten.SetWindowSize(500, 500)
	game := &MazeAstarGame{
		maze: NewMaze(Width, Height),
	}

	fmt.Println(game.maze)

	ebiten.RunGame(game)


	//path := AStar(Pos{}, Pos{}, maze)
	//fmt.Println(path)
}

