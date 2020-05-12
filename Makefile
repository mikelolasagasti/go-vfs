GOLANGCI_LINT_VERSION=1.26.0

.PHONY: nothing
nothing:

.PHONY: coverage.out
coverage.out:
	go test -coverprofile=$@ ./...

.PHONY: format
format:
	find . -name \*.go | xargs gofumports -local github.com/twpayne/go-vfs -w

.PHONY: html-coverage
html-coverage:
	go tool cover -html=coverage.out

.PHONY: install-tools
install-tools:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- v1.26.0
	GO111MODULE=off go get -u \
		github.com/mattn/goveralls \
		golang.org/x/tools/cmd/cover \
		mvdan.cc/gofumpt/gofumports

.PHONY: lint
lint:
	go vet ./...
	./bin/golangci-lint run

.PHONY: test
test:
	go test ./...