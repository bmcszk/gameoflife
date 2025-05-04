.PHONY: all build run test bench clean

# Default target
all: build

# Build the application
build:
	@echo "Building Game of Life..."
	go build -o gameoflife

# Run the application
run: build
	@echo "Running Game of Life..."
	./gameoflife

# Run unit tests with verbose output
test:
	@echo "Running unit tests..."
	go test ./game -v

# Run benchmarks with memory allocation stats
bench:
	@echo "Running benchmarks..."
	go test ./game -bench=. -benchmem

# Run both tests and benchmarks
test-all: test bench

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f gameoflife
	go clean

# Help target
help:
	@echo "Available targets:"
	@echo "  all       - Build the application (default)"
	@echo "  build     - Build the application"
	@echo "  run       - Build and run the application"
	@echo "  test      - Run unit tests"
	@echo "  bench     - Run benchmarks"
	@echo "  test-all  - Run both tests and benchmarks"
	@echo "  clean     - Remove build artifacts"
	@echo "  help      - Show this help message" 
