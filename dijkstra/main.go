package main

import (
	"fmt"

	"github.com/mateoferrari97/pathfinder/dijkstra/internal"
	"github.com/mateoferrari97/pathfinder/internal/maze"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	m := maze.NewMaze()
	pathFinder := internal.NewDijkstraPathFinder(m)

	path, err := pathFinder.Find(m.GetPosition("0,0"), m.GetPosition("15,24"))
	if err != nil {
		return err
	}

	for _, p := range path {
		fmt.Println(p)
	}

	return nil
}
