.PHONY: help build test lint coverage fmt tools clean setup-hooks pre-commit lint-strict

# Variables
BINARY_NAME=relayer-sdk
GO=go
GOFLAGS=-v
COVERAGE_FILE=coverage.out

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the library
	$(GO) build $(GOFLAGS) ./...

test: ## Run tests
	$(GO) test $(GOFLAGS) -race ./...

coverage: ## Generate coverage report
	$(GO) test -race -coverprofile=$(COVERAGE_FILE) -covermode=atomic ./...
	$(GO) tool cover -html=$(COVERAGE_FILE) -o coverage.html
	@echo "Coverage report: coverage.html"

lint: ## Run golangci-lint
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

fmt: ## Format code
	$(GO) fmt ./...
	@which goimports > /dev/null && goimports -w . || echo "Install goimports: go install golang.org/x/tools/cmd/goimports@latest"

tools: ## Install development tools
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install golang.org/x/tools/cmd/goimports@latest

clean: ## Clean build artifacts
	rm -f $(COVERAGE_FILE) coverage.html
	$(GO) clean -cache -testcache

bench: ## Run benchmarks
	$(GO) test -bench=. -benchmem ./...

mod-tidy: ## Tidy go.mod
	$(GO) mod tidy
	$(GO) mod verify

setup-hooks: ## Install and configure pre-commit hooks
	@echo "Installing pre-commit..."
	@which pre-commit > /dev/null || pip install pre-commit
	@pre-commit install
	@echo "Pre-commit hooks installed successfully"

pre-commit: ## Run pre-commit checks manually
	pre-commit run --all-files

lint-strict: ## Run golangci-lint with strict config
	golangci-lint run --config=.golangci.yml ./...

.DEFAULT_GOAL := help
