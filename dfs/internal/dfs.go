package internal

import (
	"github.com/mateoferrari97/pathfinder/internal/maze"
)

type position struct {
	parent  *position
	current *maze.Position
}

type DFSPathFinder struct {
	stack *stack
	maze  *maze.Maze
}

func NewDFSPathFinder(maze *maze.Maze) *DFSPathFinder {
	stack := NewStack()
	return &DFSPathFinder{
		stack: stack,
		maze:  maze,
	}
}

func (f *DFSPathFinder) Find(startPosition *maze.Position, endPosition *maze.Position) ([]string, error) {
	return f.find(startPosition, endPosition)
}

func (f *DFSPathFinder) find(startPosition *maze.Position, endPosition *maze.Position) ([]string, error) {
	var positions []*maze.Position
	previousPositions := make(map[string]struct{})

	f.stack.push(&position{
		parent:  nil,
		current: startPosition,
	})

	for !f.stack.isEmpty() {
		pos := f.stack.pop()
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

		neighbours := pos.current.GetNeighbours()
		for _, neighbour := range neighbours {
			if _, seen := previousPositions[neighbour.String()]; !seen {
				next := &position{
					parent:  pos,
					current: neighbour,
				}

				f.stack.push(next)
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
