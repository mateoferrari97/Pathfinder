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
	m := maze.NewMaze()
	pathFinder := internal.NewBFSPathFinder(m)

	from, _ := m.GetPosition("5,0")
	to, _ := m.GetPosition("24,5")
	path, err := pathFinder.Find(from, to)
	if err != nil {
		return err
	}

	for _, p := range path {
		fmt.Println(p)
	}

	return nil
}
