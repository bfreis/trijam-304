package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaze(t *testing.T) {
	maze := NewMaze(10, 10)
	_ = maze
}

func TestMazeString(t *testing.T) {
	t.Run("empty maze", func(t *testing.T) {
		maze := NewMaze(3, 3)
		expected := "+--+--+--+\n" +
			"|  |  |  |\n" +
			"+--+--+--+\n" +
			"|  |  |  |\n" +
			"+--+--+--+\n" +
			"|  |  |  |\n" +
			"+--+--+--+\n"
		assert.Equal(t, expected, maze.String(), "maze string representation should match expected output")
	})

	t.Run("maze with some walls removed", func(t *testing.T) {
		maze := NewMaze(2, 2)
		maze.RemoveWall(0, 0, East)
		maze.RemoveWall(0, 1, South)
		maze.RemoveWall(1, 0, South)
		maze.RemoveWall(1, 1, East)

		expected := "+--+--+\n" +
			"|     |\n" +
			"+--+  +\n" +
			"|  |   \n" +
			"+  +--+\n"
		assert.Equal(t, expected, maze.String(), "maze string representation should match expected output")
	})
}

func TestParseMaze(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantWidth   int
		wantHeight  int
		wantErr     bool
		description string
	}{
		{
			name: "1x1 maze with all walls",
			input: `+--+
|  |
+--+`,
			wantWidth:   1,
			wantHeight:  1,
			wantErr:     false,
			description: "Smallest possible valid maze",
		},
		{
			name: "2x2 maze with some walls removed",
			input: `+--+--+
|     |
+  +--+
|  |  |
+--+--+`,
			wantWidth:   2,
			wantHeight:  2,
			wantErr:     false,
			description: "Maze with some walls removed",
		},
		{
			name:        "Empty string",
			input:       "",
			wantErr:     true,
			description: "Should fail on empty input",
		},
		{
			name: "Invalid line length",
			input: `+--+
|  |  |
+--+`,
			wantErr:     true,
			description: "Inconsistent line lengths",
		},
		{
			name: "Missing bottom border",
			input: `+--+
|  |`,
			wantErr:     true,
			description: "Incomplete maze structure",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMaze(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMaze() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			if got.Width != tt.wantWidth {
				t.Errorf("ParseMaze() width = %v, want %v", got.Width, tt.wantWidth)
			}
			if got.Height != tt.wantHeight {
				t.Errorf("ParseMaze() height = %v, want %v", got.Height, tt.wantHeight)
			}
		})
	}
}

func TestMazeRoundTrip(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		modifyFn func(*Maze) // Function to modify the maze before testing
	}{
		{
			name:   "1x1 maze unmodified",
			width:  1,
			height: 1,
		},
		{
			name:   "2x2 maze with removed walls",
			width:  2,
			height: 2,
			modifyFn: func(m *Maze) {
				m.RemoveWall(0, 0, East)
				m.RemoveWall(0, 1, South)
			},
		},
		{
			name:   "3x3 maze with complex pattern",
			width:  3,
			height: 3,
			modifyFn: func(m *Maze) {
				// Create a simple path through the maze
				m.RemoveWall(0, 0, East)
				m.RemoveWall(1, 0, South)
				m.RemoveWall(1, 1, East)
				m.RemoveWall(2, 1, South)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create original maze
			original := NewMaze(tt.width, tt.height)
			if tt.modifyFn != nil {
				tt.modifyFn(original)
			}

			// Convert to string
			str := original.String()

			// Parse back to maze
			parsed, err := ParseMaze(str)
			if err != nil {
				t.Fatalf("Failed to parse maze: %v", err)
			}

			// Convert back to string
			roundTrip := parsed.String()

			// Compare strings
			if str != roundTrip {
				t.Errorf("Round trip conversion failed\nOriginal:\n%s\nGot:\n%s", str, roundTrip)
			}

			// Compare dimensions
			if original.Width != parsed.Width || original.Height != parsed.Height {
				t.Errorf("Dimensions mismatch after round trip: original(%dx%d) != parsed(%dx%d)",
					original.Width, original.Height, parsed.Width, parsed.Height)
			}

			// Compare wall structure
			for y := 0; y < original.Height; y++ {
				for x := 0; x < original.Width; x++ {
					for d := North; d <= West; d++ {
						if original.HasWall(x, y, d) != parsed.HasWall(x, y, d) {
							t.Errorf("Wall mismatch at (%d,%d) direction %s", x, y, d)
						}
					}
				}
			}
		})
	}
}
