##-----------------------
## Available make targets
##-----------------------
##

default: help
help: ## Display this message
	@grep -E '(^[a-zA-Z0-9_.-]+:.*?##.*$$)|(^##)' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

##
## ----------------------
## Q.A
## ----------------------
##

qa: lint test ## Run all QA process

lint: ## Lint source code
	golangci-lint run -v

test: ## Run tests
	go test -race -v ./...
