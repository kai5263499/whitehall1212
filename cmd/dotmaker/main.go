package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"

	"github.com/kai5263499/whitehall1212"
	"github.com/kai5263499/whitehall1212/types"

	"github.com/awalterschulze/gographviz"
)

type config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

var (
	cfg config
)

func checkError(msg string, err error) {
	if err != nil {
		logrus.WithError(err).Error(msg)
	}
}

const (
	// GRAPH_NAME sets the graph name in dot
	GRAPH_NAME = "scotlandyard"
)

func main() {
	var err error
	cfg = config{}
	err = env.Parse(&cfg)
	checkError("parse config", err)

	outfilename := flag.String("out", "", "Filename to write output to")
	flag.Parse()

	g := gographviz.NewGraph()

	err = g.SetName(GRAPH_NAME)
	checkError("set name", err)

	err = g.SetDir(true)
	checkError("set dir", err)

	scotlandyard := whitehall1212.New()

	means := []types.Transportation{}

	vertices, err := scotlandyard.PossibleMoves(types.Vertex(1), 200, means)
	checkError("possible moves", err)

	for _, vertex := range vertices {

		err = g.AddNode(GRAPH_NAME, fmt.Sprintf("%d", vertex), nil)
		checkError("add node", err)

		edges, err := scotlandyard.GetEdges(vertex, means)
		checkError("get edges", err)

		for _, edge := range edges {
			err = g.AddEdge(
				fmt.Sprintf("%d", vertex),
				fmt.Sprintf("%d", edge.Destination),
				true,
				map[string]string{"label": fmt.Sprintf("\"%s\"", edge.Transportation)},
			)
			checkError("add edge", err)
		}
	}

	outfile, err := os.Create(*outfilename)
	checkError("outfile open", err)
	defer outfile.Close()
	fmt.Fprint(outfile, g.String())
}
