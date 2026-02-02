.PHONY: build run test clean docker install lint fmt help

# Variables
BINARY_NAME=fiscal-reminders
BUILD_DIR=bin
CMD_PATH=./cmd/fiscal-reminders
DOCKER_IMAGE=fiscal-reminders

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet

## help: Show this help
help:
	@echo "Fiscal Reminders - Portuguese Tax Deadline Notifications"
	@echo ""
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Targets:"
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## /  /'

## build: Build the binary
build:
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_PATH)
	@echo "✅ Built: $(BUILD_DIR)/$(BINARY_NAME)"

## run: Run the application
run:
	$(GORUN) $(CMD_PATH)

## test: Run tests
test:
	$(GOTEST) -v ./...

## test-cover: Run tests with coverage
test-cover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report: coverage.html"

## lint: Run linters
lint:
	$(GOVET) ./...
	@echo "✅ Vet passed"

## fmt: Format code
fmt:
	$(GOFMT) ./...
	@echo "✅ Formatted"

## clean: Clean build artifacts
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	@echo "✅ Cleaned"

## docker: Build Docker image
docker:
	docker build -f deployments/docker/Dockerfile -t $(DOCKER_IMAGE) .

## docker-run: Run in Docker
docker-run:
	docker run --rm -e NTFY_TOPIC=fiscal-reminders $(DOCKER_IMAGE)

## docker-compose: Run with docker-compose
docker-compose:
	docker-compose -f deployments/docker/docker-compose.yml up -d

## install: Install systemd service
install: build
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	sudo cp deployments/systemd/fiscal-reminders.service /etc/systemd/system/
	sudo systemctl daemon-reload
	sudo systemctl enable fiscal-reminders
	@echo "✅ Installed. Start with: sudo systemctl start fiscal-reminders"

# Build for all platforms
## build-all: Build for all platforms (Linux, macOS, Windows)
build-all:
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(CMD_PATH)
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(CMD_PATH)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(CMD_PATH)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(CMD_PATH)
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(CMD_PATH)
	@echo "✅ Built all platforms in $(BUILD_DIR)/"

