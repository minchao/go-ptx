SHELL=/bin/bash -o pipefail

.PHONY: generate
generate:
	swagger generate client -f ./oas.basic.v2.json -t ./basic/v2
	swagger generate client -f ./oas.air.v2.json -t ./air/v2
	swagger generate client -f ./oas.bus.v2.json -t ./bus/v2
	swagger generate client -f ./oas.rail.v2.json -t ./rail/v2
	swagger generate client -f ./oas.rail.v3.json -t ./rail/v3
	swagger generate client -f ./oas.bike.v2.json -t ./bike/v2
	swagger generate client -f ./oas.tourism.v2.json -t ./tourism/v2

.PHONY: lint
lint:
	golangci-lint run -E gofmt ./...

.PHONY: spec
spec:
	go run spec/main.go
