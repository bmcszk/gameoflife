package game

import (
	"fmt"
	"testing"
)

func TestNewGame(t *testing.T) {
	g := NewGame()

	// Verify grid dimensions by checking cell access
	for y := 0; y < GridHeight; y++ {
		for x := 0; x < GridWidth; x++ {
			// Should be able to get and set cells
			g.SetCell(x, y, true)
			if !g.GetCell(x, y) {
				t.Errorf("Failed to set cell at (%d,%d)", x, y)
			}
		}
	}

	// Verify initial generation is 0
	if g.GetGeneration() != 0 {
		t.Errorf("Expected initial generation to be 0, got %d", g.GetGeneration())
	}
}

func TestGenerationCounter(t *testing.T) {
	g := NewGame()

	// Test initial generation
	if g.GetGeneration() != 0 {
		t.Errorf("Expected initial generation to be 0, got %d", g.GetGeneration())
	}

	// Test generation increment
	g.NextGeneration()
	if g.GetGeneration() != 1 {
		t.Errorf("Expected generation to be 1, got %d", g.GetGeneration())
	}

	// Test multiple generations
	for i := 0; i < 5; i++ {
		g.NextGeneration()
	}
	if g.GetGeneration() != 6 {
		t.Errorf("Expected generation to be 6, got %d", g.GetGeneration())
	}

	// Test generation reset on clear
	g.Clear()
	if g.GetGeneration() != 0 {
		t.Errorf("Expected generation to be 0 after clear, got %d", g.GetGeneration())
	}
}

func TestCountNeighbors(t *testing.T) {
	g := NewGame()

	// Test case 1: All cells dead
	if count := g.countNeighbors(1, 1); count != 0 {
		t.Errorf("Expected 0 neighbors, got %d", count)
	}

	// Test case 2: All neighbors alive
	g.Clear()
	g.SetCell(0, 0, true)
	g.SetCell(0, 1, true)
	g.SetCell(0, 2, true)
	g.SetCell(1, 0, true)
	g.SetCell(1, 2, true)
	g.SetCell(2, 0, true)
	g.SetCell(2, 1, true)
	g.SetCell(2, 2, true)
	if count := g.countNeighbors(1, 1); count != 8 {
		t.Errorf("Expected 8 neighbors, got %d", count)
	}

	// Test case 3: Some neighbors alive
	g.Clear()
	g.SetCell(0, 0, true)
	g.SetCell(1, 2, true)
	g.SetCell(2, 1, true)
	if count := g.countNeighbors(1, 1); count != 3 {
		t.Errorf("Expected 3 neighbors, got %d", count)
	}

	// Test case 4: Edge wrapping
	g.Clear()
	// Set corners
	g.SetCell(0, 0, true)                      // top-left
	g.SetCell(GridWidth-1, 0, true)            // top-right
	g.SetCell(0, GridHeight-1, true)           // bottom-left
	g.SetCell(GridWidth-1, GridHeight-1, true) // bottom-right

	// Test top-left corner wrapping
	count := g.countNeighbors(0, 0)
	if count != 3 {
		t.Errorf("Expected 3 neighbors for top-left corner (0,0), got %d", count)
	}

	// Test top-right corner wrapping
	count = g.countNeighbors(GridWidth-1, 0)
	if count != 3 {
		t.Errorf("Expected 3 neighbors for top-right corner (%d,0), got %d", GridWidth-1, count)
	}

	// Test bottom-left corner wrapping
	count = g.countNeighbors(0, GridHeight-1)
	if count != 3 {
		t.Errorf("Expected 3 neighbors for bottom-left corner (0,%d), got %d", GridHeight-1, count)
	}

	// Test bottom-right corner wrapping
	count = g.countNeighbors(GridWidth-1, GridHeight-1)
	if count != 3 {
		t.Errorf("Expected 3 neighbors for bottom-right corner (%d,%d), got %d", GridWidth-1, GridHeight-1, count)
	}
}

func TestNextGeneration(t *testing.T) {
	tests := []struct {
		name   string
		setup  func(*Game)
		verify func(*Game) bool
	}{
		{
			name: "Block (still life)",
			setup: func(g *Game) {
				g.Clear()
				g.SetCell(1, 1, true)
				g.SetCell(1, 2, true)
				g.SetCell(2, 1, true)
				g.SetCell(2, 2, true)
			},
			verify: func(g *Game) bool {
				return g.GetCell(1, 1) && g.GetCell(1, 2) && g.GetCell(2, 1) && g.GetCell(2, 2)
			},
		},
		{
			name: "Blinker (oscillator)",
			setup: func(g *Game) {
				g.Clear()
				g.SetCell(1, 2, true)
				g.SetCell(2, 2, true)
				g.SetCell(3, 2, true)
			},
			verify: func(g *Game) bool {
				return g.GetCell(2, 1) && g.GetCell(2, 2) && g.GetCell(2, 3)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGame()
			tt.setup(g)
			initialGen := g.GetGeneration()
			g.NextGeneration()
			if !tt.verify(g) {
				t.Error("Pattern did not evolve as expected")
			}
			if g.GetGeneration() != initialGen+1 {
				t.Errorf("Expected generation to be %d, got %d", initialGen+1, g.GetGeneration())
			}
		})
	}
}

func TestComputeNGenerations(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*Game)
		generations   int
		expectedGen   int
		verifyPattern func(*Game) bool
	}{
		{
			name: "Block (still life) - multiple generations",
			setup: func(g *Game) {
				g.Clear()
				g.SetCell(1, 1, true)
				g.SetCell(1, 2, true)
				g.SetCell(2, 1, true)
				g.SetCell(2, 2, true)
			},
			generations: 5,
			expectedGen: 5,
			verifyPattern: func(g *Game) bool {
				return g.GetCell(1, 1) && g.GetCell(1, 2) && g.GetCell(2, 1) && g.GetCell(2, 2)
			},
		},
		{
			name: "Blinker (oscillator) - even generations",
			setup: func(g *Game) {
				g.Clear()
				g.SetCell(1, 2, true)
				g.SetCell(2, 2, true)
				g.SetCell(3, 2, true)
			},
			generations: 4,
			expectedGen: 4,
			verifyPattern: func(g *Game) bool {
				return g.GetCell(1, 2) && g.GetCell(2, 2) && g.GetCell(3, 2)
			},
		},
		{
			name: "Blinker (oscillator) - odd generations",
			setup: func(g *Game) {
				g.Clear()
				g.SetCell(1, 2, true)
				g.SetCell(2, 2, true)
				g.SetCell(3, 2, true)
			},
			generations: 3,
			expectedGen: 3,
			verifyPattern: func(g *Game) bool {
				return g.GetCell(2, 1) && g.GetCell(2, 2) && g.GetCell(2, 3)
			},
		},
		{
			name: "Zero generations",
			setup: func(g *Game) {
				g.Clear()
				g.SetCell(1, 1, true)
			},
			generations: 0,
			expectedGen: 0,
			verifyPattern: func(g *Game) bool {
				return g.GetCell(1, 1)
			},
		},
		{
			name: "Negative generations",
			setup: func(g *Game) {
				g.Clear()
				g.SetCell(1, 1, true)
			},
			generations: -5,
			expectedGen: 0,
			verifyPattern: func(g *Game) bool {
				return g.GetCell(1, 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGame()
			tt.setup(g)
			initialGen := g.GetGeneration()

			computed := g.ComputeNGenerations(tt.generations)

			if computed != tt.expectedGen {
				t.Errorf("Expected %d generations computed, got %d", tt.expectedGen, computed)
			}

			if g.GetGeneration() != initialGen+tt.expectedGen {
				t.Errorf("Expected generation to be %d, got %d", initialGen+tt.expectedGen, g.GetGeneration())
			}

			if !tt.verifyPattern(g) {
				t.Error("Pattern did not evolve as expected")
			}
		})
	}
}

func setupGlider(g *Game) {
	g.Clear()
	centerX, centerY := GridWidth/2, GridHeight/2

	// Glider pattern
	g.SetCell(centerX, centerY-1, true)
	g.SetCell(centerX+1, centerY, true)
	g.SetCell(centerX-1, centerY+1, true)
	g.SetCell(centerX, centerY+1, true)
	g.SetCell(centerX+1, centerY+1, true)
}

func BenchmarkComputeNGenerations(b *testing.B) {
	// Test different numbers of generations
	generations := []int{10, 100, 1000}

	for _, n := range generations {
		b.Run(fmt.Sprintf("Compute%dGenerations", n), func(b *testing.B) {
			g := NewGame()
			setupGlider(g)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Reset the game state before each iteration
				g.Clear()
				setupGlider(g)
				g.ComputeNGenerations(n)
			}
		})
	}
}
