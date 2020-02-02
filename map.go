package whitehall1212

//go:generate mapgen -in data/game_board.csv -out map_data.go

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
	"github.com/kai5263499/whitehall1212/interfaces"
	"github.com/kai5263499/whitehall1212/types"
)

var _ interfaces.Map = (*Map)(nil)

func New() *Map {
	return &Map{}
}

type Map struct {
	vertices map[types.Vertex][]types.Edge
}

func (m *Map) GetEdges(vertex types.Vertex, means []types.Transportation) ([]types.Edge, error) {
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

func (m *Map) PossibleMoves(start types.Vertex, depth int) ([]types.Edge, error) {
	path := make([]types.Edge, 0)

	spew.Dump(m.vertices)

	for i := 0; i < depth; i++ {

	}

	return path, nil
}
