package internal

import (
	"container/heap"
	"fmt"
	"github.com/mateoferrari97/pathfinder/internal/maze"
)

const infinity = 1e10

type Randomize interface {
	Intn(n int) int
}

type vertex struct {
	position   *maze.Position
	neighbours map[string]*vertex
}

type edge struct {
	from     *vertex
	to       *vertex
	distance int
}

type graph struct {
	rand     Randomize
	vertices map[string]*vertex
	edges    map[string]*edge
}

func newGraph(maze *maze.Maze, rand Randomize) *graph {
	graph := &graph{
		rand:     rand,
		vertices: make(map[string]*vertex),
		edges:    make(map[string]*edge),
	}

	positions := maze.GetPositions()
	for _, position := range positions {
		var currentVertex *vertex
		if vx := graph.getVertex(position.String()); vx == nil {
			newVertex := &vertex{position: position, neighbours: make(map[string]*vertex)}
			currentVertex = newVertex
			graph.addVertex(newVertex)
		} else {
			currentVertex = vx
		}

		neighbours := position.GetNeighbours()
		for _, neighbour := range neighbours {
			var neighbourVertex *vertex
			if vx := graph.getVertex(neighbour.String()); vx == nil {
				newVertex := &vertex{position: neighbour, neighbours: make(map[string]*vertex)}
				neighbourVertex = newVertex
				graph.addVertex(newVertex)
			} else {
				neighbourVertex = vx
			}

			currentVertex.neighbours[neighbour.String()] = neighbourVertex

			if edge := graph.getEdge(currentVertex, neighbourVertex); edge == nil {
				graph.addEdge(currentVertex, neighbourVertex)
			}
		}
	}

	return graph
}

func (g *graph) getVertex(key string) *vertex {
	return g.vertices[key]
}

func (g *graph) addVertex(vertex *vertex) {
	g.vertices[vertex.position.String()] = vertex
}

func (g *graph) getEdge(from *vertex, to *vertex) *edge {
	key := fmt.Sprintf("%s-%s", from.position.String(), to.position.String())
	return g.edges[key]
}

func (g *graph) addEdge(from *vertex, to *vertex) {
	fromTo := fmt.Sprintf("%s-%s", from.position.String(), to.position.String())
	toFrom := fmt.Sprintf("%s-%s", to.position.String(), from.position.String())

	edge := &edge{
		from:     from,
		to:       to,
		distance: g.rand.Intn(100),
	}

	g.edges[fromTo] = edge
	g.edges[toFrom] = edge
}

type DijkstraPathFinder struct {
	maze          *maze.Maze
	graph         *graph
	priorityQueue priorityQueue
	weightTable   map[string]*item
}

func NewDijkstraPathFinder(maze *maze.Maze, rand Randomize) *DijkstraPathFinder {
	return &DijkstraPathFinder{
		maze:          maze,
		graph:         newGraph(maze, rand),
		priorityQueue: newPriorityQueue(),
		weightTable:   make(map[string]*item),
	}
}

func (f *DijkstraPathFinder) setup(fromPosition *maze.Position) {
	var index int
	for _, currentVertex := range f.graph.vertices {
		item := &item{
			parent:   nil,
			current:  currentVertex,
			priority: infinity,
			index:    index,
		}

		if fromPosition.Equal(*currentVertex.position) {
			item.priority = 0
		}

		f.priorityQueue = append(f.priorityQueue, item)
		f.weightTable[currentVertex.position.String()] = item
		index++
	}

	heap.Init(&f.priorityQueue)
}

func (f *DijkstraPathFinder) Find(fromPosition *maze.Position, toPosition *maze.Position) ([]string, error) {
	f.setup(fromPosition)
	return f.find(toPosition)
}

func (f *DijkstraPathFinder) find(toPosition *maze.Position) ([]string, error) {
	previousEdge := make(map[string]struct{})

	for f.priorityQueue.Len() != 0 {
		item := heap.Pop(&f.priorityQueue).(*item)
		currentVertex := item.current

		for _, neighbourVertex := range currentVertex.neighbours {
			fromTo := fmt.Sprintf("%s-%s", currentVertex.position.String(), neighbourVertex.position.String())
			toFrom := fmt.Sprintf("%s-%s", neighbourVertex.position.String(), currentVertex.position.String())

			if _, seen := previousEdge[fromTo]; seen {
				continue
			}

			if _, seen := previousEdge[toFrom]; seen {
				continue
			}

			previousEdge[fromTo] = struct{}{}
			previousEdge[toFrom] = struct{}{}

			currentNeighbourEdge := f.graph.getEdge(currentVertex, neighbourVertex)
			weightToNeighbour := item.priority + currentNeighbourEdge.distance
			neighbourItem := f.weightTable[neighbourVertex.position.String()]
			if weightToNeighbour < neighbourItem.priority {
				neighbourItem.parent = currentVertex
				newWeight := weightToNeighbour
				f.priorityQueue.update(neighbourItem, newWeight)
			}
		}
	}

	var path []string
	pos := f.weightTable[toPosition.String()]
	for pos.parent != nil {
		pos.current.position.WithValue(1)
		path = append(path, f.maze.String())
		pos = f.weightTable[pos.parent.position.String()]
	}

	pos.current.position.WithValue(1)
	path = append(path, f.maze.String())

	return path, nil
}
