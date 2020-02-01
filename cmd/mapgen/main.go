package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kai5263499/whitehall1212/types"
)

var (
	vertices [][][]interface{}
)

func checkError(msg string, err error) {
	if err != nil {
		fmt.Errorf("%s err=%+#v", msg, err)
	}
}

func initMap() {
	vertices = make([][][]interface{}, 201)
}

func parseLine(line string) {
	fields := strings.Split(line, ",")
	if len(fields) < 3 {
		return
	}

	source, err := strconv.Atoi(fields[0])
	checkError("atoi source", err)

	destination, err := strconv.Atoi(fields[1])
	checkError("atoi destination", err)

	transportation := fields[2]

	s := types.Vertex(source)
	t := types.Transportation(transportation)
	d := types.Vertex(destination)

	if vertices[s] == nil {
		vertices[s] = make([][]interface{}, 201)
	}

	if vertices[d] == nil {
		vertices[d] = make([][]interface{}, 201)
	}

	if vertices[s][d] == nil {
		vertices[s][d] = make([]interface{}, 0)
	}

	vertices[s][d] = append(vertices[s][d], t)
	vertices[d][s] = append(vertices[d][s], t)
}

func main() {
	infilename := flag.String("in", "", "Filename of nquad gameboard file")
	outfilename := flag.String("out", "", "Filename to write output to")
	flag.Parse()

	outfile, err := os.Create(*outfilename)
	defer outfile.Close()

	fmt.Fprintf(outfile, `package whitehall1212

// This file generated with mapgen -in %s -out %s; DO NOT EDIT.

import (
	"github.com/kai5263499/whitehall1212/types"
)

// InitializeMap builds the in-memory game board representation
func (m *Map) InitializeMap() {
	m.vertices = make(map[types.Vertex]types.Edges, 201)
	for pos := range m.vertices {
		m.vertices[pos] = make([]types.Edge, 0)
	}
`, *infilename, *outfilename)

	infile, err := os.Open(*infilename)
	checkError("file open", err)
	defer infile.Close()

	initMap()

	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		parseLine(scanner.Text())
	}

	err = scanner.Err()
	checkError("scanner err", err)

	for i := 0; i < 201; i++ {
		writeOut := false
		var b bytes.Buffer

		fmt.Fprintf(&b, `
	m.vertices[types.Vertex(%d)] = types.Edges{`, i)
		for d, d1 := range vertices[i] {
			writeOut = true
			for _, r := range d1 {
				fmt.Fprintf(&b, `
		{
			Transportation: types.Transportation("%s"),
			Destination:    types.Vertex(%d),
		},`, r, d)
			}
		}

		fmt.Fprintf(&b, `
	}
`)

		if writeOut {
			fmt.Fprintf(outfile, b.String())
		}
	}

	fmt.Fprintf(outfile, `
}
`)
}
