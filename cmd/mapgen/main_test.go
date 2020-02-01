package main

import (
	"strings"
	"testing"

	"github.com/kai5263499/whitehall1212/types"
)

func TestImport(t *testing.T) {
	data := `1,8,Taxi
1,9,Taxi
1,58,Bus
1,46,Bus
1,46,Underground`

	lines := strings.Split(data, "\n")

	initMap()

	for _, line := range lines {
		parseLine(line)
	}

	routes := vertices[1][58]
	if len(routes) < 1 {
		t.Fatal("1<->58 no routes found")
	}

	if routes[0] != types.Transportation("Bus") {
		t.Fatal("1<->58 Bus route not found")
	}
}
