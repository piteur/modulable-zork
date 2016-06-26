# Rules
.DEFAULT: help
.PHONY: help build

help: ## Display usage
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Compile Go binary for the current OS
	@echo "Compiling... " && go build -o "$(PWD)/bin/modulable-zork" main.go

build-linux: ## Compile Go binary for Linux
	@echo "Compiling for Unix... " && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "$(PWD)/bin/modulable-zork" main.go

build-windows: ## Compile Go binary for Windows
	@echo "Compiling for Windows... " && GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o "$(PWD)/bin/modulable-zork.exe" main.go

build-mac: ## Compile Go binary for Mac
	@echo "Compiling for Mac... " && GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o "$(PWD)/bin/modulable-zork.exe" main.go

build-all: build-linux build-windows build-mac ## Compile all possible binaries
	@echo "Done !"

run: build
	@"$(PWD)/bin/modulable-zork"
