package internal

import (
	"github.com/mateoferrari97/pathfinder/internal/maze"
)

type position struct {
	parent  *position
	current *maze.Position
}

type BFSPathFinder struct {
	queue *queue
	maze  *maze.Maze
}

func NewBFSPathFinder(maze *maze.Maze) *BFSPathFinder {
	return &BFSPathFinder{
		queue: NewQueue(),
		maze:  maze,
	}
}

func (f *BFSPathFinder) Find(startPosition *maze.Position, endPosition *maze.Position) ([]string, error) {
	var positions []*maze.Position
	previousPositions := make(map[string]struct{})

	previousPositions[startPosition.String()] = struct{}{}
	f.queue.push(&position{
		parent:  nil,
		current: startPosition,
	})

	for !f.queue.isEmpty() {
		pos := f.queue.pop()

		if endPosition.Equal(*pos.current) {
			currentPosition := pos
			for currentPosition.parent != nil {
				positions = append(positions, currentPosition.current)
				currentPosition = currentPosition.parent
			}

			positions = append(positions, currentPosition.current)
			break
		}

		neighbours := pos.current.GetNeighbours()
		for _, neighbour := range neighbours {
			if _, seen := previousPositions[neighbour.String()]; !seen {
				previousPositions[neighbour.String()] = struct{}{}

				next := &position{
					parent:  pos,
					current: neighbour,
				}

				f.queue.push(next)
			}
		}
	}

	path := make([]string, len(positions)+1)
	for _, position := range positions {
		position.WithValue(1)
		path = append(path, f.maze.String())
	}

	return path, nil
}