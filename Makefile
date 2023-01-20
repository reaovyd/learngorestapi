GO ?= go
TEST_DIR ?= ./api/test
AIR_CONFIG_FILE ?= .air.toml
COVERAGE_DIR ?= ./coverage

all: build

build:
	MODE=prod $(GO) build -o bin/server api/main.go

dev:
	@if [ -f "$(AIR_CONFIG_FILE)" ]; then \
		air; \
	else \
		echo "MISSING AIR CONFIGURATION FILE"; \
	fi
test:
	@mkdir -p $(COVERAGE_DIR) 
	$(GO) test $(TEST_DIR) -v -coverpkg=./... -coverprofile=$(COVERAGE_DIR)/coverage.out 
	$(GO) tool cover -o $(COVERAGE_DIR)/coverage.html -html=$(COVERAGE_DIR)/coverage.out 