package internal

import (
	"github.com/mateoferrari97/pathfinder/internal/maze"
)

type position struct {
	parent  *position
	current *maze.Position
}

type DFSPathFinder struct {
	queue *Queue
	maze  *maze.Maze
}

func NewBFSPathFinder(maze *maze.Maze) *DFSPathFinder {
	stack := NewQueue()
	return &DFSPathFinder{
		queue: stack,
		maze:  maze,
	}
}

func (f *DFSPathFinder) Find(startPosition *maze.Position, endPosition *maze.Position) ([]string, error) {
	if err := f.maze.Validate(*startPosition, *endPosition); err != nil {
		return nil, err
	}

	return f.find(startPosition, endPosition)
}

func (f *DFSPathFinder) find(startPosition *maze.Position, endPosition *maze.Position) ([]string, error) {
	var positions []*maze.Position
	previousPositions := make(map[string]struct{})

	f.queue.Push(&position{
		parent:  nil,
		current: startPosition,
	})

	for !f.queue.IsEmpty() {
		pos := f.queue.Pop()

		if _, seen := previousPositions[pos.current.String()]; seen {
			continue
		}

		previousPositions[pos.current.String()] = struct{}{}

		if endPosition.Equal(*pos.current) {
			currentPosition := pos
			for currentPosition.parent != nil {
				positions = append(positions, currentPosition.current)
				currentPosition = currentPosition.parent
			}

			positions = append(positions, currentPosition.current)
			break
		}

		neighbours := f.maze.Neighbours(pos.current)
		for _, neighbour := range neighbours {
			next := &position{
				parent:  pos,
				current: neighbour,
			}

			f.queue.Push(next)
		}
	}

	path := make([]string, len(positions)+1)
	for _, position := range positions {
		position.WithValue(1)
		path = append(path, f.maze.String())
	}

	return path, nil
}
