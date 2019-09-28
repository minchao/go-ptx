SHELL=/bin/bash -o pipefail

SPECS=$$(find . -maxdepth 1 -type f -name "oas.*")

.PHONY: spec
spec:
	go run spec/main.go

.PHONY: validate
validate:
	@for SPEC in ${SPECS}; do \
		echo "Specification: $${SPEC}"; \
		swagger validate $${SPEC} --skip-warnings --stop-on-error; \
	done

.PHONY: generate
generate:
	@for SPEC in ${SPECS}; do \
		IFS='.' read -r -a substrings <<< "$${SPEC}"; \
		CLIENT_DIR="./$${substrings[2]}/$${substrings[3]}"; \
		rm -r "$${CLIENT_DIR}"; \
		mkdir -p "$${CLIENT_DIR}"; \
        swagger generate client -f "$${SPEC}" -t "$${CLIENT_DIR}"; \
	done

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -v ./pkg/...

.PHONY: integration
integration:
	go test -v ./test/integration/...
