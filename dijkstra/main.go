package main

import (
	"fmt"
	"math/rand"

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
	r := rand.New(rand.NewSource(100))
	pathFinder := internal.NewDijkstraPathFinder(m, r)

	from, _ := m.GetPosition("1,1")
	to, _ := m.GetPosition("22,24")
	path, err := pathFinder.Find(from, to)
	if err != nil {
		return err
	}

	for _, p := range path {
		fmt.Println(p)
	}

	return nil
}
