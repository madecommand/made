



install:
	go install ./cmd/made

test:
	@go test ./...


deps: ## Install development dependencies
	@go get -u github.com/mitchellh/gox
	@go get -u github.com/tcnksm/ghr



help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.DEFAULT_GOAL := help 
.PHONY: clean all help





