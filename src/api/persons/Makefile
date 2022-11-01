.PHONY: help
help: ## Run this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Run all test
	go test ./... -v -coverprofile=coverage.out

test-with-coverage: ## Run all tests with coverage report
	go test ./... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out

test-with-benchmark: ## Run all tests with benchmark data
	go test ./... -bench=.
