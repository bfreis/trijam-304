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
