package maze

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	maxRow = 5
	maxCol = 5
)

type Maze struct {
	matrix    [maxRow][maxCol]*Position
	positions map[string]*Position
}

func NewMaze() *Maze {
	maze := &Maze{
		positions: make(map[string]*Position),
	}

	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			position := &Position{
				maze:  maze,
				key:   fmt.Sprintf("%d,%d", row, col),
				row:   row,
				col:   col,
				value: 0,
			}

			maze.matrix[row][col] = position
			maze.positions[position.String()] = position
		}
	}

	return maze
}

func (m *Maze) Validate(startPosition Position, endPosition Position) error {
	if startPosition.row < 0 {
		return errors.New("start row must be positive")
	}

	if startPosition.col < 0 {
		return errors.New("start column must be positive")
	}

	if endPosition.row > maxRow {
		return errors.New("end row must be less than max row")
	}

	if endPosition.col > maxCol {
		return errors.New("end column must be less than max column")
	}

	return nil
}

func (m *Maze) String() string {
	var s []string
	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			current := m.matrix[row][col]

			s = append(s, strconv.Itoa(current.value))
		}

		s = append(s, "\n")
	}

	return strings.Join(s, "")
}

func (m *Maze) Neighbours(position *Position) []*Position {
	return position.Neighbours()
}

func (m *Maze) Position(key string) *Position {
	return m.positions[key]
}

func (m *Maze) Positions() []*Position {
	var positions []*Position
	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			positions = append(positions, m.matrix[row][col])
		}
	}

	return positions
}

type Position struct {
	maze  *Maze
	key   string
	value int
	row   int
	col   int
}

func (p *Position) Neighbours() []*Position {
	var neighbours []*Position
	startRow := p.row
	endRow := p.row
	startCol := p.col
	endCol := p.col

	if p.row-1 >= 0 {
		startRow = p.row - 1
	}

	if p.row+1 < maxRow {
		endRow = p.row + 1
	}

	if p.col-1 >= 0 {
		startCol = p.col - 1
	}

	if p.col+1 < maxCol {
		endCol = p.col + 1
	}

	for row := startRow; row <= endRow; row++ {
		for col := startCol; col <= endCol; col++ {
			if row == p.row && col == p.col {
				continue
			}

			neighbours = append(neighbours, p.maze.matrix[row][col])
		}
	}

	return neighbours
}

func (p *Position) String() string {
	return fmt.Sprintf("%d,%d", p.row, p.col)
}

func (p *Position) Equal(candidate Position) bool {
	return p.row == candidate.row && p.col == candidate.col
}

func (p *Position) WithValue(value int) {
	p.value = value
}
