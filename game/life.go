package game

const (
	GridWidth  = 25
	GridHeight = 25
)

// Game represents the state of the Game of Life
type Game struct {
	grids      [2][GridHeight][GridWidth]bool // this improved performance, no memory allocation for each generation
	activeGrid int
	generation int
}

// NewGame creates a new Game of Life
func NewGame() *Game {
	return &Game{
		activeGrid: 0,
		generation: 0,
	}
}

// GetGeneration returns the current generation number
func (g *Game) GetGeneration() int {
	return g.generation
}

// SetCell sets the state of a cell at the given coordinates
func (g *Game) SetCell(x, y int, alive bool) {
	if x >= 0 && x < GridWidth && y >= 0 && y < GridHeight {
		g.grids[g.activeGrid][y][x] = alive
	}
}

// GetCell returns the state of a cell at the given coordinates
func (g *Game) GetCell(x, y int) bool {
	if x >= 0 && x < GridWidth && y >= 0 && y < GridHeight {
		return g.grids[g.activeGrid][y][x]
	}
	return false
}

// Clear clears all cells in the grid and resets generation counter
func (g *Game) Clear() {
	for y := 0; y < GridHeight; y++ {
		for x := 0; x < GridWidth; x++ {
			g.grids[0][y][x] = false
			g.grids[1][y][x] = false
		}
	}
	g.activeGrid = 0
	g.generation = 0
}

// NextGeneration computes the next state of the game
// TODO: for better performance, we can detect cycles in the generations computing, and cache them for later use.
func (g *Game) NextGeneration() {
	nextGrid := 1 - g.activeGrid // Toggle between 0 and 1

	for y := 0; y < GridHeight; y++ {
		for x := 0; x < GridWidth; x++ {
			neighbors := g.countNeighbors(x, y)
			cell := g.grids[g.activeGrid][y][x]

			// Apply Conway's Game of Life rules
			if cell {
				// Any live cell with fewer than two live neighbors dies
				// Any live cell with two or three live neighbors lives
				// Any live cell with more than three live neighbors dies
				g.grids[nextGrid][y][x] = neighbors == 2 || neighbors == 3
			} else {
				// Any dead cell with exactly three live neighbors becomes a live cell
				g.grids[nextGrid][y][x] = neighbors == 3
			}
		}
	}

	g.activeGrid = nextGrid
	g.generation++
}

// ComputeNGenerations computes n generations ahead and returns the final state
// without displaying intermediate states. Returns the number of generations computed.
func (g *Game) ComputeNGenerations(n int) int {
	if n <= 0 {
		return 0
	}

	// Store initial generation
	initialGen := g.generation

	// Compute n generations
	for i := 0; i < n; i++ {
		g.NextGeneration()
	}

	// Return the number of generations computed
	return g.generation - initialGen
}

// countNeighbors returns the number of live neighbors for a cell
func (g *Game) countNeighbors(x, y int) int {
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

			if g.grids[g.activeGrid][ny][nx] {
				count++
			}
		}
	}
	return count
}
