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
	m := maze.NewMaze()
	pathFinder := internal.NewDFSPathFinder(m)

	from, _ := m.GetPosition("0,0")
	to, _ := m.GetPosition("2,1")
	path, err := pathFinder.Find(from, to)
	if err != nil {
		return err
	}

	for _, p := range path {
		fmt.Println(p)
	}

	return nil
}
