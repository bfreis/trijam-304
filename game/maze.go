package game

import (
	"fmt"
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

// ParseMaze converts a string representation back into a Maze struct
func ParseMaze(s string) (*Maze, error) {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	if len(lines) < 3 { // Need at least top border + one cell row + bottom border
		return nil, fmt.Errorf("invalid maze string: too few lines")
	}

	// Calculate dimensions
	width := (len(lines[0]) - 1) / 3 // Each cell is 3 chars wide including walls
	height := (len(lines) - 1) / 2   // Each cell is 2 lines tall including walls

	if width < 1 || height < 1 {
		return nil, fmt.Errorf("invalid maze dimensions: width=%d, height=%d", width, height)
	}

	maze := NewMaze(width, height)

	// Process each cell
	for y := 0; y < height; y++ {
		// Check vertical walls (in the cell content line)
		cellLine := lines[y*2+1]
		if len(cellLine) != width*3+1 {
			return nil, fmt.Errorf("invalid line length at y=%d", y)
		}

		for x := 0; x < width; x++ {
			// Check west wall
			if x == 0 {
				if cellLine[0] != '|' {
					maze.RemoveWall(x, y, West)
				}
			}

			// Check east wall
			if cellLine[x*3+3] == ' ' {
				maze.RemoveWall(x, y, East)
			}
		}

		// Check horizontal walls (in the wall line)
		if y < height {
			wallLine := lines[y*2+2]
			if len(wallLine) != width*3+1 {
				return nil, fmt.Errorf("invalid wall line length at y=%d", y)
			}

			for x := 0; x < width; x++ {
				// Check south wall
				if wallLine[x*3+1:x*3+3] == "  " {
					maze.RemoveWall(x, y, South)
				}
			}
		}

		// Check north walls for first row
		if y == 0 {
			topLine := lines[0]
			for x := 0; x < width; x++ {
				if topLine[x*3+1:x*3+3] == "  " {
					maze.RemoveWall(x, y, North)
				}
			}
		}
	}

	return maze, nil
}
