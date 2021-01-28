.PHONY: help

help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build:  ## Build binary
	env GOOS=linux go build -ldflags="-s -w" -o ./bin/app main.go

clean:  ## Clean-up 
	rm -rf ./bin 

test: ## Test the api
	go test