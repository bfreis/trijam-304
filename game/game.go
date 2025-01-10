package game

import (
	"image/color"

	"github.com/bfreis/ebitentools/ebitenwrap"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	mazeCellDisplayHeight = 32
	mazeCellDisplayWidth  = 32
	wallThickness         = 2.0
)

type Game struct {
	maze *Maze
}

func NewGame() (*Game, error) {
	const mazeString = `+--+--+--+
|  |  |  |
+--+--+--+
|     |  |
+--+--+--+
|  |  |  |
+--+--+--+`
	maze, err := ParseMaze(mazeString)
	if err != nil {
		return nil, err
	}

	return &Game{
		maze: maze,
	}, nil
}

// Draw implements ebitenwrap.Game interface
func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen with a dark background
	screen.Fill(color.RGBA{40, 40, 40, 255})

	// Use constant cell sizes
	cellWidth := float64(mazeCellDisplayWidth)
	cellHeight := float64(mazeCellDisplayHeight)

	// Calculate total maze dimensions
	mazeWidth := float64(g.maze.Width) * cellWidth
	mazeHeight := float64(g.maze.Height) * cellHeight

	// Calculate offsets to center the maze
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	offsetX := (float64(sw) - mazeWidth) / 2
	offsetY := (float64(sh) - mazeHeight) / 2

	// Draw each cell's walls
	for y := 0; y < g.maze.Height; y++ {
		for x := 0; x < g.maze.Width; x++ {
			// Calculate cell position with offset for centering
			px := offsetX + float64(x)*cellWidth
			py := offsetY + float64(y)*cellHeight

			// Draw walls if they exist
			if g.maze.HasWall(x, y, North) {
				vector.StrokeLine(screen, float32(px), float32(py), float32(px+cellWidth), float32(py), float32(wallThickness), color.White, false)
			}
			if g.maze.HasWall(x, y, East) {
				vector.StrokeLine(screen, float32(px+cellWidth), float32(py), float32(px+cellWidth), float32(py+cellHeight), float32(wallThickness), color.White, false)
			}
			if g.maze.HasWall(x, y, South) {
				vector.StrokeLine(screen, float32(px), float32(py+cellHeight), float32(px+cellWidth), float32(py+cellHeight), float32(wallThickness), color.White, false)
			}
			if g.maze.HasWall(x, y, West) {
				vector.StrokeLine(screen, float32(px), float32(py), float32(px), float32(py+cellHeight), float32(wallThickness), color.White, false)
			}
		}
	}
}

func (g *Game) Update(tick ebitenwrap.Tick) error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
