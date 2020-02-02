package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kai5263499/whitehall1212"
	"github.com/kai5263499/whitehall1212/types"
	"github.com/sirupsen/logrus"
)

func checkError(err error) {
	if err != nil {
		logrus.WithError(err).Errorf("encountered unrecoverable err")
		panic(err)
	}
}

func main() {
	outfilename := flag.String("out", "", "Filename to write output to")
	flag.Parse()

	outfile, err := os.Create(*outfilename)
	checkError(err)
	defer outfile.Close()

	scotlandyard := whitehall1212.New()
	scotlandyard.InitializeMap()

	for i := 1; i <= 199; i++ {
		pos := types.Vertex(i)

		fmt.Fprintf(outfile, "<_:%[1]d> <position> \"%[1]d\" .\n", pos)

		edges, err := scotlandyard.GetEdges(pos, []types.Transportation{})
		checkError(err)

		for _, e := range edges {
			fmt.Fprintf(outfile, "<_:%d> <can%sTo> <_:%d> .\n", pos, e.Transportation, e.Destination)
		}
	}
}
