package main

import (
	"fmt"
	//"math/rand"
	//"time"
)


type MazeAstarGame struct {}

func main()  {
	//rand.Seed(time.Now().UnixNano())
	maze := NewMaze(10, 10)
	fmt.Println(maze)
}

