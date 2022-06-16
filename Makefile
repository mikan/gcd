.DEFAULT_GOAL := help

.PHONY: test
test: ## Run go test
	go test -cover .

.PHONY: lint
lint: ## Run go vet and staticcheck
	go vet ./...
	staticcheck -checks inherit,ST1021,ST1022 ./...

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
