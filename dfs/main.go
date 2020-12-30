package main

import (
	"fmt"

	"github.com/mateoferrari97/pathfinder/dfs/internal"
	"github.com/mateoferrari97/pathfinder/internal/maze"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	maze := maze.NewMaze()
	pathFinder := internal.NewDFSPathFinder(maze)

	path, err := pathFinder.Find(maze.GetPosition("1,1"), maze.GetPosition("0,0"))
	if err != nil {
		return err
	}

	for _, p := range path {
		fmt.Println(p)
	}

	return nil
}
