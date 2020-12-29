package main

import (
	"fmt"

	"github.com/mateoferrari97/pathfinder/bfs/internal"
	"github.com/mateoferrari97/pathfinder/internal/maze"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	maze := maze.NewMaze()
	pathFinder := internal.NewBFSPathFinder(maze)

	path, err := pathFinder.Find(maze.Position("0,0"), maze.Position("24,24"))
	if err != nil {
		return err
	}

	for _, p := range path {
		fmt.Println(p)
	}

	return nil
}
