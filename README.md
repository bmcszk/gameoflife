# Game of Life

A Go implementation of Conway's Game of Life with a console-based user interface.

## Features

- 25x25 grid with wrapping edges
- Double-buffered grid implementation for optimal performance
- Glider pattern initialization
- Generation counter
- Efficient computation of multiple generations
- Comprehensive test suite
- Benchmark tests

## Requirements

- Go 1.16 or later

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/gameoflife.git
cd gameoflife
```

2. Build the application:
```bash
make build
```

## Usage

### Running the Game

To run the game with the default glider pattern:
```bash
make run
```

The game will start with a glider pattern in the center of the grid. The simulation will continue until you stop it (Ctrl+C).

### Testing

Run all tests:
```bash
make test
```

Run benchmarks:
```bash
make bench
```

Run both tests and benchmarks:
```bash
make test-all
```

### Available Make Commands

- `make build` - Build the application
- `make run` - Build and run the application
- `make test` - Run unit tests
- `make bench` - Run benchmarks
- `make test-all` - Run both tests and benchmarks
- `make clean` - Remove build artifacts
- `make help` - Show all available commands

## Implementation Details

### Grid Implementation

The game uses a double-buffered grid implementation for optimal performance:
- Two 25x25 boolean arrays are used alternately
- No memory allocations during generation computation
- Wrapping edges for continuous simulation

### Patterns

The implementation includes several test patterns:
- Block (still life)
- Blinker (oscillator)
- Glider (spaceship)

### Performance

Benchmark results (on Intel Core i5-8365U):
- 10 generations: ~191,021 ns/op
- 100 generations: ~1,948,836 ns/op
- 1000 generations: ~19,087,484 ns/op

Zero memory allocations during generation computation.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request 
