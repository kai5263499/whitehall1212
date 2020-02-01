package interfaces

import "github.com/kai5263499/whitehall1212/types"

type Map interface {
	InitializeMap()
	PossibleMoves(types.Vertex, int) ([]*types.Edges, error)
}
