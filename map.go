package whitehall1212

//go:generate mapgen -in data/game_board.csv -out map_data.go

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/kai5263499/whitehall1212/interfaces"
	"github.com/kai5263499/whitehall1212/types"
)

var _ interfaces.Map = (*Map)(nil)

func New() *Map {
	return &Map{}
}

type Map struct {
	vertices map[types.Vertex]types.Edges
}

func (m *Map) PossibleMoves(start types.Vertex, depth int) ([]*types.Edges, error) {
	path := make([]*types.Edges, 0)

	spew.Dump(m.vertices)

	for i := 0; i < depth; i++ {

	}

	return path, nil
}
