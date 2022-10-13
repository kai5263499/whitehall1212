package whitehall1212

//go:generate mapgen -in data/game_board.csv -out map_data.go

import (
	"errors"
	"fmt"

	"container/heap"

	"github.com/kai5263499/whitehall1212/types"
)

// Map is the interface to game map data
type Map interface {
	GetEdges(types.Vertex, []types.Transportation) ([]types.Edge, error)
	PossibleMoves(types.Vertex, int, []types.Transportation) ([]types.Vertex, error)
	ShortestPath(types.Vertex, types.Vertex, []types.Transportation) (int, []types.Edge, error)
}

var _ Map = (*gameMap)(nil)

const infinity = int(^uint(0) >> 1)

// New returns a Map instance intialized with game board data
func New() Map {
	m := &gameMap{}

	// The map needs to be initialized with the game board data before we can use it
	m.initializeMap()

	return m
}

// gameMap represents a simple in-memory graph
type gameMap struct {
	vertices map[types.Vertex][]types.Edge
}

// GetEdges returns
func (m *gameMap) GetEdges(vertex types.Vertex, means []types.Transportation) ([]types.Edge, error) {
	v, found := m.vertices[vertex]

	if !found {
		return nil, errors.New("Vertex not found")
	}

	if len(means) > 0 {
		e := make([]types.Edge, 0)

		for _, t := range means {
			for _, v2 := range v {
				if t == v2.Transportation {
					e = append(e, v2)
				}
			}
		}

		return e, nil
	}

	return v, nil
}

// PossibleMoves returns a slice of vertices that are depth steps away from the start Vertex
func (m *gameMap) PossibleMoves(start types.Vertex, depth int, means []types.Transportation) ([]types.Vertex, error) {
	positions := make([]types.Vertex, 0)

	visited := make(map[string]bool, 0)

	q := &moveQueue{}
	heap.Init(q)

	q.Push(queueEntry{
		vertex:   start,
		distance: 0,
	})

	for q.Len() > 0 {
		qi := heap.Pop(q).(queueEntry)

		if qi.distance == depth {

			visitedKey := fmt.Sprintf("r-%d-%d", qi.distance, qi.vertex)
			if _, found := visited[visitedKey]; found {
				continue
			}

			positions = append(positions, qi.vertex)

			visited[visitedKey] = true

			continue
		}

		edges, err := m.GetEdges(qi.vertex, means)
		if err != nil {
			return nil, err
		}

		for _, edge := range edges {

			distance := qi.distance + 1

			dupeKey := fmt.Sprintf("e-%d-%d", distance, edge.Destination)
			if _, found := visited[dupeKey]; found {
				continue
			}

			q.Push(queueEntry{
				vertex:   edge.Destination,
				distance: distance,
			})

			visited[dupeKey] = true
		}
	}

	return positions, nil
}

// ShortestPath returns the shortest path as a slice of edges
// between two vertices using a priority queue
func (m *gameMap) ShortestPath(start types.Vertex, finish types.Vertex, means []types.Transportation) (int, []types.Edge, error) {
	q := &pathQueue{}
	heap.Init(q)

	// Seed our queue with edges from the start vertex
	edges, err := m.GetEdges(start, means)
	if err != nil {
		return 0, nil, err
	}

	for _, e := range edges {
		heap.Push(q, pathEntry{
			path:     []types.Edge{e},
			distance: 1,
		})
	}

	// Track each visited vertex and its cost
	visited := make(map[types.Vertex]int, 0)

	for q.Len() > 0 {
		qi := heap.Pop(q).(pathEntry)

		v := qi.path[len(qi.path)-1]

		if v.Destination == finish {
			return qi.distance, qi.path, nil
		}

		// Reject any path that goes back through the start node
		// This is mostly to cover the edge case where we're processing
		// the first set of edges
		if v.Destination == start {
			continue
		}

		edges, err := m.GetEdges(v.Destination, means)
		if err != nil {
			return 0, nil, err
		}

		for _, e := range edges {
			if e.Destination == start {
				continue
			}

			newDistance := qi.distance + 1

			distance, found := visited[e.Destination]

			// Only add a new path entry to the queue if the next vertex is either unknown or
			// if this is a cheaper way to get to the vertex
			if !found || newDistance < distance {
				newPath := append(qi.path, e)

				heap.Push(q, pathEntry{
					distance: newDistance,
					path:     newPath,
				})

				visited[v.Destination] = newDistance
			}
		}
	}

	return 0, nil, fmt.Errorf("Unable to find a route between %d and %d", start, finish)
}
