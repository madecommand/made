SOURCES=$(wildcard  *.proto)
TARGETS=$(SOURCES:.proto=.pb.go)


all: ## asdf \
 asdf

%.pb.go: %.proto
	protoc --go_out=plugins=grpc:. $?

build: $(TARGETS) ## Creates all the pb.go files

clean: ## Remove all the .pb.go
	@rm *.pb.go 2> /dev/null || true


patatas:
	@echo Muy ricas

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.DEFAULT_GOAL := help 
.PHONY: clean all help





