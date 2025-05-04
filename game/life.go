package game

const (
	GridWidth  = 25
	GridHeight = 25
)

// Game represents the state of the Game of Life
type Game struct {
	grid [GridHeight][GridWidth]bool
}

// NewGame creates a new Game of Life
func NewGame() Game {
	return Game{}
}

// SetCell sets the state of a cell at the given coordinates and returns a new Game
func (g Game) SetCell(x, y int, alive bool) Game {
	if x < 0 || x >= GridWidth || y < 0 || y >= GridHeight {
		return g
	}

	newGame := g
	newGame.grid[y][x] = alive
	return newGame
}

// GetCell returns the state of a cell at the given coordinates
func (g Game) GetCell(x, y int) bool {
	if x >= 0 && x < GridWidth && y >= 0 && y < GridHeight {
		return g.grid[y][x]
	}
	return false
}

// Clear returns a new Game with all cells cleared
func (g Game) Clear() Game {
	return NewGame()
}

// NextGeneration computes the next state of the game and returns a new Game instance
func (g Game) NextGeneration() Game {
	newGame := g // Create a copy of the current game state
	newGrid := [GridHeight][GridWidth]bool{}

	for y := 0; y < GridHeight; y++ {
		for x := 0; x < GridWidth; x++ {
			neighbors := g.countNeighbors(x, y)
			cell := g.grid[y][x]

			// Apply Conway's Game of Life rules
			if cell {
				// Any live cell with fewer than two live neighbors dies
				// Any live cell with two or three live neighbors lives
				// Any live cell with more than three live neighbors dies
				newGrid[y][x] = neighbors == 2 || neighbors == 3
			} else {
				// Any dead cell with exactly three live neighbors becomes a live cell
				newGrid[y][x] = neighbors == 3
			}
		}
	}

	newGame.grid = newGrid
	return newGame
}

// ComputeNGenerations computes n generations ahead and returns the final state
// Returns the final game state and number of generations actually computed
func (g Game) ComputeNGenerations(n int) (Game, int) {
	if n <= 0 {
		return g, 0
	}

	currentGame := g
	// Compute n generations
	for i := 0; i < n; i++ {
		currentGame = currentGame.NextGeneration()
	}

	return currentGame, n
}

// countNeighbors returns the number of live neighbors for a cell
func (g Game) countNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			nx, ny := x+dx, y+dy
			// Handle wrapping around the edges
			if nx < 0 {
				nx = GridWidth - 1
			} else if nx >= GridWidth {
				nx = 0
			}
			if ny < 0 {
				ny = GridHeight - 1
			} else if ny >= GridHeight {
				ny = 0
			}

			if g.grid[ny][nx] {
				count++
			}
		}
	}
	return count
}
