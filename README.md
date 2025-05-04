# Game of Life

A Go implementation of Conway's Game of Life with a console-based user interface.

## Features

- 25x25 grid with wrapping edges
- Immutable design for better predictability and thread safety
- Zero memory allocations during game evolution
- Efficient fixed-size array implementation
- Simple console UI with Unicode block characters
- Comprehensive test suite
- Built-in glider pattern initialization

## Requirements

- Go 1.16 or later

## Installation

1. Clone the repository (using SSH):
```bash
git clone git@github.com:bmcszk/gameoflife.git
cd gameoflife
```

Or using HTTPS:
```bash
git clone https://github.com/bmcszk/gameoflife.git
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

### Immutable Design

The game uses an immutable design where each state change returns a new `Game` instance. This approach:
- Makes state transitions explicit and predictable
- Eliminates the need for synchronization in concurrent scenarios
- Simplifies testing and reasoning about the code
- Keeps the core game logic pure and focused

### Data Structures

- Fixed-size arrays for the grid to avoid memory allocations
- Simple struct with just the essential game state
- No generation counter in the core game logic (handled by UI layer)

### Game Rules

The implementation follows Conway's Game of Life rules:
1. Any live cell with fewer than two live neighbors dies (underpopulation)
2. Any live cell with two or three live neighbors lives
3. Any live cell with more than three live neighbors dies (overpopulation)
4. Any dead cell with exactly three live neighbors becomes a live cell (reproduction)

### Patterns

The implementation includes several test patterns:
- Block (still life)
- Blinker (oscillator)
- Glider (spaceship)

### Performance

Benchmark results (on Intel Core i5-8365U):
- 10 generations: ~245,041 ns/op
- 100 generations: ~2,434,309 ns/op
- 1000 generations: ~24,220,062 ns/op

Zero memory allocations during generation computation.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request 
