package maze

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	maxRow = 25
	maxCol = 25
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

func (m *Maze) GetPosition(key string) (*Position, error) {
	if position, exist := m.positions[key]; exist {
		return position, nil
	}

	return nil, errors.New("maze: position not found")
}

func (m *Maze) GetPositions() []*Position {
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

func (p *Position) GetNeighbours() []*Position {
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

func (p *Position) Distance(q Position) int {
	x := p.row - q.row
	y := p.col - q.col

	return int(math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2)))
}
