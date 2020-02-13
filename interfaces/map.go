package interfaces

import "github.com/kai5263499/whitehall1212/types"

// Map is the interface to game map data
type Map interface {
	GetEdges(types.Vertex, []types.Transportation) ([]types.Edge, error)
	PossibleMoves(types.Vertex, int, []types.Transportation) ([]types.Vertex, error)
	ShortestPath(types.Vertex, types.Vertex, []types.Transportation) (int, []types.Edge, error)
}
