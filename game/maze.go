package game

import (
	"strings"
)

// MazeDirection represents a direction in the maze
type MazeDirection int

const (
	North MazeDirection = iota
	East
	South
	West
)

// String returns the string representation of a direction
func (d MazeDirection) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "Unknown"
	}
}

// Opposite returns the opposite direction
func (d MazeDirection) Opposite() MazeDirection {
	switch d {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		return d
	}
}

// Cell represents a single cell in the maze
type Cell struct {
	// Walls indicates whether there are walls in each direction
	// The index represents: 0=North, 1=East, 2=South, 3=West
	Walls [4]bool
}

// Maze represents a 2D grid maze
type Maze struct {
	Width  int
	Height int
	Grid   [][]Cell
}

// NewMaze creates a new maze with the specified dimensions
// Initially, all cells have walls on all sides
func NewMaze(width, height int) *Maze {
	maze := &Maze{
		Width:  width,
		Height: height,
		Grid:   make([][]Cell, height),
	}

	// Initialize all cells with walls
	for y := 0; y < height; y++ {
		maze.Grid[y] = make([]Cell, width)
		for x := 0; x < width; x++ {
			maze.Grid[y][x] = Cell{
				Walls: [4]bool{true, true, true, true},
			}
		}
	}

	return maze
}

// IsValidPosition checks if the given coordinates are within the maze bounds
func (m *Maze) IsValidPosition(x, y int) bool {
	return x >= 0 && x < m.Width && y >= 0 && y < m.Height
}

// HasWall checks if there's a wall in the specified direction at the given position
func (m *Maze) HasWall(x, y int, direction MazeDirection) bool {
	if !m.IsValidPosition(x, y) {
		return true // Out of bounds is considered a wall
	}
	return m.Grid[y][x].Walls[direction]
}

// RemoveWall removes a wall in the specified direction at the given position
// This also removes the corresponding wall from the adjacent cell
func (m *Maze) RemoveWall(x, y int, direction MazeDirection) {
	if !m.IsValidPosition(x, y) {
		return
	}

	// Remove wall from current cell
	m.Grid[y][x].Walls[direction] = false

	// Remove wall from adjacent cell
	var adjX, adjY int

	switch direction {
	case North:
		adjX, adjY = x, y-1
	case East:
		adjX, adjY = x+1, y
	case South:
		adjX, adjY = x, y+1
	case West:
		adjX, adjY = x-1, y
	}

	if m.IsValidPosition(adjX, adjY) {
		m.Grid[adjY][adjX].Walls[direction.Opposite()] = false
	}
}

// AddWall adds a wall in the specified direction at the given position
func (m *Maze) AddWall(x, y int, direction MazeDirection) {
	if !m.IsValidPosition(x, y) {
		return
	}

	// Add wall to current cell
	m.Grid[y][x].Walls[direction] = true

	// Add wall to adjacent cell
	var adjX, adjY int

	switch direction {
	case North:
		adjX, adjY = x, y-1
	case East:
		adjX, adjY = x+1, y
	case South:
		adjX, adjY = x, y+1
	case West:
		adjX, adjY = x-1, y
	}

	if m.IsValidPosition(adjX, adjY) {
		m.Grid[adjY][adjX].Walls[direction.Opposite()] = true
	}
}

// String returns an ASCII representation of the maze
func (m *Maze) String() string {
	var result strings.Builder

	// Write top border
	for x := 0; x < m.Width; x++ {
		result.WriteString("+--")
	}
	result.WriteString("+\n")

	// For each row
	for y := 0; y < m.Height; y++ {
		// First line: vertical walls and spaces
		result.WriteString("|")
		for x := 0; x < m.Width; x++ {
			// Write two spaces for the cell
			result.WriteString("  ")
			// Write east wall (if present)
			if m.HasWall(x, y, East) {
				result.WriteString("|")
			} else {
				result.WriteString(" ")
			}
		}
		result.WriteString("\n")

		// Second line: horizontal walls and corners
		for x := 0; x < m.Width; x++ {
			result.WriteString("+")
			// Write south wall (if present)
			if m.HasWall(x, y, South) {
				result.WriteString("--")
			} else {
				result.WriteString("  ")
			}
		}
		result.WriteString("+\n")
	}

	return result.String()
}