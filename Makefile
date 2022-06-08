
install:
	go install ./cmd/made

test:
	@go test ./...


deps: ## Install development dependencies
	@go get -u github.com/tcnksm/ghr

release: dist tag-release upload-release ## Generates a release for the given TAG

dist: ## Generates dist files
	@rm -rf dist
	@mkdir dist
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o dist/made.exe ./cmd/made
	@CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o dist/made-linux ./cmd/made
	@CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o dist/made-darwin ./cmd/made

tag-release: ## Generate git TAG for release
	git tag $(TAG)
	git push --tags
	
upload-release: ## Upload release TAG
	 ~/go/bin/ghr -t $(GITHUB_TOKEN) -u guillermo -r made  --replace --draft  $(TAG) dist/

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.DEFAULT_GOAL := help 
.PHONY: clean all help dist





