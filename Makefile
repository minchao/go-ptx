SHELL=/bin/bash -o pipefail

.PHONY: generate
generate:
	swagger generate client -f ./oas.basic.json -t ./basic
	swagger generate client -f ./oas.air.json -t ./air
	swagger generate client -f ./oas.bus.json -t ./bus
	swagger generate client -f ./oas.rail.json -t ./rail
	swagger generate client -f ./oas.bike.json -t ./bike
	swagger generate client -f ./oas.tourism.json -t ./tourism

.PHONY: lint
lint:
	golangci-lint run -E gofmt ./...
