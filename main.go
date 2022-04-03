package main

import (
	"fmt"
	"math/rand"
	"time"
)

var maze = NewMaze(10, 10)

func randWalls() []Wall {
	walls := []Wall{N, S, E, W}
	rand.Shuffle(4, func(i, j int) {
		walls[i], walls[j] = walls[j], walls[i]
	})
	return walls
}

func carve(cx, cy int) {
	for _, wall := range randWalls() {
		nx, ny := wall.Delta(cx, cy)
		if !(maze.WithinBounds(nx, ny) && maze[ny][nx] == 0) {
			continue
		}

		maze[cy][cx] |= wall
		maze[ny][nx] |= Opposite[wall]
		carve(nx, ny)
	}
}

func main()  {
	rand.Seed(time.Now().UnixNano())
	carve(0, 0)
	fmt.Println(maze)
}

