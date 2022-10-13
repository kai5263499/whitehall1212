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
	destination := flag.Int("end", 1, "Destination position")
	flag.Parse()

	scotlandyard = whitehall1212.New()
	means := []types.Transportation{}
	distance, path, err := scotlandyard.ShortestPath(types.Vertex(*start), types.Vertex(*destination), means)
	checkError("shortest path", err)

	fmt.Printf("the distance between %d and %d is %d\npath: %+#v\n", *start, *destination, distance, path)
}
