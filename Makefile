# Rules
.DEFAULT: help
.PHONY: help build

help: ## Display usage
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Compile Go binary for the current OS
	@echo "Compiling... " && GOGC=off go build -o "$(PWD)/bin/modulable-zork" main.go
	@echo "Executing..."
	@echo "-------------"
	@echo
	@$(PWD)/bin/modulable-zork

build-linux: ## Compile Go binary for Linux
	@echo "Compiling for Unix... " && GOGC=off GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "$(PWD)/bin/modulable-zork" main.go

build-windows: ## Compile Go binary for Windows
	@echo "Compiling for Windows... " && GOGC=off GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o "$(PWD)/bin/modulable-zork.exe" main.go
