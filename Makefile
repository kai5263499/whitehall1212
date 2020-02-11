
map_data.go:
	go install github.com/kai5263499/whitehall1212/cmd/mapgen
	go generate

graphviz:
	go install github.com/kai5263499/whitehall1212/cmd/dotmaker
	dotmaker --out sy.dot
	dot -Tpdf sy.dot -o sy.pdf