package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Pos struct {
	X int
	Y int
}

// Run is directional, pointing away from 0
type Run struct {
	Start Pos
	End   Pos
}

type Input struct {
	Width  int
	Height int
	Tiles  []Pos
}

// -------------- Helpers ----------------------------

// Establish a grid with holes in correct places
func createGrid(in Input) [][]int {
	var grid = make([][]int, in.Width)
	for i := range grid {
		grid[i] = make([]int, in.Height)
	}
	for _, tile := range in.Tiles {
		grid[tile.X][tile.Y] = 1
	}

	return grid
}

// Check for collisions between two runs
// True if collision
func checkCollision(r1 Run, r2 Run) bool {
	// fmt.Println("Logic")
	// fmt.Println(r1.End.X < r2.Start.X)
	// fmt.Println(r1.Start.X > r2.End.X)
	// fmt.Println(r1.End.Y < r2.Start.Y)
	// fmt.Println(r1.Start.Y > r2.End.Y)
	// fmt.Println(r1)
	// fmt.Println(r2)
	return !((r1.End.X < r2.Start.X || r1.Start.X > r2.End.X) || (r1.End.Y < r2.Start.Y || r1.Start.Y > r2.End.Y))
}

// Gets the number of available Tiles in grid
func getSize(in Input) int {
	size := in.Width * in.Height
	size -= len(in.Tiles)
	return size
}

// -------------- Main Functions -------------------------------------

// Find all runs of n length, either horizontally or vertically
// Assumes grid is has all rows equal size, and all columns equal size
func getRuns(grid [][]int, n int) []Run {
	var runs []Run
	var horizontal = 0
	var vertical = 0

	// Corner case
	if len(grid) == 0 {
		return nil
	}

	// Big-O(2hv) for the two nested for loops
	// This is a more verbose solution, but much more understandable than a reduced solutions
	// One reduced solution would be to turn the grid into a square, and the scan both [i][j] and [j][i]
	// at the same time for hor and ver runs

	// Horizontal
	for i := 0; i < len(grid); i++ {
		// Vertical
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				vertical++
				if vertical >= n {
					runs = append(runs, Run{Pos{i, j - (n - 1)}, Pos{i, j}})
				}
			} else {
				vertical = 0
				continue
			}
		}
		vertical = 0
	}

	// Vertical
	for j := 0; j < len(grid[0]); j++ {
		// Horizontal
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == 0 {
				horizontal++
				if horizontal >= n {
					runs = append(runs, Run{Pos{i - (n - 1), j}, Pos{i, j}})
				}
			} else {
				horizontal = 0
				continue
			}
		}
		horizontal = 0
	}

	return runs
}

// Exhausts all combinations of runs, or until a solution is found
// Assumes size 3 of blocks
// Adds together
func coverGrid(size int, stack []Run, collector []Run) []Run {

	tempCollector := make([]Run, len(collector))
	tempStack := make([]Run, len(stack))

	// Create a copy of the collector and stack in case no matches for the first stack item
	copy(tempCollector, collector)
	copy(tempStack, stack)

	// Base case
	if (len(collector) * 3) == size {
		return collector
	}

	// Non-starter Corner case and no solution check
	if size%3 != 0 {
		return nil
	}

	// Add the first item to the collector for comparison
	if collector == nil {
		collector = append(collector, stack[0])
		stack = stack[1:]
	}

Top:
	// If there was a collision with all remaining items, skip stack[0]
	if len(stack) == 0 {
		return nil
	}

	for _, r := range collector {
		// If there is a collision, pop the item from the stack
		if checkCollision(r, stack[0]) {
			stack = stack[1:]
			goto Top
		}
	}

	// If no collisions for stack[i], move item from stack to collector and recur
	var item = stack[0]
	stack = stack[1:]
	ret := coverGrid(size, stack, append(collector, item))

	if ret == nil && len(tempStack) != 0 {
		return coverGrid(size, tempStack[1:], tempCollector)
	}

	return ret
}

// --------------- Logistics ---------------------

// Basic file reading boilerplate
func readFromFile(filePath string) (*Input, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	raw, fErr := io.ReadAll(file)
	if fErr != nil {
		return nil, fErr
	}

	var in Input

	mErr := json.Unmarshal(raw, &in)
	if mErr != nil {
		return nil, mErr
	}

	return &in, nil
}

func main() {
	filePath := os.Args[1]

	in, err := readFromFile(filePath)
	if err != nil {
		fmt.Println("error reading from file at path")
		panic(err)
	}

	grid := createGrid(*in)
	var runs = getRuns(grid, 3)
	final := coverGrid(getSize(*in), runs, nil)

	output, err := json.Marshal(final)
	if err != nil {
		fmt.Println("error marshalling output")
		panic(err)
	}

	fmt.Println(string(output))
}
