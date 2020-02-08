package interfaces

import "github.com/kai5263499/whitehall1212/types"

type Map interface {
	InitializeMap()
	GetEdges(types.Vertex, []types.Transportation) ([]types.Edge, error)
	PossibleMoves(types.Vertex, int, []types.Transportation) ([]types.Vertex, error)
}
