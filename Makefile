# Rules
.DEFAULT: help
.PHONY: help build

help: ## Display usage
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Compile Go binary for the current OS
	@echo "Compiling... " && go build -o "$(PWD)/bin/modulable-zork" main.go

run: build
	@"$(PWD)/bin/modulable-zork"

install: ## install dependencies with glide
	@glide --version >/dev/null 2>&1 || { echo >&2 "Glide is required.  Aborting."; exit 1; }
	@glide install --strip-vcs --update-vendored

test: ## run go tests
	@go test $(shell go list github.com/piteur/modulable-zork/... | grep -v /vendor)
