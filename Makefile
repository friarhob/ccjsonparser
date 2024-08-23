BINARY_NAME=ccjsonparser
BUILD_DIR=./cmd/ccjsonparser
GO_VERSION=$(go version | cut -d' ' -f3)

# Default target
all: build test

# Build target
build:
	@echo "Building the project..."
	go build -o $(BINARY_NAME) $(BUILD_DIR)
	chmod +x $(BINARY_NAME)
	@echo "Build complete."

# Clean target
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	@echo "Clean complete."

# Test target
test:
	@echo "Running tests..."
	go test ./...
	@echo "Tests complete."

# Run target
run: build
	@echo "Running the application..."
	./$(BINARY_NAME) --help

# Install target
install: build
	@echo "Installing the application..."
	go install $(BUILD_DIR)

# Phony targets
.PHONY: all build clean test run install