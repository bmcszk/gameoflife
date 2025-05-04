package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"gameoflife/game"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printGrid(g game.Game, generation int) {
	fmt.Printf("Generation: %d\n\n", generation)
	for y := 0; y < game.GridHeight; y++ {
		for x := 0; x < game.GridWidth; x++ {
			if g.GetCell(x, y) {
				fmt.Print("■ ")
			} else {
				fmt.Print("□ ")
			}
		}
		fmt.Println()
	}
}

func initializeGlider(g game.Game) game.Game {
	g = g.Clear()

	// Place the glider in the middle
	centerX, centerY := game.GridWidth/2, game.GridHeight/2

	// Glider pattern
	g = g.SetCell(centerX, centerY-1, true)
	g = g.SetCell(centerX+1, centerY, true)
	g = g.SetCell(centerX-1, centerY+1, true)
	g = g.SetCell(centerX, centerY+1, true)
	g = g.SetCell(centerX+1, centerY+1, true)

	return g
}

func main() {
	// Create a new game
	g := game.NewGame()
	g = initializeGlider(g)
	generation := 0

	// Run the game loop
	for {
		clearScreen()
		printGrid(g, generation)
		g = g.NextGeneration()
		generation++
		time.Sleep(100 * time.Millisecond)
	}
}
