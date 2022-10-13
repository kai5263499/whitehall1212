which = $(shell which $1 2> /dev/null || echo $1)

GO_PATH := $(call which,go)
$(GO_PATH):
	$(error Missing go)

TEST_ARGS := -cover -v -count=1
test: 
	@$(GO_PATH) test $(TEST_ARGS) ./... -tags=!intg

map_data.go:
	go install github.com/kai5263499/whitehall1212/cmd/mapgen
	go generate

graphviz:
	go install github.com/kai5263499/whitehall1212/cmd/dotmaker
	dotmaker --out sy.dot
	dot -Tpdf sy.dot -o sy.pdf

PUML_PATH := $(call which,plantuml)
$(PUML_PATH):
	$(error Missing plantuml: https://plantuml.com/starting)
puml:
	$(PUML_PATH) -t png -o . $(wildcard doc/*.puml)

doc: puml
.PHONY: doc

LINTER_PATH := $(call which,golangci-lint)
$(LINTER_PATH):
	$(error Missing golangci: https://golangci-lint.run/usage/install)
lint: setup
	export GOMODCACHE=./vendor
	$(LINTER_PATH) run

setup:
	@rm -rf ./vendor
	@$(GO_PATH) mod vendor