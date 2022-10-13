package main

import (
	"flag"
	"fmt"

	"github.com/kai5263499/whitehall1212"
	"github.com/kai5263499/whitehall1212/types"
	"github.com/sirupsen/logrus"
)

var (
	scotlandyard whitehall1212.Map
)

func checkError(msg string, err error) {
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Errorf(msg)
	}
}

func main() {
	var err error

	start := flag.Int("start", 1, "Starting position")
	depth := flag.Int("depth", 1, "How far to travel in the graph")
	flag.Parse()

	scotlandyard = whitehall1212.New()

	means := []types.Transportation{}

	possibilities, err := scotlandyard.PossibleMoves(types.Vertex(*start), *depth, means)
	checkError("possible moves", err)

	fmt.Printf("%d has %d possible positions after %d moves using %s means of transportation\nmoves: %d\n", *start, len(possibilities), *depth, means, possibilities)
}
