package whitehall1212

import (
	"github.com/kai5263499/whitehall1212/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type neighborTest struct {
	start     int
	neighbors []types.Edge
	means     []types.Transportation
	expectErr bool
}

type possibilitiesTest struct {
	start         int
	depth         int
	possibilities []types.Vertex
	means         []types.Transportation
	expectErr     bool
}

type shortestPathTest struct {
	start     int
	end       int
	means     []types.Transportation
	cost      int
	path      []types.Edge
	expectErr bool
}

var _ = Describe("map", func() {

	var (
		gameMap Map
	)

	BeforeEach(func() {
		gameMap = New()
	})

	It("should return the right neighbors at the right edges", func() {
		tests := []neighborTest{
			{start: 1, neighbors: []types.Edge{
				{
					Transportation: "Taxi",
					Destination:    8,
				},
				{
					Transportation: "Taxi",
					Destination:    9,
				},
				{
					Transportation: "Bus",
					Destination:    46,
				},
				{
					Transportation: "Underground",
					Destination:    46,
				},
				{
					Transportation: "Bus",
					Destination:    58,
				},
			}, means: []types.Transportation{}, expectErr: false},
			{start: 300, neighbors: []types.Edge{}, means: []types.Transportation{}, expectErr: true},
		}

		for _, t := range tests {
			neighbors, err := gameMap.GetEdges(types.Vertex(t.start), t.means)
			if t.expectErr {
				Expect(err).ToNot(BeNil())
				continue
			} else {
				Expect(err).To(BeNil())
			}

			Expect(neighbors).To(Equal(t.neighbors))
		}
	})

	It("should return the right neighbors at the right depth", func() {
		tests := []possibilitiesTest{
			{start: 1, depth: 1, possibilities: []types.Vertex{8, 58, 46, 9}, means: []types.Transportation{}, expectErr: false},
			{start: 20, depth: 2, possibilities: []types.Vertex{46, 19, 1, 32, 21, 10, 20}, means: []types.Transportation{}, expectErr: false},
		}

		for _, t := range tests {
			possibilities, err := gameMap.PossibleMoves(types.Vertex(t.start), t.depth, t.means)
			if t.expectErr {
				Expect(err).ToNot(BeNil())
				continue
			} else {
				Expect(err).To(BeNil())
			}

			Expect(possibilities).To(Equal(t.possibilities))
		}
	})

	It("should return the shortest path between two vertices", func() {
		tests := []shortestPathTest{
			{start: 1, end: 8, means: []types.Transportation{}, cost: 1, path: []types.Edge{
				{
					Destination:    8,
					Transportation: "Taxi",
				},
			}, expectErr: false},
			{start: 1, end: 20, means: []types.Transportation{}, cost: 2, path: []types.Edge{
				{
					Destination:    9,
					Transportation: "Taxi",
				},
				{
					Destination:    20,
					Transportation: "Taxi",
				},
			}, expectErr: false},
			{start: 1, end: 94, means: []types.Transportation{}, cost: 3, path: []types.Edge{
				{
					Destination:    58,
					Transportation: "Bus",
				},
				{
					Destination:    74,
					Transportation: "Taxi",
				},
				{
					Destination:    94,
					Transportation: "Bus",
				},
			}, expectErr: false},
		}

		for _, t := range tests {
			cost, path, err := gameMap.ShortestPath(types.Vertex(t.start), types.Vertex(t.end), t.means)
			if t.expectErr {
				Expect(err).ToNot(BeNil())
				continue
			} else {
				Expect(err).To(BeNil())
			}

			Expect(cost).To(Equal(t.cost))
			Expect(path).To(Equal(t.path))
		}
	})
})
