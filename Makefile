
all: test run

build:
	@echo "Building..."
	@go build -o ./bin/zarinpal.out ./cmd/server/main.go

run:
	@go run ./cmd/server/main.go


test:
	@go clean -testcache
	@if ! command -v tparse > /dev/null; then \
		go install github.com/mfridman/tparse@latest; \
		echo "tparse installed successfully."; \
	fi;
	@echo "Testing..."; \
	go test ./... -json | tparse -all; 

clean:
	@echo "Cleaning...";
	@rm -rf ./bin/;

.PHONY: all build run test clean