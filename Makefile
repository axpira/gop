.DEFAULT_GOAL := all

.PHONY: help fmt

help:  ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

fmt: ## Format all code with gofmt
	gofmt -s -w .

doc: ## Start a go doc server, need to have installed go tools: go get -u golang.org/x/tools/...
	godoc -http $(DOC_ADDR)
