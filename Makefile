SHELL=/bin/bash -o pipefail

SPECS=$$(find . -maxdepth 1 -type f -name "oas.*")

.PHONY: help
## help: print this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: spec
## spec: fetche all openapi specifications
spec:
	go run spec/main.go

.PHONY: validate
## validate: validate the openapi specifications
validate:
	@for SPEC in ${SPECS}; do \
		echo "Specification: $${SPEC}"; \
		swagger validate $${SPEC} --skip-warnings --stop-on-error; \
	done

.PHONY: generate
## generate: generate the go client library
generate:
	@for SPEC in ${SPECS}; do \
		IFS='.' read -r -a substrings <<< "$${SPEC}"; \
		CLIENT_DIR="./$${substrings[2]}/$${substrings[3]}"; \
		rm -r "$${CLIENT_DIR}"; \
		mkdir -p "$${CLIENT_DIR}"; \
        swagger generate client -f "$${SPEC}" -t "$${CLIENT_DIR}"; \
	done

.PHONY: lint
## lint: run linters
lint:
	golangci-lint run ./...

.PHONY: test
## test: run unit tests
test:
	go test -v ./pkg/...

.PHONY: test-integration
## test-integration: run integration tests
test-integration:
	go test -v ./test/integration/...
